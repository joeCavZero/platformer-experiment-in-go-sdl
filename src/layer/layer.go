package layer

import "github.com/veandco/go-sdl2/sdl"

type Layer struct {
	layerType byte
	assetPath string
}

func NewLayer(layer_type byte, asset_path string) *Layer {
	return &Layer{
		layerType: layer_type,
		assetPath: asset_path,
	}
}

func NewEntityLayer() *Layer {
	return &Layer{
		layerType: 'e',
		assetPath: "",
	}
}

func NewTilemapLayer(asset_path string) *Layer {
	return &Layer{
		layerType: 't',
		assetPath: asset_path,
	}
}

func (l *Layer) GetLayerType() byte {
	return l.layerType
}

func (l *Layer) GetAssetPath() string {
	return l.assetPath
}

func (l *Layer) RenderTilemap(renderer *sdl.Renderer) {
	// TODO
}
