/*
A binding of SDL and SDL_image.

The binding is works in pretty much the same way as it does in C, although
some of the functions have been altered to give them an object-oriented
flavor (eg. Rather than sdl.Flip(surface) it's surface.Flip() )
*/
package sdl


// struct private_hwdata{};
// struct SDL_BlitMap{};
// #define map _map
//
// #include <SDL/SDL.h>
// #include <SDL/SDL_image.h>
// static void SetError(const char* description){SDL_SetError("%s",description);}
// static int __SDL_RWseek(SDL_RWops *ctx, int offset, int whence) { return SDL_RWseek(ctx, offset, whence); }
// static int __SDL_RWtell(SDL_RWops *ctx) { return SDL_RWseek(ctx, 0, RW_SEEK_CUR); }
// static int __SDL_RWread(SDL_RWops *ctx, void *ptr, int size, int n) { return SDL_RWread(ctx, ptr, size, n); }
// static int __SDL_RWwrite(SDL_RWops *ctx, const void *ptr, int size, int n) { return SDL_RWwrite(ctx, ptr, size, n); }
// static int __SDL_RWclose(SDL_RWops *ctx) { return SDL_RWclose(ctx); }
// static int __SDL_SaveBMP(SDL_Surface *surface, const char *file) { return SDL_SaveBMP(surface, file); }
import "C"
import "unsafe"
import "os"

type cast unsafe.Pointer

// General

// Initializes SDL.
func Init(flags uint32) int { return int(C.SDL_Init(C.Uint32(flags))) }

// Shuts down SDL
func Quit() { C.SDL_Quit() }

// Initializes subsystems.
func InitSubSystem(flags uint32) int { return int(C.SDL_InitSubSystem(C.Uint32(flags))) }

// Shuts down a subsystem.
func QuitSubSystem(flags uint32) { C.SDL_QuitSubSystem(C.Uint32(flags)) }

// Checks which subsystems are initialized.
func WasInit(flags uint32) int { return int(C.SDL_WasInit(C.Uint32(flags))) }

// Error Handling

// Gets SDL error string
func GetError() string { return C.GoString(C.SDL_GetError()) }

// Set a string describing an error to be submitted to the SDL Error system.
func SetError(description string) {
	cdescription := C.CString(description)
	C.SetError(cdescription)
	C.free(unsafe.Pointer(cdescription))
}

// TODO SDL_Error

// Clear the current SDL error
func ClearError() { C.SDL_ClearError() }

// Video

// Sets up a video mode with the specified width, height, bits-per-pixel and
// returns a corresponding surface.  You don't need to call the Free method
// of the returned surface, as it will be done automatically by sdl.Quit.
func SetVideoMode(w int, h int, bpp int, flags uint32) *Surface {
	var screen = C.SDL_SetVideoMode(C.int(w), C.int(h), C.int(bpp), C.Uint32(flags))
	return (*Surface)(cast(screen))
}

// Returns a pointer to the current display surface.
func GetVideoSurface() *Surface { return (*Surface)(cast(C.SDL_GetVideoSurface())) }

// Checks to see if a particular video mode is supported.  Returns 0 if not
// supported, or the bits-per-pixel of the closest available mode.
func VideoModeOK(width int, height int, bpp int, flags uint32) int {
	return int(C.SDL_VideoModeOK(C.int(width), C.int(height), C.int(bpp), C.Uint32(flags)))
}

func ListModes(format *PixelFormat, flags uint32) []Rect {
	modes := C.SDL_ListModes((*C.SDL_PixelFormat)(cast(format)), C.Uint32(flags))
	if modes == nil { //no modes available
		return make([]Rect, 0)
	}
	var any int
	*((***C.SDL_Rect)(unsafe.Pointer(&any))) = modes
	if any == -1 { //any dimension is ok
		return nil
	}

	var count int
	ptr := *modes //first element in the list
	for count = 0; ptr != nil; count++ {
		ptr = *(**C.SDL_Rect)(unsafe.Pointer(uintptr(unsafe.Pointer(modes)) + uintptr(count * int(unsafe.Sizeof(ptr)) )))
	}
	var ret = make([]Rect, count-1)

	*((***C.SDL_Rect)(unsafe.Pointer(&ret))) = modes // TODO
	return ret
}

