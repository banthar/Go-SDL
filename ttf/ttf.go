/*
A binding of SDL_ttf.

You use this binding pretty much the same way you use SDL_ttf, although commands
that work with loaded fonts are changed to have a more object-oriented feel.
(eg. Rather than ttf.GetFontStyle(f) it's f.GetFontStyle() )
*/
package ttf

// #cgo pkg-config: sdl
// #cgo LDFLAGS: -lSDL_ttf
// #include <SDL_ttf.h>
import "C"

import (
	"github.com/0xe2-0x9a-0x9b/Go-SDL/sdl"
	"sync"
	"unsafe"
)

// The version of Go-SDL TTF bindings.
// The version descriptor changes into a new unique string
// after a semantically incompatible Go-SDL update.
//
// The returned value can be checked by users of this package
// to make sure they are using a version with the expected semantics.
//
// If Go adds some kind of support for package versioning, this function will go away.
func GoSdlVersion() string {
	return "⚛SDL TTF bindings 1.0"
}

func wrap(cSurface *C.SDL_Surface) *sdl.Surface {
	var s *sdl.Surface

	if cSurface != nil {
		var surface sdl.Surface
		surface.SetCSurface(unsafe.Pointer(cSurface))
		s = &surface
	} else {
		s = nil
	}

	return s
}

// A ttf or otf font.
type Font struct {
	cfont *C.TTF_Font
	mutex sync.RWMutex
}

// Initializes SDL_ttf.
func Init() int {
	sdl.GlobalMutex.Lock()
	status := int(C.TTF_Init())
	sdl.GlobalMutex.Unlock()
	return status
}

// Checks to see if SDL_ttf is initialized.  Returns 1 if true, 0 if false.
func WasInit() int {
	sdl.GlobalMutex.Lock()
	status := int(C.TTF_WasInit())
	sdl.GlobalMutex.Unlock()
	return status
}

// Shuts down SDL_ttf.
func Quit() {
	sdl.GlobalMutex.Lock()
	C.TTF_Quit()
	sdl.GlobalMutex.Unlock()
}

// Loads a font from a file at the specified point size.
func OpenFont(file string, ptsize int) *Font {
	sdl.GlobalMutex.Lock()

	cfile := C.CString(file)
	cfont := C.TTF_OpenFont(cfile, C.int(ptsize))
	C.free(unsafe.Pointer(cfile))

	sdl.GlobalMutex.Unlock()

	if cfont == nil {
		return nil
	}

	return &Font{cfont: cfont}
}

// Loads a font from a file containing multiple font faces at the specified
// point size.
func OpenFontIndex(file string, ptsize, index int) *Font {
	sdl.GlobalMutex.Lock()

	cfile := C.CString(file)
	cfont := C.TTF_OpenFontIndex(cfile, C.int(ptsize), C.long(index))
	C.free(unsafe.Pointer(cfile))

	sdl.GlobalMutex.Unlock()

	if cfont == nil {
		return nil
	}

	return &Font{cfont: cfont}
}

// Frees the pointer to the font.
func (f *Font) Close() {
	sdl.GlobalMutex.Lock()
	f.mutex.Lock()

	C.TTF_CloseFont(f.cfont)

	f.mutex.Unlock()
	sdl.GlobalMutex.Unlock()
}

// Renders Latin-1 text in the specified color and returns an SDL surface.  Solid
// rendering is quick, although not as smooth as the other rendering types.
func RenderText_Solid(font *Font, text string, color sdl.Color) *sdl.Surface {
	sdl.GlobalMutex.Lock() // Because 'C.TTF_Render*' uses 'C.SDL_CreateRGBSurface'
	font.mutex.Lock()      // Use a write lock, because 'C.TTF_Render*' may update font's internal caches

	ctext := C.CString(text)
	ccol := C.SDL_Color{C.Uint8(color.R), C.Uint8(color.G), C.Uint8(color.B), C.Uint8(color.Unused)}
	surface := C.TTF_RenderText_Solid(font.cfont, ctext, ccol)
	C.free(unsafe.Pointer(ctext))

	font.mutex.Unlock()
	sdl.GlobalMutex.Unlock()

	return wrap(surface)
}

// Renders UTF-8 text in the specified color and returns an SDL surface.  Solid
// rendering is quick, although not as smooth as the other rendering types.
func RenderUTF8_Solid(font *Font, text string, color sdl.Color) *sdl.Surface {
	sdl.GlobalMutex.Lock() // Because 'C.TTF_Render*' uses 'C.SDL_CreateRGBSurface'
	font.mutex.Lock()      // Use a write lock, because 'C.TTF_Render*' may update font's internal caches

	ctext := C.CString(text)
	ccol := C.SDL_Color{C.Uint8(color.R), C.Uint8(color.G), C.Uint8(color.B), C.Uint8(color.Unused)}
	surface := C.TTF_RenderUTF8_Solid(font.cfont, ctext, ccol)
	C.free(unsafe.Pointer(ctext))

	font.mutex.Unlock()
	sdl.GlobalMutex.Unlock()

	return wrap(surface)
}

