package terminus

import (
	"github.com/gdamore/tcell"
)

// IScene is the interface through which custom
// implementations of scene can be created
type IScene interface {
	Init()
	Update(delta float64)
	Draw()
	Entities() []IEntity
}

// Scene is an abstraction and expansion of the
// game's tcell Screen. It controls fg and bg color
// and offers several methods for custom implementations
type Scene struct {
	game *Game

	foreground tcell.Color
	background tcell.Color

	entities []IEntity
	style    tcell.Style
	redraw   bool
}

// NewScene creates a new Scene to be used by a Game
func NewScene(game *Game) *Scene {

	scene := &Scene{
		game,
		White,
		Black,
		[]IEntity{},
		tcell.StyleDefault,
		true,
	}

	return scene

}

// NewSceneCustom creates a new Scene with custom
// foreground and background colors
func NewSceneCustom(game *Game, fg, bg tcell.Color) *Scene {

	scene := &Scene{
		game,
		fg,
		bg,
		[]IEntity{},
		tcell.StyleDefault,
		true,
	}

	return scene

}

// Init fires during game.Init and can be overridden
func (scene *Scene) Init() {

	screen := scene.game.screen

	screenStyle := tcell.StyleDefault.
		Foreground(scene.foreground).
		Background(scene.background)

	screen.SetStyle(screenStyle)
	scene.style = screenStyle

}

// Update fires on each pass through the game loop and
// can be overridden. delta is passed in as a parameter,
// it is the time elapsed since the last pass through the loop
func (scene *Scene) Update(delta float64) {}

// Draw is fired after the scene updates on each pass through
// the game loop. It can be overridden
func (scene *Scene) Draw() {

	// only redraw when changes are tracked
	if false == scene.redraw {
		return
	}

	game := scene.game

	if len(scene.entities) > 0 {

		for _, entity := range scene.entities {
			entity.Draw()
		}

	}

	game.screen.Show()
	game.screen.Clear()

	scene.redraw = false

}

// Add adds the given entity to the scene. It should be noted
// that entities are initialized, updated, and drawn in the
// order that they have been added to the Scene
func (scene *Scene) Add(entity IEntity) {

	entity.AddEntityToScene(scene)
	scene.entities = append(scene.entities, entity)

}

// Game returns the Game associated with the scene
func (scene *Scene) Game() *Game {

	return scene.game

}

// Entities returns the slice containing all entities
// in the scene
func (scene *Scene) Entities() []IEntity {

	return scene.entities

}

// SetRedraw allows you to tell a specific scene to
// redraw (true) or not (false) on the next frame
func (scene Scene) SetRedraw(redraw bool) {

	scene.redraw = redraw

}
