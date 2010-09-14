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
import "C"
import "unsafe"
import "sync"

type cast unsafe.Pointer

var globalMutex sync.Mutex

// TODO: This should NOT be a public function,
// but since we need it in the package "ttf" ... the Go language is failing here
func Wrap(intSurface *InternalSurface) *Surface {
	var s *Surface

	if intSurface != nil {
		var surface Surface
		surface.intSurface = intSurface
		surface.reload()
		s = &surface
	} else {
		s = nil
	}

	return s
}

// Pull data from the internal-surface. Make sure to use this when
// the internal-surface might have been changed.
func (s *Surface) reload() {
	s.Flags = s.intSurface.Flags
	s.Format = (*PixelFormat)(cast(s.intSurface.Format))
	s.W = s.intSurface.W
	s.H = s.intSurface.H
	s.Pitch = s.intSurface.Pitch
	s.Pixels = s.intSurface.Pixels
	s.Offset = s.intSurface.Offset
}

func (s *Surface) destroy() {
	s.intSurface = nil
	s.Format = nil
	s.Pixels = nil
}


// =======
// General
// =======

// The version of Go-SDL bindings.
// The version descriptor changes into a new unique string
// after a semantically incompatible Go-SDL update.
//
// The returned value can be checked by users of this package
// to make sure they are using a version with the expected semantics.
//
// If Go adds some kind of support for package versioning, this function will go away.
func GoSdlVersion() string {
	return "⚛SDL bindings 1.0"
}

// Initializes SDL.
func Init(flags uint32) int {
	globalMutex.Lock()
	status := int(C.SDL_Init(C.Uint32(flags)))
	globalMutex.Unlock()
	return status
}

// Shuts down SDL
func Quit() {
	globalMutex.Lock()

	if currentVideoSurface != nil {
		currentVideoSurface.destroy()
		currentVideoSurface = nil
	}

	C.SDL_Quit()

	globalMutex.Unlock()
}

// Initializes subsystems.
func InitSubSystem(flags uint32) int {
	globalMutex.Lock()
	status := int(C.SDL_InitSubSystem(C.Uint32(flags)))
	globalMutex.Unlock()
	return status
}

// Shuts down a subsystem.
func QuitSubSystem(flags uint32) {
	globalMutex.Lock()
	C.SDL_QuitSubSystem(C.Uint32(flags))
	globalMutex.Unlock()
}

// Checks which subsystems are initialized.
func WasInit(flags uint32) int {
	globalMutex.Lock()
	status := int(C.SDL_WasInit(C.Uint32(flags)))
	globalMutex.Unlock()
	return status
}


// ==============
// Error Handling
// ==============

// Gets SDL error string
func GetError() string {
	globalMutex.Lock()
	s := C.GoString(C.SDL_GetError())
	globalMutex.Unlock()
	return s
}

// Set a string describing an error to be submitted to the SDL Error system.
func SetError(description string) {
	globalMutex.Lock()

	cdescription := C.CString(description)
	C.SetError(cdescription)
	C.free(unsafe.Pointer(cdescription))

	globalMutex.Unlock()
}

// Clear the current SDL error
func ClearError() {
	globalMutex.Lock()
	C.SDL_ClearError()
	globalMutex.Unlock()
}


// ======
// Video
// ======

var currentVideoSurface *Surface = nil

// Sets up a video mode with the specified width, height, bits-per-pixel and
// returns a corresponding surface.  You don't need to call the Free method
// of the returned surface, as it will be done automatically by sdl.Quit.
func SetVideoMode(w int, h int, bpp int, flags uint32) *Surface {
	globalMutex.Lock()
	var screen = C.SDL_SetVideoMode(C.int(w), C.int(h), C.int(bpp), C.Uint32(flags))
	currentVideoSurface = Wrap((*InternalSurface)(cast(screen)))
	globalMutex.Unlock()
	return currentVideoSurface
}

