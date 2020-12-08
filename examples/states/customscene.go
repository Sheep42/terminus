package main

import (
	t "terminus"

	"github.com/gdamore/tcell"
)

type CustomScene struct {
	*t.Scene
	collidables  []t.IEntity
	player       *Moveable
	stateManager *t.StateManager
	runState     *RunState
	pauseState   *PauseState
}

func NewCustomScene(g *t.Game, fg, bg tcell.Color) *CustomScene {

	cs := &CustomScene{
		// NewSceneCustom is like NewScene, but allows
		// custom foreground and background colors
		Scene: t.NewSceneCustom(g, fg, bg),
	}

	cs.runState = NewRunState(cs)
	cs.pauseState = NewPauseState(cs)
	cs.stateManager = t.NewStateManager(cs.runState)

	return cs

}

func (cs *CustomScene) Setup() {

	cs.Scene.Setup() // super

	// Define and add collidables to the scene
	cs.collidables = []t.IEntity{
		t.NewSpriteEntity(25, 5, '#'), t.NewSpriteEntity(26, 5, '#'), t.NewSpriteEntity(27, 5, '#'), t.NewSpriteEntity(28, 5, '#'), t.NewSpriteEntity(29, 5, '#'),
		t.NewSpriteEntity(25, 6, '#'), t.NewSpriteEntity(26, 6, '*'), t.NewSpriteEntity(27, 6, '*'), t.NewSpriteEntity(28, 6, '*'), t.NewSpriteEntity(29, 6, '#'),
		t.NewSpriteEntity(25, 7, '#'), t.NewSpriteEntity(26, 7, '*'), t.NewSpriteEntity(27, 7, '*'), t.NewSpriteEntity(28, 7, '*'), t.NewSpriteEntity(29, 7, '#'),
		t.NewSpriteEntity(25, 8, '#'), t.NewSpriteEntity(26, 8, '*'), t.NewSpriteEntity(27, 8, '*'), t.NewSpriteEntity(28, 8, '*'), t.NewSpriteEntity(29, 8, '#'),
		t.NewSpriteEntity(25, 9, '#'), t.NewSpriteEntity(26, 9, '#'), t.NewSpriteEntity(27, 9, '#'), t.NewSpriteEntity(28, 9, '#'), t.NewSpriteEntity(29, 9, '#'),
	}

	for _, c := range cs.collidables {

		cs.Add(c)

	}

	// define player
	cs.player = NewMoveable(2, 2, '@')

	// attach collidables to player
	cs.player.SetCollidables(cs.collidables)

	// add player to the scene
	cs.Add(cs.player)
	cs.Add(t.NewText(0, 0, "Press ESC to quit", t.White, t.Black))

}

func (cs *CustomScene) Update(delta float64) {

	cs.stateManager.Update(delta)

}