type VideoInfo struct {
	HW_available bool         "Flag: Can you create hardware surfaces?"
	WM_available bool         "Flag: Can you talk to a window manager?"
	Blit_hw      bool         "Flag: Accelerated blits HW --> HW"
	Blit_hw_CC   bool         "Flag: Accelerated blits with Colorkey"
	Blit_hw_A    bool         "Flag: Accelerated blits with Alpha"
	Blit_sw      bool         "Flag: Accelerated blits SW --> HW"
	Blit_sw_CC   bool         "Flag: Accelerated blits with Colorkey"
	Blit_sw_A    bool         "Flag: Accelerated blits with Alpha"
	Blit_fill    bool         "Flag: Accelerated color fill"
	Video_mem    uint32       "The total amount of video memory (in K)"
	Vfmt         *PixelFormat "Value: The format of the video surface"
	Current_w    int32        "Value: The current video mode width"
	Current_h    int32        "Value: The current video mode height"
}

func GetVideoInfo() *VideoInfo {
	vinfo := (*internalVideoInfo)(cast(C.SDL_GetVideoInfo()))

	flags := vinfo.Flags

	return &VideoInfo{
		HW_available: flags&(1<<0) != 0,
		WM_available: flags&(1<<1) != 0,
		Blit_hw:      flags&(1<<9) != 0,
		Blit_hw_CC:   flags&(1<<10) != 0,
		Blit_hw_A:    flags&(1<<11) != 0,
		Blit_sw:      flags&(1<<12) != 0,
		Blit_sw_CC:   flags&(1<<13) != 0,
		Blit_sw_A:    flags&(1<<14) != 0,
		Blit_fill:    flags&(1<<15) != 0,
		Video_mem:    vinfo.Video_mem,
		Vfmt:         vinfo.Vfmt,
		Current_w:    vinfo.Current_w,
		Current_h:    vinfo.Current_h,
	}
}

// Makes sure the given area is updated on the given screen.  If x, y, w, and
// h are all 0, the whole screen will be updated.
func (screen *Surface) UpdateRect(x int32, y int32, w uint32, h uint32) {
	C.SDL_UpdateRect((*C.SDL_Surface)(cast(screen)), C.Sint32(x), C.Sint32(y), C.Uint32(w), C.Uint32(h))
}

func (screen *Surface) UpdateRects(rects []Rect) {
	if len(rects) > 0 {
		C.SDL_UpdateRects((*C.SDL_Surface)(cast(screen)), C.int(len(rects)), (*C.SDL_Rect)(cast(&rects[0])))
	}
}

// Gets the window title and icon name.
func WM_GetCaption() (title, icon string) {
	//SDL seems to free these strings.  TODO: Check to see if that's the case
	var ctitle, cicon *C.char
	C.SDL_WM_GetCaption(&ctitle, &cicon)
	title = C.GoString(ctitle)
	icon = C.GoString(cicon)
	return
}

// Sets the window title and icon name.
func WM_SetCaption(title, icon string) {
	ctitle := C.CString(title)
	cicon := C.CString(icon)
	C.SDL_WM_SetCaption(ctitle, cicon)
	C.free(unsafe.Pointer(ctitle))
	C.free(unsafe.Pointer(cicon))
}

// Sets the icon for the display window.
func WM_SetIcon(icon *Surface, mask *uint8) {
	C.SDL_WM_SetIcon((*C.SDL_Surface)(cast(icon)), (*C.Uint8)(mask))
}

// Minimizes the window
func WM_IconifyWindow() int { return int(C.SDL_WM_IconifyWindow()) }

// Toggles fullscreen mode
func WM_ToggleFullScreen(surface *Surface) int {
	return int(C.SDL_WM_ToggleFullScreen((*C.SDL_Surface)(cast(surface))))
}

// Swaps OpenGL framebuffers/Update Display.
func GL_SwapBuffers() { C.SDL_GL_SwapBuffers() }

func GL_SetAttribute(attr int, value int) int {
	return int(C.SDL_GL_SetAttribute(C.SDL_GLattr(attr), C.int(value)))
}

// Swaps screen buffers.
func (screen *Surface) Flip() int { return int(C.SDL_Flip((*C.SDL_Surface)(cast(screen)))) }

// Frees (deletes) a Surface
func (screen *Surface) Free() { C.SDL_FreeSurface((*C.SDL_Surface)(cast(screen))) }

// Locks a surface for direct access.
func (screen *Surface) Lock() int {
	return int(C.SDL_LockSurface((*C.SDL_Surface)(cast(screen))))
}

// Unlocks a previously locked surface.
func (screen *Surface) Unlock() int { return int(C.SDL_Flip((*C.SDL_Surface)(cast(screen)))) }

func (dst *Surface) Blit(dstrect *Rect, src *Surface, srcrect *Rect) int {
	var ret = C.SDL_UpperBlit(
		(*C.SDL_Surface)(cast(src)),
		(*C.SDL_Rect)(cast(srcrect)),
		(*C.SDL_Surface)(cast(dst)),
		(*C.SDL_Rect)(cast(dstrect)))

	return int(ret)
}

