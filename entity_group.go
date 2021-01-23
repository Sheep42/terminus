package terminus

import "github.com/gdamore/tcell"

// EntityGroup represents a set of entities that
// are grouped together within a specified boundary
type EntityGroup struct {
	*Entity

	width    int
	height   int
	entities []IEntity
}

// NewEntityGroup creates a new EntityGroup
func NewEntityGroup(x, y, width, height int, entities []IEntity, colors ...tcell.Color) *EntityGroup {

	eg := &EntityGroup{
		Entity:   NewEntity(x, y),
		width:    width,
		height:   height,
		entities: entities,
	}

	eg.colors = colors

	return eg

}

// Init fires during game.Init and can be overridden
func (eg *EntityGroup) Init() {
	eg.Entity.Init() // super
}

// Update fires after the scene update on each pass
// through the game loop, and can be overridden
func (eg *EntityGroup) Update(delta float64) {
	eg.Entity.Update(delta) // super
}

// Draw fires during scene.Draw and can be overridden
func (eg *EntityGroup) Draw() {

	// override Entity.Draw
	screen := eg.Entity.game.screen
	currentScene := eg.game.CurrentScene()

	var style tcell.Style

	if len(eg.colors) == 2 {

		style = tcell.StyleDefault.
			Foreground(eg.colors[0]).
			Background(eg.colors[1])

	} else {

		style = currentScene.style

	}

	for _, eInterface := range eg.entities {

		e := eInterface.GetEntity()

		// Don't allow entities outside of the
		// boundaries of the group
		if e.x < 0 || e.x > eg.width {
			continue
		}

		if e.y < 0 || e.y > eg.height {
			continue
		}

		// Draw the entity to the screen offset
		// by the position of the group
		screen.SetContent((eg.x + e.x), (eg.y + e.y), rune(e.GetSprite()), nil, style)

	}

}

// SetScene adds the Entity to the given scene
func (eg *EntityGroup) SetScene(scene *Scene) {

	eg.Entity.game = scene.game
	eg.Entity.scene = scene

	for _, e := range eg.entities {

		e.GetEntity().SetScene(scene)

	}

	scene.redraw = true

}

// GetEntity returns the entity used for positioning
// the group
func (eg *EntityGroup) GetEntity() *Entity {
	return eg.Entity
}

// GetEntities returns all entities contained in
// the group
func (eg *EntityGroup) GetEntities() []IEntity {
	return eg.entities
}

// GetEntityAt returns the Entity at the specified index
// if idx is out of bounds ok returns false
func (eg *EntityGroup) GetEntityAt(idx int) (*Entity, bool) {

	if idx < 0 || idx >= len(eg.entities) {
		return nil, false
	}

	return eg.entities[idx].GetEntity(), true

}

// SetWidth sets the width of the EntityGroup
func (eg *EntityGroup) SetWidth(width int) {
	eg.width = width
	eg.scene.redraw = true
}

// SetHeight sets the height of the EntityGroup
func (eg *EntityGroup) SetHeight(height int) {
	eg.height = height
	eg.scene.redraw = true
}

// GetDimensions returns the width and height
// of the EntityGroup
func (eg *EntityGroup) GetDimensions() (int, int) {
	return eg.width, eg.height
}

// SetEntities sets the EntityGroup's child entities
// and re-adds the entities to the parent's scene
func (eg *EntityGroup) SetEntities(entities []IEntity) {

	eg.entities = entities

	if nil != eg.scene {

		for _, e := range eg.entities {

			e.GetEntity().SetScene(eg.scene)

		}

	}

	eg.scene.redraw = true

}
