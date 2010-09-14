package main

import (
	"⚛sdl"
	"⚛sdl/ttf"
	"⚛sdl/mixer"
	"math"
	"fmt"
	"time"
)

type Point struct {
	x int
	y int
}

func (a Point) add(b Point) Point { return Point{a.x + b.x, a.y + b.y} }

func (a Point) sub(b Point) Point { return Point{a.x - b.x, a.y - b.y} }

func (a Point) length() float64 { return math.Sqrt(float64(a.x*a.x + a.y*a.y)) }

func (a Point) mul(b float64) Point {
	return Point{int(float64(a.x) * b), int(float64(a.y) * b)}
}

func worm(in <-chan Point, out chan<- Point, draw chan<- Point) {

	t := Point{0, 0}

	for {
		p := (<-in).sub(t)

		if p.length() > 48 {
			t = t.add(p.mul(0.1))
		}

		draw <- t
		out <- t
	}
}

func main() {

	if sdl.Init(sdl.INIT_EVERYTHING) != 0 {
		panic(sdl.GetError())
	}

	if ttf.Init() != 0 {
		panic(sdl.GetError())
	}

	if mixer.OpenAudio(mixer.DEFAULT_FREQUENCY, mixer.DEFAULT_FORMAT,
		mixer.DEFAULT_CHANNELS, 4096) != 0 {
		panic(sdl.GetError())
	}

	var screen = sdl.SetVideoMode(640, 480, 32, sdl.RESIZABLE)

	if screen == nil {
		panic(sdl.GetError())
	}

	var video_info = sdl.GetVideoInfo()

	println("HW_available = ", video_info.HW_available)
	println("WM_available = ", video_info.WM_available)
	println("Video_mem = ", video_info.Video_mem, "kb")

	sdl.EnableUNICODE(1)

	sdl.WM_SetCaption("Go-SDL SDL Test", "")

	image := sdl.Load("test.png")

	if image == nil {
		panic(sdl.GetError())
	}

	sdl.WM_SetIcon(image, nil)

	running := true

	font := ttf.OpenFont("Fontin Sans.otf", 72)

	if font == nil {
		panic(sdl.GetError())
	}

	font.SetStyle(ttf.STYLE_UNDERLINE)
	white := sdl.Color{255, 255, 255, 0}
	text := ttf.RenderText_Blended(font, "Test (with music)", white)
	music := mixer.LoadMUS("test.ogg")

	if music == nil {
		panic(sdl.GetError())
	}

	music.PlayMusic(-1)

	if sdl.GetKeyName(270) != "[+]" {
		panic("GetKeyName broken")
	}

	worm_in := make(chan Point)
	draw := make(chan Point, 64)

	var out chan Point
	var in chan Point

	out = worm_in

	in = out
	out = make(chan Point)
	go worm(in, out, draw)

	ticker := time.NewTicker(1e9/50 /*50Hz*/)

	// Note: The following SDL code is highly ineffective.
	//       It is eating too much CPU. If you intend to use Go-SDL,
	//       you should to do better than this.

	for running {
		select {
		case <-ticker.C:
			screen.FillRect(nil, 0x302019)
			screen.Blit(&sdl.Rect{0, 0, 0, 0}, text, nil)

			loop: for {
				select {
				case p := <-draw:
					screen.Blit(&sdl.Rect{int16(p.x), int16(p.y), 0, 0}, image, nil)

				case <-out:
				default:
					break loop
				}
			}

			var p Point
			sdl.GetMouseState(&p.x, &p.y)
			worm_in <- p

			screen.Flip()

		case _event := <-sdl.Events:
			switch e := _event.(type) {
			case sdl.QuitEvent:
				running = false

			case sdl.KeyboardEvent:
				println("")
				println(e.Keysym.Sym, ": ", sdl.GetKeyName(sdl.Key(e.Keysym.Sym)))

				if e.Keyboard().Keysym.Sym == 27 {
					running = false
				}

				fmt.Printf("%04x ", e.Type)

				for i := 0; i < len(e.Pad0); i++ {
					fmt.Printf("%02x ", e.Pad0[i])
				}
				println()

				fmt.Printf("Type: %02x Which: %02x State: %02x Pad: %02x\n", e.Type, e.Which, e.State, e.Pad0[0])
				fmt.Printf("Scancode: %02x Sym: %08x Mod: %04x Unicode: %04x\n", e.Keysym.Scancode, e.Keysym.Sym, e.Keysym.Mod, e.Keysym.Unicode)
			case sdl.MouseButtonEvent:
				if e.Type == sdl.MOUSEBUTTONDOWN {
					println("Click:", e.X, e.Y)
					in = out
					out = make(chan Point)
					go worm(in, out, draw)
				}

			case sdl.ResizeEvent:
				println("resize screen ", e.W, e.H)
				fmt.Printf("Type: %02x Which: %02x State: %02x Pad: %02x\n", k.Type, k.Which, k.State, k.Pad0[0])
				fmt.Printf("Scancode: %02x Sym: %08x Mod: %04x Unicode: %04x\n", k.Keysym.Scancode, k.Keysym.Sym, k.Keysym.Mod, k.Keysym.Unicode)
			case sdl.MOUSEBUTTONDOWN:
				println("Click:", e.MouseButton().X, e.MouseButton().Y)
				in = out
				out = make(chan Point)
				go worm(in, out, draw)
			case sdl.VIDEORESIZE:
				println("resize screen ", e.Resize().W, e.Resize().H)

				screen = sdl.SetVideoMode(int(e.W), int(e.H), 32, sdl.RESIZABLE)

				if screen == nil {
					panic(sdl.GetError())
				}
			}
		}
	}

	image.Free()
	music.Free()
	font.Close()

	ttf.Quit()
	sdl.Quit()
}