// This function performs a fast fill of the given rectangle with some color.
func (dst *Surface) FillRect(dstrect *Rect, color uint32) int {
	var ret = C.SDL_FillRect(
		(*C.SDL_Surface)(cast(dst)),
		(*C.SDL_Rect)(cast(dstrect)),
		C.Uint32(color))

	return int(ret)
}

// Adjusts the alpha properties of a Surface.
func (s *Surface) SetAlpha(flags uint32, alpha uint8) int {
	return int(C.SDL_SetAlpha((*C.SDL_Surface)(cast(s)), C.Uint32(flags), C.Uint8(alpha)))
}

// Sets the color key (transparent pixel)  in  a  blittable  surface  and
// enables or disables RLE blit acceleration.
func (s *Surface) SetColorKey(flags uint32, ColorKey uint32) int {
	return int(C.SDL_SetColorKey((*C.SDL_Surface)(cast(s)),
		C.Uint32(flags), C.Uint32(ColorKey)))
}

// Gets the clipping rectangle for a surface.
func (s *Surface) GetClipRect(r *Rect) {
	C.SDL_GetClipRect((*C.SDL_Surface)(cast(s)), (*C.SDL_Rect)(cast(r)))
	return
}

// Sets the clipping rectangle for a surface.
func (s *Surface) SetClipRect(r *Rect) {
	C.SDL_SetClipRect((*C.SDL_Surface)(cast(s)), (*C.SDL_Rect)(cast(r)))
	return
}

// Map a RGB color value to a pixel format.
func MapRGB(format *PixelFormat, r, g, b uint8) uint32 {
	return (uint32)(C.SDL_MapRGB((*C.SDL_PixelFormat)(cast(format)), (C.Uint8)(r), (C.Uint8)(g), (C.Uint8)(b)))
}

// Gets RGB values from a pixel in the specified pixel format.
func GetRGB(color uint32, format *PixelFormat, r, g, b *uint8) {
	C.SDL_GetRGB(C.Uint32(color), (*C.SDL_PixelFormat)(cast(format)), (*C.Uint8)(r), (*C.Uint8)(g), (*C.Uint8)(b))
}

// Map a RGBA color value to a pixel format.
func MapRGBA(format *PixelFormat, r, g, b, a uint8) uint32 {
	return (uint32)(C.SDL_MapRGBA((*C.SDL_PixelFormat)(cast(format)), (C.Uint8)(r), (C.Uint8)(g), (C.Uint8)(b), (C.Uint8)(a)))
}

// Gets RGBA values from a pixel in the specified pixel format.
func GetRGBA(color uint32, format *PixelFormat, r, g, b, a *uint8) {
	C.SDL_GetRGBA(C.Uint32(color), (*C.SDL_PixelFormat)(cast(format)), (*C.Uint8)(r), (*C.Uint8)(g), (*C.Uint8)(b), (*C.Uint8)(a))
}

// Loads Surface from file (using IMG_Load).
func Load(file string) *Surface {
	cfile := C.CString(file)
	var screen = C.IMG_Load(cfile)
	C.free(unsafe.Pointer(cfile))
	return (*Surface)(cast(screen))
}

func (src *Surface) SaveBMP(file string) int {
	cfile := C.CString(file)
	res := C.__SDL_SaveBMP( (*C.SDL_Surface)(cast(src)), cfile)
	C.free(unsafe.Pointer(cfile))
	return int(res)
}

// Creates an empty Surface.
func CreateRGBSurface(flags uint32, width int, height int, bpp int, Rmask uint32, Gmask uint32, Bmask uint32, Amask uint32) *Surface {
	p := C.SDL_CreateRGBSurface(C.Uint32(flags), C.int(width), C.int(height), C.int(bpp),
		C.Uint32(Rmask), C.Uint32(Gmask), C.Uint32(Bmask), C.Uint32(Amask))
	return (*Surface)(cast(p))
}

// Converts a surface to the display format
func DisplayFormat(src *Surface) *Surface {
	p := C.SDL_DisplayFormat( (*C.SDL_Surface)(cast(src)) )
	return (*Surface)(cast(p))
}

// Converts a surface to the display format with alpha
func DisplayFormatAlpha(src *Surface) *Surface {
	p := C.SDL_DisplayFormatAlpha((*C.SDL_Surface)(cast(src)))
	return (*Surface)(cast(p))
}

// Events

// Enables UNICODE translation.
func EnableUNICODE(enable int) int { return int(C.SDL_EnableUNICODE(C.int(enable))) }

