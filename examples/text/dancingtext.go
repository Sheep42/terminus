package main

import (
	t "github.com/Sheep42/terminus"

	"github.com/gdamore/tcell"
)

type DancingText struct {
	*t.Text
	elapsed float64
	mod     int
}

func NewDancingText(x, y int, text string, colors ...tcell.Color) *DancingText {

	dt := &DancingText{
		t.NewText(x, y, text),
		0,
		0,
	}

	if len(colors) == 2 {

		dt.SetColor(colors[0], colors[1])

	}

	return dt

}

func (dt *DancingText) Init() {

	_, h := dt.GetDimensions()
	dt.SetHeight(h * 2)

	for i, t := range dt.GetEntities() {

		if i%2 == 0 {

			te := t.GetEntity()
			te.SetPosition(te.GetX(), te.GetY()+1)

		}

	}

	dt.mod = 1

}

func (dt *DancingText) Update(delta float64) {

	dt.elapsed += delta

	if dt.elapsed > 0.5 {

		for i, t := range dt.GetEntities() {

			if i%2 == dt.mod {

				te := t.GetEntity()
				te.SetPosition(te.GetX(), te.GetY()+1)

			} else {

				te := t.GetEntity()
				te.SetPosition(te.GetX(), te.GetY()-1)

			}

		}

		dt.elapsed = 0

		if 0 == dt.mod {
			dt.mod = 1
		} else {
			dt.mod = 0
		}

	}

}
