package main

import (
	"fmt"
	"math"

	resolv "github.com/SolarLune/resolv/resolv"
	tiled "github.com/lafriks/go-tiled"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type World struct {
	Player            *Player
	FloatingPlatform  *resolv.Line
	FloatingPlatformY float64
}

type WorldInterface interface {
	Create()
	Update()
	Draw()
	Destroy()
}

type alttile struct {
	id      int32
	srcRect *sdl.Rect
	dstRect *sdl.Rect
	shape   *resolv.Rectangle
}

var tilemap *tiled.Map
var tiles []alttile

func (w *World) Create() {

	var err error

	space.Clear()

	tilemap, err = tiled.LoadFromFile("assets/tilemap.tmx")
	if err != nil {
		fmt.Printf("couldn't read tmx tilemap: %v", err)
		return
	}

	tiles = createShapesFromTilemap(*tilemap)
	addTileShapes(tiles)

	// Player := MakePlayer()
	// w.Player = Player

	w.Player = MakePlayer()

	space.AddShape(w.Player.Rect)

	// space.AddShape(resolv.NewRectangle(0, 0, 20, screenHeight)) // L
	// // space.AddShape(resolv.NewRectangle(screenWidth-20, 0, 20, screenHeight)) // R
	// // space.AddShape(resolv.NewRectangle(0, 0, screenWidth, 20))               // T
	// space.AddShape(resolv.NewRectangle(0, screenHeight-20, screenWidth, 20)) // B

	c := int32(20)

	// space.AddShape(resolv.NewRectangle(c*4, screenHeight-c*4, c*3, c))

	for _, shape := range space {
		shape.SetTags("solid")
	}

	// // A ramp
	// line := resolv.NewLine(c*5, screenHeight-c, c*6, screenHeight-c-8)
	// line.SetTags("ramp")
	// space.AddShape(line)

	// line = resolv.NewLine(c*6, screenHeight-c-8, c*7, screenHeight-c-8)
	// line.SetTags("ramp")

	// space.AddShape(line)

	// rect := resolv.NewRectangle(c*7, screenHeight-c-8, c*2, 8)
	// rect.SetTags("solid")
	// space.AddShape(rect)

	// line = resolv.NewLine(c*9, screenHeight-c-8, c*11, screenHeight-c)
	// line.SetTags("ramp")
	// space.AddShape(line)

	// line = resolv.NewLine(c*13, screenHeight-c*4, c*17, screenHeight-c*6)
	// line.SetTags("ramp")
	// space.AddShape(line)

	// line = resolv.NewLine(c*6, screenHeight-c*7, c*7, screenHeight-c*7)
	// line.SetTags("ramp")
	// space.AddShape(line)

	w.FloatingPlatform = resolv.NewLine(c*8, screenHeight-c*7, c*9, screenHeight-c*6)
	w.FloatingPlatform.SetTags("ramp")
	space.AddShape(w.FloatingPlatform)
	w.FloatingPlatformY = float64(w.FloatingPlatform.Y)

}

func (w *World) Update() {

	w.FloatingPlatformY += math.Sin(float64(sdl.GetTicks()/1000)) * .5

	w.FloatingPlatform.Y = int32(w.FloatingPlatformY)
	w.FloatingPlatform.Y2 = int32(w.FloatingPlatformY) - 20

	solids := space.FilterByTags("solid")
	ramps := space.FilterByTags("ramp")

	w.Player.Update(solids, ramps)

}

func (w *World) Draw() {

	drawTiles(tiles)

	for _, shape := range space {

		rect, ok := shape.(*resolv.Rectangle)

		if ok {

			if rect == w.Player.Rect {
				w.Player.Draw(renderer)
				renderer.SetDrawColor(0, 128, 255, 255)
			} else {
				renderer.SetDrawColor(255, 255, 255, 255)
			}

			renderer.DrawRect(&sdl.Rect{rect.X, rect.Y, rect.W, rect.H})

		}

		line, ok := shape.(*resolv.Line)

		if ok {

			renderer.DrawLine(line.X, line.Y, line.X2, line.Y2)

		}

	}

	// if drawHelpText {
	// 	DrawText(0, 0,
	// 		"Platformer test",
	// 		"Use the arrow keys to move",
	// 		"Press X to jump",
	// 		"You can jump through lines or ramps")
	// }

}

func (w *World) Destroy() {
	space.Clear()
	w.Player = nil
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
