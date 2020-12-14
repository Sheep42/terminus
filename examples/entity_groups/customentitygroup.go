package main

import (
	t "github.com/Sheep42/terminus"
)

type CustomEntityGroup struct {
	*t.EntityGroup
	elapsed float64
}

func NewCustomEntityGroup(eg *t.EntityGroup) *CustomEntityGroup {

	return &CustomEntityGroup{
		eg,
		0,
	}

}

func (ceg *CustomEntityGroup) Update(delta float64) {

	g := ceg.GetGame()
	ceg.elapsed += delta

	if ceg.elapsed >= 1 {

		// Move an entire entity group as one entity
		ceg.SetPosition(ceg.GetX()+1, ceg.GetY())

		// move a single entity within a group
		eI := ceg.GetEntities()[0] // the IEntity
		e := eI.GetEntity()        // the Entity itself

		e.SetPosition(5, 5)

		ceg.elapsed = 0

	}

	sx, _ := g.ScreenSize()

	if ceg.GetX() >= sx {

		ceg.SetPosition(0, ceg.GetY())

	}

}