// Renders Latin-1 text in the specified color (and with the specified background
// color) and returns an SDL surface.  Shaded rendering is slower than solid
// rendering and the text is in a solid box, but it's better looking.
func RenderText_Shaded(font *Font, text string, color, bgcolor sdl.Color) *sdl.Surface {
	sdl.GlobalMutex.Lock() // Because 'C.TTF_Render*' uses 'C.SDL_CreateRGBSurface'
	font.mutex.Lock()      // Use a write lock, because 'C.TTF_Render*' may update font's internal caches

	ctext := C.CString(text)
	ccol := C.SDL_Color{C.Uint8(color.R), C.Uint8(color.G), C.Uint8(color.B), C.Uint8(color.Unused)}
	cbgcol := C.SDL_Color{C.Uint8(bgcolor.R), C.Uint8(bgcolor.G), C.Uint8(bgcolor.B), C.Uint8(bgcolor.Unused)}
	surface := C.TTF_RenderText_Shaded(font.cfont, ctext, ccol, cbgcol)
	C.free(unsafe.Pointer(ctext))

	font.mutex.Unlock()
	sdl.GlobalMutex.Unlock()

	return wrap(surface)
}

// Renders UTF-8 text in the specified color (and with the specified background
// color) and returns an SDL surface.  Shaded rendering is slower than solid
// rendering and the text is in a solid box, but it's better looking.
func RenderUTF8_Shaded(font *Font, text string, color, bgcolor sdl.Color) *sdl.Surface {
	sdl.GlobalMutex.Lock() // Because 'C.TTF_Render*' uses 'C.SDL_CreateRGBSurface'
	font.mutex.Lock()      // Use a write lock, because 'C.TTF_Render*' may update font's internal caches

	ctext := C.CString(text)
	ccol := C.SDL_Color{C.Uint8(color.R), C.Uint8(color.G), C.Uint8(color.B), C.Uint8(color.Unused)}
	cbgcol := C.SDL_Color{C.Uint8(bgcolor.R), C.Uint8(bgcolor.G), C.Uint8(bgcolor.B), C.Uint8(bgcolor.Unused)}
	surface := C.TTF_RenderUTF8_Shaded(font.cfont, ctext, ccol, cbgcol)
	C.free(unsafe.Pointer(ctext))

	font.mutex.Unlock()
	sdl.GlobalMutex.Unlock()

	return wrap(surface)
}

// Renders Latin-1 text in the specified color and returns an SDL surface.
// Blended rendering is the slowest of the three methods, although it produces
// the best results, especially when blitted over another image.
func RenderText_Blended(font *Font, text string, color sdl.Color) *sdl.Surface {
	sdl.GlobalMutex.Lock() // Because 'C.TTF_Render*' uses 'C.SDL_CreateRGBSurface'
	font.mutex.Lock()      // Use a write lock, because 'C.TTF_Render*' may update font's internal caches

	ctext := C.CString(text)
	ccol := C.SDL_Color{C.Uint8(color.R), C.Uint8(color.G), C.Uint8(color.B), C.Uint8(color.Unused)}
	surface := C.TTF_RenderText_Blended(font.cfont, ctext, ccol)
	C.free(unsafe.Pointer(ctext))

	font.mutex.Unlock()
	sdl.GlobalMutex.Unlock()

	return wrap(surface)
}

// Renders UTF-8 text in the specified color and returns an SDL surface.
// Blended rendering is the slowest of the three methods, although it produces
// the best results, especially when blitted over another image.
func RenderUTF8_Blended(font *Font, text string, color sdl.Color) *sdl.Surface {
	sdl.GlobalMutex.Lock() // Because 'C.TTF_Render*' uses 'C.SDL_CreateRGBSurface'
	font.mutex.Lock()      // Use a write lock, because 'C.TTF_Render*' may update font's internal caches

	ctext := C.CString(text)
	ccol := C.SDL_Color{C.Uint8(color.R), C.Uint8(color.G), C.Uint8(color.B), C.Uint8(color.Unused)}
	surface := C.TTF_RenderUTF8_Blended(font.cfont, ctext, ccol)
	C.free(unsafe.Pointer(ctext))

	font.mutex.Unlock()
	sdl.GlobalMutex.Unlock()

	return wrap(surface)
}

// Returns the rendering style of the font.
func (f *Font) GetStyle() int {
	f.mutex.RLock()
	result := int(C.TTF_GetFontStyle(f.cfont))
	f.mutex.RUnlock()
	return result
}

// Sets the rendering style of the font.
func (f *Font) SetStyle(style int) {
	sdl.GlobalMutex.Lock()
	f.mutex.Lock()

	C.TTF_SetFontStyle(f.cfont, C.int(style))

	f.mutex.Unlock()
	sdl.GlobalMutex.Unlock()
}

