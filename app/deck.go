package app

import (
	"perun.network/perun-examples/app-channel/app/util"
)

type Deck struct {
	MainCardPile Pile
	HandCards    Pile
}

func (d *Deck) ToByte() []byte {
	dataBytes := append([]byte{}, d.MainCardPile.ToByte()...)
	dataBytes = append(dataBytes, d.HandCards.ToByte()...)
	return append([]byte{byte(len(dataBytes))}, dataBytes...)
}

func (d *Deck) Of(dataBytes []byte) {
	mainCardPileSize := dataBytes[0]
	d.MainCardPile.Of(dataBytes[1 : mainCardPileSize+1])

	handCardSize := dataBytes[mainCardPileSize+1]
	d.HandCards.Of(dataBytes[mainCardPileSize+2 : mainCardPileSize+2+handCardSize])
}

// DrawOneCard draw one card from main card pile and add it to hand Cards
func (d *Deck) DrawOneCard(seed []byte) error {
	card, err := d.MainCardPile.DrawCardBasedOnSeed(seed)

	if err != nil {
		return util.ForwardError(util.ErrorConstDECK, "DrawOneCard", err)
	}

	err = d.HandCards.AddCard(card)
	if err != nil {
		return util.ForwardError(util.ErrorConstDECK, "DrawOneCard", err)
	}

	return nil
}

// InitialHandDrawn show if player already drawn the correct amount of Cards
func (d *Deck) isInitialHandDrawn() bool {
	return len(d.HandCards.Cards) >= util.InitialHandSize
}

func (d *Deck) isAllowedToDraw() bool {
	return d.isInitialHandDrawn()
}
