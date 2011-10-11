/*
A binding of SDL and SDL_image.

The binding is works in pretty much the same way as it does in C, although
some of the functions have been altered to give them an object-oriented
flavor (eg. Rather than sdl.Flip(surface) it's surface.Flip() )
*/
package sdl

// #cgo pkg-config: sdl
// #cgo LDFLAGS: -lSDL_image
// struct private_hwdata{};
// struct SDL_BlitMap{};
// #define map _map
//
// #include "SDL.h"
// #include "SDL_image.h"
// static void SetError(const char* description){SDL_SetError("%s",description);}
//
// static int vectorLength(void** start)
// {
//     void **ptr=start;
//     for(;*ptr!=NULL;ptr++);
//
//     return ptr-start;
// }
//
// static int RWseek(SDL_RWops *rw, int os, int w){return SDL_RWseek(rw, os, w);}
// static int RWread(SDL_RWops *rw, void *d, int size, int max){return SDL_RWread(rw, d, size, max);}
// static int RWwrite(SDL_RWops *rw, void *d, int size, int num){return SDL_RWwrite(rw, d, size, num);}
// static int RWclose(SDL_RWops *rw){return SDL_RWclose(rw);}
// static int __SDL_SaveBMP(SDL_Surface *surface, const char *file) { return SDL_SaveBMP(surface, file); }
import "C"
import "unsafe"
import "image"
import "image/draw"
import "image/color"
import "io"
import "io/ioutil"
import "os"
import "fmt"
import "reflect"

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
	defer C.free(unsafe.Pointer(cdescription))
	C.SetError(cdescription)
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

func ListModes(format *PixelFormat, flags uint32) []*Rect {

	modes := C.SDL_ListModes((*C.SDL_PixelFormat)(cast(format)), C.Uint32(flags))

	if modes == nil { //modes == 0, no modes available
		return make([]*Rect, 0)
	}

	if ^uintptr(unsafe.Pointer(modes)) == 0 { //modes == -1, any dimension is ok
		return nil
	}

	var ret = make([]*Rect, C.vectorLength((*unsafe.Pointer)(unsafe.Pointer(modes))))
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
func (screen *Surface) Unlock() { C.SDL_UnlockSurface((*C.SDL_Surface)(cast(screen))) }

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

func (img *Surface) pixPtr(x, y int) reflect.Value {
	imglen := int(img.W * img.H)
	h := reflect.SliceHeader{
		uintptr(unsafe.Pointer(img.Pixels)),
		imglen,
		imglen,
	}

	switch img.Format.BitsPerPixel {
	case 8:
		pix := &(*(*[]uint8)(unsafe.Pointer(&h)))[(y*int(img.W))+x]
		return reflect.ValueOf(pix).Elem()
	case 16:
		pix := &(*(*[]uint16)(unsafe.Pointer(&h)))[(y*int(img.W))+x]
		return reflect.ValueOf(pix).Elem()
	case 32:
		pix := &(*(*[]uint32)(unsafe.Pointer(&h)))[(y*int(img.W))+x]
		return reflect.ValueOf(pix).Elem()
	case 64:
		pix := &(*(*[]uint64)(unsafe.Pointer(&h)))[(y*int(img.W))+x]
		return reflect.ValueOf(pix).Elem()
	}

	panic(fmt.Errorf("Image has unexpected BPP: %v", img.Format.BitsPerPixel))
}

func (img *Surface) ColorModel() color.Model {
	switch img.Format.BitsPerPixel {
	case 8:
		return img.Format.Palette.GoPalette()
	case 32:
		return color.NRGBAModel
	case 64:
		return color.NRGBA64Model
	}

	return color.NRGBAModel
}

func (img *Surface) Bounds() image.Rectangle {
	return image.Rect(0, 0, int(img.W), int(img.H))
}

func (img *Surface) At(x, y int) color.Color {
	var r, g, b, a uint8
	GetRGBA(uint32(img.pixPtr(x, y).Uint()), img.Format, &r, &g, &b, &a)

	return img.ColorModel().Convert(color.NRGBA{r, g, b, a})
}

func (img *Surface) Set(x, y int, c color.Color) {
	img.Lock()
	defer img.Unlock()

	r, g, b, a := c.RGBA()

	pix := img.pixPtr(x, y)
	pix.SetUint(uint64(MapRGBA(img.Format, uint8(r), uint8(g), uint8(b), uint8(a))))
}

func ColorFromGoColor(in color.Color) Color {
	r, g, b, _ := in.RGBA()

	return Color{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
	}
}

func (c Color) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R)
	r |= r << 8
	g = uint32(c.G)
	g |= g << 8
	b = uint32(c.B)
	b |= b << 8
	a = 255

	return
}

