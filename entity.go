package terminus

type Entity struct {

	x int
	y int
	sprite rune

}

func NewEntity( x, y int, sprite rune ) *Entity {

	entity := &Entity{
		x,
		y,
		sprite,
	}

	return entity

}