package terminus

import (
	"github.com/gdamore/tcell"
)

// IScene is the interface through which custom
// implementations of scene can be created
type IScene interface {
	Setup()
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
		false,
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
		false,
	}

	return scene

}

// Setup fires ONLY during game.Init and it can be overridden
func (scene *Scene) Setup() {}

// Init fires during game.Init and when the scene is first
// entered. It can be overridden
func (scene *Scene) Init() {

	screen := scene.game.screen

	screenStyle := tcell.StyleDefault.
		Foreground(scene.foreground).
		Background(scene.background)

	screen.SetStyle(screenStyle)
	scene.style = screenStyle
	scene.redraw = true

}

// Update fires on each pass through the game loop and
// can be overridden. delta is passed in as a parameter,
// it is the time elapsed since the last pass through the loop
func (scene *Scene) Update(delta float64) {}

// Draw is fired after the scene updates on each pass through
// the game loop. It can be overridden
func (scene *Scene) Draw() {

	// only redraw when changes are tracked
	if true != scene.redraw {
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
	scene.redraw = true

}

// Remove removes the given entity from the scene. This preserves
// previous entity order
func (scene *Scene) Remove(entity IEntity) {

	for i, e := range scene.entities {

		if e.GetEntity() == entity.GetEntity() {

			copy(scene.entities[i:], scene.entities[i+1:])
			scene.entities[len(scene.entities)-1] = nil
			scene.entities = scene.entities[:len(scene.entities)-1]
			break

		}

	}

	scene.redraw = true

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
func (scene *Scene) SetRedraw(redraw bool) {

	scene.redraw = redraw

}
