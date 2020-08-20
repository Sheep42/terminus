package terminus

// EntityGroup represents a set of entities that
// are grouped together within a specified boundary
type EntityGroup struct {
	*Entity

	width    int
	height   int
	entities []IEntity
}

// NewEntityGroup creates a new EntityGroup
func NewEntityGroup(x, y, width, height int, entities []IEntity) *EntityGroup {
	return &EntityGroup{
		Entity:   NewEntity(x, y),
		width:    width,
		height:   height,
		entities: entities,
	}
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
	style := eg.Entity.scene.style

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

// GetEntity returns the first entity in the group
func (eg *EntityGroup) GetEntity() *Entity {
	return eg.Entity
}

// GetEntities returns all entities contained in
// the group
func (eg *EntityGroup) GetEntities() []IEntity {
	return eg.entities
}

// SetWidth sets the width of the EntityGroup
func (eg *EntityGroup) SetWidth(width int) {
	eg.width = width
}

// SetHeight sets the height of the EntityGroup
func (eg *EntityGroup) SetHeight(height int) {
	eg.height = height
}

// GetDimensions returns the width and height
// of the EntityGroup
func (eg *EntityGroup) GetDimensions() (int, int) {
	return eg.width, eg.height
}
