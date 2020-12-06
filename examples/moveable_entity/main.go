package main

import (
	t "terminus"
)

func main() {

	// Create the Game
	g := t.NewGame()

	// Create the Scene
	s := t.NewScene(g)

	// Add a new Moveable to the Scene at x: 6, y: 3
	// The Entity's sprite will be @
	s.Add(NewMoveable(6, 3, '@'))

	// g.Init takes a slice of scenes
	ss := []t.IScene{s}

	// Init the Game
	g.Init(ss)

	// Start the Game
	g.Start()

}
