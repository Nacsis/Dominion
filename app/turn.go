package app

import "log"

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
	dataBytes = append(dataBytes, byte(len(t.possibleActions)))
	var list = make([]byte, 0)
	/*// TODO Works without this
	for k, v := range t.possibleActions {
		if v {
			list = append(list, byte(k))
		}
	}*/
	dataBytes = append(dataBytes, list...)
	dataBytes[0] = byte(len(dataBytes))
	return dataBytes
}

func (t *Turn) Of(dataBytes []byte) {
	log.Println("bytes")
	log.Println(dataBytes)
	t.nextActor = dataBytes[0]
	t.performedAction = ActionTypes(dataBytes[1])
	t.possibleActions = map[ActionTypes]bool{RngCommit: true, EndTurn: true}
	// TODO REMOVE THAT
	for i := 3; i < int(dataBytes[2]); i++ {
		//	t.possibleActions[ActionTypes(dataBytes[i])] = true
	}

	/*
		for _, v := range dataBytes[3:size] {
			t.possibleActions[ActionTypes(v)] = true
		}*/
}

func (t *Turn) allowed(at ActionTypes) bool {
	return true //TODO REMOVE  t.possibleActions[at]
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
