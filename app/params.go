package app

import "perun.network/perun-examples/dominion-cli/app/util"

type Params struct {
	MainTarget       util.CardType
	SecondLvlTarget  util.CardType
	SecondLvlIndices []uint8
}

// Init sets up initial Params state
func (p *Params) Init() {
}

// Init sets up initial Params state
func (p *Params) Clean() Params {
	return Params{
		MainTarget:      util.NONE,
		SecondLvlTarget: util.NONE,
	}
}

// ToByte create a byte representation of Params
func (p *Params) ToByte() []byte {
	dataBytes := append([]byte{}, byte(p.MainTarget))
	dataBytes = append(dataBytes, byte(p.SecondLvlTarget))
	dataBytes = append(dataBytes, p.SecondLvlIndices...)

	return util.AppendLength(dataBytes)
}

// Of create Params out of a bytes
func (p *Params) Of(dataBytes []byte) {
	p.MainTarget = util.CardType(dataBytes[0])
	p.SecondLvlTarget = util.CardType(dataBytes[1])
	p.SecondLvlIndices = make([]uint8, len(dataBytes)-2)
	copy(p.SecondLvlIndices, dataBytes[2:])
	// for i, b := range dataBytes[2:] {
	// 	p.SecondLvlIndices[i] = b
	// }
}
