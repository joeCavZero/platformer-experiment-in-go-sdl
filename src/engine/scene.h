#ifndef SCENE_H
#define SCENE_H

#include "SDL.h"
#include <vector>
#include "entities/entity.h"

class Scene {
    public:
        std::vector<Entity*> entities;

        Scene();
        ~Scene();


};

#endif