package main

import (
	resolv "github.com/SolarLune/resolv/resolv"
	"github.com/veandco/go-sdl2/sdl"
)

type keyboardMover struct {
	container *element
}

func newKeyboardMover(container *element) *keyboardMover {
	return &keyboardMover{
		container: container,
	}
}

func (mover *keyboardMover) onUpdate() error {

	solids := space.FilterByTags("solid")
	ramps := space.FilterByTags("ramp")

	p := mover.container

	p.speedy += 0.4

	friction := float32(0.5)
	accel := 0.4 + friction

	maxSpd := float32(3)

	if p.speedx > friction {
		p.speedx -= friction
	} else if p.speedx < -friction {
		p.speedx += friction
	} else {
		p.speedx = 0
	}

	if keyboard.KeyDown(sdl.K_RIGHT) {
		p.speedx += accel
	}

	if keyboard.KeyDown(sdl.K_LEFT) {
		p.speedx -= accel
	}

	if p.speedx > maxSpd {
		p.speedx = maxSpd
	}

	if p.speedx < -maxSpd {
		p.speedx = -maxSpd
	}

	// JUMP

	// Check for a collision downwards by just attempting a resolution downwards and seeing if it collides with something.
	down := space.Resolve(p.rect, 0, 4)
	onGround := down.Colliding()

	if keyboard.KeyPressed(sdl.K_SPACE) && onGround {
		p.speedy = -8
	}

	x := int32(p.speedx)
	y := int32(p.speedy)

	// X-movement. We only want to collide with solid objects (not ramps) because we want to be able to move up them
	// and don't need to be inhibited on the x-axis when doing so.
	if res := solids.Resolve(p.rect, x, 0); res.Colliding() {
		x = res.ResolveX
		p.speedx = 0
	}

	p.rect.X += x

	// Y movement. We check for ramp collision first; if we find it, then we just automatically will
	// slide up the ramp because the player is moving into it.

	// We look for ramps a little aggressively downwards because when walking down them, we want to stick to them.
	// If we didn't do this, then you would "bob" when walking down the ramp as the Player moves too quickly out into
	// space for gravity to push back down onto the ramp.
	res := ramps.Resolve(p.rect, 0, y+4)

	if y < 0 || (res.Teleporting && res.ResolveY < -p.rect.H/2) {
		res = resolv.Collision{}
	}

	if !res.Colliding() {
		res = solids.Resolve(p.rect, 0, y)
	}

	if res.Colliding() {
		y = res.ResolveY
		p.speedy = 0
	}

	p.rect.Y += y

	return nil
}

func (mover *keyboardMover) onDraw(renderer *sdl.Renderer) error {
	return nil
}
