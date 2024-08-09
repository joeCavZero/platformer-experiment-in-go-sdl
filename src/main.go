package main

import (
	"project/src/engine/engine"
	"project/src/entities/player"
)

func main() {
	engine := engine.NewEngine()
	engine.InitCore()
	engine.Scene.AddEntity(player.NewPlayer(engine))

	engine.Run()
}