// Returns the maximum height of all the glyphs of the font.
func (f *Font) Height() int {
	f.mutex.RLock()
	result := int(C.TTF_FontHeight(f.cfont))
	f.mutex.RUnlock()
	return result
}

// Returns the maximum pixel ascent (from the baseline) of all the glyphs
// of the font.
func (f *Font) Ascent() int {
	f.mutex.RLock()
	result := int(C.TTF_FontAscent(f.cfont))
	f.mutex.RUnlock()
	return result
}

// Returns the maximum pixel descent (from the baseline) of all the glyphs
// of the font.
func (f *Font) Descent() int {
	f.mutex.RLock()
	result := int(C.TTF_FontDescent(f.cfont))
	f.mutex.RUnlock()
	return result
}

// Returns the recommended pixel height of a rendered line of text.
func (f *Font) LineSkip() int {
	f.mutex.RLock()
	result := int(C.TTF_FontLineSkip(f.cfont))
	f.mutex.RUnlock()
	return result
}

// Returns the number of available faces (sub-fonts) in the font.
func (f *Font) Faces() int {
	f.mutex.RLock()
	result := int(C.TTF_FontFaces(f.cfont))
	f.mutex.RUnlock()
	return result
}

// Returns >0 if the font's currently selected face is fixed width
// (i.e. monospace), 0 if not.
func (f *Font) IsFixedWidth() int {
	f.mutex.RLock()
	result := int(C.TTF_FontFaceIsFixedWidth(f.cfont))
	f.mutex.RUnlock()
	return result
}

// Returns the family name of the font's currently selected face,
// or a blank string if unavailable.
func (f *Font) FamilyName() string {
	var s string
	if f != nil {
		f.mutex.RLock()

		p := C.TTF_FontFaceFamilyName(f.cfont)
		if p != nil {
			s = C.GoString(p)
		} else {
			s = ""
		}

		f.mutex.RUnlock()
	} else {
		s = "nil"
	}

	return s
}

// Returns the style name of the font's currently selected face,
// or a blank string if unavailable.
func (f *Font) StyleName() string {
	var s string
	if f != nil {
		f.mutex.RLock()

		p := C.TTF_FontFaceStyleName(f.cfont)
		if p != nil {
			s = C.GoString(p)
		} else {
			s = ""
		}

		f.mutex.RUnlock()
	} else {
		s = "nil"
	}

	return s
}

// Returns the metrics (dimensions) of a glyph.
//
// Return values are:
//   minx, maxx, miny, maxy, advance, err
//
// The last return value (err) is 0 for success, -1 for any error (for example
// if the glyph is not available in this font).
//
// For more information on glyph metrics, visit
// http://freetype.sourceforge.net/freetype2/docs/tutorial/step2.html
func (f *Font) GlyphMetrics(ch uint16) (int, int, int, int, int, int) {
	sdl.GlobalMutex.Lock() // Because the underlying C code is fairly complex
	f.mutex.Lock()         // Use a write lock, because 'C.TTF_GlyphMetrics' may update font's internal caches

	minx := C.int(0)
	maxx := C.int(0)
	miny := C.int(0)
	maxy := C.int(0)
	advance := C.int(0)
	err := C.TTF_GlyphMetrics(f.cfont, C.Uint16(ch), &minx, &maxx, &miny, &maxy, &advance)

	sdl.GlobalMutex.Unlock()
	f.mutex.Unlock()

	return int(minx), int(maxx), int(miny), int(maxy), int(advance), int(err)
}

// Return the width and height of the rendered Latin-1 text.
//
// Return values are (width, height, err) where err is 0 for success, -1 on any error.
func (f *Font) SizeText(text string) (int, int, int) {
	sdl.GlobalMutex.Lock() // Because the underlying C code is fairly complex
	f.mutex.Lock()         // Use a write lock, because 'C.TTF_Size*' may update font's internal cache

	w := C.int(0)
	h := C.int(0)
	s := C.CString(text)
	err := C.TTF_SizeText(f.cfont, s, &w, &h)

	sdl.GlobalMutex.Unlock()
	f.mutex.Unlock()

	return int(w), int(h), int(err)
}

// Return the width and height of the rendered UTF-8 text.
//
// Return values are (width, height, err) where err is 0 for success, -1 on any error.
func (f *Font) SizeUTF8(text string) (int, int, int) {
	sdl.GlobalMutex.Lock() // Because the underlying C code is fairly complex
	f.mutex.Lock()         // Use a write lock, because 'C.TTF_Size*' may update font's internal caches

	w := C.int(0)
	h := C.int(0)
	s := C.CString(text)
	err := C.TTF_SizeUTF8(f.cfont, s, &w, &h)

	sdl.GlobalMutex.Unlock()
	f.mutex.Unlock()

	return int(w), int(h), int(err)
}
