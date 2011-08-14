/*
  pygame - Python Game Library
  Copyright (C) 2000-2001  Pete Shinners
  Copyright (C) 2007  Rene Dudfield, Richard Goedeken 

  This library is free software; you can redistribute it and/or
  modify it under the terms of the GNU Library General Public
  License as published by the Free Software Foundation; either
  version 2 of the License, or (at your option) any later version.

  This library is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
  Library General Public License for more details.

  You should have received a copy of the GNU Library General Public
  License along with this library; if not, write to the Free
  Foundation, Inc., 59 Temple Place, Suite 330, Boston, MA  02111-1307  USA

  Pete Shinners
  pete@shinners.org
*/

/*
 *  surface transformations for pygame
 */


/* this function implements an area-averaging shrinking filter in the X-dimension */
#include <math.h>
#include <string.h>
#include "transform.h"

typedef void (* SMOOTHSCALE_FILTER_P)(Uint8 *, Uint8 *, int, int, int, int, int);
struct _module_state {
    const char *filter_type;
    SMOOTHSCALE_FILTER_P filter_shrink_X;
    SMOOTHSCALE_FILTER_P filter_shrink_Y;
    SMOOTHSCALE_FILTER_P filter_expand_X;
    SMOOTHSCALE_FILTER_P filter_expand_Y;
};

static struct _module_state *st, g_st;

void filter_shrink_X_ONLYC(Uint8 *, Uint8 *, int, int, int, int, int);
void filter_shrink_Y_ONLYC(Uint8 *, Uint8 *, int, int, int, int, int);
void filter_expand_X_ONLYC(Uint8 *, Uint8 *, int, int, int, int, int);
void filter_expand_Y_ONLYC(Uint8 *, Uint8 *, int, int, int, int, int);

static SDL_Surface* newsurf_fromsurf (SDL_Surface* surf, int width, int height);

static void convert_24_32(Uint8 *srcpix, int srcpitch, Uint8 *dstpix, int dstpitch, int width, int height)
{
    int srcdiff = srcpitch - (width * 3);
    int dstdiff = dstpitch - (width * 4);
    int x, y;

    for (y = 0; y < height; y++)
    {
        for (x = 0; x < width; x++)
        {
            *dstpix++ = *srcpix++;
            *dstpix++ = *srcpix++;
            *dstpix++ = *srcpix++;
            *dstpix++ = 0xff;
        }
        srcpix += srcdiff;
        dstpix += dstdiff;
    }
}

static void convert_32_24(Uint8 *srcpix, int srcpitch, Uint8 *dstpix, int dstpitch, int width, int height)
{
    int srcdiff = srcpitch - (width * 4);
    int dstdiff = dstpitch - (width * 3);
    int x, y;

    for (y = 0; y < height; y++)
    {
        for (x = 0; x < width; x++)
        {
            *dstpix++ = *srcpix++;
            *dstpix++ = *srcpix++;
            *dstpix++ = *srcpix++;
            srcpix++;
        }
        srcpix += srcdiff;
        dstpix += dstdiff;
    }
}

static void smoothscale_init ()
{
	if (SDL_HasSSE ())
	{
	    st->filter_type = "SSE";
	    st->filter_shrink_X = filter_shrink_X_SSE;
	    st->filter_shrink_Y = filter_shrink_Y_SSE;
	    st->filter_expand_X = filter_expand_X_SSE;
	    st->filter_expand_Y = filter_expand_Y_SSE;
	}
	else if (SDL_HasMMX ())
	{
	    st->filter_type = "MMX";
	    st->filter_shrink_X = filter_shrink_X_MMX;
	    st->filter_shrink_Y = filter_shrink_Y_MMX;
	    st->filter_expand_X = filter_expand_X_MMX;
	    st->filter_expand_Y = filter_expand_Y_MMX;
	}
	else
	{
	    st->filter_type = "GENERIC";
	    st->filter_shrink_X = filter_shrink_X_ONLYC;
	    st->filter_shrink_Y = filter_shrink_Y_ONLYC;
	    st->filter_expand_X = filter_expand_X_ONLYC;
	    st->filter_expand_Y = filter_expand_Y_ONLYC;
	}
}

