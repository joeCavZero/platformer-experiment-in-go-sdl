package main

import (
	"project/src/engine/engine"
	"project/src/levels"
)

func main() {
	engine := engine.NewEngine()

	level1 := levels.NewLevel1(engine)

	engine.SetScene(level1)

	engine.Run()
}
