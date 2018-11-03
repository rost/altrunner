package main

import (
	"fmt"

	resolv "github.com/SolarLune/resolv/resolv"
	"github.com/veandco/go-sdl2/gfx"
	img "github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 1280
	screenHeight = 200
)

var cell int32 = 4

var space resolv.Space
var window *sdl.Window
var renderer *sdl.Renderer
var avgFramerate int

var debugMode = true

func main() {

	sdl.Init(sdl.INIT_EVERYTHING)
	defer sdl.Quit()

	imgFlags := img.Init(img.INIT_PNG)
	if imgFlags == 0 {
		fmt.Printf("couldn't init img flags")
		return
	}

	window, renderer, _ = sdl.CreateWindowAndRenderer(screenWidth, screenHeight, sdl.WINDOW_SHOWN|sdl.WINDOW_RESIZABLE)
	defer window.Destroy()

	window.SetResizable(true)

	renderer.SetLogicalSize(screenWidth, screenHeight)

	fpsManager := &gfx.FPSmanager{}
	gfx.InitFramerate(fpsManager)
	gfx.SetFramerate(fpsManager, 60)

	// world

	var world WorldInterface = &World{}

	world.Create()

	running := true

	var frameCount int
	var framerateDelay uint32

	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyboardEvent:
				keyboard.ReportEvent(event.(*sdl.KeyboardEvent))
			}
		}

		keyboard.Update()

		if keyboard.KeyPressed(sdl.K_r) {
			world.Create()
		}

		if keyboard.KeyPressed(sdl.K_q) {
			return
		}

		world.Update()

		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		world.Draw()

		framerateDelay += gfx.FramerateDelay(fpsManager)

		if framerateDelay >= 1000 {
			avgFramerate = frameCount
			framerateDelay -= 1000
			frameCount = 0
			fmt.Println(avgFramerate, " FPS")
		}

		frameCount++

		renderer.Present()
	}
}
