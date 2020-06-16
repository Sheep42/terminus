package terminus

import (

	"github.com/gdamore/tcell"
	// "os"
	// "time"

)

type Game struct {

	screen tcell.Screen
	scenes []IScene
	scene_index int
	exit_key tcell.Key
	input *tcell.EventKey
	chan_key_press chan *tcell.EventKey

}

func NewGame() *Game {

	game := &Game{}

	return game

}

func (game *Game) Init( scenes []IScene ) {
	
	// TODO: Error checking
	screen, _ := tcell.NewScreen()

	game.screen = screen
	game.exit_key = KEY_ESC

	game.scene_index = 0
	game.scenes = scenes

	game.screen.Init()
	game.scenes[game.scene_index].Init()

	if len( game.scenes[game.scene_index].Entities() ) > 0 {

		for _, entity := range game.scenes[game.scene_index].Entities() {
			entity.Init()
		}

	}

	game.chan_key_press = make(chan *tcell.EventKey)

}

func (game *Game) getInput() {
		
	screen := game.screen

	for {

		game.input = nil
		
		ev := screen.PollEvent()

		switch ev := ev.(type) {
			case *tcell.EventResize:
				screen.Sync()
			case *tcell.EventKey:
				select{
					case game.chan_key_press <-ev:
				}
			default:
		}

	}

}

func (game *Game) handleInput() {

	select {
		case game.input = <-game.chan_key_press:
		default:
	}

}

func (game *Game) Start() {

	screen := game.screen
	// clock := time.Now()

	go game.getInput();

game_loop:
	for {

		//TODO: Need to pass tick time to update
		
		// update := time.Now()
		// delta := update.Sub(clock).Seconds()
		// clock = update
		// 
		
		screen.Clear()

		game.handleInput()

		if game.input != nil && game.input.Key() == game.exit_key {
			screen.Fini()
			break game_loop
		}

		scene := game.scenes[game.scene_index]
		scene.Update()

		if len( scene.Entities() ) > 0 {

			for _, entity := range scene.Entities() {
				entity.Update()
			}

		}

		scene.Draw()

		screen.Show()

	}

}

func (game *Game) NextScene() {

	if game.scene_index < len( game.scenes ) - 1 {
		game.scene_index += 1

		game.scenes[game.scene_index].Init()
	}

}

func (game *Game) PrevScene() {

	if game.scene_index > 0 {
		game.scene_index -= 1

		game.scenes[game.scene_index].Init()
	}
	
}

func (game *Game) Screen() tcell.Screen {
	return game.screen
}

func (game *Game) ExitKey() tcell.Key {
	return game.exit_key
}

func (game *Game) SetExitKey( exit_key tcell.Key ) {
	game.exit_key = exit_key
}

func (game *Game) Input() *tcell.EventKey {
	return game.input
}