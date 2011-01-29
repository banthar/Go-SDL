package transform

// #include "transform.h"
import "C"

import (
	"sdl"
	"unsafe"
)

func SmoothScaleBackend() string {
	return C.GoString(C.get_smoothscale_backend())
}

func SmoothScale(s, d *sdl.Surface) {
	C.scalesmooth( (*C.SDL_Surface)(unsafe.Pointer(s)), (*C.SDL_Surface)(unsafe.Pointer(d)) )
}

func b2ci(b bool) C.int {
	if b { return 1 }
	return 0
}

func Flip(s *sdl.Surface, x, y bool) *sdl.Surface {
	return (*sdl.Surface)(unsafe.Pointer( C.surf_flip( (*C.SDL_Surface)(unsafe.Pointer(s)), b2ci(x), b2ci(y)) ))
}