const char* get_smoothscale_backend() {
	if (st == NULL) { st = &g_st; smoothscale_init(); }
	return st->filter_type;
}

static void _scalesmooth(SDL_Surface *src, SDL_Surface *dst)
{
	if (st == NULL) { st = &g_st; smoothscale_init(); }

    Uint8* srcpix = (Uint8*)src->pixels;
    Uint8* dstpix = (Uint8*)dst->pixels;
    Uint8* dst32 = NULL;
    int srcpitch = src->pitch;
    int dstpitch = dst->pitch;

    int srcwidth = src->w;
    int srcheight = src->h;
    int dstwidth = dst->w;
    int dstheight = dst->h;

    int bpp = src->format->BytesPerPixel;

    Uint8 *temppix = NULL;
    int tempwidth=0, temppitch=0, tempheight=0;

    /* convert to 32-bit if necessary */
    if (bpp == 3)
    {
        int newpitch = srcwidth * 4;
        Uint8 *newsrc = (Uint8 *) malloc(newpitch * srcheight);
        if (!newsrc)
            return;
        convert_24_32(srcpix, srcpitch, newsrc, newpitch, srcwidth, srcheight);
        srcpix = newsrc;
        srcpitch = newpitch;
        /* create a destination buffer for the 32-bit result */
        dstpitch = dstwidth << 2;
        dst32 = (Uint8 *) malloc(dstpitch * dstheight);
        if (dst32 == NULL)
        {
            free(srcpix);
            return;
        }
        dstpix = dst32;
    }

    /* Create a temporary processing buffer if we will be scaling both X and Y */
    if (srcwidth != dstwidth && srcheight != dstheight)
    {
        tempwidth = dstwidth;
        temppitch = tempwidth << 2;
        tempheight = srcheight;
        temppix = (Uint8 *) malloc(temppitch * tempheight);
        if (temppix == NULL)
        {
            if (bpp == 3)
            {
                free(srcpix);
                free(dstpix);
            }
            return;
        }
    }

    /* Start the filter by doing X-scaling */
    if (dstwidth < srcwidth) /* shrink */
    {
        if (srcheight != dstheight)
            st->filter_shrink_X(srcpix, temppix, srcheight, srcpitch, temppitch, srcwidth, dstwidth);
        else
            st->filter_shrink_X(srcpix, dstpix, srcheight, srcpitch, dstpitch, srcwidth, dstwidth);
    }
    else if (dstwidth > srcwidth) /* expand */
    {
        if (srcheight != dstheight)
            st->filter_expand_X(srcpix, temppix, srcheight, srcpitch, temppitch, srcwidth, dstwidth);
        else
            st->filter_expand_X(srcpix, dstpix, srcheight, srcpitch, dstpitch, srcwidth, dstwidth);
    }
    /* Now do the Y scale */
    if (dstheight < srcheight) /* shrink */
    {
        if (srcwidth != dstwidth)
            st->filter_shrink_Y(temppix, dstpix, tempwidth, temppitch, dstpitch, srcheight, dstheight);
        else
            st->filter_shrink_Y(srcpix, dstpix, srcwidth, srcpitch, dstpitch, srcheight, dstheight);
    }
    else if (dstheight > srcheight)  /* expand */
    {
        if (srcwidth != dstwidth)
            st->filter_expand_Y(temppix, dstpix, tempwidth, temppitch, dstpitch, srcheight, dstheight);
        else
            st->filter_expand_Y(srcpix, dstpix, srcwidth, srcpitch, dstpitch, srcheight, dstheight);
    }

    /* Convert back to 24-bit if necessary */
    if (bpp == 3)
    {
        convert_32_24(dst32, dstpitch, (Uint8*)dst->pixels, dst->pitch, dstwidth, dstheight);
        free(dst32);
        dst32 = NULL;
        free(srcpix);
        srcpix = NULL;
    }
    /* free temporary buffer if necessary */
    if (temppix != NULL)
        free(temppix);

}

