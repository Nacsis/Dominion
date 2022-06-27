package app

import (
	"perun.network/perun-examples/dominion-cli/app/util"
)

type Deck struct {
	MainPile      Pile // main deck
	HandPile      Pile // Hand cards
	DiscardedPile Pile // deck of discarded cards
	PlayedPile    Pile // deck of played cards
	Resources     [util.DeckResourcesCount]uint8
}

// ToByte create a byte representation of Deck
func (d *Deck) ToByte() []byte {
	dataBytes := append([]byte{}, d.MainPile.ToByte()...)
	dataBytes = append(dataBytes, d.HandPile.ToByte()...)
	dataBytes = append(dataBytes, d.DiscardedPile.ToByte()...)
	dataBytes = append(dataBytes, d.PlayedPile.ToByte()...)

	actionValuesBytes := []uint8{byte(len(d.Resources))}

	for _, value := range d.Resources {
		actionValuesBytes = append(actionValuesBytes, value)
	}

	dataBytes = append(dataBytes, actionValuesBytes...)

	return append([]byte{byte(len(dataBytes))}, dataBytes...)
}

// Of create Deck out of a bytes
func (d *Deck) Of(dataBytes []byte) {
	mainCardPileSize := dataBytes[0]
	d.MainPile.Of(dataBytes[1 : mainCardPileSize+1])
	dataBytes = dataBytes[mainCardPileSize+1:]

	handCardSize := dataBytes[0]
	d.HandPile.Of(dataBytes[1 : handCardSize+1])
	dataBytes = dataBytes[handCardSize+1:]

	discardPileSize := dataBytes[0]
	d.DiscardedPile.Of(dataBytes[1 : discardPileSize+1])
	dataBytes = dataBytes[discardPileSize+1:]

	playedPileSize := dataBytes[0]
	d.PlayedPile.Of(dataBytes[1 : playedPileSize+1])
	dataBytes = dataBytes[playedPileSize+1:]

	actionValuesSize := dataBytes[0]
	for i, b := range dataBytes[1 : actionValuesSize+1] {
		d.Resources[i] = b
	}
}

// Init sets up initial Deck state
func (d *Deck) Init(initialMainPile Pile) error {
	d.MainPile = initialMainPile
	d.HandPile.Init()
	d.DiscardedPile.Init()
	d.PlayedPile.Init()
	d.Resources[util.DrawableCards] = util.InitialDrawResources
	d.Resources[util.PlayableCards] = util.InitialPlayResources
	d.Resources[util.PurchasableCards] = util.InitialBuyResources
	d.Resources[util.SpendableMoney] = util.InitialMoneyResources
	return nil
}
