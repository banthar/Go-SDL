package sdlgui

import (
	"errors"
	"exp/gui"
	"github.com/banthar/Go-SDL/sdl"
	"image"
	"image/draw"
	"runtime"
	"time"
)

type window struct {
	screen *sdl.Surface

	ec     chan interface{}
	events bool
}

func (win *window) eventLoop() {
	if win.ec == nil {
		win.ec = make(chan interface{})
	}

eloop:
	for win.events {
		for ev := sdl.PollEvent(); ev != nil; ev = sdl.PollEvent() {
			switch e := ev.(type) {
			case *sdl.KeyboardEvent:
				switch e.Type {
				case sdl.KEYUP:
					win.ec <- gui.KeyEvent{int(-e.Keysym.Sym)}
				case sdl.KEYDOWN:
					win.ec <- gui.KeyEvent{int(e.Keysym.Sym)}
				}
			case *sdl.MouseMotionEvent:
				win.ec <- gui.MouseEvent{
					Buttons: int(e.State),
					Loc:     image.Pt(int(e.X), int(e.Y)),
					Nsec:    time.Nanoseconds(),
				}
			case *sdl.MouseButtonEvent:
				win.ec <- gui.MouseEvent{
					Buttons: int(sdl.GetMouseState(nil, nil)),
					Loc:     image.Pt(int(e.X), int(e.Y)),
					Nsec:    time.Nanoseconds(),
				}
			case *sdl.ResizeEvent:
				win.ec <- gui.ConfigEvent{image.Config{
					win.Screen().ColorModel(),
					int(e.W),
					int(e.H),
				}}
			case *sdl.QuitEvent:
				break eloop
			}
		}
	}

	close(win.ec)
}

func (win *window) Screen() draw.Image {
	return win.screen
}

func (win *window) FlushImage() {
	win.screen.Flip()
}

func (win *window) EventChan() <-chan interface{} {
	return win.ec
}

func (win *window) Close() error {
	win.events = false

	win.screen.Free()

	initdec()

	return nil
}

var initnum uint

func initinc() error {
	if initnum == 0 {
		errn := sdl.Init(sdl.INIT_VIDEO)
		if errn < 0 {
			return errors.New(sdl.GetError())
		}
	}

	initnum++

	return nil
}

func initdec() {
	initnum--

	if initnum == 0 {
		sdl.Quit()
	}
}

func NewWindow(w, h, bpp int, flags uint32) (gui.Window, error) {
	win := new(window)

	err := initinc()
	if err != nil {
		return nil, err
	}

	win.screen = sdl.SetVideoMode(w, h, bpp, flags)
	if win.screen == nil {
		return nil, errors.New(sdl.GetError())
	}

	win.ec = make(chan interface{})
	win.events = true
	go win.eventLoop()

	runtime.SetFinalizer(win, func(subwin *window) {
		subwin.Close()
	})

	return win, nil
}
