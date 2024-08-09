package entities

import (
	"project/src/settings"

	"github.com/veandco/go-sdl2/sdl"
)

type Player struct {
	Entity
	speed     float32
	velocity  sdl.FPoint
	jumpForce float32
}

func NewPlayer() *Player {
	return &Player{
		Entity: Entity{
			Position: sdl.FPoint{
				X: 0,
				Y: 0,
			},
			Size: sdl.FPoint{
				X: 32,
				Y: 32,
			},
		},
		speed: 1.0,
		velocity: sdl.FPoint{
			X: 0,
			Y: 0,
		},
		jumpForce: 3.0,
	}
}

func (p *Player) Update(keyboard *[]uint8) {
	if (*keyboard)[sdl.SCANCODE_D] == 1 {
		p.velocity.X = p.speed
	} else if (*keyboard)[sdl.SCANCODE_A] == 1 {
		p.velocity.X = -p.speed
	} else {
		p.velocity.X = 0
	}

	if (*keyboard)[sdl.SCANCODE_W] == 1 && p.Position.Y+p.Size.Y == settings.CANVAS_HEIGHT {
		p.velocity.Y = -p.jumpForce
	}

	p.velocity.Y += 0.1

	p.moveAndCollide()
}

func (p *Player) moveAndCollide() {
	p.Position.X += p.velocity.X
	p.Position.Y += p.velocity.Y

	if p.Position.Y+p.Size.Y > settings.CANVAS_HEIGHT {
		p.Position.Y = settings.CANVAS_HEIGHT - p.Size.Y
	}

	if p.Position.X < 0 {
		p.Position.X = 0
	} else if p.Position.X+p.Size.X > settings.CANVAS_WIDTH {
		p.Position.X = settings.CANVAS_WIDTH - p.Size.X
	}
}

func (p *Player) Draw(renderer *sdl.Renderer) {
	renderer.DrawRectF(
		&sdl.FRect{
			X: p.Position.X,
			Y: p.Position.Y,
			W: p.Size.X,
			H: p.Size.Y,
		},
	)
}
