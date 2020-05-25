package terminus

type Entity struct {

	x int
	y int
	width int
	height int
	sprite rune

}

func NewEntity( x, y, width, height int, sprite rune ) *Entity {

	entity := &Entity{
		x,
		y,
		width,
		height,
		sprite,
	}

	return entity

}