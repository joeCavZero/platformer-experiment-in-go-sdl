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

func (l *Layer) GetTilemap() *[settings.TILE_QUANTITY]tilemap.Tile {
	return l.tilemap
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

	var byte_map map[byte]uint8 = map[byte]uint8{
		'0': 0,
		'1': 1,
		'2': 2,
	}
	for i := 0; i < settings.TILE_QUANTITY; i++ {
		char := content[i]
		switch char {
		case '0':
			l.tilemap[i] = tilemap.Tile{
				TileType: 0,
				Position: sdl.Point{X: 0, Y: 0},
			}
		default:
			l.tilemap[i] = tilemap.Tile{
				TileType: byte_map[char],
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

		switch tile.TileType {
		case 1:
			renderer.FillRect(
				&sdl.Rect{
					X: tile.Position.X, Y: tile.Position.Y,
					W: 32, H: 32,
				},
			)
			renderer.Copy(
				l.texture,
				&sdl.Rect{
					X: 16, Y: 16,
					W: 16, H: 16,
				},
				&sdl.Rect{
					X: tile.Position.X, Y: tile.Position.Y,
					W: 32, H: 32,
				},
			)
		case 2:
			renderer.Copy(
				l.texture,
				&sdl.Rect{
					X: 16, Y: 0,
					W: 16, H: 16,
				},
				&sdl.Rect{
					X: tile.Position.X, Y: tile.Position.Y,
					W: 32, H: 32,
				},
			)
		}
	}
}
