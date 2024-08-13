package layer

import (
	"bufio"
	"os"
	assetManager "project/src/engine/assetManager"
	"project/src/engine/tilemap"
	"project/src/settings"

	"github.com/veandco/go-sdl2/sdl"
)

type Layer struct {
	layerType byte
	data_path string
	texture   *sdl.Texture
	tilemap   *[settings.TILE_QUANTITY]tilemap.Tile
}

func NewEntityLayer() *Layer {
	return &Layer{
		layerType: 'e',
		data_path: "",
		texture:   nil,
		tilemap:   nil,
	}
}

func NewTilemapLayer(data_path string, asset_path string, renderer *sdl.Renderer) *Layer {
	layer := Layer{
		layerType: 't',
		data_path: data_path,
		texture:   assetManager.GetTexture(asset_path, renderer),
		tilemap:   &[settings.TILE_QUANTITY]tilemap.Tile{},
	}

	layer.LoadTilemap()

	return &layer
}

func (l *Layer) GetLayerType() byte {
	return l.layerType
}

func (l *Layer) LoadTilemap() {
	/*
		for y_index := 0; y_index < settings.TILE_Y_QUANTITY; y_index++ {
			for x_index := 0; x_index < settings.TILE_X_QUANTITY; x_index++ {
				l.tilemap[(y_index*settings.TILE_X_QUANTITY)+x_index] = tilemap.Tile{
					TileType: 0,
					Position: sdl.Point{
						X: int32(x_index * 32),
						Y: int32(y_index * 32),
					},
				}
			}
		}
	*/

	file, err := os.Open("data/level.data")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	content := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content += scanner.Text()
	}
	for i := 0; i < settings.TILE_QUANTITY; i++ {
		char := content[i]
		switch char {
		case '0':
			l.tilemap[i] = tilemap.Tile{
				TileType: 0,
				Position: sdl.Point{X: 0, Y: 0},
			}
		case '1':
			l.tilemap[i] = tilemap.Tile{
				TileType: int(char),
				Position: sdl.Point{
					X: int32(
						(i % settings.TILE_X_QUANTITY) * 32,
					),
					Y: int32(
						int(i/settings.TILE_X_QUANTITY) * 32,
					),
				},
			}
		}

	}
}

func (l *Layer) RenderTilemap(renderer *sdl.Renderer) {
	for _, tile := range l.tilemap {
		if tile.TileType == 0 {
			continue
		} else {
			renderer.Copy(
				l.texture,
				&sdl.Rect{
					X: 0, Y: 0,
					W: 32, H: 32,
				},
				&sdl.Rect{
					X: tile.Position.X, Y: tile.Position.Y,
					W: 32, H: 32,
				},
			)
		}
	}
}
