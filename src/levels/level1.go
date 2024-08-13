package levels

import (
	"project/src/engine/engine"
	"project/src/engine/scene"
	"project/src/entities/player"
	"project/src/layer"
)

type Level1 struct {
	scene.Scene
}

func NewLevel1(engine *engine.Engine) *Level1 {
	level := Level1{}
	level.AddLayer(
		layer.NewTilemapLayer("data/level.data"),
	)
	level.AddLayer(
		layer.NewEntityLayer(),
	)

	level.AddEntity(
		player.NewPlayer(),
		1,
	)
	return &level
}
