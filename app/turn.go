package app

import (
	"perun.network/perun-examples/dominion-cli/app/util"
)

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

	actionCounter := 0
	for _, v := range t.PossibleActions {
		if v {
			actionCounter++
		}
	}

	dataBytes = append(dataBytes, byte(actionCounter)) // byte sufficient as length here!

	for k, v := range t.PossibleActions {
		if v {
			dataBytes = append(dataBytes, byte(k))
		}
	}
	dataBytes = append(dataBytes, t.Params.ToByte()...)
	return util.AppendLength(dataBytes)
}

// Of create Turn out of a bytes
func (t *Turn) Of(dataBytes []byte) {
	t.NextActor = dataBytes[0]
	t.PerformedAction = util.GeneralTypesOfActions(dataBytes[1])
	t.MandatoryPartFulfilled = util.ByteToBool(dataBytes[2])

	t.PossibleActions = [util.GeneralTypesOfActionsCount]bool{}
	lengthActions := dataBytes[3] // byte exceptionally used as legth here (see above)!

	for _, k := range dataBytes[4 : 4+lengthActions] {
		t.PossibleActions[k] = true
	}
	_, dataBytes = util.PopLength(dataBytes[4+lengthActions:])
	t.Params.Of(dataBytes)
}