SDL_Surface* scalesmooth(SDL_Surface *surfobj2, SDL_Surface *surf, int width, int height)
{
	int bpp;
	SDL_Surface *newsurf;

    if (width < 0 || height < 0)
        return NULL;


    bpp = surf->format->BytesPerPixel;
    if(bpp < 3 || bpp > 4)
		return NULL;

	
    if (!surfobj2)
    {
        newsurf = newsurf_fromsurf (surf, width, height); //FIXME: Free this surface on failure.
        if (!newsurf)
            return NULL;
    } else {
		newsurf = surfobj2;
	}

    if (newsurf->w != width || newsurf->h != height)
		return NULL;


    if(((width * bpp + 3) >> 2) > newsurf->pitch)
        return NULL;


    if(width && height)
    {
        SDL_LockSurface(newsurf);

        /* handle trivial case */
        if (surf->w == width && surf->h == height) {
            int y;
            for (y = 0; y < height; y++) {
                memcpy((Uint8*)newsurf->pixels + y * newsurf->pitch, 
                       (Uint8*)surf->pixels + y * surf->pitch, width * bpp);
            }
        }
        else {
           _scalesmooth(surf, newsurf);
        }

        SDL_UnlockSurface(newsurf);
    }

    if (surfobj2)
    {
        return surfobj2;
    }
    else
        return newsurf;

}


static SDL_Surface*
newsurf_fromsurf (SDL_Surface* surf, int width, int height)
{
    SDL_Surface* newsurf;
    int result;

    if (surf->format->BytesPerPixel <= 0 || surf->format->BytesPerPixel > 4)
        return NULL;

    newsurf = SDL_CreateRGBSurface (surf->flags, width, height,
                                    surf->format->BitsPerPixel,
                                    surf->format->Rmask, surf->format->Gmask,
                                    surf->format->Bmask, surf->format->Amask);
    if (!newsurf)
        return NULL;

    /* Copy palette, colorkey, etc info */
    if (surf->format->BytesPerPixel==1 && surf->format->palette)
        SDL_SetColors (newsurf, surf->format->palette->colors, 0,
                       surf->format->palette->ncolors);
    if (surf->flags & SDL_SRCCOLORKEY)
        SDL_SetColorKey (newsurf, (surf->flags&SDL_RLEACCEL) | SDL_SRCCOLORKEY,
                         surf->format->colorkey);

    if (surf->flags&SDL_SRCALPHA)
    {
        result = SDL_SetAlpha (newsurf, surf->flags, surf->format->alpha);
        if (result == -1)
            return NULL;
    }
    return newsurf;
}

/*
 * smooth scale functions.
 */


/* this function implements an area-averaging shrinking filter in the X-dimension */
void filter_shrink_X_ONLYC(Uint8 *srcpix, Uint8 *dstpix, int height, int srcpitch, int dstpitch, int srcwidth, int dstwidth)
{
    int srcdiff = srcpitch - (srcwidth * 4);
    int dstdiff = dstpitch - (dstwidth * 4);
    int x, y;

    int xspace = 0x10000 * srcwidth / dstwidth; /* must be > 1 */
    int xrecip = (int) (0x100000000LL / xspace);
    for (y = 0; y < height; y++)
    {
        Uint16 accumulate[4] = {0,0,0,0};
        int xcounter = xspace;
        for (x = 0; x < srcwidth; x++)
        {
            if (xcounter > 0x10000)
            {
                accumulate[0] += (Uint16) *srcpix++;
                accumulate[1] += (Uint16) *srcpix++;
                accumulate[2] += (Uint16) *srcpix++;
                accumulate[3] += (Uint16) *srcpix++;
                xcounter -= 0x10000;
            }
            else
            {
                int xfrac = 0x10000 - xcounter;
                /* write out a destination pixel */
                *dstpix++ = (Uint8) (((accumulate[0] + ((srcpix[0] * xcounter) >> 16)) * xrecip) >> 16);
                *dstpix++ = (Uint8) (((accumulate[1] + ((srcpix[1] * xcounter) >> 16)) * xrecip) >> 16);
                *dstpix++ = (Uint8) (((accumulate[2] + ((srcpix[2] * xcounter) >> 16)) * xrecip) >> 16);
                *dstpix++ = (Uint8) (((accumulate[3] + ((srcpix[3] * xcounter) >> 16)) * xrecip) >> 16);
                /* reload the accumulator with the remainder of this pixel */
                accumulate[0] = (Uint16) ((*srcpix++ * xfrac) >> 16);
                accumulate[1] = (Uint16) ((*srcpix++ * xfrac) >> 16);
                accumulate[2] = (Uint16) ((*srcpix++ * xfrac) >> 16);
                accumulate[3] = (Uint16) ((*srcpix++ * xfrac) >> 16);
                xcounter = xspace - xfrac;
            }
        }
        srcpix += srcdiff;
        dstpix += dstdiff;
    }
}

