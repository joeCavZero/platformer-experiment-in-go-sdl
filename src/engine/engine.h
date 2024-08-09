#ifndef ENGINE_H
#define ENGINE_H

#include "SDL.h"
#include "controllers/texture_controller.h"

class Engine {
    public:

        SDL_Window* window;
        SDL_Renderer* renderer;
        bool isRunning;
    
        SDL_Texture* canvas;

        TextureController* textureController;



        Engine();
        ~Engine();

        void initCore();
        void run();

        void handleInput();
        void process();
        void render();

        void close();

    private:
        void renderCanvasOnScreen();
};

#endif