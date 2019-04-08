package main

import (
	resolv "github.com/SolarLune/resolv/resolv"
)

func initShapes() {
	platform := &element{}

	c := int32(20)

	platform.FloatingPlatform = resolv.NewLine(c*8, screenHeight-c*7, c*9, screenHeight-c*6)
	platform.FloatingPlatform.AddTags("ramp")
	space.Add(platform.FloatingPlatform)
	platform.FloatingPlatformY = float64(platform.FloatingPlatform.Y)

	sr := newShapeRenderer(platform, renderer)
	platform.addComponent(sr)

	elements = append(elements, platform)
}
