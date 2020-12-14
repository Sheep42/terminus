package main

import (
	t "github.com/Sheep42/terminus"
)

type EndState struct {
	*t.State
	scene   *CustomScene
	endText *t.Text
}

func NewEndState(cs *CustomScene) *EndState {

	return &EndState{
		State: t.NewState(),
		scene: cs,
	}

}

func (es *EndState) OnEnter() {

	scene := es.scene
	sx, sy := scene.Game().ScreenSize()

	es.endText = t.NewText(sx/2, sy/2, "GAME OVER")

	scene.Add(es.endText)

	tw, _ := es.endText.GetDimensions()
	es.endText.SetPosition(es.endText.GetX()-tw/2, es.endText.GetY())

}

func (es *EndState) OnExit() {

	es.scene.Remove(es.endText)

}

func (es *EndState) Tick(delta float64) {

	// Notice we removed the call to Scene Update()
	// This will pause the game since no entities
	// in the scene will be updated

	g := es.scene.Game()
	i := g.Input()

	if nil != i {

		if t.KeyEnter == i.Key() {

			// Change to RunState
			es.scene.stateManager.ChangeState(es.scene.runState)

		}

	}

}
