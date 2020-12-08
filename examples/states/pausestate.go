package main

import (
	t "terminus"
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

	g := ps.cs.Game()
	i := g.Input()

	if nil != i {

		if 'p' == i.Rune() {

			ps.cs.stateManager.ChangeState(ps.cs.runState)

		}

	}

}