// Sets keyboard repeat rate.
func EnableKeyRepeat(delay, interval int) int {
	return int(C.SDL_EnableKeyRepeat(C.int(delay), C.int(interval)))
}

// Gets keyboard repeat rate.
func GetKeyRepeat() (int, int) {

	var delay int
	var interval int

	C.SDL_GetKeyRepeat((*C.int)(cast(&delay)), (*C.int)(cast(&interval)))

	return delay, interval
}

// Gets a snapshot of the current keyboard state
func GetKeyState() []uint8 {
	var numkeys C.int
	array := C.SDL_GetKeyState(&numkeys)

	var ptr = make([]uint8, numkeys)

	*((**C.Uint8)(unsafe.Pointer(&ptr))) = array // TODO

	return ptr

}

// Modifier
type Mod C.int

// Key
type Key C.int

// Retrieves the current state of the mouse.
func GetMouseState(x, y *int) uint8 {
	return uint8(C.SDL_GetMouseState((*C.int)(cast(x)), (*C.int)(cast(y))))
}

// Retrieves the current state of the mouse relative to the last time this
// function was called.
func GetRelativeMouseState(x, y *int) uint8 {
	return uint8(C.SDL_GetRelativeMouseState((*C.int)(cast(x)), (*C.int)(cast(y))))
}

// Gets the state of modifier keys
func GetModState() Mod { return Mod(C.SDL_GetModState()) }

// Sets the state of modifier keys
func SetModState(modstate Mod) { C.SDL_SetModState(C.SDLMod(modstate)) }

// Gets the name of an SDL virtual keysym
func GetKeyName(key Key) string { return C.GoString(C.SDL_GetKeyName(C.SDLKey(key))) }

// Events

// Waits indefinitely for the next available event
func (event *Event) Wait() bool {
	var ret = C.SDL_WaitEvent((*C.SDL_Event)(cast(event)))
	return ret != 0
}

// Push the event onto the event queue
func (event *Event) Push() bool {
	var ret = C.SDL_PushEvent((*C.SDL_Event)(cast(event)))
	return ret != 0
}

// Polls for currently pending events
func (event *Event) Poll() bool {
	var ret = C.SDL_PollEvent((*C.SDL_Event)(cast(event)))
	return ret != 0
}

// Returns KeyboardEvent or nil if event has other type
func (event *Event) Keyboard() *KeyboardEvent {
	if event.Type == KEYUP || event.Type == KEYDOWN {
		return (*KeyboardEvent)(cast(event))
	}

	return nil
}

// Returns MouseButtonEvent or nil if event has other type
func (event *Event) MouseButton() *MouseButtonEvent {
	if event.Type == MOUSEBUTTONDOWN || event.Type == MOUSEBUTTONUP {
		return (*MouseButtonEvent)(cast(event))
	}

	return nil
}

// Returns MouseMotion or nil if event has other type
func (event *Event) MouseMotion() *MouseMotionEvent {
	if event.Type == MOUSEMOTION {
		return (*MouseMotionEvent)(cast(event))
	}

	return nil
}

// Returns ActiveEvent or nil if event has other type
func (event *Event) Active() *ActiveEvent {
	if event.Type == ACTIVEEVENT {
		return (*ActiveEvent)(cast(event))
	}

	return nil
}

// Returns ResizeEvent or nil if event has other type
func (event *Event) Resize() *ResizeEvent {
	if event.Type == VIDEORESIZE {
		return (*ResizeEvent)(cast(event))
	}

	return nil
}

// Time

// Gets the number of milliseconds since the SDL library initialization.
func GetTicks() uint32 { return uint32(C.SDL_GetTicks()) }

// Waits a specified number of milliseconds before returning.
func Delay(ms uint32) { C.SDL_Delay(C.Uint32(ms)) }

func ShowCursor(toggle int) int {
	return int(C.SDL_ShowCursor(C.int(toggle)))
}

type RWops struct {
	ptr *C.SDL_RWops
}

func (rwops *RWops) Free() {
	C.SDL_FreeRW(rwops.ptr)
}

func RWFromFile(file, mode string) *RWops {
	ptr := C.SDL_RWFromFile(C.CString(file), C.CString(mode))
	return &RWops{ptr: ptr}
}

// FIXME(salviati): It'd be nice if we could get rid of this
// extra mode parameter.
func RWFromFP(f *os.File, autoclose bool, mode string) *RWops {
	fd := C.int(f.Fd())
	fp := C.fdopen(fd, C.CString(mode))
	autocloseInt := 0
	if autoclose { autocloseInt = 1 }
	ptr := C.SDL_RWFromFP(fp, C.int(autocloseInt))
	return &RWops{ptr: ptr}
}

