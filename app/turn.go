package app

type Turn struct {
	nextActor       uint8
	performedAction ActionTypes
	possibleActions map[ActionTypes]bool
}

func InitTurn(firstActor uint8) Turn {
	return Turn{
		nextActor:       firstActor,
		performedAction: GameInit,
		possibleActions: map[ActionTypes]bool{RngCommit: true, EndTurn: true},
	}
}

func (t *Turn) ToByte() []byte {
	var dataBytes = make([]byte, 1)
	dataBytes = append(dataBytes, t.nextActor)
	dataBytes = append(dataBytes, byte(t.performedAction))

	for k, v := range t.possibleActions {
		if v {
			dataBytes = append(dataBytes, byte(k))
		}
	}
	dataBytes[0] = byte(len(dataBytes))
	return dataBytes
}

func (t *Turn) Of(dataBytes []byte) {
	var size = dataBytes[0]
	t.nextActor = dataBytes[1]
	t.performedAction = ActionTypes(dataBytes[2])
	t.possibleActions = map[ActionTypes]bool{}

	for _, v := range dataBytes[2:size] {
		t.possibleActions[ActionTypes(v)] = true
	}
}

func (t *Turn) allowed(at ActionTypes) bool {
	return t.possibleActions[at]
}

func (t *Turn) SetAllowed(possibleActions ...ActionTypes) {
	t.possibleActions = map[ActionTypes]bool{}
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
