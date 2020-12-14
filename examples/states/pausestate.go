package main

import (
	t "github.com/Sheep42/terminus"
)

type PauseState struct {
	*t.State
	cs        *CustomScene
	pauseText *t.Text
}

func NewPauseState(cs *CustomScene) *PauseState {

	return &PauseState{
		State: t.NewState(),
		cs:    cs,
	}

}

func (ps *PauseState) OnEnter() {

	cs := ps.cs
	sx, sy := cs.Game().ScreenSize()

	ps.pauseText = t.NewText(sx/2, sy/2, "PAUSED")

	cs.Add(ps.pauseText)

}

func (ps *PauseState) OnExit() {

	ps.cs.Remove(ps.pauseText)

}

func (ps *PauseState) Tick(delta float64) {

	// Notice we removed the call to Scene Update()
	// This will pause the game since no entities
	// in the scene will be updated

	g := ps.cs.Game()
	i := g.Input()

	if nil != i {

		if 'p' == i.Rune() {

			// Change to RunState
			ps.cs.stateManager.ChangeState(ps.cs.runState)

		}

	}

}