func PaletteFromGoPalette(in color.Palette) *Palette {
	col := make([]Color, len(in))
	for i := range col {
		col[i] = ColorFromGoColor(in[i])
	}

	pal := &Palette{Ncolors: int32(len(in))}

	size := C.size_t(len(in) * 4)
	pal.Colors = (*Color)(C.malloc(size))
	C.memcpy(unsafe.Pointer(pal.Colors), unsafe.Pointer(&col[0]), size)

	return pal
}

func (pal *Palette) GoPalette() color.Palette {
	if (pal == nil) || (pal.Ncolors == 0) {
		return nil
	}

	s := *(*[]Color)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(pal.Colors)),
		Len:  int(pal.Ncolors),
		Cap:  int(pal.Ncolors),
	}))

	out := make(color.Palette, pal.Ncolors)
	for i := range out {
		out[i] = s[i]
	}

	return out
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
	res := C.__SDL_SaveBMP((*C.SDL_Surface)(cast(src)), cfile)
	C.free(unsafe.Pointer(cfile))
	return int(res)
}

// Loads Surface from RWops (using IMG_Load_RW).
func Load_RW(rw *RWops, ac bool) *Surface {
	acArg := C.int(0)
	if ac {
		acArg = 1
	}

	return (*Surface)(unsafe.Pointer(C.IMG_Load_RW((*C.SDL_RWops)(rw), acArg)))
}

// Loads Surface of type t from RWops (using IMG_LoadTyped_RW).
func LoadTyped_RW(rw *RWops, ac bool, t string) *Surface {
	ct := C.CString(t)
	defer C.free(unsafe.Pointer(ct))

	acArg := C.int(0)
	if ac {
		acArg = 1
	}

	return (*Surface)(unsafe.Pointer(C.IMG_LoadTyped_RW((*C.SDL_RWops)(rw), acArg, ct)))
}

// Create new sdl.Surface from image.Image
func CreateSurfaceFromImage(img image.Image) *Surface {
	r := img.Bounds()

	var bpp int
	var pix []uint8
	var pal *Palette
	switch c := img.(type) {
	case *image.NRGBA:
		pix = c.Pix
		bpp = 32
	case *image.RGBA:
		pix = c.Pix
		bpp = 32
	case *image.Paletted:
		pix = c.Pix
		bpp = 8
		pal = PaletteFromGoPalette(c.Palette)
	default:
		nrgba := image.NewNRGBA(r)
		draw.Draw(nrgba, r, img, image.ZP, draw.Src)
		pix = nrgba.Pix
		bpp = 32
	}

	s := CreateRGBSurface(SWSURFACE, r.Dx(), r.Dy(), bpp, 0x000000ff, 0x0000ff00, 0x00ff0000, 0xff000000)

	if pal != nil {
		s.Format.Palette = pal
	}

	s.Lock()
	defer s.Unlock()
	C.memcpy(unsafe.Pointer(s.Pixels), unsafe.Pointer(&pix[0]), C.size_t(len(pix)))

	return s
}

// Creates an empty Surface.
func CreateRGBSurface(flags uint32, width int, height int, bpp int, Rmask uint32, Gmask uint32, Bmask uint32, Amask uint32) *Surface {
	p := C.SDL_CreateRGBSurface(C.Uint32(flags), C.int(width), C.int(height), C.int(bpp),
		C.Uint32(Rmask), C.Uint32(Gmask), C.Uint32(Bmask), C.Uint32(Amask))
	return (*Surface)(cast(p))
}

// Creates a Surface from existing pixel data. It expects pix to be a slice, array, or pointer.
func CreateRGBSurfaceFrom(pix interface{}, w, h, d, p int, rm, gm, bm, am uint32) *Surface {
	var ptr unsafe.Pointer
	switch v := reflect.ValueOf(pix); v.Kind() {
	case reflect.Ptr, reflect.UnsafePointer, reflect.Slice:
		ptr = unsafe.Pointer(v.Pointer())
	default:
		panic("Don't know how to handle type: " + v.Kind().String())
	}

	s := C.SDL_CreateRGBSurfaceFrom(ptr, C.int(w), C.int(h), C.int(d), C.int(p), C.Uint32(rm), C.Uint32(gm), C.Uint32(bm), C.Uint32(am))

	return (*Surface)(cast(s))
}

