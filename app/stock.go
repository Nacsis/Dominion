package app

import (
	"perun.network/perun-examples/app-channel/app/util"
)

type Stock struct {
	CardAmounts [util.CardTypeCount]uint8
}

// ToByte create a byte representation of Stock
func (s *Stock) ToByte() []byte {
	dataBytes := make([]byte, 0)
	for _, stock := range s.CardAmounts {
		dataBytes = append(dataBytes, stock)
	}
	return append([]byte{byte(len(dataBytes))}, dataBytes...)
}

// Of create Stock out of a bytes
func (s *Stock) Of(dataBytes []byte) {
	s.CardAmounts = [util.CardTypeCount]uint8{}

	for i := 0; i < len(s.CardAmounts); i++ {
		s.CardAmounts[i] = dataBytes[i]
	}
}

// Init sets up initial Stock state
func (s *Stock) Init() {
	s.CardAmounts[util.Copper] = util.CopperInitialStock
	s.CardAmounts[util.Silver] = util.SilverInitialStock
	s.CardAmounts[util.Gold] = util.GoldInitialStock
	s.CardAmounts[util.VictorySmall] = util.VictorySmallInitialStock
	s.CardAmounts[util.VictoryMid] = util.VictoryMidInitialStock
	s.CardAmounts[util.VictoryBig] = util.VictoryBigInitialStock
}
