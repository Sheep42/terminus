package terminus

import (
	"os"
	"time"

	"github.com/gdamore/tcell"
)

// TODO: Implement Logging
// TODO: Implement Debug Mode

// Game is collection of properties used to
// abstract interaction with a tcell Screen
type Game struct {
	screen       tcell.Screen
	width        int
	height       int
	scenes       []IScene
	sceneIndex   int
	exitKey      tcell.Key
	input        *tcell.EventKey
	chanKeyPress chan *tcell.EventKey
	fps          float64
}

// NewGame creates a game
func NewGame() *Game {

	game := &Game{}

	return game

}

// Init takes an array of scenes, and sets up the game
// before the loop is started
func (game *Game) Init(scenes []IScene) {

	screen, err := tcell.NewScreen()
	if err != nil {
		// TODO: Log error
		os.Exit(2)
	}

	game.screen = screen
	game.exitKey = KeyEsc

	game.sceneIndex = 0
	game.scenes = scenes

	if game.fps == 0 {
		game.fps = 60
	}

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
		game.input = nil
	}

}

// Start begins listening for input and starts the game loop
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

		game.handleInput()

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

		// enforce fps
		select {
		case <-time.After(time.Duration((update.Sub(time.Now()).Seconds()*1000.0)+1000.0/game.fps) * time.Millisecond):
			continue
		}
	}

}

// NextScene increments the game sceneIndex if
// we are not already at the last scene
func (game *Game) NextScene() {

	if game.sceneIndex < len(game.scenes)-1 {
		game.sceneIndex++

		game.scenes[game.sceneIndex].Init()
	}

}

// PrevScene decrements the game sceneIndex if
// we are not already at the first scene
func (game *Game) PrevScene() {

	if game.sceneIndex > 0 {
		game.sceneIndex--

		game.scenes[game.sceneIndex].Init()
	}

}

// ExitKey gets the assigned exit key
func (game *Game) ExitKey() tcell.Key {
	return game.exitKey
}

// SetExitKey sets the game's exit key
func (game *Game) SetExitKey(exitKey tcell.Key) {
	game.exitKey = exitKey
}

// GetFPS gets the game's target FPS
func (game *Game) GetFPS() float64 {
	return game.fps
}

// SetFPS sets the game's target FPS
func (game *Game) SetFPS(fps float64) {
	game.fps = fps
}

// Input gets the current input as an EventKey
func (game *Game) Input() *tcell.EventKey {
	return game.input
}

// ScreenSize returns the screen size - (width, height)
func (game *Game) ScreenSize() (int, int) {

	return game.width, game.height

}
