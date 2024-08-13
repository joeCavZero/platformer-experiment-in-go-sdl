package scene

import (
	"project/src/entities/entity"
	"project/src/layer"

	"github.com/veandco/go-sdl2/sdl"
)

type SceneInterface interface {
	Process(keyboard *[]uint8)
	Render(renderer *sdl.Renderer)

	AddEntity(ent entity.EntityInterface, layer uint8)

	GetEntities() []entity.EntityInterface
	GetLayers() []*layer.Layer
}
