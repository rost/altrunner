package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type tileRenderer struct {
	container *element
	texid     int32
	tex       *sdl.Texture
}

func newTileRenderer(container *element, renderer *sdl.Renderer, texture *sdl.Texture) *tileRenderer {
	return &tileRenderer{
		container: container,
		tex:       texture,
	}
}

func (tr *tileRenderer) onUpdate() error {
	return nil
}

func (tr *tileRenderer) onDraw(renderer *sdl.Renderer) error {
	tile := tr.container

	tileDstRectX := tile.dstRect.X - camX

	dstRect := sdl.Rect{X: tileDstRectX, Y: tile.dstRect.Y, W: tile.dstRect.W, H: tile.dstRect.H}

	renderer.Copy(tr.tex, tile.srcRect, &dstRect)

	if debugMode == true {
		if tile.rect != nil {
			renderer.SetDrawColor(255, 255, 255, 255)
			renderer.DrawRect(&dstRect)
		}
	}

	return nil
}
