package main

import (
	t "terminus"
)

func main() {

	// Create the Game
	g := t.NewGame()

	// Create the Scene
	s := t.NewScene(g)

	// override scene color
	s.Add(t.NewText(15, 15, "This example will crash, view terminus.log to see logged text..."))

	// g.Init takes a slice of IScenes
	ss := []t.IScene{s}

	// Init the Game
	g.Init(ss)

	l := g.GetLogger()

	l.Println("Hello Logs")
	l.Printf("This is formatted %s containing numbers: %d, %d, %d", "log text", 1, 2, 3)
	l.Println("You can print as many lines as you want for debugging purposes")
	l.Fatalf("Or you can crash the game and print a formatted error like this: %s Code: %d", "Oh no a spooky error occurred!", 12345)

	// Start the Game
	g.Start()

}
