package main

import (
	"sdl"
)

func loadImage(name string) *sdl.Surface {
	image := sdl.Load(name)

	if image == nil {
		panic(sdl.GetError())
	}

	return image

}

func main() {

	if sdl.Init(sdl.INIT_EVERYTHING) != 0 {
		panic(sdl.GetError())
	}

	var screen = sdl.SetVideoMode(640, 480, 32, 0)

	if screen == nil {
		panic(sdl.GetError())
	}

	sdl.WM_SetCaption("Template", "")

	running := true

	for running {

		e := &sdl.Event{}

		for e.Poll() {
			switch e.Type {
			case sdl.QUIT:
				running = false
				break
			default:
			}
		}

		screen.FillRect(nil, 0x000000)

		//screen.Blit(&sdl.Rect{x,y, 0, 0}, image, nil)

		screen.Flip()
		sdl.Delay(25)

	}

	sdl.Quit()

}
