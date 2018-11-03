package main

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type playerRenderer struct {
	container *element
	texid     int32
	tex       *sdl.Texture
}

func newPlayerRenderer(container *element, renderer *sdl.Renderer, filename string) *playerRenderer {
	return &playerRenderer{
		container: container,
		tex:       textureFromTileset(renderer, filename),
	}
}

func (pr *playerRenderer) onUpdate() error {
	return nil
}

func (pr *playerRenderer) onDraw(renderer *sdl.Renderer) error {

	x := pr.container.rect.X
	y := pr.container.rect.Y

	tileSrcX := int32(2)
	tileSrcY := int32(2)

	tileW := int32(20)
	tileH := int32(20)

	tileDestX := x
	tileDestY := y

	srcRect := sdl.Rect{X: tileSrcX * tileW, Y: tileSrcY * tileH, W: tileW, H: tileH}
	dstRect := sdl.Rect{X: tileDestX, Y: tileDestY, W: tileW, H: tileH}

	renderer.Copy(pr.tex, &srcRect, &dstRect)

	// draw debug rectangle
	if debugMode == true {
		rect := pr.container.rect
		renderer.SetDrawColor(0, 128, 255, 255)
		renderer.DrawRect(&sdl.Rect{X: rect.X, Y: rect.Y, W: rect.W, H: rect.H})
	}

	return nil
}

func textureFromTileset(renderer *sdl.Renderer, filename string) *sdl.Texture {
	texture, _ := img.LoadTexture(renderer, filename)
	return texture
}
