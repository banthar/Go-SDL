package gfx

// #include <SDL/SDL_gfxBlitFunc.h>
import "C"

import "sdl"

func BlitRGBA(dst *sdl.Surface, dstrect *sdl.Rect, src *sdl.Surface, srcrect *sdl.Rect) int {
	return int( C.SDL_gfxBlitRGBA(
		(*C.SDL_Surface)(ptr(src)),
		(*C.SDL_Rect)(ptr(srcrect)),
		(*C.SDL_Surface)(ptr(dst)),
		(*C.SDL_Rect)(ptr(dstrect))) )
}

func SetAlpha(src *sdl.Surface, a uint8) int {
	return int( C.SDL_gfxSetAlpha((*C.SDL_Surface)(ptr(src)), C.Uint8(a)) )
}

func MultiplyAlpha(src *sdl.Surface, a uint8) int {
	return int( C.SDL_gfxMultiplyAlpha((*C.SDL_Surface)(ptr(src)), C.Uint8(a)) )
}
