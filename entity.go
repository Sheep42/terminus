package terminus

import (

	"github.com/gdamore/tcell"

)

type IEntity interface {

	Init()
	Update( tcell.Screen )
	Draw( tcell.Screen, tcell.Style )

}

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

func (entity *Entity) Init() {}

func (entity *Entity) Update( screen tcell.Screen ) { }

func (entity *Entity) Draw( screen tcell.Screen, style tcell.Style ) {

	screen.SetContent( entity.x, entity.y, entity.sprite, nil, style )

}