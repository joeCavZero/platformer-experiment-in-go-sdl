package assetmanager

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

var textures map[string]*sdl.Texture = make(map[string]*sdl.Texture)

func GetTexture(asset_path string, renderer *sdl.Renderer) *sdl.Texture {
	if textures[asset_path] == nil {
		new_texture, err := img.LoadTexture(renderer, asset_path)
		if err != nil {
			fmt.Println("Error loading texture: ", asset_path)
			panic(err)
		}
		textures[asset_path] = new_texture
		return textures[asset_path]
	} else {
		return textures[asset_path]
	}
}
