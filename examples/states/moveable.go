package main

import (
	t "terminus"
)

type Moveable struct {
	*t.Entity
	collidables []t.IEntity
}

func NewMoveable(x, y int, sprite rune) *Moveable {

	m := &Moveable{
		Entity: t.NewSpriteEntity(x, y, sprite),
	}

	return m
}

func (m *Moveable) Update(delta float64) {

	// super
	m.Entity.Update(delta)

	game := m.GetGame()
	input := game.Input()

	// Screen Wrap
	gw, gh := game.ScreenSize()

	if m.GetX() >= gw {
		m.SetPosition(0, m.GetY())
	} else if m.GetX() < 0 {
		m.SetPosition(gw-1, m.GetY())
	}

	if m.GetY() >= gh {
		m.SetPosition(m.GetX(), 0)
	} else if m.GetY() < 0 {
		m.SetPosition(m.GetX(), gh-1)
	}

	// Moveable movement with collision
	if nil != input {

		moveX, moveY := 0, 0

		if t.KeyLeft == input.Key() {

			moveX = -1

		} else if t.KeyRight == input.Key() {

			moveX = 1

		} else if t.KeyUp == input.Key() {

			moveY = -1

		} else if t.KeyDown == input.Key() {

			moveY = 1

		}

		collided := false

		for _, c := range m.collidables {

			ce := c.GetEntity()
			ceX, ceY := ce.GetPosition()

			if m.CheckDir('x', moveX, ceX) && m.CheckDir('y', moveY, ceY) {
				collided = true
				break
			}

		}

		if false == collided {

			m.SetPosition(m.GetX()+moveX, m.GetY()+moveY)

		}

	}

}

func (m *Moveable) SetCollidables(collidables []t.IEntity) {

	m.collidables = collidables

}
