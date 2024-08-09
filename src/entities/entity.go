package entities

import "github.com/veandco/go-sdl2/sdl"

type Entity struct {
	Position sdl.FPoint
	Size     sdl.FPoint
	Sprite   *sdl.Texture
}

type EntityInterface interface {
	Update(*[]uint8)
	Draw(*sdl.Renderer)
}