func RWFromMem(mem []byte) *RWops {
	ptr := C.SDL_RWFromMem( unsafe.Pointer(&mem[0]), C.int(len(mem)) )
	return &RWops{ptr: ptr}
}

func RWFromConstMem(mem []byte) *RWops {
	ptr := C.SDL_RWFromConstMem( unsafe.Pointer(&mem[0]), C.int(len(mem)) )
	return &RWops{ptr: ptr}
}

// Do we need this one?
func AllocRW() *RWops {
	ptr := C.SDL_AllocRW()
	return &RWops{ptr: ptr}
}

const (
	RW_SEEK_SET	= 0	/* Seek from the beginning of data */
	RW_SEEK_CUR	= 1	/* Seek relative to current read point */
	RW_SEEK_END = 2	/* Seek relative to the end of data */
)

func (rwops *RWops) Seek(offset, whence int) int {
	return int( C.__SDL_RWseek(rwops.ptr, C.int(offset), C.int(whence)) )
}

func (rwops *RWops) Tell() int {
	return int(C.__SDL_RWtell(rwops.ptr))
}

func (rwops *RWops) Read(mem []byte, size int) int {
	return int( C.__SDL_RWread(rwops.ptr, unsafe.Pointer(&mem[0]), C.int(size), C.int(len(mem))) )
}

func (rwops *RWops) Write(mem []byte, size int) int {
	return int( C.__SDL_RWwrite(rwops.ptr, unsafe.Pointer(&mem[0]), C.int(size), C.int(len(mem))) )
}

func (rwops *RWops) Close() int {
	return int(C.__SDL_RWclose(rwops.ptr))
}

func (rwops *RWops) Load(freesrc bool) *Surface {
	freesrcInt := 0
	if freesrc { freesrcInt = 1 }
	im := C.IMG_Load_RW(rwops.ptr, C.int(freesrcInt))
	return (*Surface)(cast(im))
}

func Putenv(s string) int {
	cstr := C.CString(s)
	ret := C.SDL_putenv(cstr)
	C.free(unsafe.Pointer(cstr))
	return int(ret)
}

func (rwops *RWops) ReadLE16() uint16 { return uint16( C.SDL_ReadLE16(rwops.ptr) ) }
func (rwops *RWops) ReadBE16() uint16 { return uint16( C.SDL_ReadBE16(rwops.ptr) ) }
func (rwops *RWops) ReadLE32() uint32 { return uint32( C.SDL_ReadLE32(rwops.ptr) ) }
func (rwops *RWops) ReadBE32() uint32 { return uint32( C.SDL_ReadBE32(rwops.ptr) ) }
func (rwops *RWops) ReadLE64() uint64 { return uint64( C.SDL_ReadLE64(rwops.ptr) ) }
func (rwops *RWops) ReadBE64() uint64 { return uint64( C.SDL_ReadBE64(rwops.ptr) ) }

func (rwops *RWops) WriteLE16(value uint16) int { return int( C.SDL_WriteLE16(rwops.ptr, C.Uint16(value)) ) }
func (rwops *RWops) WriteBE16(value uint16) int { return int( C.SDL_WriteBE16(rwops.ptr, C.Uint16(value)) ) }
func (rwops *RWops) WriteLE32(value uint32) int { return int( C.SDL_WriteLE32(rwops.ptr, C.Uint32(value)) ) }
func (rwops *RWops) WriteBE32(value uint32) int { return int( C.SDL_WriteBE32(rwops.ptr, C.Uint32(value)) ) }
func (rwops *RWops) WriteLE64(value uint64) int { return int( C.SDL_WriteLE64(rwops.ptr, C.Uint64(value)) ) }
func (rwops *RWops) WriteBE64(value uint64) int { return int( C.SDL_WriteBE64(rwops.ptr, C.Uint64(value)) ) }

func HasRDTSC() bool { return C.SDL_HasRDTSC() != 0 }
func HasMMX() bool { return C.SDL_HasMMX() != 0 }
func HasMMXExt() bool { return C.SDL_HasMMXExt() != 0 }
func Has3DNow() bool { return C.SDL_Has3DNow() != 0 }
func Has3DNowExt() bool { return C.SDL_Has3DNowExt() != 0 }
func HasSSE() bool { return C.SDL_HasSSE() != 0 }
func HasSSE2() bool { return C.SDL_HasSSE2() != 0 }
func HasAltiVec() bool { return C.SDL_HasAltiVec() != 0 }
