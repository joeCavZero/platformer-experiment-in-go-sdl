package entity

import "github.com/veandco/go-sdl2/sdl"

type Entity struct {
	Position sdl.FPoint
	Size     sdl.FPoint
	Sprite   *sdl.Texture
}
