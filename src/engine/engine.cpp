
#include "engine.h"
#include <iostream>
#include "settings.h"
#include "controllers/texture_controller.h"

Engine::Engine() {
}

Engine::~Engine() {
}

void Engine::initCore() {
    if( SDL_Init( SDL_INIT_EVERYTHING ) != 0 ) {
        SDL_Log( "SDL nao inicializado com sucesso:\n--> %s", SDL_GetError() );
    }

    this->window = SDL_CreateWindow(
        "Engine", 
        SDL_WINDOWPOS_CENTERED, SDL_WINDOWPOS_CENTERED, 
        CANVAS_WIDTH, CANVAS_HEIGHT,
        SDL_WINDOW_RESIZABLE
    );
    if( this->window == nullptr ) {
        SDL_Log( "Janela (Window) nao inicializado com sucesso:\n--> %s", SDL_GetError() );
    }

    this->renderer = SDL_CreateRenderer(
        this->window,
        -1,
        0
    );
    if( this->renderer == nullptr ) {
        SDL_Log( "Renderizador (Renderer) nao inicializado com sucesso:\n--> %s", SDL_GetError() );
    }

    this->canvas = SDL_CreateTexture(
        this->renderer,
        SDL_PIXELFORMAT_RGBA8888,
        SDL_TEXTUREACCESS_TARGET,
        CANVAS_WIDTH, CANVAS_HEIGHT
    );

    this->textureController = new TextureController(this->renderer);
}

void Engine::run() {
    this->initCore();
    this->isRunning = true;

    while( this->isRunning ){
        this->handleInput();
        this->process();
        this->render();
    }

    this->close();
}

void Engine::handleInput() {
    SDL_Event event;
    while( SDL_PollEvent( &event ) ) {
        switch( event.type ) {
            case SDL_QUIT:
                this->isRunning = false;
                break;
            case SDL_KEYDOWN:
                switch( event.key.keysym.sym ) {
                    case SDLK_p:
                        this->isRunning = false;
                        break;
                    default:
                        break;
                }
                break;
            default:
                break;
        }
    }
}

void Engine::process() {

}

void Engine::render() {
    SDL_SetRenderDrawColor( this->renderer, 0,0,0,255);
    SDL_RenderClear( this->renderer );

    SDL_SetRenderTarget( this->renderer, this->canvas );
        {
            SDL_SetRenderDrawColor( this->renderer, 255,255,25,255);
            SDL_RenderClear( this->renderer );
        }
    this->renderCanvasOnScreen();
    SDL_RenderPresent( this->renderer );
}

void Engine::close() {
    delete this->textureController;
    SDL_DestroyTexture( this->canvas );
    SDL_DestroyRenderer( this->renderer );
    SDL_DestroyWindow( this->window );
    SDL_Quit();
}

void Engine::renderCanvasOnScreen() {
    float window_width = SDL_GetWindowSurface( this->window )->w;
    float window_height = SDL_GetWindowSurface( this->window )->h;

    float delta_x = window_width / CANVAS_WIDTH;
    float delta_y = window_height / CANVAS_HEIGHT;
    
    float scale = 1;
    
    if( delta_x > delta_y ){
        scale = delta_y;
    } else {
        scale = delta_x;
    }

    int diff_x = window_width - ( CANVAS_WIDTH * scale );
    int diff_y = window_height - ( CANVAS_HEIGHT * scale );

    SDL_SetRenderTarget( this->renderer, nullptr );
    SDL_Rect* aux_dstrect = new SDL_Rect{
        diff_x/2, diff_y/2,
        //CANVAS_WIDTH , CANVAS_HEIGHT
        (int)(CANVAS_WIDTH * scale) , (int)(CANVAS_HEIGHT * scale)
    };

    SDL_RenderCopy(
        this->renderer,
        this->canvas,
        nullptr,
        aux_dstrect
    );

    delete aux_dstrect;
}