package sdl

type PixelFormat struct {
	Palette       *Palette
	BitsPerPixel  uint8
	BytesPerPixel uint8
	Rloss         uint8
	Gloss         uint8
	Bloss         uint8
	Aloss         uint8
	Rshift        uint8
	Gshift        uint8
	Bshift        uint8
	Ashift        uint8
	Pad0          [2]byte
	Rmask         uint32
	Gmask         uint32
	Bmask         uint32
	Amask         uint32
	Colorkey      uint32
	Alpha         uint8
	Pad1          [3]byte
}

type Rect struct {
	X int16
	Y int16
	W uint16
	H uint16
}

type Color struct {
	R      uint8
	G      uint8
	B      uint8
	Unused uint8
}

type Palette struct {
	Ncolors int32
	Colors  *Color
}

type internalVideoInfo struct {
	Flags     uint32
	Video_mem uint32
	Vfmt      *PixelFormat
	Current_w int32
	Current_h int32
}

type Overlay struct {
	Format  uint32
	W       int32
	H       int32
	Planes  int32
	Pitches *uint16
	Pixels  **uint8
	Hwfuncs *[0]byte /* sprivate_yuvhwfuncs */
	Hwdata  *[0]byte /* sprivate_yuvhwdata */
	Pad0    [4]byte
}

type ActiveEvent struct {
	Type  uint8
	Gain  uint8
	State uint8
}

type KeyboardEvent struct {
	Type   uint8
	Which  uint8
	State  uint8
	Pad0   [1]byte
	Keysym Keysym
}

type MouseMotionEvent struct {
	Type  uint8
	Which uint8
	State uint8
	Pad0  [1]byte
	X     uint16
	Y     uint16
	Xrel  int16
	Yrel  int16
}

type MouseButtonEvent struct {
	Type   uint8
	Which  uint8
	Button uint8
	State  uint8
	X      uint16
	Y      uint16
}

type JoyAxisEvent struct {
	Type  uint8
	Which uint8
	Axis  uint8
	Pad0  [1]byte
	Value int16
}

type JoyBallEvent struct {
	Type  uint8
	Which uint8
	Ball  uint8
	Pad0  [1]byte
	Xrel  int16
	Yrel  int16
}

type JoyHatEvent struct {
	Type  uint8
	Which uint8
	Hat   uint8
	Value uint8
}

type JoyButtonEvent struct {
	Type   uint8
	Which  uint8
	Button uint8
	State  uint8
}

type ResizeEvent struct {
	Type uint8
	Pad0 [3]byte
	W    int32
	H    int32
}

type ExposeEvent struct {
	Type uint8
}

type QuitEvent struct {
	Type uint8
}

type UserEvent struct {
	Type  uint8
	Pad0  [3]byte
	Code  int32
	Data1 *byte
	Data2 *byte
}

type SysWMmsg struct{}

type SysWMEvent struct {
	Type uint8
	Pad0 [3]byte
	Msg  *SysWMmsg
}

type Event struct {
	Type uint8
	Pad0 [19]byte
}

type Keysym struct {
	Scancode uint8
	Pad0     [3]byte
	Sym      uint32
	Mod      uint32
	Unicode  uint16
}

type BallDelta struct {
	Dx int
	Dy int
}

// type Joystick struct {
// 	Index uint8		// Device index
// 	Name	*int8 // Joystick name - system dependent
// 	Naxes int		// Number of axis controls on the joystick
// 	Axes *int16		// Current axis states
// 	Nhats int		// Number of hats on the joystick
// 	Hats *uint8		// Current hat states
// 	Nballs int		// Number of trackballs on the joystick
// 	Balls *BallDelta // Current ball motion deltas	
// 	Nbuttons int		// Number of buttons on the joystick
// 	Buttons *uint8		// Current button states
// 	Hwdata  *[0]byte // Driver dependent information
// 	Ref_count int		// Reference count for multiple opens
// }
