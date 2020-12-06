package terminus

// IState is the interface through which custom
// implementations of State can be created
type IState interface {
	OnEnter()
	OnExit()
	Tick(delta float64)
}

// State is an abstract struct meant to be
// extended for use with any game or entity,
// and to be managed by a StateManager
type State struct {
}

// NewState creates a new State to be used
// by a StateManager
func NewState() *State {

	s := &State{}

	return s

}

// OnEnter is fired when the state is entered.
// This should be overridden
func (state *State) OnEnter() {}

// OnExit is fired when the state is exited.
// This should be overridden
func (state *State) OnExit() {}

// Tick is fired each frame.
// This should be overridden
func (state *State) Tick() {}
