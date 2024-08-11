package main

import (
	"project/src/engine/engine"
	"project/src/levels"
)

func main() {
	engine := engine.NewEngine()
	engine.InitCore()

	level1 := levels.NewLevel1(engine)
	engine.AddScene(level1)

	engine.Run()
}
