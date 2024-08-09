package scene

import (
	"project/src/engine/tilemap"
	"project/src/entities/entity"
	"project/src/settings"

	"github.com/veandco/go-sdl2/sdl"
)

type Scene struct {
	entities []entity.EntityInterface
	tilemap  [settings.TILE_QUANTITY]tilemap.Tile
}

func NewScene() *Scene {
	return &Scene{
		entities: make([]entity.EntityInterface, 0),
		tilemap:  [settings.TILE_QUANTITY]tilemap.Tile{},
	}
}

func (s *Scene) AddEntity(ent entity.EntityInterface) {
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
