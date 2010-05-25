// godefs -g sdl constants.c

// MACHINE GENERATED - DO NOT EDIT.

package sdl

// Constants
const (
	INIT_TIMER          = 0x1
	INIT_AUDIO          = 0x10
	INIT_VIDEO          = 0x20
	INIT_CDROM          = 0x100
	INIT_JOYSTICK       = 0x200
	INIT_NOPARACHUTE    = 0x100000
	INIT_EVENTTHREAD    = 0x1000000
	INIT_EVERYTHING     = 0xffff
	SWSURFACE           = 0
	HWSURFACE           = 0x1
	ASYNCBLIT           = 0x4
	ANYFORMAT           = 0x10000000
	HWPALETTE           = 0x20000000
	DOUBLEBUF           = 0x40000000
	FULLSCREEN          = 0x80000000
	OPENGL              = 0x2
	OPENGLBLIT          = 0xa
	RESIZABLE           = 0x10
	NOFRAME             = 0x20
	HWACCEL             = 0x100
	SRCCOLORKEY         = 0x1000
	RLEACCELOK          = 0x2000
	RLEACCEL            = 0x4000
	SRCALPHA            = 0x10000
	PREALLOC            = 0x1000000
	YV12_OVERLAY        = 0x32315659
	IYUV_OVERLAY        = 0x56555949
	YUY2_OVERLAY        = 0x32595559
	UYVY_OVERLAY        = 0x59565955
	YVYU_OVERLAY        = 0x55595659
	LOGPAL              = 0x1
	PHYSPAL             = 0x2
	NOEVENT             = 0
	ACTIVEEVENT         = 0x1
	KEYDOWN             = 0x2
	KEYUP               = 0x3
	MOUSEMOTION         = 0x4
	MOUSEBUTTONDOWN     = 0x5
	MOUSEBUTTONUP       = 0x6
	JOYAXISMOTION       = 0x7
	JOYBALLMOTION       = 0x8
	JOYHATMOTION        = 0x9
	JOYBUTTONDOWN       = 0xa
	JOYBUTTONUP         = 0xb
	QUIT                = 0xc
	SYSWMEVENT          = 0xd
	EVENT_RESERVEDA     = 0xe
	EVENT_RESERVEDB     = 0xf
	VIDEORESIZE         = 0x10
	VIDEOEXPOSE         = 0x11
	EVENT_RESERVED2     = 0x12
	EVENT_RESERVED3     = 0x13
	EVENT_RESERVED4     = 0x14
	EVENT_RESERVED5     = 0x15
	EVENT_RESERVED6     = 0x16
	EVENT_RESERVED7     = 0x17
	USEREVENT           = 0x18
	NUMEVENTS           = 0x20
	ACTIVEEVENTMASK     = 0x2
	KEYDOWNMASK         = 0x4
	KEYUPMASK           = 0x8
	KEYEVENTMASK        = 0xc
	MOUSEMOTIONMASK     = 0x10
	MOUSEBUTTONDOWNMASK = 0x20
	MOUSEBUTTONUPMASK   = 0x40
	MOUSEEVENTMASK      = 0x70
	JOYAXISMOTIONMASK   = 0x80
	JOYBALLMOTIONMASK   = 0x100
	JOYHATMOTIONMASK    = 0x200
	JOYBUTTONDOWNMASK   = 0x400
	JOYBUTTONUPMASK     = 0x800
	JOYEVENTMASK        = 0xf80
	VIDEORESIZEMASK     = 0x10000
	VIDEOEXPOSEMASK     = 0x20000
	QUITMASK            = 0x1000
	SYSWMEVENTMASK      = 0x2000
	K_UNKNOWN           = 0
	K_FIRST             = 0
	K_BACKSPACE         = 0x8
	K_TAB               = 0x9
	K_CLEAR             = 0xc
	K_RETURN            = 0xd
	K_PAUSE             = 0x13
	K_ESCAPE            = 0x1b
	K_SPACE             = 0x20
	K_EXCLAIM           = 0x21
	K_QUOTEDBL          = 0x22
	K_HASH              = 0x23
	K_DOLLAR            = 0x24
	K_AMPERSAND         = 0x26
	K_QUOTE             = 0x27
	K_LEFTPAREN         = 0x28
	K_RIGHTPAREN        = 0x29
	K_ASTERISK          = 0x2a
	K_PLUS              = 0x2b
	K_COMMA             = 0x2c
	K_MINUS             = 0x2d
	K_PERIOD            = 0x2e
	K_SLASH             = 0x2f
	K_0                 = 0x30
	K_1                 = 0x31
	K_2                 = 0x32
	K_3                 = 0x33
	K_4                 = 0x34
	K_5                 = 0x35
	K_6                 = 0x36
	K_7                 = 0x37
	K_8                 = 0x38
	K_9                 = 0x39
	K_COLON             = 0x3a
	K_SEMICOLON         = 0x3b
	K_LESS              = 0x3c
	K_EQUALS            = 0x3d
	K_GREATER           = 0x3e
	K_QUESTION          = 0x3f
	K_AT                = 0x40
	K_LEFTBRACKET       = 0x5b
	K_BACKSLASH         = 0x5c
	K_RIGHTBRACKET      = 0x5d
	K_CARET             = 0x5e
	K_UNDERSCORE        = 0x5f
	K_BACKQUOTE         = 0x60
	K_a                 = 0x61
	K_b                 = 0x62
	K_c                 = 0x63
	K_d                 = 0x64
	K_e                 = 0x65
	K_f                 = 0x66
	K_g                 = 0x67
	K_h                 = 0x68
	K_i                 = 0x69
	K_j                 = 0x6a
	K_k                 = 0x6b
	K_l                 = 0x6c
	K_m                 = 0x6d
	K_n                 = 0x6e
	K_o                 = 0x6f
	K_p                 = 0x70
	K_q                 = 0x71
	K_r                 = 0x72
	K_s                 = 0x73
	K_t                 = 0x74
	K_u                 = 0x75
	K_v                 = 0x76
	K_w                 = 0x77
	K_x                 = 0x78
	K_y                 = 0x79
	K_z                 = 0x7a
	K_DELETE            = 0x7f
	K_WORLD_0           = 0xa0
	K_WORLD_1           = 0xa1
	K_WORLD_2           = 0xa2
	K_WORLD_3           = 0xa3
	K_WORLD_4           = 0xa4
	K_WORLD_5           = 0xa5
	K_WORLD_6           = 0xa6
	K_WORLD_7           = 0xa7
	K_WORLD_8           = 0xa8
	K_WORLD_9           = 0xa9
	K_WORLD_10          = 0xaa
	K_WORLD_11          = 0xab
	K_WORLD_12          = 0xac
	K_WORLD_13          = 0xad
	K_WORLD_14          = 0xae
	K_WORLD_15          = 0xaf
	K_WORLD_16          = 0xb0
	K_WORLD_17          = 0xb1
	K_WORLD_18          = 0xb2
	K_WORLD_19          = 0xb3
	K_WORLD_20          = 0xb4
	K_WORLD_21          = 0xb5
	K_WORLD_22          = 0xb6
	K_WORLD_23          = 0xb7
	K_WORLD_24          = 0xb8
	K_WORLD_25          = 0xb9
	K_WORLD_26          = 0xba
	K_WORLD_27          = 0xbb
	K_WORLD_28          = 0xbc
	K_WORLD_29          = 0xbd
	K_WORLD_30          = 0xbe
	K_WORLD_31          = 0xbf
	K_WORLD_32          = 0xc0
	K_WORLD_33          = 0xc1
	K_WORLD_34          = 0xc2
	K_WORLD_35          = 0xc3
	K_WORLD_36          = 0xc4
	K_WORLD_37          = 0xc5
	K_WORLD_38          = 0xc6
	K_WORLD_39          = 0xc7
	K_WORLD_40          = 0xc8
	K_WORLD_41          = 0xc9
	K_WORLD_42          = 0xca
	K_WORLD_43          = 0xcb
	K_WORLD_44          = 0xcc
	K_WORLD_45          = 0xcd
	K_WORLD_46          = 0xce
	K_WORLD_47          = 0xcf
	K_WORLD_48          = 0xd0
	K_WORLD_49          = 0xd1
	K_WORLD_50          = 0xd2
	K_WORLD_51          = 0xd3
	K_WORLD_52          = 0xd4
	K_WORLD_53          = 0xd5
	K_WORLD_54          = 0xd6
	K_WORLD_55          = 0xd7
	K_WORLD_56          = 0xd8
	K_WORLD_57          = 0xd9
	K_WORLD_58          = 0xda
	K_WORLD_59          = 0xdb
	K_WORLD_60          = 0xdc
	K_WORLD_61          = 0xdd
	K_WORLD_62          = 0xde
	K_WORLD_63          = 0xdf
	K_WORLD_64          = 0xe0
	K_WORLD_65          = 0xe1
	K_WORLD_66          = 0xe2
	K_WORLD_67          = 0xe3
	K_WORLD_68          = 0xe4
	K_WORLD_69          = 0xe5
	K_WORLD_70          = 0xe6
	K_WORLD_71          = 0xe7
	K_WORLD_72          = 0xe8
	K_WORLD_73          = 0xe9
	K_WORLD_74          = 0xea
	K_WORLD_75          = 0xeb
	K_WORLD_76          = 0xec
	K_WORLD_77          = 0xed
	K_WORLD_78          = 0xee
	K_WORLD_79          = 0xef
	K_WORLD_80          = 0xf0
	K_WORLD_81          = 0xf1
	K_WORLD_82          = 0xf2
	K_WORLD_83          = 0xf3
	K_WORLD_84          = 0xf4
	K_WORLD_85          = 0xf5
	K_WORLD_86          = 0xf6
	K_WORLD_87          = 0xf7
	K_WORLD_88          = 0xf8
	K_WORLD_89          = 0xf9
	K_WORLD_90          = 0xfa
	K_WORLD_91          = 0xfb
	K_WORLD_92          = 0xfc
	K_WORLD_93          = 0xfd
	K_WORLD_94          = 0xfe
	K_WORLD_95          = 0xff
	K_KP0               = 0x100
	K_KP1               = 0x101
	K_KP2               = 0x102
	K_KP3               = 0x103
	K_KP4               = 0x104
	K_KP5               = 0x105
	K_KP6               = 0x106
	K_KP7               = 0x107
	K_KP8               = 0x108
	K_KP9               = 0x109
	K_KP_PERIOD         = 0x10a
	K_KP_DIVIDE         = 0x10b
	K_KP_MULTIPLY       = 0x10c
	K_KP_MINUS          = 0x10d
	K_KP_PLUS           = 0x10e
	K_KP_ENTER          = 0x10f
	K_KP_EQUALS         = 0x110
	K_UP                = 0x111
	K_DOWN              = 0x112
	K_RIGHT             = 0x113
	K_LEFT              = 0x114
	K_INSERT            = 0x115
	K_HOME              = 0x116
	K_END               = 0x117
	K_PAGEUP            = 0x118
	K_PAGEDOWN          = 0x119
	K_F1                = 0x11a
	K_F2                = 0x11b
	K_F3                = 0x11c
	K_F4                = 0x11d
	K_F5                = 0x11e
	K_F6                = 0x11f
	K_F7                = 0x120
	K_F8                = 0x121
	K_F9                = 0x122
	K_F10               = 0x123
	K_F11               = 0x124
	K_F12               = 0x125
	K_F13               = 0x126
	K_F14               = 0x127
	K_F15               = 0x128
	K_NUMLOCK           = 0x12c
	K_CAPSLOCK          = 0x12d
	K_SCROLLOCK         = 0x12e
	K_RSHIFT            = 0x12f
	K_LSHIFT            = 0x130
	K_RCTRL             = 0x131
	K_LCTRL             = 0x132
	K_RALT              = 0x133
	K_LALT              = 0x134
	K_RMETA             = 0x135
	K_LMETA             = 0x136
	K_LSUPER            = 0x137
	K_RSUPER            = 0x138
	K_MODE              = 0x139
	K_COMPOSE           = 0x13a
	K_HELP              = 0x13b
	K_PRINT             = 0x13c
	K_SYSREQ            = 0x13d
	K_BREAK             = 0x13e
	K_MENU              = 0x13f
	K_POWER             = 0x140
	K_EURO              = 0x141
	K_UNDO              = 0x142
	KMOD_NONE           = 0
	KMOD_LSHIFT         = 0x1
	KMOD_RSHIFT         = 0x2
	KMOD_LCTRL          = 0x40
	KMOD_RCTRL          = 0x80
	KMOD_LALT           = 0x100
	KMOD_RALT           = 0x200
	KMOD_LMETA          = 0x400
	KMOD_RMETA          = 0x800
	KMOD_NUM            = 0x1000
	KMOD_CAPS           = 0x2000
	KMOD_MODE           = 0x4000
	KMOD_RESERVED       = 0x8000
)

// Types

type Surface struct {
	Flags          uint32
	Format         *PixelFormat
	W              int32
	H              int32
	Pitch          uint16
	Pad0           [2]byte
	Pixels         *byte
	Offset         int32
	Hwdata         *[0]byte /* sprivate_hwdata */
	Clip_rect      Rect
	Unused1        uint32
	Locked         uint32
	Map            *[0]byte /* sSDL_BlitMap */
	Format_version uint32
	Refcount       int32
}

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
