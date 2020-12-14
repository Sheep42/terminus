package main

import (
	t "github.com/Sheep42/terminus"
)

func main() {

	// Create the Game
	g := t.NewGame()

	// Create the Scene
	s := NewCustomScene(g, t.DarkGreen, t.Black)

	// g.Init takes a slice of IScenes
	ss := []t.IScene{s}

	// Init the Game
	g.Init(ss)

	// Start the Game
	g.Start()

}
