package app

import (
	"perun.network/perun-examples/dominion-cli/app/util"
)

type Stock struct {
	CardAmounts [util.CardTypeCount]uint8
	Trash       [util.CardTypeCount]uint8
}

// ToByte create a byte representation of Stock
func (s *Stock) ToByte() []byte {
	dataBytes := make([]byte, 0)
	for _, stock := range s.CardAmounts {
		dataBytes = append(dataBytes, stock)
	}

	for _, trash := range s.Trash {
		dataBytes = append(dataBytes, trash)
	}

	return util.AppendLength(dataBytes)
}

// Of create Stock out of a bytes
func (s *Stock) Of(dataBytes []byte) {
	s.CardAmounts = [util.CardTypeCount]uint8{}

	for i := 0; i < len(s.CardAmounts); i++ {
		s.CardAmounts[i] = dataBytes[i]
	}
	for i := 0; i < len(s.Trash); i++ {
		s.Trash[i] = dataBytes[util.CardTypeCount+i]
	}
}

// Init sets up initial Stock state
func (s *Stock) Init() {
	s.CardAmounts[util.Copper-1] = util.CopperInitialStock
	s.CardAmounts[util.Silver-1] = util.SilverInitialStock
	s.CardAmounts[util.Gold-1] = util.GoldInitialStock
	s.CardAmounts[util.VictorySmall-1] = util.VictorySmallInitialStock
	s.CardAmounts[util.VictoryMid-1] = util.VictoryMidInitialStock
	s.CardAmounts[util.VictoryBig-1] = util.VictoryBigInitialStock

	for i := 6; i < util.CardTypeCount; i++ {
		s.CardAmounts[i] = 10
	}
}

func (s *Stock) DecrementBy(ct util.CardType, amount uint8) {
	s.CardAmounts[int(ct)-1] -= amount // shifted by 1 because CardType NONE(0) has no amount
}

func (s *Stock) Decrement(ct util.CardType) {
	s.DecrementBy(ct, 1)
}

func (s *Stock) GetAmount(ct util.CardType) (amount uint8) {
	return s.CardAmounts[int(ct)-1] // shifted by 1 because CardType NONE(0) has no amount
}
