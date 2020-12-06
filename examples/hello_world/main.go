package main

import (
	t "terminus"
)

func main() {

	// Create the Game
	g := t.NewGame()

	// Create the Scene
	s := t.NewSceneCustom(g, t.Black, t.Gray)

	// Add some text

	// override scene color
	s.Add(t.NewText(2, 2, "Press ESC to quit", t.White, t.Black))

	// Inherit scene color
	s.Add(t.NewText(5, 5, "Hello World"))

	// g.Init takes a slice of IScenes
	ss := []t.IScene{s}

	// Init the Game
	g.Init(ss)

	// Start the Game
	g.Start()

}
