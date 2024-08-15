package player

import (
	"math"
	assetmanager "project/src/engine/assetManager"
	"project/src/engine/engine"
	"project/src/entities/entity"

	"github.com/veandco/go-sdl2/sdl"
)

type Player struct {
	entity.Entity
	speed     float32
	velocity  sdl.FPoint
	jumpForce float32

	texture         *sdl.Texture
	animation_index float32
	collision_rect  sdl.FRect
	is_mirrored     bool

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
		texture:   assetmanager.GetTexture("assets/player-sheet.png", engine.Renderer),
		collision_rect: sdl.FRect{
			X: 6, Y: 0,
			W: 20, H: 32,
		},
		is_mirrored: false,
		engine:      engine,
	}
}

func (p *Player) SetEngine(engine *engine.Engine) {
	p.engine = engine
}

func (p *Player) Update(keyboard *[]uint8) {
	if (*keyboard)[sdl.SCANCODE_D] == 1 {
		p.velocity.X = p.speed
		p.is_mirrored = false
	} else if (*keyboard)[sdl.SCANCODE_A] == 1 {
		p.velocity.X = -p.speed
		p.is_mirrored = true
	} else {
		p.velocity.X = 0
	}

	if (*keyboard)[sdl.SCANCODE_W] == 1 && p.isOnFloor() {
		p.velocity.Y = -p.jumpForce
	}

	if p.velocity.X != 0 {
		p.animation_index += 0.3
	}
	p.velocity.Y += 0.1

	if p.animation_index > 8.0 {
		p.animation_index = 0
	}

	p.moveAndCollide()
}

func (p *Player) moveAndCollide() {
	p.Position.X += p.velocity.X
	p.Position.Y += p.velocity.Y

	for _, tile := range p.engine.GetScene().GetLayers()[0].GetTilemap() {
		if tile.CheckCollision(int32(p.Position.X+p.collision_rect.X+p.velocity.X), int32(p.Position.Y+p.collision_rect.Y), int32(p.collision_rect.W), int32(p.collision_rect.H)) {
			p.Position.X -= p.velocity.X
			p.velocity.X = 0
		}
		if tile.CheckCollision(int32(p.Position.X+p.collision_rect.X), int32(p.Position.Y+p.collision_rect.Y+redDirection(p.velocity.Y)), int32(p.collision_rect.W), int32(p.collision_rect.H)) {
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
			int32(p.Position.X+p.collision_rect.X+foot_offset),
			int32(p.Position.Y+p.collision_rect.Y+p.Size.Y),
			int32(p.collision_rect.W-foot_offset*2),
			int32(foot_offset*2),
		) {
			return true
		}
	}
	return false
}

func (p *Player) Draw(renderer *sdl.Renderer) {
	var delta sdl.RendererFlip = 0

	if p.is_mirrored {
		delta = sdl.FLIP_HORIZONTAL
	}

	renderer.CopyExF(
		p.texture,
		&sdl.Rect{
			X: int32(math.Floor(float64(p.animation_index))) * 32,
			Y: 0,
			W: 32,
			H: 32,
		},
		&sdl.FRect{
			X: p.Position.X,
			Y: p.Position.Y,
			W: p.Size.X,
			H: p.Size.Y,
		},
		0,
		&sdl.FPoint{
			X: 0,
			Y: 0,
		},
		delta,
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
