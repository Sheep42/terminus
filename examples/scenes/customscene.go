package main

import (
	t "terminus"

	"github.com/gdamore/tcell"
)

type CustomScene struct {
	*t.Scene
	title *t.Text
}

func NewCustomScene(g *t.Game, fg, bg tcell.Color, title string) *CustomScene {

	cs := &CustomScene{
		// NewSceneCustom is like NewScene, but allows
		// custom foreground and background colors
		t.NewSceneCustom(g, fg, bg),
		t.NewText(0, 0, title),
	}

	return cs

}

func (cs *CustomScene) Setup() {

	cs.Scene.Setup() // super

	cs.Add(cs.title)

}

func (cs *CustomScene) Init() {

	cs.Scene.Init()

	game := cs.Game()

	screenWidth, screenHeight := game.ScreenSize()
	textWidth, textHeight := cs.title.GetDimensions()

	cs.title.SetPosition(screenWidth/2-textWidth/2, screenHeight/2-textHeight/2)

}

func (cs *CustomScene) Update(delta float64) {

	// Use z and x to cycle through scenes
	game := cs.Game()
	input := game.Input()

	if nil != input {

		if 'z' == input.Rune() {

			game.PrevScene()

		} else if 'x' == input.Rune() {

			game.NextScene()

		}

	}

}