/* this function implements an area-averaging shrinking filter in the Y-dimension */
void filter_shrink_Y_ONLYC(Uint8 *srcpix, Uint8 *dstpix, int width, int srcpitch, int dstpitch, int srcheight, int dstheight)
{
    Uint16 *templine;
    int srcdiff = srcpitch - (width * 4);
    int dstdiff = dstpitch - (width * 4);
    int x, y;
    int yspace = 0x10000 * srcheight / dstheight; /* must be > 1 */
    int yrecip = (int) (0x100000000LL / yspace);
    int ycounter = yspace;

    /* allocate and clear a memory area for storing the accumulator line */
    templine = (Uint16 *) malloc(dstpitch * 2);
    if (templine == NULL) return;
    memset(templine, 0, dstpitch * 2);

    for (y = 0; y < srcheight; y++)
    {
        Uint16 *accumulate = templine;
        if (ycounter > 0x10000)
        {
            for (x = 0; x < width; x++)
            {
                *accumulate++ += (Uint16) *srcpix++;
                *accumulate++ += (Uint16) *srcpix++;
                *accumulate++ += (Uint16) *srcpix++;
                *accumulate++ += (Uint16) *srcpix++;
            }
            ycounter -= 0x10000;
        }
        else
        {
            int yfrac = 0x10000 - ycounter;
            /* write out a destination line */
            for (x = 0; x < width; x++)
            {
                *dstpix++ = (Uint8) (((*accumulate++ + ((*srcpix++ * ycounter) >> 16)) * yrecip) >> 16);
                *dstpix++ = (Uint8) (((*accumulate++ + ((*srcpix++ * ycounter) >> 16)) * yrecip) >> 16);
                *dstpix++ = (Uint8) (((*accumulate++ + ((*srcpix++ * ycounter) >> 16)) * yrecip) >> 16);
                *dstpix++ = (Uint8) (((*accumulate++ + ((*srcpix++ * ycounter) >> 16)) * yrecip) >> 16);
            }
            dstpix += dstdiff;
            /* reload the accumulator with the remainder of this line */
            accumulate = templine;
            srcpix -= 4 * width;
            for (x = 0; x < width; x++)
            {
                *accumulate++ = (Uint16) ((*srcpix++ * yfrac) >> 16);
                *accumulate++ = (Uint16) ((*srcpix++ * yfrac) >> 16);
                *accumulate++ = (Uint16) ((*srcpix++ * yfrac) >> 16);
                *accumulate++ = (Uint16) ((*srcpix++ * yfrac) >> 16);
            }
            ycounter = yspace - yfrac;
        }
        srcpix += srcdiff;
    } /* for (int y = 0; y < srcheight; y++) */

    /* free the temporary memory */
    free(templine);
}

