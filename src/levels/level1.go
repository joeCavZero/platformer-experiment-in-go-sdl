package levels

import (
	"project/src/engine/engine"
	"project/src/engine/scene"
	"project/src/entities/player"
)

type Level1 struct {
	scene.Scene
}

func NewLevel1(engine *engine.Engine) *Level1 {
	level := Level1{}
	level.AddEntity(player.NewPlayer(engine))
	return &level
}
