package main

import (
	"fmt"

	resolv "github.com/SolarLune/resolv/resolv"
	tiled "github.com/lafriks/go-tiled"
	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/img"
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

var debugMode = true

func main() {

	sdl.Init(sdl.INIT_EVERYTHING)
	defer sdl.Quit()

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

	tilemap, err := tiled.LoadFromFile("assets/tilemap.tmx")
	if err != nil {
		fmt.Printf("couldn't read tmx tilemap: %v", err)
		return
	}

	imgFlags := img.Init(img.INIT_PNG)
	if imgFlags == 0 {
		fmt.Printf("couldn't init img: %v", err)
		return
	}

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

		world.Update()

		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()
		drawTileMap(renderer, *tilemap)

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

func drawTileMap(renderer *sdl.Renderer, tilemap tiled.Map) {
	tileW := int32(tilemap.TileWidth)
	tileH := int32(tilemap.TileHeight)

	tilemapImageW := int32(80)
	// tilemapImageH := int32(60)

	tilemapImage, _ := img.LoadTexture(renderer, "assets/tileset.png")

	for _, layer := range tilemap.Layers {

		for tileIndex, tile := range layer.Tiles {

			id := int32(tile.ID)
			tileDestX := int32(tileIndex%tilemap.Width) * tileW
			tileDestY := int32((int32(tileIndex) / int32(tilemap.Width)) * tileH)

			var tileSrcX = id % (tilemapImageW / tileW)
			var tileSrcY = int32(id / (tilemapImageW / tileW))

			srcRect := sdl.Rect{X: tileSrcX * tileW, Y: tileSrcY * tileH, W: tileW, H: tileH}
			dstRect := sdl.Rect{X: tileDestX, Y: tileDestY, W: tileW, H: tileH}

			renderer.Copy(tilemapImage, &srcRect, &dstRect)
		}
	}
}

func textureFromBMP(renderer *sdl.Renderer, filename string) *sdl.Texture {
	surface, err := img.Load(filename)
	if err != nil {
		panic(fmt.Errorf("loading %v: %v", filename, err))
	}
	defer surface.Free()

	tex, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		panic(fmt.Errorf("creating texture from: %v: %v", filename, err))
	}
	defer tex.Destroy()

	return tex
}
