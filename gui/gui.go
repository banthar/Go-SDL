package sdlgui

import (
	"os"
	"sdl"
	"time"
	"image"
	"runtime"
	"exp/gui"
	"image/draw"
)

type window struct {
	screen *surfimg

	ec     chan interface{}
	events bool
}

func (win *window) eventLoop() {
	if win.ec == nil {
		win.ec = make(chan interface{})
	}

eloop:
	for win.events {
		var ev sdl.Event
		for ev.Poll() {
			switch ev.Type {
			case sdl.KEYUP:
				key := ev.Keyboard().Keysym.Sym
				win.ec <- gui.KeyEvent{int(-key)}
			case sdl.KEYDOWN:
				key := ev.Keyboard().Keysym.Sym
				win.ec <- gui.KeyEvent{int(key)}
			case sdl.MOUSEMOTION:
				m := ev.MouseMotion()
				win.ec <- gui.MouseEvent{
					Buttons: int(m.State),
					Loc:     image.Pt(int(m.X), int(m.Y)),
					Nsec:    time.Nanoseconds(),
				}
			case sdl.MOUSEBUTTONUP, sdl.MOUSEBUTTONDOWN:
				m := ev.MouseButton()
				win.ec <- gui.MouseEvent{
					Buttons: int(sdl.GetMouseState(nil, nil)),
					Loc:     image.Pt(int(m.X), int(m.Y)),
					Nsec:    time.Nanoseconds(),
				}
			case sdl.VIDEORESIZE:
				r := ev.Resize()
				win.ec <- gui.ConfigEvent{image.Config{
					win.Screen().ColorModel(),
					int(r.W),
					int(r.H),
				}}
			case sdl.QUIT:
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

func (win *window) Close() os.Error {
	win.events = false

	win.screen.Free()

	initdec()

	return nil
}

var initnum uint

func initinc() os.Error {
	if initnum == 0 {
		errn := sdl.Init(sdl.INIT_VIDEO)
		if errn < 0 {
			return os.NewError(sdl.GetError())
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

func NewWindow(w, h, bpp int, flags uint32) (gui.Window, os.Error) {
	win := new(window)

	err := initinc()
	if err != nil {
		return nil, err
	}

	win.screen = &surfimg{sdl.SetVideoMode(w, h, bpp, flags)}
	if win.screen == nil {
		return nil, os.NewError(sdl.GetError())
	}

	win.ec = make(chan interface{})
	win.events = true
	go win.eventLoop()

	runtime.SetFinalizer(win, func(subwin *window) {
		subwin.Close()
	})

	return win, nil
}
