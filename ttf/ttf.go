
package ttf

// #include <SDL/SDL_ttf.h>
import "C";
import "sdl";
import "unsafe";


func Init() int {
    return int(C.TTF_Init());
}

func Quit() {
    C.TTF_Quit();
}

func OpenFont(file string, ptsize int) *C.TTF_Font {
    cfile := C.CString(file);
    cfont := C.TTF_OpenFont(cfile, C.int(ptsize));
    C.free(unsafe.Pointer(cfile));
    return cfont;
}

func CloseFont(font *C.TTF_Font) {
    C.TTF_CloseFont(font);
}

func RenderText_Blended(cfont *C.TTF_Font, text string, color sdl.Color) *sdl.Surface {
    ctext := C.CString(text);
    ccol := C.SDL_Color{C.Uint8(color.R),C.Uint8(color.G),C.Uint8(color.B), C.Uint8(color.Unused)};
    surface := C.TTF_RenderText_Blended(cfont, ctext, ccol);
    C.free(unsafe.Pointer(ctext));
    return (*sdl.Surface)(unsafe.Pointer(surface));
}
