// godefs -g sdl sdl.c

// MACHINE GENERATED - DO NOT EDIT.

package sdl

// Constants
const (
	INIT_AUDIO = 0x10;
	INIT_VIDEO = 0x20;
	INIT_CDROM = 0x100;
	INIT_JOYSTICK = 0x200;
	INIT_NOPARACHUTE = 0x100000;
	INIT_EVENTTHREAD = 0x1000000;
	INIT_EVERYTHING = 0xffff;
	SWSURFACE = 0;
	HWSURFACE = 0x1;
	ASYNCBLIT = 0x4;
	ANYFORMAT = 0x10000000;
	HWPALETTE = 0x20000000;
	DOUBLEBUF = 0x40000000;
	FULLSCREEN = 0x80000000;
	OPENGL = 0x2;
	OPENGLBLIT = 0xa;
	RESIZABLE = 0x10;
	NOFRAME = 0x20;
	HWACCEL = 0x100;
	SRCCOLORKEY = 0x1000;
	RLEACCELOK = 0x2000;
	RLEACCEL = 0x4000;
	SRCALPHA = 0x10000;
	PREALLOC = 0x1000000;
	YV12_OVERLAY = 0x32315659;
	IYUV_OVERLAY = 0x56555949;
	YUY2_OVERLAY = 0x32595559;
	UYVY_OVERLAY = 0x59565955;
	YVYU_OVERLAY = 0x55595659;
	LOGPAL = 0x1;
	PHYSPAL = 0x2;
	NOEVENT = 0;
	ACTIVEEVENT = 0x1;
	KEYDOWN = 0x2;
	KEYUP = 0x3;
	MOUSEMOTION = 0x4;
	MOUSEBUTTONDOWN = 0x5;
	MOUSEBUTTONUP = 0x6;
	JOYAXISMOTION = 0x7;
	JOYBALLMOTION = 0x8;
	JOYHATMOTION = 0x9;
	JOYBUTTONDOWN = 0xa;
	JOYBUTTONUP = 0xb;
	QUIT = 0xc;
	SYSWMEVENT = 0xd;
	EVENT_RESERVEDA = 0xe;
	EVENT_RESERVEDB = 0xf;
	VIDEORESIZE = 0x10;
	VIDEOEXPOSE = 0x11;
	EVENT_RESERVED2 = 0x12;
	EVENT_RESERVED3 = 0x13;
	EVENT_RESERVED4 = 0x14;
	EVENT_RESERVED5 = 0x15;
	EVENT_RESERVED6 = 0x16;
	EVENT_RESERVED7 = 0x17;
	USEREVENT = 0x18;
	NUMEVENTS = 0x20;
	ACTIVEEVENTMASK = 0x2;
	KEYDOWNMASK = 0x4;
	KEYUPMASK = 0x8;
	KEYEVENTMASK = 0xc;
	MOUSEMOTIONMASK = 0x10;
	MOUSEBUTTONDOWNMASK = 0x20;
	MOUSEBUTTONUPMASK = 0x40;
	MOUSEEVENTMASK = 0x70;
	JOYAXISMOTIONMASK = 0x80;
	JOYBALLMOTIONMASK = 0x100;
	JOYHATMOTIONMASK = 0x200;
	JOYBUTTONDOWNMASK = 0x400;
	JOYBUTTONUPMASK = 0x800;
	JOYEVENTMASK = 0xf80;
	VIDEORESIZEMASK = 0x10000;
	VIDEOEXPOSEMASK = 0x20000;
	QUITMASK = 0x1000;
	SYSWMEVENTMASK = 0x2000;
	SDLK_UNKNOWN = 0;
	SDLK_FIRST = 0;
	SDLK_BACKSPACE = 0x8;
	SDLK_TAB = 0x9;
	SDLK_CLEAR = 0xc;
	SDLK_RETURN = 0xd;
	SDLK_PAUSE = 0x13;
	SDLK_ESCAPE = 0x1b;
	SDLK_SPACE = 0x20;
	SDLK_EXCLAIM = 0x21;
	SDLK_QUOTEDBL = 0x22;
	SDLK_HASH = 0x23;
	SDLK_DOLLAR = 0x24;
	SDLK_AMPERSAND = 0x26;
	SDLK_QUOTE = 0x27;
	SDLK_LEFTPAREN = 0x28;
	SDLK_RIGHTPAREN = 0x29;
	SDLK_ASTERISK = 0x2a;
	SDLK_PLUS = 0x2b;
	SDLK_COMMA = 0x2c;
	SDLK_MINUS = 0x2d;
	SDLK_PERIOD = 0x2e;
	SDLK_SLASH = 0x2f;
	SDLK_0 = 0x30;
	SDLK_1 = 0x31;
	SDLK_2 = 0x32;
	SDLK_3 = 0x33;
	SDLK_4 = 0x34;
	SDLK_5 = 0x35;
	SDLK_6 = 0x36;
	SDLK_7 = 0x37;
	SDLK_8 = 0x38;
	SDLK_9 = 0x39;
	SDLK_COLON = 0x3a;
	SDLK_SEMICOLON = 0x3b;
	SDLK_LESS = 0x3c;
	SDLK_EQUALS = 0x3d;
	SDLK_GREATER = 0x3e;
	SDLK_QUESTION = 0x3f;
	SDLK_AT = 0x40;
	SDLK_LEFTBRACKET = 0x5b;
	SDLK_BACKSLASH = 0x5c;
	SDLK_RIGHTBRACKET = 0x5d;
	SDLK_CARET = 0x5e;
	SDLK_UNDERSCORE = 0x5f;
	SDLK_BACKQUOTE = 0x60;
	SDLK_a = 0x61;
	SDLK_b = 0x62;
	SDLK_c = 0x63;
	SDLK_d = 0x64;
	SDLK_e = 0x65;
	SDLK_f = 0x66;
	SDLK_g = 0x67;
	SDLK_h = 0x68;
	SDLK_i = 0x69;
	SDLK_j = 0x6a;
	SDLK_k = 0x6b;
	SDLK_l = 0x6c;
	SDLK_m = 0x6d;
	SDLK_n = 0x6e;
	SDLK_o = 0x6f;
	SDLK_p = 0x70;
	SDLK_q = 0x71;
	SDLK_r = 0x72;
	SDLK_s = 0x73;
	SDLK_t = 0x74;
	SDLK_u = 0x75;
	SDLK_v = 0x76;
	SDLK_w = 0x77;
	SDLK_x = 0x78;
	SDLK_y = 0x79;
	SDLK_z = 0x7a;
	SDLK_DELETE = 0x7f;
	SDLK_WORLD_0 = 0xa0;
	SDLK_WORLD_1 = 0xa1;
	SDLK_WORLD_2 = 0xa2;
	SDLK_WORLD_3 = 0xa3;
	SDLK_WORLD_4 = 0xa4;
	SDLK_WORLD_5 = 0xa5;
	SDLK_WORLD_6 = 0xa6;
	SDLK_WORLD_7 = 0xa7;
	SDLK_WORLD_8 = 0xa8;
	SDLK_WORLD_9 = 0xa9;
	SDLK_WORLD_10 = 0xaa;
	SDLK_WORLD_11 = 0xab;
	SDLK_WORLD_12 = 0xac;
	SDLK_WORLD_13 = 0xad;
	SDLK_WORLD_14 = 0xae;
	SDLK_WORLD_15 = 0xaf;
	SDLK_WORLD_16 = 0xb0;
	SDLK_WORLD_17 = 0xb1;
	SDLK_WORLD_18 = 0xb2;
	SDLK_WORLD_19 = 0xb3;
	SDLK_WORLD_20 = 0xb4;
	SDLK_WORLD_21 = 0xb5;
	SDLK_WORLD_22 = 0xb6;
	SDLK_WORLD_23 = 0xb7;
	SDLK_WORLD_24 = 0xb8;
	SDLK_WORLD_25 = 0xb9;
	SDLK_WORLD_26 = 0xba;
	SDLK_WORLD_27 = 0xbb;
	SDLK_WORLD_28 = 0xbc;
	SDLK_WORLD_29 = 0xbd;
	SDLK_WORLD_30 = 0xbe;
	SDLK_WORLD_31 = 0xbf;
	SDLK_WORLD_32 = 0xc0;
	SDLK_WORLD_33 = 0xc1;
	SDLK_WORLD_34 = 0xc2;
	SDLK_WORLD_35 = 0xc3;
	SDLK_WORLD_36 = 0xc4;
	SDLK_WORLD_37 = 0xc5;
	SDLK_WORLD_38 = 0xc6;
	SDLK_WORLD_39 = 0xc7;
	SDLK_WORLD_40 = 0xc8;
	SDLK_WORLD_41 = 0xc9;
	SDLK_WORLD_42 = 0xca;
	SDLK_WORLD_43 = 0xcb;
	SDLK_WORLD_44 = 0xcc;
	SDLK_WORLD_45 = 0xcd;
	SDLK_WORLD_46 = 0xce;
	SDLK_WORLD_47 = 0xcf;
	SDLK_WORLD_48 = 0xd0;
	SDLK_WORLD_49 = 0xd1;
	SDLK_WORLD_50 = 0xd2;
	SDLK_WORLD_51 = 0xd3;
	SDLK_WORLD_52 = 0xd4;
	SDLK_WORLD_53 = 0xd5;
	SDLK_WORLD_54 = 0xd6;
	SDLK_WORLD_55 = 0xd7;
	SDLK_WORLD_56 = 0xd8;
	SDLK_WORLD_57 = 0xd9;
	SDLK_WORLD_58 = 0xda;
	SDLK_WORLD_59 = 0xdb;
	SDLK_WORLD_60 = 0xdc;
	SDLK_WORLD_61 = 0xdd;
	SDLK_WORLD_62 = 0xde;
	SDLK_WORLD_63 = 0xdf;
	SDLK_WORLD_64 = 0xe0;
	SDLK_WORLD_65 = 0xe1;
	SDLK_WORLD_66 = 0xe2;
	SDLK_WORLD_67 = 0xe3;
	SDLK_WORLD_68 = 0xe4;
	SDLK_WORLD_69 = 0xe5;
	SDLK_WORLD_70 = 0xe6;
	SDLK_WORLD_71 = 0xe7;
	SDLK_WORLD_72 = 0xe8;
	SDLK_WORLD_73 = 0xe9;
	SDLK_WORLD_74 = 0xea;
	SDLK_WORLD_75 = 0xeb;
	SDLK_WORLD_76 = 0xec;
	SDLK_WORLD_77 = 0xed;
	SDLK_WORLD_78 = 0xee;
	SDLK_WORLD_79 = 0xef;
	SDLK_WORLD_80 = 0xf0;
	SDLK_WORLD_81 = 0xf1;
	SDLK_WORLD_82 = 0xf2;
	SDLK_WORLD_83 = 0xf3;
	SDLK_WORLD_84 = 0xf4;
	SDLK_WORLD_85 = 0xf5;
	SDLK_WORLD_86 = 0xf6;
	SDLK_WORLD_87 = 0xf7;
	SDLK_WORLD_88 = 0xf8;
	SDLK_WORLD_89 = 0xf9;
	SDLK_WORLD_90 = 0xfa;
	SDLK_WORLD_91 = 0xfb;
	SDLK_WORLD_92 = 0xfc;
	SDLK_WORLD_93 = 0xfd;
	SDLK_WORLD_94 = 0xfe;
	SDLK_WORLD_95 = 0xff;
	SDLK_KP0 = 0x100;
	SDLK_KP1 = 0x101;
	SDLK_KP2 = 0x102;
	SDLK_KP3 = 0x103;
	SDLK_KP4 = 0x104;
	SDLK_KP5 = 0x105;
	SDLK_KP6 = 0x106;
	SDLK_KP7 = 0x107;
	SDLK_KP8 = 0x108;
	SDLK_KP9 = 0x109;
	SDLK_KP_PERIOD = 0x10a;
	SDLK_KP_DIVIDE = 0x10b;
	SDLK_KP_MULTIPLY = 0x10c;
	SDLK_KP_MINUS = 0x10d;
	SDLK_KP_PLUS = 0x10e;
	SDLK_KP_ENTER = 0x10f;
	SDLK_KP_EQUALS = 0x110;
	SDLK_UP = 0x111;
	SDLK_DOWN = 0x112;
	SDLK_RIGHT = 0x113;
	SDLK_LEFT = 0x114;
	SDLK_INSERT = 0x115;
	SDLK_HOME = 0x116;
	SDLK_END = 0x117;
	SDLK_PAGEUP = 0x118;
	SDLK_PAGEDOWN = 0x119;
	SDLK_F1 = 0x11a;
	SDLK_F2 = 0x11b;
	SDLK_F3 = 0x11c;
	SDLK_F4 = 0x11d;
	SDLK_F5 = 0x11e;
	SDLK_F6 = 0x11f;
	SDLK_F7 = 0x120;
	SDLK_F8 = 0x121;
	SDLK_F9 = 0x122;
	SDLK_F10 = 0x123;
	SDLK_F11 = 0x124;
	SDLK_F12 = 0x125;
	SDLK_F13 = 0x126;
	SDLK_F14 = 0x127;
	SDLK_F15 = 0x128;
	SDLK_NUMLOCK = 0x12c;
	SDLK_CAPSLOCK = 0x12d;
	SDLK_SCROLLOCK = 0x12e;
	SDLK_RSHIFT = 0x12f;
	SDLK_LSHIFT = 0x130;
	SDLK_RCTRL = 0x131;
	SDLK_LCTRL = 0x132;
	SDLK_RALT = 0x133;
	SDLK_LALT = 0x134;
	SDLK_RMETA = 0x135;
	SDLK_LMETA = 0x136;
	SDLK_LSUPER = 0x137;
	SDLK_RSUPER = 0x138;
	SDLK_MODE = 0x139;
	SDLK_COMPOSE = 0x13a;
	SDLK_HELP = 0x13b;
	SDLK_PRINT = 0x13c;
	SDLK_SYSREQ = 0x13d;
	SDLK_BREAK = 0x13e;
	SDLK_MENU = 0x13f;
	SDLK_POWER = 0x140;
	SDLK_EURO = 0x141;
	SDLK_UNDO = 0x142;
	KMOD_NONE = 0;
	KMOD_LSHIFT = 0x1;
	KMOD_RSHIFT = 0x2;
	KMOD_LCTRL = 0x40;
	KMOD_RCTRL = 0x80;
	KMOD_LALT = 0x100;
	KMOD_RALT = 0x200;
	KMOD_LMETA = 0x400;
	KMOD_RMETA = 0x800;
	KMOD_NUM = 0x1000;
	KMOD_CAPS = 0x2000;
	KMOD_MODE = 0x4000;
	KMOD_RESERVED = 0x8000;
)

