package main

import "sdl"
import "fmt"

func listModes(flags uint32) {

	modes := sdl.ListModes(nil, flags)

	if modes == nil {
		fmt.Printf("\tany mode\n")
		return
	}

	for _, mode := range modes {
		fmt.Printf("\t%dx%d\n", mode.W, mode.H)
	}
}

func main() {

	sdl.Init(sdl.INIT_VIDEO)

	fmt.Printf("fullscreen modes:\n")
	listModes(sdl.FULLSCREEN)

}
