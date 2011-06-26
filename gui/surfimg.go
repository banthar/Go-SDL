package sdlgui

import(
	"fmt"
	"sdl"
	"image"
	"unsafe"
	"reflect"
	"image/draw"
)

type surfimg struct {
	*sdl.Surface
}

func (img *surfimg)pixPtr(x, y int) (reflect.Value) {
	imglen := int(img.W * img.H)
	h := reflect.SliceHeader{
		uintptr(unsafe.Pointer(img.Pixels)),
		imglen,
		imglen,
	}

	switch img.Format.BitsPerPixel {
		case 8:
			pix := (*(*[]uint8)(unsafe.Pointer(&h)))[(y*int(img.W))+x]
			return reflect.ValueOf(&pix).Elem()
		case 16:
			pix := (*(*[]uint16)(unsafe.Pointer(&h)))[(y*int(img.W))+x]
			return reflect.ValueOf(&pix).Elem()
		case 32:
			pix := (*(*[]uint32)(unsafe.Pointer(&h)))[(y*int(img.W))+x]
			return reflect.ValueOf(&pix).Elem()
		case 64:
			pix := (*(*[]uint64)(unsafe.Pointer(&h)))[(y*int(img.W))+x]
			return reflect.ValueOf(&pix).Elem()
	}

	panic(fmt.Errorf("Image has unexpected BPP: %v", img.Format.BitsPerPixel))
}

func (img *surfimg)ColorModel() (image.ColorModel) {
	return image.RGBAColorModel
}

func (img *surfimg)Bounds() (image.Rectangle) {
	return image.Rect(0, 0, int(img.W), int(img.H))
}

func (img *surfimg)At(x, y int) (image.Color) {
	var r, g, b, a uint8
	sdl.GetRGBA(uint32(img.pixPtr(x, y).Uint()), img.Format, &r, &g, &b, &a)

	return img.ColorModel().Convert(image.RGBAColor{r, g, b, a})
}

func (img *surfimg)Set(x, y int, c image.Color) {
	img.Lock()
	defer img.Unlock()

	r, g, b, a := c.RGBA()

	pix := img.pixPtr(x, y)
	pix.SetUint(uint64(sdl.MapRGBA(img.Format, uint8(r), uint8(g), uint8(b), uint8(a))))
}

func SurfaceToImage(s *sdl.Surface) (draw.Image) {
	return &surfimg{s}
}
