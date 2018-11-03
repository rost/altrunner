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

func (p *Player) Update(solids, ramps resolv.Space) {

	p.SpeedY += 0.4

	friction := float32(0.5)
	accel := 0.4 + friction

	maxSpd := float32(3)

	if p.SpeedX > friction {
		p.SpeedX -= friction
	} else if p.SpeedX < -friction {
		p.SpeedX += friction
	} else {
		p.SpeedX = 0
	}

	if keyboard.KeyDown(sdl.K_RIGHT) {
		p.SpeedX += accel
	}

	if keyboard.KeyDown(sdl.K_LEFT) {
		p.SpeedX -= accel
	}

	if p.SpeedX > maxSpd {
		p.SpeedX = maxSpd
	}

	if p.SpeedX < -maxSpd {
		p.SpeedX = -maxSpd
	}

	// JUMP

	// Check for a collision downwards by just attempting a resolution downwards and seeing if it collides with something.
	down := space.Resolve(p.Rect, 0, 4)
	onGround := down.Colliding()

	if keyboard.KeyPressed(sdl.K_SPACE) && onGround {
		p.SpeedY = -8
	}

	x := int32(p.SpeedX)
	y := int32(p.SpeedY)

	// X-movement. We only want to collide with solid objects (not ramps) because we want to be able to move up them
	// and don't need to be inhibited on the x-axis when doing so.

	if res := solids.Resolve(p.Rect, x, 0); res.Colliding() {
		x = res.ResolveX
		p.SpeedX = 0
	}

	p.Rect.X += x

	// Y movement. We check for ramp collision first; if we find it, then we just automatically will
	// slide up the ramp because the player is moving into it.

	// We look for ramps a little aggressively downwards because when walking down them, we want to stick to them.
	// If we didn't do this, then you would "bob" when walking down the ramp as the Player moves too quickly out into
	// space for gravity to push back down onto the ramp.
	res := ramps.Resolve(p.Rect, 0, y+4)

	if y < 0 || (res.Teleporting && res.ResolveY < -p.Rect.H/2) {
		res = resolv.Collision{}
	}

	if !res.Colliding() {
		res = solids.Resolve(p.Rect, 0, y)
	}

	if res.Colliding() {
		y = res.ResolveY
		p.SpeedY = 0
	}

	p.Rect.Y += y
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
