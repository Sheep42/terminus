package main

import (
	t "github.com/Sheep42/terminus"
)

func main() {

	// Create the Game
	g := t.NewGame()

	// shared text between both scenes, hard-coded color
	quit := t.NewText(2, 2, "Press ESC to quit", t.White, t.Black)

	// Create the Scenes
	s := NewCustomScene(g, t.White, t.Blue, "Scene 1 - Press 'x' to see Scene 2")
	s2 := NewCustomScene(g, t.Black, t.Gray, "Scane 2 - Press 'z' to see Scene 1")

	// Add the quit text to both scenes
	s.Add(quit)
	s2.Add(quit)

	// g.Init takes a slice of IScenes
	ss := []t.IScene{s, s2}

	// Init the Game
	g.Init(ss)

	// Start the Game
	g.Start()

}
