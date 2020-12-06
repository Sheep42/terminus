package main

import (
	t "terminus"
)

type Moveable struct {
	*t.Entity
}

func NewMoveable(x, y int, sprite rune) *Moveable {

	m := &Moveable{
		t.NewSpriteEntity(x, y, sprite),
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

	// Moveable movement
	if nil != input {

		if t.KeyLeft == input.Key() {

			m.SetPosition(m.GetX()-1, m.GetY())

		} else if t.KeyRight == input.Key() {

			m.SetPosition(m.GetX()+1, m.GetY())

		} else if t.KeyUp == input.Key() {

			m.SetPosition(m.GetX(), m.GetY()-1)

		} else if t.KeyDown == input.Key() {

			m.SetPosition(m.GetX(), m.GetY()+1)

		}

	}

}
