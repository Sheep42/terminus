package terminus

import (
	"time"

	"github.com/gdamore/tcell"
)

type Game struct {
	screen       tcell.Screen
	width        int
	height       int
	scenes       []IScene
	sceneIndex   int
	exitKey      tcell.Key
	input        *tcell.EventKey
	chanKeyPress chan *tcell.EventKey
}

func NewGame() *Game {

	game := &Game{}

	return game

}

func (game *Game) Init(scenes []IScene) {

	// TODO: Error checking
	screen, _ := tcell.NewScreen()

	game.screen = screen
	game.exitKey = KeyEsc

	game.sceneIndex = 0
	game.scenes = scenes

	game.screen.Init()
	game.scenes[game.sceneIndex].Init()

	if len(game.scenes[game.sceneIndex].Entities()) > 0 {

		for _, entity := range game.scenes[game.sceneIndex].Entities() {
			entity.Init()
		}

	}

	game.chanKeyPress = make(chan *tcell.EventKey)

}

func (game *Game) getInput() {

	screen := game.screen

	var ev tcell.Event

	for {

		game.input = nil

		ev = screen.PollEvent()

		switch eventType := ev.(type) {

		case *tcell.EventResize:
			screen.Sync()
			game.width, game.height = screen.Size()

		case *tcell.EventKey:
			select {
			case game.chanKeyPress <- eventType:
			}

		default:

		}

	}

}

func (game *Game) handleInput() {

	select {
	case game.input = <-game.chanKeyPress:
	default:
	}

}

func (game *Game) Start() {

	screen := game.screen
	clock := time.Now()

	go game.getInput()

	game.width, game.height = screen.Size()

game_loop:
	for {

		update := time.Now()
		delta := update.Sub(clock).Seconds()
		clock = update

		screen.Clear()

		select {
		case <-game.chanKeyPress:
			game.handleInput()
			continue
		default:
		}

		if game.input != nil && game.input.Key() == game.exitKey {
			screen.Fini()
			break game_loop
		}

		scene := game.scenes[game.sceneIndex]
		scene.Update(delta)

		if len(scene.Entities()) > 0 {

			for _, entity := range scene.Entities() {
				entity.Update(delta)
			}

		}

		scene.Draw()
		screen.Show()

	}

}

func (game *Game) NextScene() {

	if game.sceneIndex < len(game.scenes)-1 {
		game.sceneIndex += 1

		game.scenes[game.sceneIndex].Init()
	}

}

func (game *Game) PrevScene() {

	if game.sceneIndex > 0 {
		game.sceneIndex -= 1

		game.scenes[game.sceneIndex].Init()
	}

}

func (game *Game) ExitKey() tcell.Key {
	return game.exitKey
}

func (game *Game) SetExitKey(exitKey tcell.Key) {
	game.exitKey = exitKey
}

func (game *Game) Input() *tcell.EventKey {
	return game.input
}

func (game *Game) ScreenSize() (int, int) {

	return game.width, game.height

}
