package app

import (
	log "github.com/sirupsen/logrus"
	"perun.network/perun-examples/app-channel/app/util"
)

type Deck struct {
	MainCardPile Pile
	HandCards    Pile
	DiscardPile  Pile
}

func (d *Deck) ToByte() []byte {
	dataBytes := append([]byte{}, d.MainCardPile.ToByte()...)
	dataBytes = append(dataBytes, d.HandCards.ToByte()...)
	dataBytes = append(dataBytes, d.DiscardPile.ToByte()...)
	return append([]byte{byte(len(dataBytes))}, dataBytes...)
}

func (d *Deck) Of(dataBytes []byte) {
	mainCardPileSize := dataBytes[0]
	d.MainCardPile.Of(dataBytes[1 : mainCardPileSize+1])

	handCardSize := dataBytes[mainCardPileSize+1]
	d.HandCards.Of(dataBytes[mainCardPileSize+2 : mainCardPileSize+2+handCardSize])

	discardPileSize := dataBytes[mainCardPileSize+handCardSize+2]
	d.DiscardPile.Of(dataBytes[mainCardPileSize+handCardSize+3 : mainCardPileSize+handCardSize+3+discardPileSize])
}

// DrawOneCard draw one card from main card pile and add it to hand Cards
func (d *Deck) DrawOneCard(seed []byte) error {
	if d.MainCardPile.Size() == 0 {
		d.mixDiscardPile()
	}

	card, err := d.MainCardPile.DrawCardBasedOnSeed(seed)

	if err != nil {
		return util.ForwardError(util.ErrorConstDECK, "DrawOneCard", err)
	}

	err = d.HandCards.AddCard(card)
	if err != nil {
		return util.ForwardError(util.ErrorConstDECK, "DrawOneCard", err)
	}

	log.Info("Card drawn")
	return nil
}

func (d *Deck) DiscardHandCards() error {
	handCards := d.HandCards
	d.HandCards.Clear()
	for _, card := range handCards.Cards {
		d.DiscardPile.AddCard(card)
	}
	return nil // TODO Checks
}

func (d *Deck) mixDiscardPile() error {
	discardPile := d.DiscardPile
	d.DiscardPile.Clear()
	d.MainCardPile = discardPile
	return nil // TODO Checks
}

// IsInitialHandDrawn show if player already drawn the correct amount of Cards
func (d *Deck) IsInitialHandDrawn() bool {
	return len(d.HandCards.Cards) >= util.InitialHandSize
}

func (d *Deck) isAllowedToDraw() bool {
	return d.IsInitialHandDrawn()
}
