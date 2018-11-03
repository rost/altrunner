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

	renderer.Copy(tr.tex, tile.srcRect, tile.dstRect)

	if debugMode == true {
		if tile.rect != nil {
			renderer.SetDrawColor(255, 255, 255, 255)
			renderer.DrawRect(tile.dstRect)
		}
	}

	return nil
}
