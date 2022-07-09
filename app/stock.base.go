package app

import (
	"fmt"

	"perun.network/perun-examples/dominion-cli/app/util"
)

// TakeOffOneInitialDeck returns an initial list of cards
func (s *Stock) TakeOffOneInitialDeck() (Pile, error) {
	errorInfo := util.ErrorInfo{FunctionName: "TakeOffOneInitialDeck", FileName: util.ErrorConstStock}

	var pile Pile
	pile.Init()

	for i := 0; i < util.InitialMoneyCards; i++ {
		card := Card{}
		card.Of([]byte{byte(util.Copper)})
		err := pile.AddCardToPile(card)
		if err != nil {
			return Pile{}, errorInfo.ForwardError(err)
		}
	}

	for i := 0; i < util.InitialVictoryCards; i++ {
		card := Card{}
		card.Of([]byte{byte(util.VictorySmall)})
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
	errorInfo := util.ErrorInfo{FunctionName: "TakeOffCard", FileName: util.ErrorConstStock}
	if s.CardAmounts[cardType] <= 0 {
		return Card{}, errorInfo.ThrowError(fmt.Sprint("No more cards of Type %T available", cardType))
	}
	s.CardAmounts[cardType]--
	card := Card{}
	card.Of([]byte{byte(cardType)})
	return card, nil
}

// TrashCard trash given card
func (s *Stock) TrashCard(cardType util.CardType) error {
	s.Trash[cardType]++
	return nil
}

// EmptyCardSets count amount of empty cardTypes
func (s *Stock) EmptyCardSets() uint8 {
	var count = uint8(0)
	for _, amount := range s.CardAmounts {
		if amount == 0 {
			count++
		}
	}
	return count
}

// IsBigVictoryCardEmpty check if VictoryBig stock is empty
func (s *Stock) IsBigVictoryCardEmpty() bool {
	return s.CardAmounts[util.VictoryBig] <= 0
}
