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

	scene *Scene
	game *Game

	x int
	y int
	sprite rune

}

func NewEntity( x, y int, sprite rune ) *Entity {

	entity := &Entity{
		x: x,
		y: y,
		sprite: sprite,
	}

	return entity

}

func (entity *Entity) Init() {}

func (entity *Entity) Update() { }

func (entity *Entity) Draw() {

	screen := entity.game.screen
	style := entity.scene.style
	
	screen.SetContent( entity.x, entity.y, entity.sprite, nil, style )

}

func (entity *Entity) GetScene() *Scene {
	return entity.scene
}

func (entity *Entity) GetGame() *Game {
	return entity.game
}

func (entity *Entity) SetX( x int ) {
	entity.x = x
}

func (entity *Entity) GetX() int { 
	return entity.x 
}

func (entity *Entity) SetY( y int ) {
	entity.y = y
}

func (entity *Entity) GetY() int {
	return entity.y
}

func (entity *Entity) SetSprite( sprite rune ) {
	entity.sprite = sprite
}

func (entity *Entity) GetSprite() rune {
	return entity.sprite
}