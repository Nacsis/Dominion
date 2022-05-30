package app

import "perun.network/perun-examples/app-channel/app/util"

type Turn struct {
	nextActor              uint8
	performedAction        util.GeneralTypesOfActions
	MandatoryPartFulfilled bool
	possibleActions        [util.GameEnd]bool
}

func (t *Turn) Init(firstActor uint8) {
	possibleActions := [util.GameEnd]bool{}
	possibleActions[util.RngCommit] = true

	t.nextActor = firstActor
	t.performedAction = util.GameInit
	t.possibleActions = possibleActions
	t.MandatoryPartFulfilled = false
}

func (t *Turn) ToByte() []byte {
	dataBytes := append([]byte{}, t.nextActor)
	dataBytes = append(dataBytes, byte(t.performedAction))
	dataBytes = append(dataBytes, util.BoolToByte(t.MandatoryPartFulfilled))

	for k, v := range t.possibleActions {
		if v {
			dataBytes = append(dataBytes, byte(k))
		}
	}

	return append([]byte{byte(len(dataBytes))}, dataBytes...)
}

func (t *Turn) Of(dataBytes []byte) {
	t.nextActor = dataBytes[0]
	t.performedAction = util.GeneralTypesOfActions(dataBytes[1])
	t.MandatoryPartFulfilled = util.ByteToBool(dataBytes[2])

	t.possibleActions = [util.GameEnd]bool{}

	for _, k := range dataBytes[3:] {
		t.possibleActions[k] = true
	}
}

func (t *Turn) allowed(at util.GeneralTypesOfActions) bool {
	return t.possibleActions[at]
}

func (t *Turn) SetAllowed(possibleActions ...util.GeneralTypesOfActions) {
	t.possibleActions = [util.GameEnd]bool{}
	for _, v := range possibleActions {
		t.possibleActions[v] = true
	}
}

func (t *Turn) NextActor() {
	t.nextActor = (t.nextActor + 1) % util.NumPlayers
}