// Returns a pointer to the current display surface.
func GetVideoSurface() *Surface {
	globalMutex.Lock()
	surface := currentVideoSurface
	globalMutex.Unlock()
	return surface
}

// Checks to see if a particular video mode is supported.  Returns 0 if not
// supported, or the bits-per-pixel of the closest available mode.
func VideoModeOK(width int, height int, bpp int, flags uint32) int {
	globalMutex.Lock()
	status := int(C.SDL_VideoModeOK(C.int(width), C.int(height), C.int(bpp), C.Uint32(flags)))
	globalMutex.Unlock()
	return status
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
		ptr = *(**C.SDL_Rect)(unsafe.Pointer(uintptr(unsafe.Pointer(modes)) + uintptr(count*unsafe.Sizeof(ptr))))
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
	globalMutex.Lock()
	vinfo := (*internalVideoInfo)(cast(C.SDL_GetVideoInfo()))
	globalMutex.Unlock()

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
	globalMutex.Lock()
	screen.mutex.Lock()

	C.SDL_UpdateRect((*C.SDL_Surface)(cast(screen.intSurface)), C.Sint32(x), C.Sint32(y), C.Uint32(w), C.Uint32(h))

	screen.mutex.Unlock()
	globalMutex.Unlock()
}

func (screen *Surface) UpdateRects(rects []Rect) {
	if len(rects) > 0 {
		globalMutex.Lock()
		screen.mutex.Lock()

		C.SDL_UpdateRects((*C.SDL_Surface)(cast(screen.intSurface)), C.int(len(rects)), (*C.SDL_Rect)(cast(&rects[0])))

		screen.mutex.Unlock()
		globalMutex.Unlock()
	}
}

// Gets the window title and icon name.
func WM_GetCaption() (title, icon string) {
	globalMutex.Lock()

	// SDL seems to free these strings.  TODO: Check to see if that's the case
	var ctitle, cicon *C.char
	C.SDL_WM_GetCaption(&ctitle, &cicon)
	title = C.GoString(ctitle)
	icon = C.GoString(cicon)

	globalMutex.Unlock()

	return
}

// Sets the window title and icon name.
func WM_SetCaption(title, icon string) {
	ctitle := C.CString(title)
	cicon := C.CString(icon)

	globalMutex.Lock()
	C.SDL_WM_SetCaption(ctitle, cicon)
	globalMutex.Unlock()

	C.free(unsafe.Pointer(ctitle))
	C.free(unsafe.Pointer(cicon))
}

// Sets the icon for the display window.
func WM_SetIcon(icon *Surface, mask *uint8) {
	globalMutex.Lock()
	C.SDL_WM_SetIcon((*C.SDL_Surface)(cast(icon.intSurface)), (*C.Uint8)(mask))
	globalMutex.Unlock()
}

// Minimizes the window
func WM_IconifyWindow() int {
	globalMutex.Lock()
	status := int(C.SDL_WM_IconifyWindow())
	globalMutex.Unlock()
	return status
}

// Toggles fullscreen mode
func WM_ToggleFullScreen(surface *Surface) int {
	globalMutex.Lock()
	status := int(C.SDL_WM_ToggleFullScreen((*C.SDL_Surface)(cast(surface.intSurface))))
	globalMutex.Unlock()
	return status
}

// Swaps OpenGL framebuffers/Update Display.
func GL_SwapBuffers() {
	globalMutex.Lock()
	C.SDL_GL_SwapBuffers()
	globalMutex.Unlock()
}

func GL_SetAttribute(attr int, value int) int {
	globalMutex.Lock()
	status := int(C.SDL_GL_SetAttribute(C.SDL_GLattr(attr), C.int(value)))
	globalMutex.Unlock()
	return status
}

