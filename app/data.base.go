package app

import (
	"perun.network/go-perun/channel"
	"perun.network/perun-examples/app-channel/app/util"
)

// EndTurn ends current Turn by switching actors
func (d *DominionAppData) EndTurn(actorIdx channel.Index) error {
	errorInfo := util.ErrorInfo{FunctionName: "EndTurn", FileName: util.ErrorConstDATA}

	//------ Checks ------
	if d.Turn.NextActor != uint8(actorIdx) {
		return errorInfo.ThrowError("Wrong Actor")
	}
	if !d.Turn.IsActionAllowed(util.EndTurn) {
		return errorInfo.ThrowError("EndTurn is not an allowed action")
	}

	//------ Perform action ------
	err := d.CardDecks[actorIdx].DiscardHandCards()
	if err != nil {
		return errorInfo.ForwardError(err)
	}
	err = d.CardDecks[actorIdx].DiscardPlayedCards()
	if err != nil {
		return errorInfo.ForwardError(err)
	}
	err = d.CardDecks[actorIdx].ResetResources()
	if err != nil {
		return errorInfo.ForwardError(err)
	}

	//------ Update Turn ------
	d._TurnAfter(util.EndTurn)

	return nil
}

// _TurnAfter generate Turn state after performed action
func (d *DominionAppData) _TurnAfter(at util.GeneralTypesOfActions) {
	switch at {
	case util.RngCommit:
		d.Turn.PerformedAction = util.RngCommit
		d.Turn.SetAllowed(util.RngTouch)
		d.Turn.SetNextActor()
		break
	case util.RngTouch:
		d.Turn.PerformedAction = util.RngTouch
		d.Turn.SetAllowed(util.RngRelease)
		d.Turn.SetNextActor()
		break
	case util.RngRelease:
		d.Turn.PerformedAction = util.RngRelease
		d.Turn.SetAllowed(util.DrawCard)
		break
	case util.DrawCard:
		d.Turn.PerformedAction = util.DrawCard
		var allowedActions []util.GeneralTypesOfActions

		if !d.Turn.MandatoryPartFulfilled && d.CardDecks[d.Turn.NextActor].IsInitialHandDrawn() {
			d.Turn.MandatoryPartFulfilled = true
		}

		if d.Turn.MandatoryPartFulfilled {
			allowedActions = d._GetAllowedDeckActions()
			allowedActions = append(allowedActions, util.EndTurn)
		} else {
			allowedActions = []util.GeneralTypesOfActions{util.RngCommit}
		}

		d.Turn.SetAllowed(allowedActions...)
		break
	case util.PlayCard:
		d.Turn.PerformedAction = util.PlayCard
		allowedActions := d._GetAllowedDeckActions()
		allowedActions = append(allowedActions, util.EndTurn)
		d.Turn.SetAllowed(allowedActions...)
		break
	case util.BuyCard:
		d.Turn.PerformedAction = util.BuyCard
		allowedActions := d._GetAllowedDeckActions()
		allowedActions = append(allowedActions, util.EndTurn)
		d.Turn.SetAllowed(allowedActions...)
		break
	case util.EndTurn:
		d.Turn.PerformedAction = util.EndTurn
		d.Turn.MandatoryPartFulfilled = false
		d.Turn.SetAllowed(util.RngCommit)
		d.Turn.SetNextActor()
		break
	case util.GameEnd:
		d.Turn.PerformedAction = util.GameEnd
		d.Turn.MandatoryPartFulfilled = false
		d.Turn.SetAllowed()
		break
	}
}

// _GetAllowedDeckActions helper function to avoid duplicated code
func (d *DominionAppData) _GetAllowedDeckActions() []util.GeneralTypesOfActions {

	var allowedActions []util.GeneralTypesOfActions
	currentDeck := d.CardDecks[d.Turn.NextActor]
	if currentDeck.IsPlayActionPossible() {
		allowedActions = append(allowedActions, util.PlayCard)
	}
	if currentDeck.IsDrawActionPossible() {
		allowedActions = append(allowedActions, util.RngCommit)
	}
	if currentDeck.IsBuyActionPossible() {
		allowedActions = append(allowedActions, util.BuyCard)
	}

	return allowedActions
}

// EndGame ends game
func (d *DominionAppData) EndGame(idx channel.Index) error {
	d._TurnAfter(util.GameEnd)
	return nil
}

//------------------------ Decks ------------------------

