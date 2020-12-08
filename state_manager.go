package terminus

// IStateManager is the interface through which custom
// implementations of StateManager can be created
type IStateManager interface {
	ChangeState(s *State)
	BackToDefault()
	BackToPrevious()
	Update(delta float64)
}

// StateManager manages IStates, it can be extended
// or overridden, but should provide a basic state
// machine for most general needs.
type StateManager struct {
	defaultState  *State
	previousState *State
	currentState  *State
}

// NewStateManager creates a new StateManager
func NewStateManager(defaultState *State) *StateManager {

	sm := &StateManager{
		defaultState:  defaultState,
		previousState: defaultState,
	}

	return sm

}

// ChangeState changes the current state of
// the StateManager
func (sm *StateManager) ChangeState(s *State) {

	if nil != sm.currentState {

		sm.previousState = sm.currentState
		sm.currentState.OnExit()

	}

	sm.currentState = s
	sm.currentState.OnEnter()

}

// BackToDefault resets the StateManager to
// the default State, passed at creation
func (sm *StateManager) BackToDefault() {

	sm.ChangeState(sm.defaultState)

}

// BackToPrevious sets the StateManager's
// current State back to the previous State
func (sm *StateManager) BackToPrevious() {

	sm.ChangeState(sm.previousState)

}

// Update should be called inside of the owner
// entity, or scene's Update method. This
// powers the StateManager
func (sm *StateManager) Update(delta float64) {

	// On first pass, make sure we fire OnEnter
	if nil == sm.currentState {

		sm.ChangeState(sm.defaultState)

	}

	sm.currentState.Tick()

}
