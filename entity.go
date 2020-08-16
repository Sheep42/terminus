package terminus

// IEntity is the interface through which custom
// implementations of Entity can be created
type IEntity interface {
	Init()
	Update(delta float64)
	Draw()
	AddEntityToScene(scene *Scene)
}

// Entity represents a simple entity to be rendered
// to the game screen
type Entity struct {
	scene *Scene
	game  *Game

	x      int
	y      int
	sprite rune
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
func NewSpriteEntity(x, y int, sprite rune) *Entity {

	entity := &Entity{
		x:      x,
		y:      y,
		sprite: sprite,
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
	style := entity.scene.style

	if 0 != entity.sprite {
		screen.SetContent(entity.x, entity.y, entity.sprite, nil, style)
	}

}

// AddEntityToScene adds the Entity to the given scene
func (entity *Entity) AddEntityToScene(scene *Scene) {

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

// SetX sets the x position of the Entity
func (entity *Entity) SetX(x int) {
	entity.x = x
}

// GetX gets the current x position of the Entity
func (entity *Entity) GetX() int {
	return entity.x
}

// SetY sets the y position of the Entity
func (entity *Entity) SetY(y int) {
	entity.y = y
}

// GetY gets the current y position of the Entity
func (entity *Entity) GetY() int {
	return entity.y
}

// SetSprite sets the Entity's sprite rune
func (entity *Entity) SetSprite(sprite rune) {
	entity.sprite = sprite
}

// GetSprite returns the rune that represents the Entity
func (entity *Entity) GetSprite() rune {
	return entity.sprite
}
