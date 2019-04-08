package main

import (
	"fmt"

	resolv "github.com/SolarLune/resolv/resolv"
	tiled "github.com/lafriks/go-tiled"
	"github.com/veandco/go-sdl2/sdl"
)

var tilemap *tiled.Map

func initTiles() {
	var err error

	tilemap, err = tiled.LoadFromFile("assets/tilemap.tmx")
	if err != nil {
		fmt.Printf("couldn't read tmx tilemap: %v", err)
		return
	}

	for _, layer := range tilemap.Layers {
		for tileIndex, layerTile := range layer.Tiles {
			t := newTile(tileIndex, layerTile)
			if t.rect != nil {
				space.Add(t.rect)
			}

			filename := "assets/tileset.png"
			texture := textureFromTileset(renderer, filename)
			tr := newTileRenderer(t, renderer, texture)
			t.addComponent(tr)

			elements = append(elements, t)
		}
	}
}

func newTile(tileIndex int, layerTile *tiled.LayerTile) *element {
	tileW := int32(tilemap.TileWidth)
	tileH := int32(tilemap.TileHeight)

	tilemapImageW := int32(80)

	id := int32(layerTile.ID)
	tileDestX := int32(tileIndex%tilemap.Width) * tileW
	tileDestY := int32((int32(tileIndex) / int32(tilemap.Width)) * tileH)

	var tileSrcX = id % (tilemapImageW / tileW)
	var tileSrcY = int32(id / (tilemapImageW / tileW))

	srcRect := sdl.Rect{X: tileSrcX * tileW, Y: tileSrcY * tileH, W: tileW, H: tileH}
	dstRect := sdl.Rect{X: tileDestX, Y: tileDestY, W: tileW, H: tileH}

	if id == 3 {
		shape := resolv.NewRectangle(tileDestX, tileDestY, 20, 20)
		shape.AddTags("solid")
		t := element{tileID: id, srcRect: &srcRect, dstRect: &dstRect, rect: shape}
		return &t
	} else if id == 4 {
		shape := resolv.NewRectangle(tileDestX, tileDestY, 20, 20)
		shape.AddTags("solid")
		t := element{tileID: id, srcRect: &srcRect, dstRect: &dstRect, rect: shape}
		return &t
	} else {
		t := element{tileID: id, srcRect: &srcRect, dstRect: &dstRect}
		return &t
	}
}
