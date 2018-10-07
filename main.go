package main

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	width  = 600
	height = 800
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatal("SDL: ", err)
		return
	}

	window, err := sdl.CreateWindow("alt", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		width, height, sdl.WINDOW_OPENGL)
	if err != nil {
		log.Fatal("window:", err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		log.Fatal("renderer:", err)
		return
	}
	defer renderer.Destroy()

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return

			}
		}

		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		renderer.Present()
	}
}
