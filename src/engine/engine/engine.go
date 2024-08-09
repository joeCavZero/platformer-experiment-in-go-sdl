package engine

import (
	"project/src/engine/scene"
	"project/src/settings"

	"github.com/veandco/go-sdl2/sdl"
)

type Engine struct {
	Window    *sdl.Window
	Renderer  *sdl.Renderer
	Canvas    *sdl.Texture
	IsRunning bool
	Keyboard  []uint8
	Scene     *scene.Scene
}

func NewEngine() *Engine {
	return &Engine{}
}

func (e *Engine) Run() {

	for e.IsRunning {
		first_time := sdl.GetTicks64()

		e.handleEvents()
		e.process()
		e.render()

		elapsed_time := sdl.GetTicks64() - first_time

		// if the elapsed time is less than 16ms,
		// its 16 because 1000ms / 60fps = 16ms
		if elapsed_time < settings.TICK_PER_FRAME {
			sdl.Delay(uint32(settings.TICK_PER_FRAME - elapsed_time))
		}
	}

	e.close()

}

func (e *Engine) handleEvents() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			e.IsRunning = false
		}
	}

	e.Keyboard = sdl.GetKeyboardState()
}

func (e *Engine) process() {
	e.Scene.Process(&e.Keyboard)
}

func (e *Engine) render() {
	e.Renderer.SetDrawColor(0, 0, 0, 255)
	e.Renderer.Clear()

	e.Renderer.SetRenderTarget(e.Canvas)
	{
		e.Renderer.SetDrawColor(0, 155, 155, 255)
		e.Renderer.Clear()

		e.Renderer.SetDrawColor(255, 255, 255, 255)
		e.Renderer.DrawLine(0, 0, 640, 360)

		e.Scene.Render(e.Renderer)
	}
	e.Renderer.SetRenderTarget(nil)
	e.renderCanvas()

	e.Renderer.Present()
}

func (e *Engine) renderCanvas() {
	win_width, win_height := e.Window.GetSize()

	delta_x := float32(win_width) / settings.CANVAS_WIDTH
	delta_y := float32(win_height) / settings.CANVAS_HEIGHT

	var scale float32 = 1.0

	//the smallest of the two deltas
	if delta_x > delta_y {
		scale = delta_y
	} else {
		scale = delta_x
	}

	diff_x := float32(win_width) - (settings.CANVAS_WIDTH * scale)
	diff_y := float32(win_height) - (settings.CANVAS_HEIGHT * scale)

	e.Renderer.Copy(
		e.Canvas,
		nil,
		&sdl.Rect{
			X: int32(diff_x / 2),
			Y: int32(diff_y / 2),
			W: int32(settings.CANVAS_WIDTH * scale),
			H: int32(settings.CANVAS_HEIGHT * scale),
		},
	)

}

func (e *Engine) close() {
	e.Renderer.Destroy()
	e.Window.Destroy()
	sdl.Quit()
}

func (e *Engine) InitCore() {
	//==== SDL INITIALIZATION ====
	err := sdl.Init(uint32(sdl.INIT_EVERYTHING))
	if err != nil {
		panic(err)
	}

	//==== WINDOW ====
	e.Window, err = sdl.CreateWindow(
		"TITLE",
		int32(sdl.WINDOWPOS_UNDEFINED),
		int32(sdl.WINDOWPOS_UNDEFINED),
		settings.CANVAS_WIDTH,
		settings.CANVAS_HEIGHT,
		sdl.WINDOW_RESIZABLE,
	)
	if err != nil {
		panic(err)
	}

	//==== RENDERER ====
	e.Renderer, err = sdl.CreateRenderer(
		e.Window,
		-1,
		0,
	)
	if err != nil {
		panic(err)
	}

	//==== CANVAS ====
	e.Canvas, err = e.Renderer.CreateTexture(
		sdl.PIXELFORMAT_RGBA8888,
		sdl.TEXTUREACCESS_TARGET,
		settings.CANVAS_WIDTH,
		settings.CANVAS_HEIGHT,
	)
	if err != nil {
		panic(err)
	}

	//==== SCENE ====
	e.Scene = scene.NewScene()

	e.IsRunning = true
}
