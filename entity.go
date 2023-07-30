package terminus

import (
	"github.com/gdamore/tcell"
)

// IEntity is the interface through which custom
// implementations of Entity can be created
type IEntity interface {
	Init()
	Update(delta float64)
	Draw()
	SetScene(scene *Scene)
	GetEntity() *Entity
	GetEntityGroup() *EntityGroup
	SetEntityGroup(group *EntityGroup)
}

// Entity represents a simple entity to be rendered
// to the game screen
type Entity struct {
	scene *Scene
	game  *Game

	x      int
	y      int
	sprite rune

	group *EntityGroup

	colors []tcell.Color
}

// NewEntity takes an x position and a y position and
// creates an Entity
func NewEntity(x, y int) *Entity {

	entity := &Entity{
		x: x,
		y: y,
	}

	return entity

}

// NewSpriteEntity takes an x position, a y position, and a rune
// to be used as a visual representation, and creates an Entity
// colors: optional - foreground, background required if used
func NewSpriteEntity(x, y int, sprite rune, colors ...tcell.Color) *Entity {

	entity := &Entity{
		x:      x,
		y:      y,
		sprite: sprite,
		colors: colors,
	}

	return entity

}

// Init fires duting game.Init and can be overridden
func (entity *Entity) Init() {}

// Update fires after the scene update on each pass
// through the game loop, and can be overridden
func (entity *Entity) Update(delta float64) {}

// Draw fires during scene.Draw and can be overridden.
// Be careful, overridding this means that you will
// need to handle rendering on your own.
func (entity *Entity) Draw() {

	screen := entity.game.screen
	game := entity.game
	currentScene := game.CurrentScene()

	var style tcell.Style

	if len(entity.colors) == 2 {

		style = tcell.StyleDefault.
			Foreground(entity.colors[0]).
			Background(entity.colors[1])

	} else {

		style = currentScene.style

	}

	if 0 != entity.sprite {
		screen.SetContent(entity.x, entity.y, entity.sprite, nil, style)
	}

}

// GetEntity returns the entity in question
func (entity *Entity) GetEntity() *Entity {
	return entity
}

// GetEntityGroup gets the EntityGroup that the Entity belongs to
// Returns nil if not part of an EntityGroup
func (entity *Entity) GetEntityGroup() *EntityGroup {
	return entity.group
}

// SetEntityGroup sets the Entity's EntityGroup
func (entity *Entity) SetEntityGroup(group *EntityGroup) {
	entity.group = group
}

// SetScene Sets the Entity's Scene and Game
func (entity *Entity) SetScene(scene *Scene) {

	entity.game = scene.game
	entity.scene = scene

}

// GetScene gets the Scene that the Entity is associated with
func (entity *Entity) GetScene() *Scene {
	return entity.scene
}

// GetGame gets the Game the the Entity is associated with
func (entity *Entity) GetGame() *Game {
	return entity.game
}

// GetX gets the current x position of the Entity
func (entity *Entity) GetX() int {
	return entity.x
}

// GetY gets the current y position of the Entity
func (entity *Entity) GetY() int {
	return entity.y
}

// SetPosition sets the entity's x and y position
// simultaneously
func (entity *Entity) SetPosition(x, y int) {
	entity.x, entity.y = x, y
	entity.scene.redraw = true
}

// GetPosition returns the entity's current x and y
// position
func (entity *Entity) GetPosition() (int, int) {
	return entity.x, entity.y
}

// SetSprite sets the Entity's sprite rune
func (entity *Entity) SetSprite(sprite rune) {
	entity.sprite = sprite
	entity.scene.redraw = true
}

// GetSprite returns the rune that represents the Entity
func (entity *Entity) GetSprite() rune {
	return entity.sprite
}

// SetColor changes the entity's style foreground and
// background colors
func (entity *Entity) SetColor(fg, bg tcell.Color) {

	entity.colors = []tcell.Color{fg, bg}
	entity.scene.redraw = true

}

// Overlaps checks if the entity overlaps the target
// entity
func (entity *Entity) Overlaps(target IEntity) bool {
	return entity.x == target.GetEntity().x && entity.y == target.GetEntity().y
}

// OverlapsPoint checks if the entity overlaps the
// specified screen point
func (entity *Entity) OverlapsPoint(x, y int) bool {
	return entity.x == x && entity.y == y
}

// CheckDir checks if the entity is the specified
// distance away from the target point
func (entity *Entity) CheckDir(axis rune, distance, point int) bool {

	if axis == 'x' {
		return (entity.x + distance) == point
	} else if axis == 'y' {
		return (entity.y + distance) == point
	}

	return false

}

// IsLeftOf checks if the entity is directly to the
// left of the target entity
func (entity *Entity) IsLeftOf(target IEntity) bool {
	return entity.y == target.GetEntity().y && entity.CheckDir('x', 1, target.GetEntity().x)
}

// IsRightOf checks if the entity is directly to the
// right of the target entity
func (entity *Entity) IsRightOf(target IEntity) bool {
	return entity.y == target.GetEntity().y && entity.CheckDir('x', -1, target.GetEntity().x)
}

// IsAbove checks if the entity is directly above
// the target entity
func (entity *Entity) IsAbove(target IEntity) bool {
	return entity.x == target.GetEntity().x && entity.CheckDir('y', 1, target.GetEntity().y)
}

// IsBelow checks if the entity is directly below
// the target entity
func (entity *Entity) IsBelow(target IEntity) bool {
	return entity.x == target.GetEntity().x && entity.CheckDir('y', -1, target.GetEntity().y)
}
