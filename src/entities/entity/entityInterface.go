package entity

import "github.com/veandco/go-sdl2/sdl"

type EntityInterface interface {
	Update(*[]uint8)
	Draw(*sdl.Renderer)
	GetLayer() uint8
	SetLayer(uint8)
}
