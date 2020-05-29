package terminus

import (
	"os"

	"github.com/gdamore/tcell"
)

type Scene interface {

	Init()
	Update()
	Draw()
	Add( *Entity )
	Game() *Game

}

type BasicScene struct {

	game *Game

	foreground tcell.Color
	background tcell.Color

	entities []*Entity

}

func NewScene( game *Game ) *BasicScene {

	scene := &BasicScene{
		game,
		WHITE,
		BLACK,
		[]*Entity{},
	}

	return scene

}

func NewSceneCustom(game *Game, fg, bg tcell.Color) *BasicScene {

	scene := &BasicScene{
		game,
		fg,
		bg,
		[]*Entity{},
	}

	return scene

}

func (scene *BasicScene) Init() {

	screen := scene.game.screen

	screen_style := tcell.StyleDefault.
		Foreground(scene.foreground).
		Background(scene.background)

	screen.SetStyle(screen_style)

}

func (scene *BasicScene) Update() {

	// TODO: Remove test polling
	screen := scene.game.screen

	ev := screen.PollEvent()

	switch ev := ev.(type) {

	case *tcell.EventResize:
		screen.Sync()
	case *tcell.EventKey:
		if ev.Key() == tcell.KeyEscape {
			screen.Fini()
			os.Exit(0)
		}
	default:

	}

}

func (scene *BasicScene) Draw() {

	screen := scene.game.screen

	screen.Clear()
	screen.Show()

}

func (scene *BasicScene) Add(entity *Entity) {

	scene.entities = append(scene.entities, entity)

}

func (scene *BasicScene) Game() *Game {

	return scene.game

}