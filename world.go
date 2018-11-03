package main

import (
	"math"

	resolv "github.com/SolarLune/resolv/resolv"
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

func (w *World) Create() {

	space.Clear()

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
