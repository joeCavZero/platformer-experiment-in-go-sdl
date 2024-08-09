package engine

import (
	"project/src/entities"
	"project/src/settings"

	"github.com/veandco/go-sdl2/sdl"
)

type Scene struct {
	entities []entities.EntityInterface
	tilemap  [settings.TILE_QUANTITY]Tile
}

func NewScene() *Scene {
	return &Scene{
		entities: make([]entities.EntityInterface, 0),
		tilemap:  [settings.TILE_QUANTITY]Tile{},
	}
}

func (s *Scene) AddEntity(ent entities.EntityInterface) {
	s.entities = append(s.entities, ent)
}

func (s *Scene) Process(keyboard *[]uint8) {
	for _, ent := range s.entities {
		ent.Update(keyboard)
	}
}

func (s *Scene) Render(renderer *sdl.Renderer) {
	for _, ent := range s.entities {
		ent.Draw(renderer)
	}
}

func loadTilemapFromFile(path string) {

}
