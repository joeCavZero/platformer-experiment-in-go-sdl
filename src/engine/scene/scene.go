package scene

import (
	"project/src/engine/tilemap"
	"project/src/entities/entity"
	"project/src/layer"
	"project/src/settings"

	"github.com/veandco/go-sdl2/sdl"
)

type Scene struct {
	entities []entity.EntityInterface
	tilemap  [settings.TILE_QUANTITY]tilemap.Tile
	layers   []*layer.Layer
}

func NewScene() *Scene {
	return &Scene{
		entities: make([]entity.EntityInterface, 0),
		tilemap:  [settings.TILE_QUANTITY]tilemap.Tile{},
		layers:   make([]*layer.Layer, 0),
	}
}

func (s *Scene) AddEntity(ent entity.EntityInterface, layer uint8) {
	ent.SetLayer(layer)
	s.entities = append(s.entities, ent)
}

func (s *Scene) Process(keyboard *[]uint8) {
	for _, ent := range s.entities {
		ent.Update(keyboard)
	}
}

func (s *Scene) Render(renderer *sdl.Renderer) {

	for lyr_index, lyr := range s.GetLayers() {

		switch lyr.GetLayerType() {
		case 't':
			lyr.RenderTilemap(renderer)

			renderer.SetDrawColor(0, 0, 0, 0)
			renderer.DrawLine(0, 0, 100, 100)
		case 'e':
			for _, ent := range s.GetEntities() {
				if ent.GetLayer() == uint8(lyr_index) {
					ent.Draw(renderer)
				}
			}
		default:
			continue
		}

	}

}

func (s *Scene) GetEntities() []entity.EntityInterface {
	return s.entities
}

func (s *Scene) GetLayers() []*layer.Layer {
	return s.layers
}

func (s *Scene) AddLayer(lyr *layer.Layer) {
	s.layers = append(s.layers, lyr)
}

func loadTilemapFromFile(path string) {

}