/* this function implements a bilinear filter in the X-dimension */
void filter_expand_X_ONLYC(Uint8 *srcpix, Uint8 *dstpix, int height, int srcpitch, int dstpitch, int srcwidth, int dstwidth)
{
    int dstdiff = dstpitch - (dstwidth * 4);
    int *xidx0, *xmult0, *xmult1;
    int x, y;
    int factorwidth = 4;

    /* Allocate memory for factors */
    xidx0 = malloc(dstwidth * 4);
    if (xidx0 == NULL) return;
    xmult0 = (int *) malloc(dstwidth * factorwidth);
    xmult1 = (int *) malloc(dstwidth * factorwidth);
    if (xmult0 == NULL || xmult1 == NULL)
    {
        free(xidx0);
        if (xmult0) free(xmult0);
        if (xmult1) free(xmult1);
    }

    /* Create multiplier factors and starting indices and put them in arrays */
    for (x = 0; x < dstwidth; x++)
    {
        xidx0[x] = x * (srcwidth - 1) / dstwidth;
        xmult1[x] = 0x10000 * ((x * (srcwidth - 1)) % dstwidth) / dstwidth;
        xmult0[x] = 0x10000 - xmult1[x];
    }

    /* Do the scaling in raster order so we don't trash the cache */
    for (y = 0; y < height; y++)
    {
        Uint8 *srcrow0 = srcpix + y * srcpitch;
        for (x = 0; x < dstwidth; x++)
        {
            Uint8 *src = srcrow0 + xidx0[x] * 4;
            int xm0 = xmult0[x];
            int xm1 = xmult1[x];
            *dstpix++ = (Uint8) (((src[0] * xm0) + (src[4] * xm1)) >> 16);
            *dstpix++ = (Uint8) (((src[1] * xm0) + (src[5] * xm1)) >> 16);
            *dstpix++ = (Uint8) (((src[2] * xm0) + (src[6] * xm1)) >> 16);
            *dstpix++ = (Uint8) (((src[3] * xm0) + (src[7] * xm1)) >> 16);
        }
        dstpix += dstdiff;
    }

    /* free memory */
    free(xidx0);
    free(xmult0);
    free(xmult1);
}

/* this function implements a bilinear filter in the Y-dimension */
void filter_expand_Y_ONLYC(Uint8 *srcpix, Uint8 *dstpix, int width, int srcpitch, int dstpitch, int srcheight, int dstheight)
{
    int x, y;

    for (y = 0; y < dstheight; y++)
    {
        int yidx0 = y * (srcheight - 1) / dstheight;
        Uint8 *srcrow0 = srcpix + yidx0 * srcpitch;
        Uint8 *srcrow1 = srcrow0 + srcpitch;
        int ymult1 = 0x10000 * ((y * (srcheight - 1)) % dstheight) / dstheight;
        int ymult0 = 0x10000 - ymult1;
        for (x = 0; x < width; x++)
        {
            *dstpix++ = (Uint8) (((*srcrow0++ * ymult0) + (*srcrow1++ * ymult1)) >> 16);
            *dstpix++ = (Uint8) (((*srcrow0++ * ymult0) + (*srcrow1++ * ymult1)) >> 16);
            *dstpix++ = (Uint8) (((*srcrow0++ * ymult0) + (*srcrow1++ * ymult1)) >> 16);
            *dstpix++ = (Uint8) (((*srcrow0++ * ymult0) + (*srcrow1++ * ymult1)) >> 16);
        }
    }
}

