#ifndef TEXTURE_CONTROLLER_H
#define TEXTURE_CONTROLLER_H

#include "SDL.h"
#include <vector>
#include <map>
#include <string>
#include <memory>

class TextureController {
    public:
        std::map<const char*, SDL_Texture* > textures;
        SDL_Renderer* renderer;

        TextureController( SDL_Renderer* renderer );
        ~TextureController();

        void LoadTexture( const char* file_path, const char* name );
        SDL_Texture* GetTexture( const char* name );
};

#endif