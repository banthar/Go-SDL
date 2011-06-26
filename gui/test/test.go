package main

import(
	"os"
	"fmt"
	"sdl"
	"image"
	"exp/gui"
	"image/draw"
	"exp/gui/sdl"
)

func main() {
	win, err := sdlgui.NewWindow(320, 240, 32, sdl.DOUBLEBUF)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//img := image.NewColorImage(image.RGBAColor{255, 0, 255, 255})
	img := sdlgui.SurfaceToImage(sdl.Load("test.png"))
	draw.Draw(win.Screen(),
		image.Rect(10, 10, 20, 20),
		img,
		image.Pt(0, 0),
		draw.Over,
	)
	win.FlushImage()

	for ev := range(win.EventChan()) {
		switch e := ev.(type) {
			case gui.KeyEvent:
				fmt.Println(e.Key)
				if e.Key == 'q' {
					win.Close()
				}
		}
	}
}
