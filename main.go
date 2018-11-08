package main

import (
	"fmt"

	resolv "github.com/SolarLune/resolv/resolv"
	"github.com/veandco/go-sdl2/gfx"
	img "github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 320
	screenHeight = 200
)

var cell int32 = 4

var space resolv.Space
var window *sdl.Window
var renderer *sdl.Renderer
var avgFramerate int

// camera position
var camX, camY int32

var debugMode = true

func main() {

	camX = 0

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

	initGame()

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
			initGame()
		}

		if keyboard.KeyPressed(sdl.K_q) {
			return
		}

		for _, elem := range elements {
			err := elem.update()
			if err != nil {
				fmt.Println("updating element:", err)
				return
			}
		}

		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		// world.Draw()

		for _, elem := range elements {
			err := elem.draw(renderer)
			if err != nil {
				fmt.Println("drawing element:", err)
				return
			}
		}

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

func initGame() {
	space.Clear()
	camX = 0
	elements = nil
	initTiles()
	initShapes()
	initPlayer()
}