// Types

type Surface struct {
	Flags uint32;
	Format *PixelFormat;
	W int32;
	H int32;
	Pitch uint16;
	Pad0 [2]byte;
	Pixels *byte;
	Offset int32;
	Hwdata *[0]byte /* sprivate_hwdata */;
	Clip_rect Rect;
	Unused1 uint32;
	Locked uint32;
	Map *[0]byte /* sSDL_BlitMap */;
	Format_version uint32;
	Refcount int32;
}

type PixelFormat struct {
	Palette *Palette;
	BitsPerPixel uint8;
	BytesPerPixel uint8;
	Rloss uint8;
	Gloss uint8;
	Bloss uint8;
	Aloss uint8;
	Rshift uint8;
	Gshift uint8;
	Bshift uint8;
	Ashift uint8;
	Pad0 [2]byte;
	Rmask uint32;
	Gmask uint32;
	Bmask uint32;
	Amask uint32;
	Colorkey uint32;
	Alpha uint8;
	Pad1 [3]byte;
}

type Rect struct {
	X int16;
	Y int16;
	W uint16;
	H uint16;
}

type Color struct {
	R uint8;
	G uint8;
	B uint8;
	Unused uint8;
}

type Palette struct {
	Ncolors int32;
	Colors *Color;
}

type VideoInfo struct {
	Pad0 [2]byte;
	UnusedBits3 uint16;
	Video_mem uint32;
	Vfmt *PixelFormat;
	Current_w int32;
	Current_h int32;
}

