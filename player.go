package main

import (
	resolv "github.com/SolarLune/resolv/resolv"
	"github.com/veandco/go-sdl2/sdl"
)

type Player struct {
	Rect   *resolv.Rectangle
	SpeedX float32
	SpeedY float32
}

func newPlayer(renderer *sdl.Renderer) *element {
	player := &element{}

	player.position = vector{
		x: 40,
		y: 40,
	}

	player.rect = resolv.NewRectangle(40, 40, 20, 20)
	player.rect.X = 40
	player.rect.Y = 40
	player.rect.W = 20
	player.rect.H = 20
	player.speedx = 0
	player.speedy = 0

	player.rect.SetTags("player", "solid")

	pr := newPlayerRenderer(player, renderer, "assets/tileset.png")
	player.addComponent(pr)

	mover := newKeyboardMover(player)
	player.addComponent(mover)

	return player
}

func initPlayer() {
	player := newPlayer(renderer)
	elements = append(elements, player)
	space.AddShape(player.rect)
}
