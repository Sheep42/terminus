package terminus

import (
	"github.com/gdamore/tcell"
)

type IScene interface {
	Init()
	Update(delta float64)
	Draw()
	Entities() []IEntity
}

type Scene struct {
	game *Game

	foreground tcell.Color
	background tcell.Color

	entities []IEntity
	style    tcell.Style
}

func NewScene(game *Game) *Scene {

	scene := &Scene{
		game,
		White,
		Black,
		[]IEntity{},
		tcell.StyleDefault,
	}

	return scene

}

func NewSceneCustom(game *Game, fg, bg tcell.Color) *Scene {

	scene := &Scene{
		game,
		fg,
		bg,
		[]IEntity{},
		tcell.StyleDefault,
	}

	return scene

}

func (scene *Scene) Init() {

	screen := scene.game.screen

	screenStyle := tcell.StyleDefault.
		Foreground(scene.foreground).
		Background(scene.background)

	screen.SetStyle(screenStyle)
	scene.style = screenStyle

}

func (scene *Scene) Update(delta float64) {}

func (scene *Scene) Draw() {

	if len(scene.entities) > 0 {

		for _, entity := range scene.entities {
			entity.Draw()
		}

	}

}

func (scene *Scene) Add(entity IEntity) {

	entity.AddEntityToScene(scene)
	scene.entities = append(scene.entities, entity)

}

func (scene *Scene) Game() *Game {

	return scene.game

}

func (scene *Scene) Entities() []IEntity {
	return scene.entities
}
