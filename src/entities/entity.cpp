#ifndef ENTITY_H
#define ENTITY_H

#include "SDL.h"

class Entity {
    public:
        SDL_FPoint position;
        SDL_FPoint size;
        SDL_FPoint velocity;

        Entity();
        Entity(float x , float y );
        ~Entity();
};

#endif