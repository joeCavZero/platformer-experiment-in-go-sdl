package tilemap

import "github.com/veandco/go-sdl2/sdl"

type Tile struct {
	Position sdl.Point
	TileType int
}
