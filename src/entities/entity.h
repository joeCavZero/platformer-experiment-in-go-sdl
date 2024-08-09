#ifndef ENTITY_H
#define ENTITY_H

#include "SDL.h"
#include "SDL_image.h"
#include <memory>

class Entity {
    public:
        SDL_FPoint position;
        SDL_FPoint size;
        SDL_FPoint velocity;
        std::shared_ptr<SDL_Texture>* texture;

        Entity(float x, float y);
        ~Entity();

        void update();
        void draw();
};

#endif