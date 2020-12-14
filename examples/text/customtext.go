package main

import (
	t "github.com/Sheep42/terminus"

	"github.com/gdamore/tcell"
)

type CustomText struct {
	*t.Text
	elapsed    float64
	colors     [][]tcell.Color
	colorIndex int
}

func NewCustomText(x, y int, text string, colors [][]tcell.Color) *CustomText {

	return &CustomText{
		t.NewText(x, y, text, colors[0][0], colors[0][1]),
		0,
		colors,
		0,
	}

}

func (ct *CustomText) Update(delta float64) {

	ct.elapsed += delta

	if ct.elapsed > 0.5 {

		if ct.colorIndex < (len(ct.colors) - 1) {
			ct.colorIndex++
		} else {
			ct.colorIndex = 0
		}

		ct.SetColor(
			ct.colors[ct.colorIndex][0],
			ct.colors[ct.colorIndex][1],
		)

		ct.elapsed = 0

	}

}
