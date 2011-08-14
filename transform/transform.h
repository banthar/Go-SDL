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

/* Pentium MMX/SSE smoothscale routines
 * Available on Win32 or GCC on a Pentium.
 * Sorry, no Win64 support yet for Visual C builds, but it can be added.
 */

#ifndef TRANSFORM_H
# define TRANSFORM_H
#include <SDL/SDL.h>
#define SCALE_MMX_SUPPORT

SDL_Surface* scalesmooth(SDL_Surface *surfobj2, SDL_Surface *surf, int width, int height);
const char* get_smoothscale_backend();
SDL_Surface * surf_flip (SDL_Surface* surf, int xaxis, int yaxis);

/*void convert_32_24(Uint8 *srcpix, int srcpitch, Uint8 *dstpix, int dstpitch, int width, int height);
void convert_24_32(Uint8 *srcpix, int srcpitch, Uint8 *dstpix, int dstpitch, int width, int height);
void filter_shrink_X_ONLYC(Uint8 *srcpix, Uint8 *dstpix, int height, int srcpitch, int dstpitch, int srcwidth, int dstwidth);
void filter_shrink_Y_ONLYC(Uint8 *srcpix, Uint8 *dstpix, int height, int srcpitch, int dstpitch, int srcwidth, int dstwidth);
void filter_expand_X_ONLYC(Uint8 *srcpix, Uint8 *dstpix, int height, int srcpitch, int dstpitch, int srcwidth, int dstwidth);
void filter_expand_Y_ONLYC(Uint8 *srcpix, Uint8 *dstpix, int height, int srcpitch, int dstpitch, int srcwidth, int dstwidth);
*/

/* These functions implement an area-averaging shrinking filter in the X-dimension.
 */
void filter_shrink_X_MMX(Uint8 *srcpix, Uint8 *dstpix, int height, int srcpitch, int dstpitch, int srcwidth, int dstwidth);

void filter_shrink_X_SSE(Uint8 *srcpix, Uint8 *dstpix, int height, int srcpitch, int dstpitch, int srcwidth, int dstwidth);

/* These functions implement an area-averaging shrinking filter in the Y-dimension.
 */
void filter_shrink_Y_MMX(Uint8 *srcpix, Uint8 *dstpix, int width, int srcpitch, int dstpitch, int srcheight, int dstheight);

void filter_shrink_Y_SSE(Uint8 *srcpix, Uint8 *dstpix, int width, int srcpitch, int dstpitch, int srcheight, int dstheight);

/* These functions implement a bilinear filter in the X-dimension.
 */
void filter_expand_X_MMX(Uint8 *srcpix, Uint8 *dstpix, int height, int srcpitch, int dstpitch, int srcwidth, int dstwidth);

void filter_expand_X_SSE(Uint8 *srcpix, Uint8 *dstpix, int height, int srcpitch, int dstpitch, int srcwidth, int dstwidth);

/* These functions implement a bilinear filter in the Y-dimension.
 */
void filter_expand_Y_MMX(Uint8 *srcpix, Uint8 *dstpix, int width, int srcpitch, int dstpitch, int srcheight, int dstheight);

void filter_expand_Y_SSE(Uint8 *srcpix, Uint8 *dstpix, int width, int srcpitch, int dstpitch, int srcheight, int dstheight);

#endif
