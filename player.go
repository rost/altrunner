package main

import (
	"math/rand"

	resolv "github.com/SolarLune/resolv/resolv"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Player struct {
	Rect   *resolv.Rectangle
	SpeedX float32
	SpeedY float32
}

func MakePlayer() *Player {
	player := &Player{Rect: resolv.NewRectangle(cell*2+rand.Int31n(screenWidth-cell*4), cell*2+rand.Int31n(screenHeight-cell*4), cell, cell),
		SpeedX: (0.5 - rand.Float32()) * 8,
		SpeedY: (0.5 - rand.Float32()) * 8}

	player.Rect.SetTags("player", "solid")

	space.AddShape(player.Rect)

	player.Rect.X = 40
	player.Rect.Y = 40
	player.Rect.W = 20
	player.Rect.H = 20
	player.SpeedX = 0
	player.SpeedY = 0

	return player
}

func (p *Player) Draw(renderer *sdl.Renderer) {
	texture, _ := img.LoadTexture(renderer, "assets/tileset.png")
	defer texture.Destroy()

	tileSrcX := int32(2)
	tileSrcY := int32(2)

	tileW := int32(20)
	tileH := int32(20)

	tileDestX := p.Rect.X
	tileDestY := p.Rect.Y

	srcRect := sdl.Rect{X: tileSrcX * tileW, Y: tileSrcY * tileH, W: tileW, H: tileH}
	dstRect := sdl.Rect{X: tileDestX, Y: tileDestY, W: tileW, H: tileH}

	renderer.Copy(texture, &srcRect, &dstRect)

}
