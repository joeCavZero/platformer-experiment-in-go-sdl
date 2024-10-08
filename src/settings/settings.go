package settings

const (
	CANVAS_WIDTH  = 640
	CANVAS_HEIGHT = 360

	FPS            = 60
	TICK_PER_FRAME = 1000 / FPS

	TILE_X_QUANTITY = CANVAS_WIDTH / 32
	TILE_Y_QUANTITY = CANVAS_HEIGHT/32 + 1
	TILE_QUANTITY   = TILE_X_QUANTITY * TILE_Y_QUANTITY
)