type Overlay struct {
	Format uint32;
	W int32;
	H int32;
	Planes int32;
	Pitches *uint16;
	Pixels **uint8;
	Hwfuncs *[0]byte /* sprivate_yuvhwfuncs */;
	Hwdata *[0]byte /* sprivate_yuvhwdata */;
	Pad0 [4]byte;
}

type ActiveEvent struct {
	Type uint8;
	Gain uint8;
	State uint8;
}

type KeyboardEvent struct {
	Type uint8;
	Which uint8;
	State uint8;
	Pad0 [1]byte;
	Keysym Keysym;
}

type MouseMotionEvent struct {
	Type uint8;
	Which uint8;
	State uint8;
	Pad0 [1]byte;
	X uint16;
	Y uint16;
	Xrel int16;
	Yrel int16;
}

type MouseButtonEvent struct {
	Type uint8;
	Which uint8;
	Button uint8;
	State uint8;
	X uint16;
	Y uint16;
}

type JoyAxisEvent struct {
	Type uint8;
	Which uint8;
	Axis uint8;
	Pad0 [1]byte;
	Value int16;
}

type JoyBallEvent struct {
	Type uint8;
	Which uint8;
	Ball uint8;
	Pad0 [1]byte;
	Xrel int16;
	Yrel int16;
}

type JoyHatEvent struct {
	Type uint8;
	Which uint8;
	Hat uint8;
	Value uint8;
}

type JoyButtonEvent struct {
	Type uint8;
	Which uint8;
	Button uint8;
	State uint8;
}

type ResizeEvent struct {
	Type uint8;
	Pad0 [3]byte;
	W int32;
	H int32;
}

type ExposeEvent struct {
	Type uint8;
}

type QuitEvent struct {
	Type uint8;
}

type UserEvent struct {
	Type uint8;
	Pad0 [3]byte;
	Code int32;
	Data1 *byte;
	Data2 *byte;
}

type SysWMmsg struct {
}

type SysWMEvent struct {
	Type uint8;
	Pad0 [3]byte;
	Msg *SysWMmsg;
}

type Event struct {
	Type uint8;
	Pad0 [19]byte;
}

type Keysym struct {
	Scancode uint8;
	Sym int;
	Mod byte;
	Unicode uint16;
	Pad1 [2]byte;
}
