package main

import "sdl"
//import "math"

import "fmt"
import "image/png"
import "os"

func main() {
	sdl.Init(sdl.INIT_VIDEO);
	sdl.TTF_Init();

	var screen = sdl.SetVideoMode(640, 480, 32, 0);
	sdl.WM_SetCaption("Go-SDL SDL Test", "");

	image := sdl.Load("test.png");
	running := true;
	var x, y int16;
	font := sdl.TTF_OpenFont("CloisterBlack.ttf", 72);
	white := sdl.Color{255,255,255,0};
	text := sdl.TTF_RenderText_Blended(font, "Test", white);

	for running {

		x++;
		y++;

		e := &sdl.Event{};

		for e.Poll() {
			switch e.Type {
			case sdl.QUIT:
				file, err := os.Open("shoot.png", os.O_CREATE|os.O_WRONLY, 0766);
				println(err);
				png.Encode(file, screen);
				file.Close();
				running = false;
				break;
			case sdl.KEYDOWN:
				println(e.Keyboard().Keysym.Sym)
			case sdl.MOUSEBUTTONDOWN:
				println("Click:", e.MouseButton().X, e.MouseButton().Y);
				x = int16(e.MouseButton().X);
				y = int16(e.MouseButton().Y);
			default:
				fmt.Printf("Event: %08x\n", e.Type)
			}
		}

		screen.FillRect(nil, 0x302019);
		screen.Blit(&sdl.Rect{x, y, 0, 0}, image, nil);
		screen.Blit(&sdl.Rect{0,0,0,0}, text, nil);
		screen.Flip();
		sdl.Delay(25);
	}

	image.Free();
	screen.Free();
	sdl.TTF_CloseFont(font);

    sdl.TTF_Quit();
	sdl.Quit();

}
