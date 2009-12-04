package main

import (
	"sdl";
	"sdl/ttf";
	"sdl/mixer";
	"math";
)

type Point struct {
	x	int;
	y	int;
}

func (a Point) add(b Point) Point	{ return Point{a.x + b.x, a.y + b.y} }

func (a Point) sub(b Point) Point	{ return Point{a.x - b.x, a.y - b.y} }

func (a Point) length() float64	{ return math.Sqrt(float64(a.x*a.x + a.y*a.y)) }

func (a Point) mul(b float64) Point {
	return Point{int(float64(a.x) * b), int(float64(a.y) * b)}
}

func worm(in <-chan Point, out chan<- Point, draw chan<- Point) {

	t := Point{0, 0};

	for {

		p := (<-in).sub(t);

		if p.length() > 48 {
			t = t.add(p.mul(0.1))
		}

		draw <- t;
		out <- t;
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

	var screen = sdl.SetVideoMode(640, 480, 32, 0);

	if screen == nil {
		panic(sdl.GetError())
	}

	sdl.WM_SetCaption("Go-SDL SDL Test", "");

	image := sdl.Load("test.png");

	if image == nil {
		panic(sdl.GetError())
	}

	running := true;

	font := ttf.OpenFont("Fontin Sans.otf", 72);

	if font == nil {
		panic(sdl.GetError())
	}

	font.SetFontStyle(ttf.STYLE_UNDERLINE);
	white := sdl.Color{255, 255, 255, 0};
	text := ttf.RenderText_Blended(font, "Test (with music)", white);
	music := mixer.LoadMUS("test.ogg");

	if music == nil {
		panic(sdl.GetError())
	}

	music.PlayMusic(-1);

	if sdl.GetKeyName(270) != "[+]" {
		panic("GetKeyName broken")
	}

	worm_in := make(chan Point);
	draw := make(chan Point, 64);

	var out chan Point;
	var in chan Point;

	out = worm_in;

	in = out;
	out = make(chan Point);
	go worm(in, out, draw);

	for running {

		e := &sdl.Event{};

		for e.Poll() {
			switch e.Type {
			case sdl.QUIT:
				running = false;
				break;
			case sdl.KEYDOWN:
				println(e.Keyboard().Keysym.Sym, ": ", sdl.GetKeyName(sdl.Key(e.Keyboard().Keysym.Sym)))
			case sdl.MOUSEBUTTONDOWN:
				println("Click:", e.MouseButton().X, e.MouseButton().Y);
				in = out;
				out = make(chan Point);
				go worm(in, out, draw);
			default:
			}
		}

		screen.FillRect(nil, 0x302019);
		screen.Blit(&sdl.Rect{0, 0, 0, 0}, text, nil);

		loop := true;

		for loop {

			select {
			case p := <-draw:
				screen.Blit(&sdl.Rect{int16(p.x), int16(p.y), 0, 0}, image, nil)
			case <-out:
			default:
				loop = false
			}

		}

		var p Point;
		sdl.GetMouseState(&p.x, &p.y);
		worm_in <- p;

		screen.Flip();
		sdl.Delay(25);
	}

	image.Free();
	music.Free();
	font.Close();

	ttf.Quit();
	sdl.Quit();
}