// Converts a surface to the display format
func DisplayFormat(src *Surface) *Surface {
	p := C.SDL_DisplayFormat((*C.SDL_Surface)(cast(src)))
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

func JoystickEventState(flag int) int {
	return int(C.SDL_JoystickEventState(C.int(flag)))
}

func NumJoysticks() int {
	return int(C.SDL_NumJoysticks())
}

func JoystickName(n int) string {
	return C.GoString(C.SDL_JoystickName(C.int(n)))
}

func JoystickOpen(n int) *Joystick {
	return (*Joystick)(unsafe.Pointer(C.SDL_JoystickOpen(C.int(n))))
}

func JoystickOpened(n int) bool {
	return C.SDL_JoystickOpened(C.int(n)) != 0
}

func (j *Joystick) Index() int {
	return int(C.SDL_JoystickIndex((*C.SDL_Joystick)(unsafe.Pointer(j))))
}

func (j *Joystick) NumAxes() int {
	return int(C.SDL_JoystickNumAxes((*C.SDL_Joystick)(unsafe.Pointer(j))))
}

func (j *Joystick) NumBalls() int {
	return int(C.SDL_JoystickNumBalls((*C.SDL_Joystick)(unsafe.Pointer(j))))
}

func (j *Joystick) NumHats() int {
	return int(C.SDL_JoystickNumHats((*C.SDL_Joystick)(unsafe.Pointer(j))))
}

func (j *Joystick) NumButtons() int {
	return int(C.SDL_JoystickNumButtons((*C.SDL_Joystick)(unsafe.Pointer(j))))
}

func JoystickUpdate() {
	C.SDL_JoystickUpdate()
}

func (j *Joystick) GetAxis(n int) int16 {
	return int16(C.SDL_JoystickGetAxis((*C.SDL_Joystick)(unsafe.Pointer(j)), C.int(n)))
}

func (j *Joystick) GetHat(n int) int8 {
	return int8(C.SDL_JoystickGetHat((*C.SDL_Joystick)(unsafe.Pointer(j)), C.int(n)))
}

func (j *Joystick) GetButton(n int) bool {
	return C.SDL_JoystickGetButton((*C.SDL_Joystick)(unsafe.Pointer(j)), C.int(n)) != 0
}

func (j *Joystick) GetBall(n int) (dx, dy int) {
	ret := C.SDL_JoystickGetBall((*C.SDL_Joystick)(unsafe.Pointer(j)),
		C.int(n),
		(*C.int)(unsafe.Pointer(&dx)),
		(*C.int)(unsafe.Pointer(&dy)),
	)

	if ret < 0 {
		dx, dy = -1, -1
	}

	return
}

func (j *Joystick) Close() os.Error {
	C.SDL_JoystickClose((*C.SDL_Joystick)(unsafe.Pointer(j)))

	return nil
}

// Events

// Waits indefinitely for the next available event
func WaitEvent() Event {
	var cev C.SDL_Event
	ret := C.SDL_WaitEvent(&cev)
	if ret == 0 {
		return nil
	}

	return goEvent((*cevent)(cast(&cev)))
}

// Push the event onto the event queue
func PushEvent(event Event) bool {
	ret := C.SDL_PushEvent((*C.SDL_Event)(cast(cEvent(event))))

	return ret != 0
}

// Polls for currently pending events
func PollEvent() Event {
	var cev C.SDL_Event
	ret := C.SDL_PollEvent(&cev)
	if ret == 0 {
		return nil
	}

	return goEvent((*cevent)(cast(&cev)))
}

func goEvent(cev *cevent) Event {
	switch cev.Type {
	case ACTIVEEVENT:
		return (*ActiveEvent)(cast(cev))
	case KEYUP, KEYDOWN:
		return (*KeyboardEvent)(cast(cev))
	case MOUSEMOTION:
		return (*MouseMotionEvent)(cast(cev))
	case MOUSEBUTTONUP, MOUSEBUTTONDOWN:
		return (*MouseButtonEvent)(cast(cev))
	case JOYAXISMOTION:
		return (*JoyAxisEvent)(cast(cev))
	case JOYBALLMOTION:
		return (*JoyBallEvent)(cast(cev))
	case JOYHATMOTION:
		return (*JoyHatEvent)(cast(cev))
	case JOYBUTTONUP, JOYBUTTONDOWN:
		return (*JoyButtonEvent)(cast(cev))
	case VIDEORESIZE:
		return (*ResizeEvent)(cast(cev))
	case VIDEOEXPOSE:
		return (*ExposeEvent)(cast(cev))
	case QUIT:
		return (*QuitEvent)(cast(cev))
	case USEREVENT:
		return (*UserEvent)(cast(cev))
	case SYSWMEVENT:
		return (*SysWMEvent)(cast(cev))
	}

	panic(fmt.Errorf("Unknown event type: %v", cev.Type))
}

func cEvent(ev Event) *cevent {
	evv := reflect.ValueOf(ev)
	return (*cevent)(cast(evv.UnsafeAddr()))
}

// Time

// Gets the number of milliseconds since the SDL library initialization.
func GetTicks() uint32 { return uint32(C.SDL_GetTicks()) }

// Waits a specified number of milliseconds before returning.
func Delay(ms uint32) { C.SDL_Delay(C.Uint32(ms)) }

func ShowCursor(toggle int) int {
	return int(C.SDL_ShowCursor(C.int(toggle)))
}

type RWops C.SDL_RWops

func AllocRW() *RWops {
	return (*RWops)(C.SDL_AllocRW())
}

func (rw *RWops) Free() {
	C.SDL_FreeRW((*C.SDL_RWops)(rw))
}

func RWFromMem(m []byte) *RWops {
	return (*RWops)(C.SDL_RWFromMem(unsafe.Pointer(&m[0]), C.int(len(m))))
}

func RWFromConstMem(m []byte) *RWops {
	return (*RWops)(C.SDL_RWFromConstMem(unsafe.Pointer(&m[0]), C.int(len(m))))
}

func RWFromReader(r io.Reader) *RWops {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		SetError(err.String())
		return nil
	}

	return RWFromConstMem(data)
}

