package app

import "perun.network/perun-examples/app-channel/app/util"

// IsActionAllowed check if given action is allowed
func (t *Turn) IsActionAllowed(action util.GeneralTypesOfActions) bool {
	return t.possibleActions[action]
}

// SetAllowed update list of allowed actions
func (t *Turn) SetAllowed(possibleActions ...util.GeneralTypesOfActions) {
	t.possibleActions = [util.GameEnd]bool{}
	for _, v := range possibleActions {
		t.possibleActions[v] = true
	}
}

// SetNextActor update next actor
func (t *Turn) SetNextActor() {
	t.nextActor = (t.nextActor + 1) % util.NumPlayers
}
