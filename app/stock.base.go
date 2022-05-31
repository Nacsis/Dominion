package app

import (
	"perun.network/perun-examples/app-channel/app/util"
)

// TakeOffOneInitialDeck returns an initial list of cards
func (s *Stock) TakeOffOneInitialDeck() (Pile, error) {
	errorInfo := util.ErrorInfo{FunctionName: "TakeOffOneInitialDeck", FileName: util.ErrorConstStock}

	var pile Pile
	pile.Init()

	for i := 0; i < util.InitialMoneyCards; i++ {
		card := Card{}
		card.Of([]byte{byte(util.Copper)}) // TODO Handle error
		err := pile.AddCardToPile(card)
		if err != nil {
			return Pile{}, errorInfo.ForwardError(err)
		}
	}

	for i := 0; i < util.InitialVictoryCards; i++ {
		card := Card{}
		card.Of([]byte{byte(util.VictorySmall)}) // TODO Handle error
		err := pile.AddCardToPile(card)
		if err != nil {
			return Pile{}, errorInfo.ForwardError(err)
		}
	}

	s.CardAmounts[util.Copper] -= util.InitialMoneyCards
	s.CardAmounts[util.VictorySmall] -= util.InitialVictoryCards
	return pile, nil
}
