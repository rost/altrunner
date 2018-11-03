package main

import (
	"fmt"
	"reflect"

	"github.com/SolarLune/resolv/resolv"
	"github.com/veandco/go-sdl2/sdl"
)

type vector struct {
	x, y int32
}

type component interface {
	onUpdate() error
	onDraw(renderer *sdl.Renderer) error
}

type element struct {
	position          vector
	rect              *resolv.Rectangle
	tex               *sdl.Texture
	tileID            int32
	FloatingPlatform  *resolv.Line
	FloatingPlatformY float64
	srcRect           *sdl.Rect
	dstRect           *sdl.Rect
	speedx            float32
	speedy            float32
	components        []component
}

func (elem *element) update() error {
	for _, comp := range elem.components {
		err := comp.onUpdate()
		if err != nil {
			return err
		}
	}
	return nil
}

func (elem *element) draw(renderer *sdl.Renderer) error {
	for _, comp := range elem.components {
		err := comp.onDraw(renderer)
		if err != nil {
			return err
		}
	}
	return nil
}

func (elem *element) addComponent(new component) {
	for _, existing := range elem.components {
		if reflect.TypeOf(new) == reflect.TypeOf(existing) {
			panic(fmt.Sprintf(
				"attempt to add new component with existing type %v",
				reflect.TypeOf(new)))
		}
	}
	elem.components = append(elem.components, new)
}

func (elem *element) getComponent(withType component) component {
	typeName := reflect.TypeOf(withType)
	for _, comp := range elem.components {
		if reflect.TypeOf(comp) == typeName {
			return comp
		}
	}
	panic(fmt.Sprintf("no component with type %v", reflect.TypeOf(withType)))
}

var elements []*element
