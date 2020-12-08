package main

import (
	t "terminus"
)

type RunState struct {
	*t.State
	cs *CustomScene
}

func NewRunState(cs *CustomScene) *RunState {

	return &RunState{
		State: t.NewState(),
		cs:    cs,
	}

}

func (rs *RunState) OnEnter() {}
func (rs *RunState) OnExit()  {}
func (rs *RunState) Tick(delta float64) {

	rs.cs.Scene.Update(delta)

	g := rs.cs.Game()
	i := g.Input()

	if nil != i {

		if 'p' == i.Rune() {

			rs.cs.stateManager.ChangeState(rs.cs.pauseState)

		}

	}

}
