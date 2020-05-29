package terminus

import (

	"github.com/gdamore/tcell"

)

type Game struct {

	screen tcell.Screen
	scenes []Scene
	scene_index int

}

func NewGame() *Game {

	game := &Game{}

	return game

}

func (game *Game) Init( scenes []Scene ) {
	
	// TODO: Error checking
	screen, _ := tcell.NewScreen()

	game.screen = screen

	game.scene_index = 0
	game.scenes = scenes

	game.screen.Init()
	game.scenes[game.scene_index].Init()

}

func (game *Game) Start() {

// game_loop:
	for {

		game.scenes[game.scene_index].Update()
		game.scenes[game.scene_index].Draw()

	}

}

func (game *Game) NextScene() {

	if game.scene_index < len( game.scenes ) {
		game.scene_index += 1
	}

}

func (game *Game) PrevScene() {

	if game.scene_index > 0 {
		game.scene_index -= 1
	}
	
}

func (game *Game) Screen() tcell.Screen {
	return game.screen
}