// DrawCard draws one card to the hand pile.
// A Rng need to be performed before.
func (d *DominionAppData) DrawCard(actorIdx channel.Index) error {
	errorInfo := util.ErrorInfo{FunctionName: "DrawCard", FileName: util.ErrorConstDATA}

	//------ Checks ------
	if d.Turn.NextActor != uint8(actorIdx) {
		return errorInfo.ThrowError("Wrong Actor")
	}
	if !d.Turn.IsActionAllowed(util.DrawCard) {
		return errorInfo.ThrowError("DrawCard is not an allowed action")
	}

	//------ Perform action ------
	value, err := d.Rng.RNGValue()
	if err != nil {
		return errorInfo.ForwardError(err)
	}

	err = d.CardDecks[actorIdx].DrawCard(value)
	if err != nil {
		return errorInfo.ForwardError(err)
	}
	// Rng was used, therefore delete it
	d.Rng = RNG{}

	//------ Update Turn ------
	d._TurnAfter(util.DrawCard)

	return nil
}

// PlayCard plays one card of the hand pile.
func (d *DominionAppData) PlayCard(actorIdx channel.Index, index uint8) error {
	errorInfo := util.ErrorInfo{FunctionName: "PlayCard", FileName: util.ErrorConstDATA}

	//------ Checks ------
	if d.Turn.NextActor != uint8(actorIdx) {
		return errorInfo.ThrowError("Wrong Actor")
	}
	if !d.Turn.IsActionAllowed(util.PlayCard) {
		return errorInfo.ThrowError("PlayCard is not an allowed action")
	}

	//------ Perform action ------
	err := d.CardDecks[actorIdx].PlayCardWithIndex(uint(index))
	if err != nil {
		return errorInfo.ForwardError(err)
	}

	//------ Update Turn ------
	d._TurnAfter(util.PlayCard)

	return nil
}

// BuyCard Buy one card for given CardType.
func (d *DominionAppData) BuyCard(actorIdx channel.Index, cardType util.CardType) error {
	errorInfo := util.ErrorInfo{FunctionName: "BuyCard", FileName: util.ErrorConstDATA}

	//------ Checks ------
	if d.Turn.NextActor != uint8(actorIdx) {
		return errorInfo.ThrowError("Wrong Actor")
	}
	if !d.Turn.IsActionAllowed(util.BuyCard) {
		return errorInfo.ThrowError("BuyCard is not an allowed action")
	}

	//------ Perform action ------
	card, err := d.Stock.TakeOffCard(cardType)
	if err != nil {
		return errorInfo.ForwardError(err)
	}

	err = d.CardDecks[d.Turn.NextActor].BoughtCard(card)
	if err != nil {
		return errorInfo.ForwardError(err)
	}

	//------ Update Turn ------
	d._TurnAfter(util.BuyCard)

	return nil
}

//------------------------ Rng ------------------------

// RngCommit set an image for Rng
// Players who want to draw a card need to start by committing to a preimage
func (d *DominionAppData) RngCommit(actorIdx channel.Index, image []byte) error {
	errorInfo := util.ErrorInfo{FunctionName: "RngCommit", FileName: util.ErrorConstDATA}

	//------ Checks ------
	if d.Turn.NextActor != uint8(actorIdx) {
		return errorInfo.ThrowError("Wrong Actor")
	}
	if !d.Turn.IsActionAllowed(util.RngCommit) {
		return errorInfo.ThrowError("RngCommit is not an allowed action")
	}

	//------ Perform action ------
	err := d.Rng.Commit(image)
	if err != nil {
		return errorInfo.ForwardError(err)
	}

	//------ Update Turn ------
	d._TurnAfter(util.RngCommit)

	return nil
}

// RngTouch set a second preimage for Rng
// Players need accept the set image by selecting a second preimage
func (d *DominionAppData) RngTouch(actorIdx channel.Index) error {
	errorInfo := util.ErrorInfo{FunctionName: "RngTouch", FileName: util.ErrorConstDATA}

	//------ Checks ------
	if d.Turn.NextActor != uint8(actorIdx) {
		return errorInfo.ThrowError("Actor can't touch for his own Rng")
	}
	if !d.Turn.IsActionAllowed(util.RngTouch) {
		return errorInfo.ThrowError("RngTouch is not an allowed action")
	}

	//------ Perform action ------
	err := d.Rng.Touch()
	if err != nil {
		return errorInfo.ForwardError(err)
	}

	//------ Update Turn ------
	d._TurnAfter(util.RngTouch)

	return nil
}

// RngRelease set preimage of set image
// Players publish their preimage of the image, s.t. a shared random value can be calculated
func (d *DominionAppData) RngRelease(actorIdx channel.Index, image []byte) error {
	errorInfo := util.ErrorInfo{FunctionName: "RngRelease", FileName: util.ErrorConstDATA}

	//------ Checks ------
	if d.Turn.NextActor != uint8(actorIdx) {
		return errorInfo.ThrowError("Wrong Actor")
	}
	if !d.Turn.IsActionAllowed(util.RngRelease) {
		return errorInfo.ThrowError("RngRelease is not an allowed action")
	}

	//------ Perform action ------
	err := d.Rng.Release(image)
	if err != nil {
		return errorInfo.ForwardError(err)
	}

	//------ Update Turn ------
	d._TurnAfter(util.RngRelease)

	return nil
}
