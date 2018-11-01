package main

import (
	"fmt"

	resolv "github.com/SolarLune/resolv/resolv"
	tiled "github.com/lafriks/go-tiled"
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

	tiles := createShapesFromTilemap(*tilemap)
	addTileShapes(tiles)

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

		drawTiles(tiles)

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

type alttile struct {
	id      int32
	srcRect *sdl.Rect
	dstRect *sdl.Rect
	shape   *resolv.Rectangle
}

func createShapesFromTilemap(tilemap tiled.Map) []alttile {

	tiles := make([]alttile, 0)

	tileW := int32(tilemap.TileWidth)
	tileH := int32(tilemap.TileHeight)

	tilemapImageW := int32(80)

	for _, layer := range tilemap.Layers {

		for tileIndex, tile := range layer.Tiles {

			id := int32(tile.ID)
			tileDestX := int32(tileIndex%tilemap.Width) * tileW
			tileDestY := int32((int32(tileIndex) / int32(tilemap.Width)) * tileH)

			var tileSrcX = id % (tilemapImageW / tileW)
			var tileSrcY = int32(id / (tilemapImageW / tileW))

			srcRect := sdl.Rect{X: tileSrcX * tileW, Y: tileSrcY * tileH, W: tileW, H: tileH}
			dstRect := sdl.Rect{X: tileDestX, Y: tileDestY, W: tileW, H: tileH}

			if id == 0 {
				shape := resolv.NewRectangle(tileDestX, tileDestY, 20, 20)
				shape.SetTags("solid")
				t := alttile{id: id, srcRect: &srcRect, dstRect: &dstRect, shape: shape}
				tiles = append(tiles, t)
			} else if id == 4 {
				shape := resolv.NewRectangle(tileDestX, tileDestY, 20, 20)
				shape.SetTags("solid")
				t := alttile{id: id, srcRect: &srcRect, dstRect: &dstRect, shape: shape}
				tiles = append(tiles, t)
			} else {
				t := alttile{id: id, srcRect: &srcRect, dstRect: &dstRect}
				tiles = append(tiles, t)
			}

		}
	}

	return tiles
}

func drawTiles(tiles []alttile) {
	tilemapImage, _ := img.LoadTexture(renderer, "assets/tileset.png")
	defer tilemapImage.Destroy()

	for _, tile := range tiles {
		renderer.Copy(tilemapImage, tile.srcRect, tile.dstRect)
	}
}

func addTileShapes(tiles []alttile) {
	for _, tile := range tiles {
		if tile.shape != nil {
			space.AddShape(tile.shape)
		}
	}
}
