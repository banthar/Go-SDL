
package ttf

// #include <SDL/SDL_ttf.h>
import "C";
import "sdl";
import "unsafe";


type Font struct {
    cfont *C.TTF_Font;
}

func Init() int {
    return int(C.TTF_Init());
}

func Quit() {
    C.TTF_Quit();
}

func OpenFont(file string, ptsize int) *Font {
    cfile := C.CString(file);
    cfont := C.TTF_OpenFont(cfile, C.int(ptsize));
    C.free(unsafe.Pointer(cfile));
    font := new(Font);
    font.cfont = cfont;
    return font;
}

func (f *Font) Close() {
    C.TTF_CloseFont(f.cfont);
}

func RenderText_Blended(font *Font, text string, color sdl.Color) *sdl.Surface {
    ctext := C.CString(text);
    ccol := C.SDL_Color{C.Uint8(color.R),C.Uint8(color.G),C.Uint8(color.B), C.Uint8(color.Unused)};
    surface := C.TTF_RenderText_Blended(font.cfont, ctext, ccol);
    C.free(unsafe.Pointer(ctext));
    return (*sdl.Surface)(unsafe.Pointer(surface));
}
