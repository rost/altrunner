package main

import (
	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	width  = 600
	height = 800
)

var window *sdl.Window
var renderer *sdl.Renderer
var fpsManager = &gfx.FPSmanager{}
var debugMode = true

func main() {

	gfx.InitFramerate(fpsManager)
	gfx.SetFramerate(fpsManager, 60)

	sdl.Init(sdl.INIT_EVERYTHING)
	defer sdl.Quit()

	window, renderer, _ = sdl.CreateWindowAndRenderer(width, height, sdl.WINDOW_RESIZABLE)
	defer window.Destroy()
	
	renderer.SetLogicalSize(width, height)

	running := true

	for running {
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