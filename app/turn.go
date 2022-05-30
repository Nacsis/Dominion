package app

type Turn struct {
	nextActor       uint8
	performedAction ActionTypes
	possibleActions [GameEnd]bool
}

func InitTurn(firstActor uint8) Turn {
	possibleActions := [GameEnd]bool{}
	possibleActions[RngCommit] = true
	possibleActions[EndTurn] = true
	return Turn{
		nextActor:       firstActor,
		performedAction: GameInit,
		possibleActions: possibleActions,
	}
}

func (t *Turn) ToByte() []byte {
	dataBytes := append([]byte{}, t.nextActor)
	dataBytes = append(dataBytes, byte(t.performedAction))

	for k, v := range t.possibleActions {
		if v {
			dataBytes = append(dataBytes, byte(k))
		}
	}

	return append([]byte{byte(len(dataBytes))}, dataBytes...)
}

func (t *Turn) Of(dataBytes []byte) {
	t.nextActor = dataBytes[0]
	t.performedAction = ActionTypes(dataBytes[1])
	t.possibleActions = [GameEnd]bool{}

	for _, k := range dataBytes[2:] {
		t.possibleActions[k] = true
	}
}

func (t *Turn) allowed(at ActionTypes) bool {
	return t.possibleActions[at]
}

func (t *Turn) SetAllowed(possibleActions ...ActionTypes) {
	t.possibleActions = [GameEnd]bool{}
	for _, v := range possibleActions {
		t.possibleActions[v] = true
	}
}

func (t *Turn) NextActor() {
	if t.nextActor == 0 {
		t.nextActor = 1
	} else {
		t.nextActor = 0
	}
}
