package tilemap

import "github.com/veandco/go-sdl2/sdl"

type Tile struct {
	Position sdl.Point
	TileType uint8
}

func (t *Tile) CheckCollision(other_x, other_y, other_w, other_h int32) bool {
	if t.Position.X < other_x+other_w &&
		t.Position.X+32 > other_x &&
		t.Position.Y < other_y+other_h &&
		t.Position.Y+32 > other_y {
		return true
	}
	return false
}
