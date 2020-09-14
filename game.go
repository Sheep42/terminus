package terminus

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/gdamore/tcell"
)

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
	logger       *log.Logger
	logFile      *os.File
	logFileName  string
	ticker       *time.Ticker
}

// NewGame creates a game
func NewGame() *Game {

	game := &Game{}

	return game

}

// Init takes an array of scenes, and sets up the game
// before the loop is started
func (game *Game) Init(scenes []IScene) {

	game.exitKey = KeyEsc
	game.logger = log.New(os.Stderr, "", log.Ldate|log.Ltime|log.LUTC|log.Lshortfile)

	if game.logFileName == "" {
		game.logFileName = "terminus.log"
	}

	baseDir, err := filepath.Abs(filepath.Dir(os.Args[0])) // baseDir = game directory
	if err != nil {
		game.logger.Fatal("Error getting baseDir: ", err)
	}

	game.logFile, err = os.OpenFile(baseDir+"/"+game.logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		game.logger.Fatal("Error opening log file: ", err)
	}

	game.logger.SetOutput(game.logFile)

	screen, err := tcell.NewScreen()
	if err != nil {
		game.logger.Fatal("Error creating screen: ", err)
	}

	game.screen = screen

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

	game.ticker = time.NewTicker(time.Duration(1000000/game.fps) * time.Microsecond)
	game.chanKeyPress = make(chan *tcell.EventKey)

	game.logger.Println("Game Init finished")
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

	game.logger.Println("Game Start running...")

	defer game.logFile.Close()

	screen := game.screen
	clock := time.Now()

	go game.getInput()

	game.width, game.height = screen.Size()

game_loop:
	for {

		update := time.Now()
		delta := update.Sub(clock).Seconds()
		clock = update

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

		// enforce fps
		select {
		case <-game.ticker.C:
			screen.Show()
			screen.Clear()
			continue
		}
	}

	game.logger.Println("Game loop exited")

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

// GetLogger gets the game's logger for
// use in your game
func (game *Game) GetLogger() *log.Logger {
	return game.logger
}

// SetLogFileName sets a custom log file name for
// logger output. Default value is 'terminus.log'
func (game *Game) SetLogFileName(filename string) {
	game.logFileName = filename
}

// Input gets the current input as an EventKey
func (game *Game) Input() *tcell.EventKey {
	return game.input
}

// ScreenSize returns the screen size - (width, height)
func (game *Game) ScreenSize() (int, int) {

	return game.width, game.height

}
