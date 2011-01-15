package gfx

// #include <SDL/SDL_rotozoom.h>
import "C"

import "unsafe"
import "sdl"

type ptr unsafe.Pointer

func b2ci(b bool) C.int {
	if b {
		return C.int(0)
	}
	return C.int(1)
}

func RotozoomSurface(src *sdl.Surface, angle, zoom float64, smooth bool) *sdl.Surface {
	return (*sdl.Surface)(ptr( C.rotozoomSurface((*C.SDL_Surface)(ptr(src)), C.double(angle), C.double(zoom), b2ci(smooth)) ))
}

func RotozoomSurfaceXY(src *sdl.Surface, angle, zoomx, zoomy float64, smooth bool) *sdl.Surface {
	return (*sdl.Surface)(ptr( C.rotozoomSurfaceXY((*C.SDL_Surface)(ptr(src)), C.double(angle), C.double(zoomx), C.double(zoomy), b2ci(smooth)) ))
}

func RotozoomSurfaceSize(width, height int, angle, zoom float64) (dstwidth, dstheight int) {
	var dw, dh C.int
	C.rotozoomSurfaceSize(C.int(width), C.int(height), C.double(angle), C.double(zoom), &dw, &dh)
	return int(dw), int(dh)
}

func RotozoomSurfaceSizeXY(width, height int, angle, zoomx, zoomy float64) (dstwidth, dstheight int) {
	var dw, dh C.int
	C.rotozoomSurfaceSizeXY(C.int(width), C.int(height), C.double(angle), C.double(zoomx), C.double(zoomy), &dw, &dh)
	return int(dw), int(dh)
}

func ZoomSurface(src *sdl.Surface, zoomx, zoomy float64, smooth bool) *sdl.Surface {
	return (*sdl.Surface)(ptr( C.zoomSurface((*C.SDL_Surface)(ptr(src)), C.double(zoomx), C.double(zoomy), b2ci(smooth)) ))
}

func ZoomSurfaceSize(width, height int, zoomx, zoomy float64) (dstwidth, dstheight int) {
	var dw, dh C.int
	C.zoomSurfaceSize(C.int(width), C.int(height), C.double(zoomx), C.double(zoomy), &dw, &dh)
	return int(dw), int(dh)
}

func ShrinkSurface(src *sdl.Surface, factorx, factory int) *sdl.Surface {
	return (*sdl.Surface)(ptr( C.shrinkSurface((*C.SDL_Surface)(ptr(src)), C.int(factorx), C.int(factory)) ))
}

func RotateSurface90Degrees(src *sdl.Surface, numClockwiseTurns int) *sdl.Surface {
	return (*sdl.Surface)(ptr( C.rotateSurface90Degrees((*C.SDL_Surface)(ptr(src)), C.int(numClockwiseTurns)) ))
}
