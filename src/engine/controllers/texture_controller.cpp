#include "texture_controller.h"
#include "SDL_image.h"
#include <iostream>
using namespace std;

TextureController::TextureController(SDL_Renderer* renderer) {
    this->renderer = renderer;
    this->textures = std::map<const char*, SDL_Texture*>();
}

TextureController::~TextureController() {
    for ( auto const& [name, texture] : this->textures ) {
        SDL_DestroyTexture( texture );
        this->textures.erase( name );
    }
    cout << this->textures.size() << endl;
    this->textures.clear();
}	

void TextureController::LoadTexture( const char* file_path , const char* name ) {
    SDL_Texture* new_texture = IMG_LoadTexture( this->renderer, file_path );
    this->textures.insert( 
        std::pair< const char*, SDL_Texture* >( name, new_texture)
    );
}

SDL_Texture* TextureController::GetTexture( const char* name ) {
    return this->textures.at( name );
}