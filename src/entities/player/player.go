package player

import (
	"project/src/engine/engine"
	"project/src/entities/entity"

	"github.com/veandco/go-sdl2/sdl"
)

type Player struct {
	entity.Entity
	speed     float32
	velocity  sdl.FPoint
	jumpForce float32

	engine *engine.Engine
}

func NewPlayer(engine *engine.Engine) *Player {
	return &Player{
		Entity: entity.Entity{
			Position: sdl.FPoint{
				X: 32,
				Y: 128,
			},
			Size: sdl.FPoint{
				X: 30,
				Y: 30,
			},
		},
		speed: 1.0,
		velocity: sdl.FPoint{
			X: 0,
			Y: 0,
		},
		jumpForce: 3.0,
		engine:    engine,
	}
}

func (p *Player) SetEngine(engine *engine.Engine) {
	p.engine = engine
}

func (p *Player) Update(keyboard *[]uint8) {
	if (*keyboard)[sdl.SCANCODE_D] == 1 {
		p.velocity.X = p.speed
	} else if (*keyboard)[sdl.SCANCODE_A] == 1 {
		p.velocity.X = -p.speed
	} else {
		p.velocity.X = 0
	}

	if (*keyboard)[sdl.SCANCODE_W] == 1 && p.isOnFloor() {
		p.velocity.Y = -p.jumpForce
	}

	p.velocity.Y += 0.1

	p.moveAndCollide()
}

func (p *Player) moveAndCollide() {
	p.Position.X += p.velocity.X
	p.Position.Y += p.velocity.Y

	for _, tile := range p.engine.GetScene().GetLayers()[0].GetTilemap() {
		if tile.CheckCollision(int32(p.Position.X+p.velocity.X), int32(p.Position.Y), int32(p.Size.X), int32(p.Size.Y)) {
			p.Position.X -= p.velocity.X
			p.velocity.X = 0
		}
		if tile.CheckCollision(int32(p.Position.X), int32(p.Position.Y+redDirection(p.velocity.Y)), int32(p.Size.X), int32(p.Size.Y)) {
			if p.velocity.Y > 0 { // if falling

				p.Position.Y -= p.velocity.Y
				p.velocity.Y = 0
			} else if p.velocity.Y < 0 { // if jumping

				p.Position.Y -= p.velocity.Y
				p.velocity.Y = 0
			}
		}
	}

}

func (p *Player) isOnFloor() bool {
	for _, tile := range p.engine.GetScene().GetLayers()[0].GetTilemap() {
		var foot_offset float32 = 1
		if tile.CheckCollision(
			int32(p.Position.X+foot_offset),
			int32(p.Position.Y+p.Size.Y),
			int32(p.Size.X-foot_offset*2),
			int32(foot_offset*2),
		) {
			return true
		}
	}
	return false
}

func (p *Player) Draw(renderer *sdl.Renderer) {
	renderer.SetDrawColor(255, 255, 255, 255)

	renderer.FillRectF(
		&sdl.FRect{
			X: p.Position.X,
			Y: p.Position.Y,
			W: p.Size.X,
			H: p.Size.Y,
		},
	)
}

func redDirection(x float32) float32 {
	if x > 0 {
		return 0.1
	} else if x < 0 {
		return -0.1
	}
	return 0
}
