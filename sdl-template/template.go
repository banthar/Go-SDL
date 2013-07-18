package main

import (
	"fmt"
	"github.com/neagix/Go-SDL/sdl"
	"math/rand"
	"time"
)

func loadImage(name string) *sdl.Surface {
	image := sdl.Load(name)

	if image == nil {
		panic(sdl.GetError())
	}

	return image

}

func main() {
	if sdl.Init(sdl.INIT_VIDEO) != 0 {
		panic(sdl.GetError())
	}

	defer sdl.Quit()

	screen := sdl.SetVideoMode(400, 300, 32, 0)
	if screen == nil {
		panic(sdl.GetError())
	}

	sdl.WM_SetCaption("Template", "")

	ticker := time.NewTicker(1e9 / 2 /*2 Hz*/ )

loop:
	for {
		select {
		case <-ticker.C:
			// Note: For better efficiency, use UpdateRects instead of Flip
			screen.FillRect(nil, /*color*/ rand.Uint32())
			//screen.Blit(&sdl.Rect{x,y, 0, 0}, image, nil)
			screen.Flip()

		case event := <-sdl.Events:
			fmt.Printf("%#v\n", event)

			switch event.(type) {
			case sdl.QuitEvent:
				break loop
			}
		}
	}
}
