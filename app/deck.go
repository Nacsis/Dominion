package app

import (
	"fmt"

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

	dataResources := make([]byte, len(d.Resources))
	copy(dataResources, d.Resources[:])
	dataResources = util.AppendLength(dataResources)

	dataBytes = append(dataBytes, dataResources...)

	return util.AppendLength(dataBytes)
}

// Of create Deck out of a bytes
func (d *Deck) Of(dataBytes []byte) {
	mainCardPileSize, dataBytes := util.PopLength(dataBytes)
	d.MainPile.Of(dataBytes[:mainCardPileSize])
	dataBytes = dataBytes[mainCardPileSize:]

	handCardSize, dataBytes := util.PopLength(dataBytes)
	d.HandPile.Of(dataBytes[:handCardSize])
	dataBytes = dataBytes[handCardSize:]

	discardPileSize, dataBytes := util.PopLength(dataBytes)
	d.DiscardedPile.Of(dataBytes[:discardPileSize])
	dataBytes = dataBytes[discardPileSize:]

	playedPileSize, dataBytes := util.PopLength(dataBytes)
	d.PlayedPile.Of(dataBytes[:playedPileSize])
	dataBytes = dataBytes[playedPileSize:]

	_, dataBytes = util.PopLength(dataBytes)
	copy(d.Resources[:], dataBytes)
}

// Init sets up initial Deck state
func (d *Deck) Init(initialMainPile Pile) error {
	d.MainPile = initialMainPile
	d.HandPile.Init()
	d.DiscardedPile.Init()
	d.PlayedPile.Init()
	d.Resources[util.DrawableCards] = util.InitialDrawResources
	d.Resources[util.PlayableCards] = util.InitialPlayResources
	d.Resources[util.BuyableCards] = util.InitialBuyResources
	d.Resources[util.SpendableMoney] = util.InitialMoneyResources
	return nil
}

func (d *Deck) Print() {
	fmt.Printf("  Hand:\t%s\n", d.HandPile.Pretty())
	fmt.Printf("  Played:\t%s\n", d.PlayedPile.Pretty())
	fmt.Printf("  Deck size:\t  Main: %v\t  Discarded: %v\t  Total: %v\n", d.MainPile.Length(), d.DiscardedPile.Length(), d.DiscardedPile.Length()+d.MainPile.Length()+d.HandPile.Length()+d.PlayedPile.Length())
	fmt.Printf("  Victory Points: %v\n", d.VictoryPointInDeck())
}
