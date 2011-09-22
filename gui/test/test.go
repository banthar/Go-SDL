package main

import (
	"os"
	"fmt"
	"sdl"
	"image"
	"exp/gui"
	"image/draw"
	"exp/gui/sdl"
)

func main() {
	win, err := sdlgui.NewWindow(320, 240, 32, sdl.DOUBLEBUF|sdl.RESIZABLE)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	screen := win.Screen()

	cimg := image.NewColorImage(image.RGBAColor{255, 0, 255, 255})
	draw.Draw(screen,
		image.Rect(10, 10, 200, 200),
		cimg,
		image.ZP,
		draw.Over,
	)

	timg := sdl.Load("test.png")
	draw.Draw(screen,
		image.Rect(50, 50, 100, 100),
		timg,
		image.ZP,
		draw.Over,
	)

	win.FlushImage()

	for ev := range win.EventChan() {
		switch e := ev.(type) {
		case gui.ConfigEvent:
			fmt.Printf("Window size: (%v, %v)\n", e.Config.Width, e.Config.Height)
		case gui.KeyEvent:
			fmt.Printf("Key pressed: %v\n", e.Key)
			if e.Key == sdl.K_ESCAPE {
				win.Close()
			}
		case gui.MouseEvent:
			if e.Buttons&(1<<0) > 0 {
				fmt.Println("The left mouse button has been pressed.")
			}
			if e.Buttons&(1<<1) > 0 {
				fmt.Println("The middle mouse button has been pressed.")
			}
			if e.Buttons&(1<<2) > 0 {
				fmt.Println("The right mouse button has been pressed.")
			}
		}
	}
}
