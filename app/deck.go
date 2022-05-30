package app

import (
	log "github.com/sirupsen/logrus"
	"perun.network/perun-examples/app-channel/app/util"
)

type Deck struct {
	MainPile      Pile // main deck
	HandPile      Pile // Hand cards
	DiscardedPile Pile // deck of discarded cards
	PlayedPile    Pile // deck of played cards
	Resources     [util.DeckResourcesCount]uint8
}

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

func (d *Deck) Init() error {
	for i := 0; i < util.InitialMoneyCards; i++ {

		card := Card{}
		card.Of([]byte{byte(MoneyCopper)}) // TODO Handle error
		d.MainPile.AddCard(card)
	}
	for i := 0; i < util.InitialVictoryCards; i++ {
		card := Card{}
		card.Of([]byte{byte(VictorySmall)}) // TODO Handle error
		d.MainPile.AddCard(card)
	}
	if d.HandPile.Size() > 0 || d.DiscardedPile.Size() > 0 || d.PlayedPile.Size() > 0 {

	}
	d.HandPile.Init()
	d.DiscardedPile.Init()
	d.PlayedPile.Init()
	d.Resources[util.DrawableCards] = util.InitialDrawResources
	d.Resources[util.PlayableCards] = util.InitialPlayResources
	d.Resources[util.PurchasableCards] = util.InitialBuyResources
	d.Resources[util.SpendableMoney] = util.InitialMoneyResources
	return nil
}

// DrawOneCard draw one card from main card pile and add it to hand Cards
func (d *Deck) DrawOneCard(seed []byte) error {
	errorDescription := util.ErrorInfo{FunctionName: "DrawOneCard", FileName: util.ErrorConstDECK}

	if d.Resources[util.DrawableCards] > 0 {
		return util.ThrowError(errorDescription, "Not enough play actions left")
	}

	if d.MainPile.Size() == 0 {
		d.mixDiscardPile()
	}

	card, err := d.MainPile.DrawCardBasedOnSeed(seed)

	if err != nil {
		return util.ForwardError(errorDescription, err)
	}

	err = d.HandPile.AddCard(card)
	if err != nil {
		return util.ForwardError(errorDescription, err)
	}

	d.Resources[util.DrawableCards]--
	log.Info("Card drawn")
	return nil
}

func (d *Deck) DiscardHandCards() error {
	handCards := d.HandPile
	d.HandPile.Clear()
	for _, card := range handCards.Cards {
		d.DiscardedPile.AddCard(card) // TODO Checks
	}
	// TODO discard playedcards
	return nil // TODO Checks
}

func (d *Deck) mixDiscardPile() error {
	discardPile := d.DiscardedPile
	d.DiscardedPile.Clear()
	d.MainPile = discardPile
	return nil // TODO Checks
}

// PlayCardWithIndex play card with given index
func (d *Deck) PlayCardWithIndex(index uint) error {
	errorDescription := util.ErrorInfo{FunctionName: "PlayCardWithIndex", FileName: util.ErrorConstDECK}

	card, err := d.HandPile.ViewCard(index)
	if err != nil {
		return util.ForwardError(errorDescription, err)
	}
	log.Println("Played card:")
	log.Println(card)
	if card.PlayCost > d.Resources[util.PlayableCards] {
		return util.ThrowError(errorDescription, "Not enough play actions left")
	}

	card, err = d.HandPile.GetAndRemoveCard(index)
	if err != nil {
		return util.ForwardError(errorDescription, err)
	}

	d.Resources[util.PlayableCards] -= card.PlayCost
	d.Resources[util.SpendableMoney] += card.Money
	d.PlayedPile.AddCard(card) // TODO Checks
	return nil
}

// IsInitialHandDrawn show if player already drawn the correct amount of Cards
func (d *Deck) IsInitialHandDrawn() bool {
	return d.Resources[util.DrawableCards] == 0 && d.PlayedPile.Size() == 0 && d.HandPile.Size() == util.InitialDrawResources
}

func (d *Deck) IsPlayActionPossible() bool {
	minCost, err := d.HandPile.MinimalPlayCostInDeck()
	if err != nil {
		return false
	}
	return d.Resources[util.PlayableCards] >= minCost
}

func (d *Deck) IsDrawActionPossible() bool {
	return d.Resources[util.DrawableCards] > 0
}

func (d *Deck) IsBuyActionPossible() bool {
	return d.Resources[util.PurchasableCards] > 0 // TODO Hier prüfen wenn shared suply
}