SDL_Surface * surf_flip (SDL_Surface* surf, int xaxis, int yaxis)
{
    SDL_Surface *newsurf;
    int loopx, loopy;
    int pixsize, srcpitch, dstpitch;
    Uint8 *srcpix, *dstpix;

    newsurf = newsurf_fromsurf (surf, surf->w, surf->h);
    if (!newsurf)
        return NULL;

    pixsize = surf->format->BytesPerPixel;
    srcpitch = surf->pitch;
    dstpitch = newsurf->pitch;

    SDL_LockSurface (newsurf);

    srcpix = (Uint8*) surf->pixels;
    dstpix = (Uint8*) newsurf->pixels;

    if (!xaxis)
    {
        if (!yaxis)
        {
            for (loopy = 0; loopy < surf->h; ++loopy)
                memcpy (dstpix + loopy * dstpitch, srcpix + loopy * srcpitch,
                        surf->w * surf->format->BytesPerPixel);
            }
            else
            {
                for (loopy = 0; loopy < surf->h; ++loopy)
                    memcpy (dstpix + loopy * dstpitch,
                            srcpix + (surf->h - 1 - loopy) * srcpitch,
                            surf->w * surf->format->BytesPerPixel);
            }
	}
	else /*if (xaxis)*/
	{
            if (yaxis)
            {
                switch (surf->format->BytesPerPixel)
                {
                case 1:
                    for (loopy = 0; loopy < surf->h; ++loopy)
                    {
                        Uint8* dst = (Uint8*) (dstpix + loopy * dstpitch);
                        Uint8* src = ((Uint8*) (srcpix + (surf->h - 1 - loopy)
                                                * srcpitch)) + surf->w - 1;
                        for (loopx = 0; loopx < surf->w; ++loopx)
                            *dst++ = *src--;
                    }
                    break;
                case 2:
                    for (loopy = 0; loopy < surf->h; ++loopy)
                    {
                        Uint16* dst = (Uint16*) (dstpix + loopy * dstpitch);
                        Uint16* src = ((Uint16*)
                                       (srcpix + (surf->h - 1 - loopy)
                                        * srcpitch)) + surf->w - 1;
                        for (loopx = 0; loopx < surf->w; ++loopx)
                            *dst++ = *src--;
                    }
                    break;
                case 4:
                    for (loopy = 0; loopy < surf->h; ++loopy)
                    {
                        Uint32* dst = (Uint32*) (dstpix + loopy * dstpitch);
                        Uint32* src = ((Uint32*)
                                       (srcpix + (surf->h - 1 - loopy)
                                        * srcpitch)) + surf->w - 1;
                        for (loopx = 0; loopx < surf->w; ++loopx)
                            *dst++ = *src--;
                    }
                    break;
                case 3:
                    for (loopy = 0; loopy < surf->h; ++loopy)
                    {
                        Uint8* dst = (Uint8*) (dstpix + loopy * dstpitch);
                        Uint8* src = ((Uint8*) (srcpix + (surf->h - 1 - loopy)
                                                * srcpitch)) + surf->w * 3 - 3;
                        for (loopx = 0; loopx < surf->w; ++loopx)
                        {
                            dst[0] = src[0];
                            dst[1] = src[1];
                            dst[2] = src[2];
                            dst += 3;
                            src -= 3;
                        }
                    }
                    break;
                }
            }
            else
            {
                switch (surf->format->BytesPerPixel)
                {
                case 1:
                    for (loopy = 0; loopy < surf->h; ++loopy)
                    {
                        Uint8* dst = (Uint8*) (dstpix + loopy * dstpitch);
                        Uint8* src = ((Uint8*) (srcpix + loopy * srcpitch)) +
                            surf->w - 1;
                        for (loopx = 0; loopx < surf->w; ++loopx)
                            *dst++ = *src--;
                    }
                    break;
                case 2:
                    for (loopy = 0; loopy < surf->h; ++loopy)
                    {
                        Uint16* dst = (Uint16*) (dstpix + loopy * dstpitch);
                        Uint16* src = ((Uint16*) (srcpix + loopy * srcpitch))
                            + surf->w - 1;
                        for (loopx = 0; loopx < surf->w; ++loopx)
                            *dst++ = *src--;
                    }
                    break;
                case 4:
                    for (loopy = 0; loopy < surf->h; ++loopy)
                    {
                        Uint32* dst = (Uint32*) (dstpix + loopy * dstpitch);
                        Uint32* src = ((Uint32*) (srcpix + loopy * srcpitch))
                            + surf->w - 1;
                        for (loopx = 0; loopx < surf->w; ++loopx)
                            *dst++ = *src--;
                    }
                    break;
                case 3:
                    for (loopy = 0; loopy < surf->h; ++loopy)
                    {
                        Uint8* dst = (Uint8*) (dstpix + loopy * dstpitch);
                        Uint8* src = ((Uint8*) (srcpix + loopy * srcpitch))
                            + surf->w * 3 - 3;
                        for (loopx = 0; loopx < surf->w; ++loopx)
                        {
                            dst[0] = src[0];
                            dst[1] = src[1];
                            dst[2] = src[2];
                            dst += 3;
                            src -= 3;
                        }
                    }
                    break;
                }
            }
	}

    SDL_UnlockSurface (newsurf);
    return newsurf;
}
