package main

import (
	t "terminus"

	"github.com/gdamore/tcell"
)

type CustomScene struct {
	*t.Scene
	stateManager *t.StateManager
	runState     *RunState
	endState     *EndState
}

func NewCustomScene(g *t.Game, fg, bg tcell.Color) *CustomScene {

	cs := &CustomScene{
		// NewSceneCustom is like NewScene, but allows
		// custom foreground and background colors
		Scene: t.NewSceneCustom(g, fg, bg),
	}

	// define the States
	cs.runState = NewRunState(cs)
	cs.endState = NewEndState(cs)

	// create the StateManager, pass RunState as default
	cs.stateManager = t.NewStateManager(cs.runState)

	return cs

}

func (cs *CustomScene) Setup() {

	cs.Scene.Setup() // super

	cs.Add(t.NewText(0, 0, "Press ESC to quit", t.White, t.Black))

}

func (cs *CustomScene) Update(delta float64) {

	// Run the StateManager
	cs.stateManager.Update(delta)

}
