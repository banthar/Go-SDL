package transform

// #include "transform.h"
import "C"

import (
	"sdl"
	"unsafe"
)

var shrink_X, shrink_Y, expand_X, expand_Y func(spixels, dpixels *byte, sh int32, spitch, dpitch uint16, sw, dw int32)

func filter_shrink_X_MMX(spixels, dpixels *byte, sh int32, spitch, dpitch uint16, sw, dw int32) {
	C.filter_shrink_X_MMX( (*C.Uint8)(spixels), (*C.Uint8)(dpixels), C.int(sh), C.int(spitch), C.int(dpitch), C.int(sw), C.int(dw))
}

func filter_shrink_Y_MMX(spixels, dpixels *byte, sh int32, spitch, dpitch uint16, sw, dw int32) {
	C.filter_shrink_Y_MMX( (*C.Uint8)(spixels), (*C.Uint8)(dpixels), C.int(sh), C.int(spitch), C.int(dpitch), C.int(sw), C.int(dw))
}

func filter_expand_X_MMX(spixels, dpixels *byte, sh int32, spitch, dpitch uint16, sw, dw int32) {
	C.filter_expand_X_MMX( (*C.Uint8)(spixels), (*C.Uint8)(dpixels), C.int(sh), C.int(spitch), C.int(dpitch), C.int(sw), C.int(dw))
}

func filter_expand_Y_MMX(spixels, dpixels *byte, sh int32, spitch, dpitch uint16, sw, dw int32) {
	C.filter_expand_Y_MMX( (*C.Uint8)(spixels), (*C.Uint8)(dpixels), C.int(sh), C.int(spitch), C.int(dpitch), C.int(sw), C.int(dw))
}

func filter_shrink_X_SSE(spixels, dpixels *byte, sh int32, spitch, dpitch uint16, sw, dw int32) {
	C.filter_shrink_X_SSE( (*C.Uint8)(spixels), (*C.Uint8)(dpixels), C.int(sh), C.int(spitch), C.int(dpitch), C.int(sw), C.int(dw))
}

func filter_shrink_Y_SSE(spixels, dpixels *byte, sh int32, spitch, dpitch uint16, sw, dw int32) {
	C.filter_shrink_Y_SSE( (*C.Uint8)(spixels), (*C.Uint8)(dpixels), C.int(sh), C.int(spitch), C.int(dpitch), C.int(sw), C.int(dw))
}

func filter_expand_X_SSE(spixels, dpixels *byte, sh int32, spitch, dpitch uint16, sw, dw int32) {
	C.filter_expand_X_SSE( (*C.Uint8)(spixels), (*C.Uint8)(dpixels), C.int(sh), C.int(spitch), C.int(dpitch), C.int(sw), C.int(dw))
}

func filter_expand_Y_SSE(spixels, dpixels *byte, sh int32, spitch, dpitch uint16, sw, dw int32) {
	C.filter_expand_Y_SSE( (*C.Uint8)(spixels), (*C.Uint8)(dpixels), C.int(sh), C.int(spitch), C.int(dpitch), C.int(sw), C.int(dw))
}

func filter_shrink_X_ONLYC(spixels, dpixels *byte, sh int32, spitch, dpitch uint16, sw, dw int32) {
// 	C.filter_shrink_X_ONLYC( (*C.Uint8)(spixels), (*C.Uint8)(dpixels), C.int(sh), C.int(spitch), C.int(dpitch), C.int(sw), C.int(dw))
}

func filter_shrink_Y_ONLYC(spixels, dpixels *byte, sh int32, spitch, dpitch uint16, sw, dw int32) {
// 	C.filter_shrink_Y_ONLYC( (*C.Uint8)(spixels), (*C.Uint8)(dpixels), C.int(sh), C.int(spitch), C.int(dpitch), C.int(sw), C.int(dw))
}

func filter_expand_X_ONLYC(spixels, dpixels *byte, sh int32, spitch, dpitch uint16, sw, dw int32) {
// 	C.filter_expand_X_ONLYC( (*C.Uint8)(spixels), (*C.Uint8)(dpixels), C.int(sh), C.int(spitch), C.int(dpitch), C.int(sw), C.int(dw))
}

func filter_expand_Y_ONLYC(spixels, dpixels *byte, sh int32, spitch, dpitch uint16, sw, dw int32) {
// 	C.filter_expand_Y_ONLYC( (*C.Uint8)(spixels), (*C.Uint8)(dpixels), C.int(sh), C.int(spitch), C.int(dpitch), C.int(sw), C.int(dw))
}

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
