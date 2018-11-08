package main

import (
	"math"

	resolv "github.com/SolarLune/resolv/resolv"
	"github.com/veandco/go-sdl2/sdl"
)

type shapeRenderer struct {
	container *element
	texid     int32
	tex       *sdl.Texture
}

func newShapeRenderer(container *element, renderer *sdl.Renderer) *shapeRenderer {
	return &shapeRenderer{
		container: container,
	}
}

func (sr *shapeRenderer) onUpdate() error {
	shape := sr.container
	shape.FloatingPlatformY += math.Sin(float64(sdl.GetTicks()/1000)) * .5

	shape.FloatingPlatform.Y = int32(shape.FloatingPlatformY)
	shape.FloatingPlatform.Y2 = int32(shape.FloatingPlatformY) - 20

	return nil
}

func (sr *shapeRenderer) onDraw(renderer *sdl.Renderer) error {
	tile := sr.container

	renderer.Copy(sr.tex, tile.srcRect, tile.dstRect)

	if debugMode == true {
		if tile.rect != nil {
			renderer.SetDrawColor(255, 255, 255, 255)
			renderer.DrawRect(tile.dstRect)
		}
	}

	for _, shape := range space {
		line, ok := shape.(*resolv.Line)
		if ok {
			renderer.DrawLine(line.X-camX, line.Y, line.X2-camX, line.Y2)
		}
	}

	return nil
}
