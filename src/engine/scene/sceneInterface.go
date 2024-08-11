package scene

import (
	"project/src/entities/entity"

	"github.com/veandco/go-sdl2/sdl"
)

type SceneInterface interface {
	Process(keyboard *[]uint8)
	Render(renderer *sdl.Renderer)

	AddEntity(ent entity.EntityInterface)
}
