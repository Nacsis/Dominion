package app

import (
	"perun.network/perun-examples/app-channel/app/util"
)

type Deck struct {
	mainCardPile Pile
	handCards    Pile
}

func (d *Deck) ToByte() []byte {
	mainCardPileBytes := d.mainCardPile.ToByte()
	handCardsBytes := d.handCards.ToByte()

	var deckByteLength = len(mainCardPileBytes) + len(handCardsBytes) + 1 // +1 is for length
	var dataBytes = make([]byte, deckByteLength)

	dataBytes[0] = byte(deckByteLength)
	dataBytes = append(dataBytes, mainCardPileBytes...)
	dataBytes = append(dataBytes, handCardsBytes...)

	return dataBytes
}

func (d *Deck) Of(dataBytes []byte) {
	mainCardPileSize := dataBytes[0]
	d.mainCardPile.Of(dataBytes[1:mainCardPileSize])

	handCardSize := dataBytes[mainCardPileSize]
	d.handCards.Of(dataBytes[handCardSize : handCardSize+mainCardPileSize])
}

// DrawOneCard draw one card from main card pile and add it to hand cards
func (d *Deck) DrawOneCard(seed []byte) error {
	card, err := d.mainCardPile.DrawCardBasedOnSeed(seed)

	if err != nil {
		return util.ForwardError(util.ErrorConstDECK, "DrawOneCard", err)
	}

	err = d.handCards.AddCard(card)
	if err != nil {
		return util.ForwardError(util.ErrorConstDECK, "DrawOneCard", err)
	}

	return nil
}
