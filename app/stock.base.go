package app

import (
	"fmt"
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

// TakeOffCard takes off card of given CardType
func (s *Stock) TakeOffCard(cardType util.CardType) (Card, error) {
	errorInfo := util.ErrorInfo{FunctionName: "TakeOffOneInitialDeck", FileName: util.ErrorConstStock}
	if s.CardAmounts[cardType] <= 0 {
		return Card{}, errorInfo.ThrowError(fmt.Sprint("No more cards of Type %T available", cardType))
	}
	s.CardAmounts[cardType]--
	card := Card{}
	card.Of([]byte{byte(cardType)})
	return card, nil
}
