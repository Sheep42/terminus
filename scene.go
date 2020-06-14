package terminus

import (
	// "os"

	"github.com/gdamore/tcell"
)

type IScene interface {

	Init()
	Update()
	Draw()
	Entities() []*Entity

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

func (scene *Scene) Update() { }

func (scene *Scene) Draw() {

	screen := scene.game.screen

	if len( scene.entities ) > 0 {

		for _, entity := range scene.entities {
			entity.Draw( screen, scene.style )
		}

	}

}

func (scene *Scene) Add(entity *Entity) {

	scene.entities = append(scene.entities, entity)

}

func (scene *Scene) Game() *Game {

	return scene.game

}

func (scene *Scene) Entities() []*Entity {
	return scene.entities
}