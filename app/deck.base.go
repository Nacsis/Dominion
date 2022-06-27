package app

import (
	"perun.network/perun-examples/dominion-cli/app/util"
)

// DrawCard draw a card from main card pile based on seed.
// Adds card to hand pile.
func (d *Deck) DrawCard(seed []byte) error {
	errorInfo := util.ErrorInfo{FunctionName: "DrawCard", FileName: util.ErrorConstDECK}

	if d.Resources[util.DrawableCards] <= 0 {
		return errorInfo.ThrowError("Not enough draw actions left")
	}

	if d.MainPile.Length() == 0 {
		err := d._MixAndReassignDiscardedPile()
		if err != nil {
			return errorInfo.ForwardError(err)
		}
	}

	card, err := d.MainPile.DrawCardBasedOnSeed(seed)
	if err != nil {
		return errorInfo.ForwardError(err)
	}

	err = d.HandPile.AddCardToPile(card)
	if err != nil {
		return errorInfo.ForwardError(err)
	}

	d.Resources[util.DrawableCards]--
	return nil
}

// DiscardHandCards discard hand cards to DiscardedPile
func (d *Deck) DiscardHandCards() error {
	errorInfo := util.ErrorInfo{FunctionName: "DiscardHandCards", FileName: util.ErrorConstDECK}

	handCards := d.HandPile
	d.HandPile.Clear()

	for _, card := range handCards.Cards {
		err := d.DiscardedPile.AddCardToPile(card)
		if err != nil {
			return errorInfo.ForwardError(err)
		}
	}
	return nil
}

// DiscardPlayedCards discard played cards to DiscardedPile
func (d *Deck) DiscardPlayedCards() error {
	errorInfo := util.ErrorInfo{FunctionName: "DiscardPlayedCards", FileName: util.ErrorConstDECK}

	playedPile := d.PlayedPile
	d.PlayedPile.Clear()

	for _, card := range playedPile.Cards {
		err := d.DiscardedPile.AddCardToPile(card)
		if err != nil {
			return errorInfo.ForwardError(err)
		}
	}
	return nil
}

// ResetResources reset resources
func (d *Deck) ResetResources() error {

	d.Resources[util.DrawableCards] = util.InitialDrawResources
	d.Resources[util.PlayableCards] = util.InitialPlayResources
	d.Resources[util.BuyableCards] = util.InitialBuyResources
	d.Resources[util.SpendableMoney] = util.InitialMoneyResources

	return nil
}

// _MixAndReassignDiscardedPile mix discardedPile and assign it to MainPile
func (d *Deck) _MixAndReassignDiscardedPile() error {
	discardPile := d.DiscardedPile
	d.DiscardedPile.Clear()
	d.MainPile = discardPile
	return nil
}

// GetHandCardWithIndex return card with given index of HandPile
func (d *Deck) GetHandCardWithIndex(index uint) (Card, error) {
	errorInfo := util.ErrorInfo{FunctionName: "GetHandCardWithIndex", FileName: util.ErrorConstDECK}

	card, err := d.HandPile.ViewCardWithIndex(index)
	if err != nil {
		return Card{}, errorInfo.ForwardError(err)
	}

	if card.PlayCost > d.Resources[util.PlayableCards] {
		return Card{}, errorInfo.ThrowError("Not enough play actions left")
	}

	card, err = d.HandPile.DrawCardWithIndex(index)
	if err != nil {
		return Card{}, errorInfo.ForwardError(err)
	}

	d.Resources[util.PlayableCards] -= card.PlayCost

	return card, nil
}

// MoveToPlayedPile //TODO
func (d *Deck) MoveToPlayedPile(card Card) error {
	errorInfo := util.ErrorInfo{FunctionName: "MoveToPlayedPile", FileName: util.ErrorConstDECK}
	err := d.PlayedPile.AddCardToPile(card)
	if err != nil {
		return errorInfo.ForwardError(err)
	}
	return nil
}

// MoveToHandPile //TODO
func (d *Deck) MoveToHandPile(card Card) error {
	errorInfo := util.ErrorInfo{FunctionName: "MoveToHandPile", FileName: util.ErrorConstDECK}
	err := d.HandPile.AddCardToPile(card)
	if err != nil {
		return errorInfo.ForwardError(err)
	}
	return nil
}

// MoveToDiscardPile //TODO
func (d *Deck) MoveToDiscardPile(card Card) error {
	errorInfo := util.ErrorInfo{FunctionName: "MoveToDiscardPile", FileName: util.ErrorConstDECK}
	err := d.DiscardedPile.AddCardToPile(card)
	if err != nil {
		return errorInfo.ForwardError(err)
	}
	return nil
}

// UpdateResourcesAfterPlayedCard  //TODO
func (d *Deck) UpdateResourcesAfterPlayedCard(card Card) error {

	d.Resources[util.SpendableMoney] += card.SpendableMoney
	d.Resources[util.DrawableCards] += card.Drawables
	d.Resources[util.BuyableCards] += card.Buyables
	d.Resources[util.PlayableCards] += card.Playables
	return nil
}

// BoughtCard buy one card
func (d *Deck) BoughtCard(card Card) error {
	errorInfo := util.ErrorInfo{FunctionName: "BoughtCard", FileName: util.ErrorConstDECK}
	if d.Resources[util.SpendableMoney] < card.BuyCost {
		return errorInfo.ThrowError("Not enough spendable money available")
	}
	if !d.IsBuyActionPossible() {
		return errorInfo.ThrowError("Not enough buy actions available")
	}

	d.Resources[util.SpendableMoney] -= card.BuyCost
	d.Resources[util.BuyableCards] -= 1

	err := d.MoveToDiscardPile(card)

	if err != nil {
		return errorInfo.ForwardError(err)
	}

	return nil
}

// IsInitialHandDrawn true if initial hand was fully drawn
func (d *Deck) IsInitialHandDrawn() bool {
	return d.Resources[util.DrawableCards] == 0 && d.PlayedPile.Length() == 0 && d.HandPile.Length() == util.InitialDrawResources
}

// IsPlayActionPossible true if another play action is possible
func (d *Deck) IsPlayActionPossible() bool {
	minCost, err := d.HandPile.MinimalPlayCostInDeck()
	if err != nil {
		return false
	}
	return d.Resources[util.PlayableCards] >= minCost
}

// IsDrawActionPossible true if another draw action is possible
func (d *Deck) IsDrawActionPossible() bool {
	return d.Resources[util.DrawableCards] > 0
}

// IsBuyActionPossible true if another buy action is possible
func (d *Deck) IsBuyActionPossible() bool {
	return d.Resources[util.BuyableCards] > 0 // TODO Add check when shared supply is available
}
