package app

import "perun.network/perun-examples/app-channel/app/util"

// IsActionAllowed check if given action is allowed
func (t *Turn) IsActionAllowed(action util.GeneralTypesOfActions) bool {
	return t.PossibleActions[action]
}

// SetAllowed update list of allowed actions
func (t *Turn) SetAllowed(possibleActions ...util.GeneralTypesOfActions) {
	t.PossibleActions = [util.GameEnd]bool{}
	for _, v := range possibleActions {
		t.PossibleActions[v] = true
	}
}

// SetNextActor update next Actor
func (t *Turn) SetNextActor() {
	t.NextActor = (t.NextActor + 1) % util.NumPlayers
}