func modeFromFlags(flag int) *C.char {
	switch flag {
	case os.O_RDONLY:
		return C.CString("r")
	case os.O_WRONLY | os.O_CREATE:
		return C.CString("w")
	case os.O_WRONLY | os.O_APPEND | os.O_CREATE:
		return C.CString("a")
	case os.O_RDWR:
		return C.CString("r+")
	case os.O_RDWR | os.O_CREATE:
		return C.CString("w+")
	case os.O_RDWR | os.O_APPEND | os.O_CREATE:
		return C.CString("a+")
	}

	SetError("Unknown mode.")
	return nil
}

func RWFromFile(file string, mode int) *RWops {
	cfile := C.CString(file)
	defer C.free(unsafe.Pointer(cfile))

	cmode := modeFromFlags(mode)
	if cmode == nil {
		return nil
	}
	defer C.free(unsafe.Pointer(cmode))

	return (*RWops)(C.SDL_RWFromFile(cfile, cmode))
}

// Causes 'SIGNONE: no trap'. Not sure why...
func RWFromFP(fp *os.File, ac bool) *RWops {
	acArg := C.int(0)
	if ac {
		acArg = 1
	}

	cmode := C.CString("r+") // Doesn't really matter, anyways. I hope.
	defer C.free(unsafe.Pointer(cmode))
	cfp := C.fdopen(C.int(fp.Fd()), cmode)

	return (*RWops)(C.SDL_RWFromFP(cfp, acArg))
}

func (rw *RWops) Tell() int64 {
	cur, err := rw.Seek(0, 1)
	if err != nil {
		SetError(err.String())
		return -1
	}

	return cur
}

func (rw *RWops) Length() int64 {
	cur := rw.Tell()
	if cur < 0 {
		return -1
	}

	eof, err := rw.Seek(0, 2)
	if err != nil {
		SetError(err.String())
		return -1
	}

	rw.Seek(cur, 0)

	return eof
}

func (rw *RWops) EOF() bool {
	cur := rw.Tell()
	eof := rw.Length()
	if (cur < 0) || (eof < 0) {
		panic(GetError())
	}

	if eof <= cur {
		return true
	}

	return false
}

func (rw *RWops) Seek(offset int64, whence int) (int64, os.Error) {
	var w C.int
	switch whence {
	case 0:
		w = C.SEEK_SET
	case 1:
		w = C.SEEK_CUR
	case 2:
		w = C.SEEK_END
	default:
		return offset, os.NewError("Bad whence.")
	}

	return int64(C.RWseek((*C.SDL_RWops)(rw), C.int(offset), w)), nil
}

func (rw *RWops) Read(buf []byte) (n int, err os.Error) {
	n = int(C.RWread((*C.SDL_RWops)(rw), unsafe.Pointer(&buf[0]), 1, C.int(len(buf))))

	if rw.EOF() {
		err = os.EOF
	}

	if n < 0 {
		n = 0
		err = os.NewError(GetError())
	}

	return
}

func (rw *RWops) Write(buf []byte) (n int, err os.Error) {
	n = int(C.RWwrite((*C.SDL_RWops)(rw), unsafe.Pointer(&buf[0]), 1, C.int(len(buf))))

	if n < 0 {
		n = 0
		err = os.NewError(GetError())
	}

	return
}

func (rw *RWops) Close() os.Error {
	if int(C.RWclose((*C.SDL_RWops)(rw))) != 0 {
		return os.NewError(GetError())
	}

	return nil
}