// Swaps screen buffers.
func (screen *Surface) Flip() int {
	globalMutex.Lock()
	screen.mutex.Lock()

	status := int(C.SDL_Flip((*C.SDL_Surface)(cast(screen.intSurface))))

	screen.mutex.Unlock()
	globalMutex.Unlock()

	return status
}

// Frees (deletes) a Surface
func (screen *Surface) Free() {
	globalMutex.Lock()
	screen.mutex.Lock()

	C.SDL_FreeSurface((*C.SDL_Surface)(cast(screen.intSurface)))

	screen.destroy()
	if screen == currentVideoSurface {
		currentVideoSurface = nil
	}

	screen.mutex.Unlock()
	globalMutex.Unlock()
}

// Locks a surface for direct access.
func (screen *Surface) Lock() int {
	screen.mutex.Lock()
	status := int(C.SDL_LockSurface((*C.SDL_Surface)(cast(screen.intSurface))))
	screen.mutex.Unlock()
	return status
}

// Unlocks a previously locked surface.
func (screen *Surface) Unlock() {
	screen.mutex.Lock()
	C.SDL_UnlockSurface((*C.SDL_Surface)(cast(screen.intSurface)))
	screen.mutex.Unlock()
}

func (dst *Surface) Blit(dstrect *Rect, src *Surface, srcrect *Rect) int {
	src.mutex.RLock()
	dst.mutex.Lock()

	var ret = C.SDL_UpperBlit(
		(*C.SDL_Surface)(cast(src.intSurface)),
		(*C.SDL_Rect)(cast(srcrect)),
		(*C.SDL_Surface)(cast(dst.intSurface)),
		(*C.SDL_Rect)(cast(dstrect)))

	dst.mutex.Unlock()
	src.mutex.RUnlock()

	return int(ret)
}

// This function performs a fast fill of the given rectangle with some color.
func (dst *Surface) FillRect(dstrect *Rect, color uint32) int {
	dst.mutex.Lock()

	var ret = C.SDL_FillRect(
		(*C.SDL_Surface)(cast(dst.intSurface)),
		(*C.SDL_Rect)(cast(dstrect)),
		C.Uint32(color))

	dst.mutex.Unlock()

	return int(ret)
}

// Adjusts the alpha properties of a Surface.
func (s *Surface) SetAlpha(flags uint32, alpha uint8) int {
	s.mutex.Lock()
	status := int(C.SDL_SetAlpha((*C.SDL_Surface)(cast(s.intSurface)), C.Uint32(flags), C.Uint8(alpha)))
	s.mutex.Unlock()
	return status
}

// Sets the color key (transparent pixel)  in  a  blittable  surface  and
// enables or disables RLE blit acceleration.
func (s *Surface) SetColorKey(flags uint32, ColorKey uint32) int {
	return int(C.SDL_SetColorKey((*C.SDL_Surface)(cast(s)),
		C.Uint32(flags), C.Uint32(ColorKey)))
}

// Gets the clipping rectangle for a surface.
func (s *Surface) GetClipRect(r *Rect) {
	s.mutex.RLock()
	C.SDL_GetClipRect((*C.SDL_Surface)(cast(s.intSurface)), (*C.SDL_Rect)(cast(r)))
	s.mutex.RUnlock()
	return
}

// Sets the clipping rectangle for a surface.
func (s *Surface) SetClipRect(r *Rect) {
	s.mutex.Lock()
	C.SDL_SetClipRect((*C.SDL_Surface)(cast(s.intSurface)), (*C.SDL_Rect)(cast(r)))
	s.mutex.Unlock()
	return
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
	globalMutex.Lock()

	cfile := C.CString(file)
	var screen = C.IMG_Load(cfile)
	C.free(unsafe.Pointer(cfile))

	globalMutex.Unlock()

	return Wrap((*InternalSurface)(cast(screen)))
}

