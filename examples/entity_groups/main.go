package main

import (
	t "terminus"
)

func main() {

	// Create the Game
	g := t.NewGame()

	// Create the Scene
	s := t.NewSceneCustom(g, t.Black, t.Gray)

	entities := []t.IEntity{
		t.NewSpriteEntity(1, 1, '*'),
		t.NewSpriteEntity(2, 1, '.'),
		t.NewSpriteEntity(3, 1, '*'),
		t.NewSpriteEntity(4, 1, '.'),
		t.NewSpriteEntity(1, 2, '*'),
		t.NewSpriteEntity(2, 5, '.'),
		t.NewSpriteEntity(3, 3, '*'),
		t.NewSpriteEntity(4, 3, '.'),
	}

	s.Add(NewCustomEntityGroup(t.NewEntityGroup(10, 10, 5, 5, entities)))
	s.Add(t.NewText(2, 2, "Press ESC to quit", t.White, t.Black))

	// g.Init takes a slice of IScenes
	ss := []t.IScene{s}

	// Init the Game
	g.Init(ss)

	// Start the Game
	g.Start()

}
