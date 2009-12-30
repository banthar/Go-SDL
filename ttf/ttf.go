/*
A binding of SDL_ttf.

You use this binding pretty much the same way you use SDL_ttf, although commands
that work with loaded fonts are changed to have a more object-oriented feel.
(eg. Rather than ttf.GetFontStyle(f) it's f.GetFontStyle() )
*/
package ttf

// #include <SDL/SDL_ttf.h>
import "C"
import "sdl"
import "unsafe"

// A ttf or otf font.
type Font struct {
	cfont *C.TTF_Font
}

// Initializes SDL_ttf.
func Init() int { return int(C.TTF_Init()) }

// Checks to see if SDL_ttf is initialized.  Returns 1 if true, 0 if false.
func WasInit() int { return int(C.TTF_WasInit()) }

// Shuts down SDL_ttf.
func Quit() { C.TTF_Quit() }

// Loads a font from a file at the specified point size.
func OpenFont(file string, ptsize int) *Font {
	cfile := C.CString(file)
	cfont := C.TTF_OpenFont(cfile, C.int(ptsize))
	C.free(unsafe.Pointer(cfile))

	if cfont == nil {
		return nil
	}

	return &Font{cfont}
}

// Loads a font from a file containing multiple font faces at the specified
// point size.
func OpenFontIndex(file string, ptsize, index int) *Font {
	cfile := C.CString(file)
	cfont := C.TTF_OpenFontIndex(cfile, C.int(ptsize), C.long(index))
	C.free(unsafe.Pointer(cfile))

	if cfont == nil {
		return nil
	}

	return &Font{cfont}
}

// Frees the pointer to the font.
func (f *Font) Close() { C.TTF_CloseFont(f.cfont) }

// Renders text in the specified color and returns an SDL surface.  Solid
// rendering is quick, although not as smooth as the other rendering types.
func RenderText_Solid(font *Font, text string, color sdl.Color) *sdl.Surface {
	ctext := C.CString(text)
	ccol := C.SDL_Color{C.Uint8(color.R), C.Uint8(color.G), C.Uint8(color.B), C.Uint8(color.Unused)}
	surface := C.TTF_RenderText_Solid(font.cfont, ctext, ccol)
	C.free(unsafe.Pointer(ctext))
	return (*sdl.Surface)(unsafe.Pointer(surface))
}

// Renders text in the specified color (and with the specified background color)
// and returns an SDL surface.  Shaded rendering is slower than solid
// rendering and the text is in a solid box, but it's better looking.
func RenderText_Shaded(font *Font, text string, color, bgcolor sdl.Color) *sdl.Surface {
	ctext := C.CString(text)
	ccol := C.SDL_Color{C.Uint8(color.R), C.Uint8(color.G), C.Uint8(color.B), C.Uint8(color.Unused)}
	cbgcol := C.SDL_Color{C.Uint8(bgcolor.R), C.Uint8(bgcolor.G), C.Uint8(bgcolor.B), C.Uint8(bgcolor.Unused)}
	surface := C.TTF_RenderText_Shaded(font.cfont, ctext, ccol, cbgcol)
	C.free(unsafe.Pointer(ctext))
	return (*sdl.Surface)(unsafe.Pointer(surface))
}

// Renders text in the specified color and returns an SDL surface.  Blended
// rendering is the slowest of the three methods, although it produces the best
// results, especially when blitted over another image.
func RenderText_Blended(font *Font, text string, color sdl.Color) *sdl.Surface {
	ctext := C.CString(text)
	ccol := C.SDL_Color{C.Uint8(color.R), C.Uint8(color.G), C.Uint8(color.B), C.Uint8(color.Unused)}
	surface := C.TTF_RenderText_Blended(font.cfont, ctext, ccol)
	C.free(unsafe.Pointer(ctext))
	return (*sdl.Surface)(unsafe.Pointer(surface))
}

// Returns the rendering style of the font.
func (f *Font) GetStyle() int { return int(C.TTF_GetFontStyle(f.cfont)) }

// Sets the rendering style of the font.
func (f *Font) SetStyle(style int) { C.TTF_SetFontStyle(f.cfont, C.int(style)) }

// Returns the maximum height of all the glyphs of the font.
func (f *Font) Height() int { return int(C.TTF_FontHeight(f.cfont)) }

// Returns the maximum pixel ascent (from the baseline) of all the glyphs
// of the font.
func (f *Font) Ascent() int { return int(C.TTF_FontAscent(f.cfont)) }

// Returns the maximum pixel descent (from the baseline) of all the glyphs
// of the font.
func (f *Font) Descent() int { return int(C.TTF_FontDescent(f.cfont)) }

// Returns the recommended pixel height of a rendered line of text.
func (f *Font) LineSkip() int { return int(C.TTF_FontLineSkip(f.cfont)) }
