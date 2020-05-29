package terminus

import (
	// "os"

	"github.com/gdamore/tcell"
)

type IScene interface {

	Init()
	Update()
	Draw()

}

type Scene struct {

	game *Game

	foreground tcell.Color
	background tcell.Color

	entities []*Entity
	style tcell.Style

}

func NewScene( game *Game ) *Scene {

	scene := &Scene{
		game,
		WHITE,
		BLACK,
		[]*Entity{},
		tcell.StyleDefault,
	}

	return scene

}

func NewSceneCustom(game *Game, fg, bg tcell.Color) *Scene {

	scene := &Scene{
		game,
		fg,
		bg,
		[]*Entity{},
		tcell.StyleDefault,
	}

	return scene

}

func (scene *Scene) Init() {

	screen := scene.game.screen

	screen_style := tcell.StyleDefault.
		Foreground(scene.foreground).
		Background(scene.background)

	screen.SetStyle(screen_style)
	scene.style = screen_style

}

func (scene *Scene) Update() {

	// TODO: Remove test polling
	screen := scene.game.screen

	// TODO: Break entity update out into sep function 
	if len( scene.entities ) > 0 {

		for _, entity := range scene.entities {
			entity.Update( screen )
		}

	}

	// ev := screen.PollEvent()

	// switch ev := ev.(type) {

	// case *tcell.EventResize:
	// 	screen.Sync()
	// case *tcell.EventKey:
	// 	if ev.Key() == tcell.KeyEscape {
	// 		screen.Fini()
	// 		os.Exit(0)
	// 	}
	// default:

	// }

}

func (scene *Scene) Draw() {

	screen := scene.game.screen

	screen.Clear()

	if len( scene.entities ) > 0 {

		for _, entity := range scene.entities {
			entity.Draw( screen, scene.style )
		}

	}

	screen.Show()

}

func (scene *Scene) Add(entity *Entity) {

	scene.entities = append(scene.entities, entity)

}

func (scene *Scene) Game() *Game {

	return scene.game

}