package entity

import "github.com/veandco/go-sdl2/sdl"

type Entity struct {
	Position sdl.FPoint
	Size     sdl.FPoint
	layer    uint8
	Sprite   *sdl.Texture
}

func (e *Entity) GetLayer() uint8 {
	return e.layer
}

func (e *Entity) SetLayer(layer uint8) {
	e.layer = layer
}
