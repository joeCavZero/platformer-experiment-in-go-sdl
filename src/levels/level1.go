package levels

import (
	"project/src/engine/engine"
	"project/src/engine/layer"
	"project/src/engine/scene"
	"project/src/entities/player"
)

type Level1 struct {
	scene.Scene
}

func NewLevel1(engine *engine.Engine) *Level1 {
	level := Level1{}
	level.AddLayer(
		layer.NewTilemapLayer("data/level.data", "assets/tilemap.png", engine.Renderer),
	)
	level.AddLayer(
		layer.NewEntityLayer(),
	)

	level.AddEntity(
		player.NewPlayer(engine),
		1,
	)
	return &level
}