// Creates an empty Surface.
func CreateRGBSurface(flags uint32, width int, height int, bpp int, Rmask uint32, Gmask uint32, Bmask uint32, Amask uint32) *Surface {
	globalMutex.Lock()

	p := C.SDL_CreateRGBSurface(C.Uint32(flags), C.int(width), C.int(height), C.int(bpp),
		C.Uint32(Rmask), C.Uint32(Gmask), C.Uint32(Bmask), C.Uint32(Amask))

	globalMutex.Unlock()

	return Wrap((*InternalSurface)(cast(p)))
}

// Converts a surface to the display format
func DisplayFormat(src *Surface) *Surface {
	p := C.SDL_DisplayFormat((*C.SDL_Surface)(cast(src)))
	return (*Surface)(cast(p))
}

// ========
// Keyboard
// ========

// Enables UNICODE translation.
func EnableUNICODE(enable int) int {
	globalMutex.Lock()
	previous := int(C.SDL_EnableUNICODE(C.int(enable)))
	globalMutex.Unlock()
	return previous
}

// Sets keyboard repeat rate.
func EnableKeyRepeat(delay, interval int) int {
	globalMutex.Lock()
	status := int(C.SDL_EnableKeyRepeat(C.int(delay), C.int(interval)))
	globalMutex.Unlock()
	return status
}

// Gets keyboard repeat rate.
func GetKeyRepeat() (int, int) {
	var delay int
	var interval int

	globalMutex.Lock()
	C.SDL_GetKeyRepeat((*C.int)(cast(&delay)), (*C.int)(cast(&interval)))
	globalMutex.Unlock()

	return delay, interval
}

// Gets a snapshot of the current keyboard state
func GetKeyState() []uint8 {
	globalMutex.Lock()

	var numkeys C.int
	array := C.SDL_GetKeyState(&numkeys)

	var ptr = make([]uint8, numkeys)

	*((**C.Uint8)(unsafe.Pointer(&ptr))) = array // TODO

	globalMutex.Unlock()

	return ptr

}

// Modifier
type Mod C.int

// Key
type Key C.int

// Gets the state of modifier keys
func GetModState() Mod {
	globalMutex.Lock()
	state := Mod(C.SDL_GetModState())
	globalMutex.Unlock()
	return state
}

// Sets the state of modifier keys
func SetModState(modstate Mod) {
	globalMutex.Lock()
	C.SDL_SetModState(C.SDLMod(modstate))
	globalMutex.Unlock()
}

// Gets the name of an SDL virtual keysym
func GetKeyName(key Key) string {
	globalMutex.Lock()
	name := C.GoString(C.SDL_GetKeyName(C.SDLKey(key)))
	globalMutex.Unlock()
	return name
}


// ======
// Events
// ======

// Polls for currently pending events
func (event *Event) poll() bool {
	globalMutex.Lock()

	var ret = C.SDL_PollEvent((*C.SDL_Event)(cast(event)))

	if ret != 0 {
		if (event.Type == VIDEORESIZE) && (currentVideoSurface != nil) {
			currentVideoSurface.reload()
		}
	}

	globalMutex.Unlock()

	return ret != 0
}

// =====
// Mouse
// =====

// Retrieves the current state of the mouse.
func GetMouseState(x, y *int) uint8 {
	globalMutex.Lock()
	state := uint8(C.SDL_GetMouseState((*C.int)(cast(x)), (*C.int)(cast(y))))
	globalMutex.Unlock()
	return state
}

// Retrieves the current state of the mouse relative to the last time this
// function was called.
func GetRelativeMouseState(x, y *int) uint8 {
	globalMutex.Lock()
	state := uint8(C.SDL_GetRelativeMouseState((*C.int)(cast(x)), (*C.int)(cast(y))))
	globalMutex.Unlock()
	return state
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

// Toggle whether or not the cursor is shown on the screen.
func ShowCursor(toggle int) int {
	globalMutex.Lock()
	state := int(C.SDL_ShowCursor((C.int)(toggle)))
	globalMutex.Unlock()
	return state
}
