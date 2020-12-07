package main

import (
	t "terminus"

	"github.com/gdamore/tcell"
)

func main() {

	// Create the Game
	g := t.NewGame()

	// Create the Scene
	s := t.NewSceneCustom(g, t.Black, t.Gray)

	// override scene color
	s.Add(t.NewText(2, 2, "Press ESC to quit", t.White, t.Black))

	// Inherit scene color
	s.Add(t.NewText(5, 5, "Hello World"))

	// Extend text functionality using composition
	s.Add(NewCustomText(10, 10, "Color Changing", [][]tcell.Color{
		{t.DarkBlue, t.Green},
		{t.Black, t.Gray},
		{t.DarkRed, t.White},
		{t.White, t.Black},
	}))

	// Extend text functionality using composition
	s.Add(NewDancingText(20, 20, "MovingText"))

	// g.Init takes a slice of IScenes
	ss := []t.IScene{s}

	// Init the Game
	g.Init(ss)

	// Start the Game
	g.Start()

}
