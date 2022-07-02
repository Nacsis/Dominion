package app

import "perun.network/perun-examples/dominion-cli/app/util"

type Turn struct {
	NextActor              uint8
	PerformedAction        util.GeneralTypesOfActions
	MandatoryPartFulfilled bool
	PossibleActions        [util.GeneralTypesOfActionsCount]bool
	Params                 Params
}

// Init sets up initial Turn state
func (t *Turn) Init(firstActor uint8) {
	possibleActions := [util.GeneralTypesOfActionsCount]bool{}
	possibleActions[util.RngCommit] = true

	t.NextActor = firstActor
	t.PerformedAction = util.GameInit
	t.PossibleActions = possibleActions
	t.MandatoryPartFulfilled = false
	t.Params.Init()
}

// ToByte create a byte representation of Turn
func (t *Turn) ToByte() []byte {
	dataBytes := append([]byte{}, t.NextActor)
	dataBytes = append(dataBytes, byte(t.PerformedAction))
	dataBytes = append(dataBytes, util.BoolToByte(t.MandatoryPartFulfilled))
	dataBytes = append(dataBytes, byte(len(t.PossibleActions)))

	for k, v := range t.PossibleActions {
		if v {
			dataBytes = append(dataBytes, byte(k))
		}
	}
	dataBytes = append(dataBytes, t.Params.ToByte()...)
	return append([]byte{byte(len(dataBytes))}, dataBytes...)
}

// Of create Turn out of a bytes
func (t *Turn) Of(dataBytes []byte) {
	t.NextActor = dataBytes[0]
	t.PerformedAction = util.GeneralTypesOfActions(dataBytes[1])
	t.MandatoryPartFulfilled = util.ByteToBool(dataBytes[2])

	t.PossibleActions = [util.GeneralTypesOfActionsCount]bool{}

	for _, k := range dataBytes[4 : 4+dataBytes[3]] {
		t.PossibleActions[k] = true
	}
	t.Params.Of(dataBytes[4+dataBytes[3]:])
}
