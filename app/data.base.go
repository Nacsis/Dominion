package app

import (
	"perun.network/go-perun/channel"
	"perun.network/perun-examples/app-channel/app/util"
)

// EndTurn ends current turn by switching actors
func (d *DominionAppData) EndTurn(actorIdx channel.Index) error {
	errorInfo := util.ErrorInfo{FunctionName: "EndTurn", FileName: util.ErrorConstDATA}

	//------ Checks ------
	if d.turn.nextActor != uint8(actorIdx) {
		return errorInfo.ThrowError("Wrong actor")
	}
	if !d.turn.IsActionAllowed(util.EndTurn) {
		return errorInfo.ThrowError("EndTurn is not IsActionAllowed")
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

	//------ Update turn ------
	d._TurnAfter(util.EndTurn)

	return nil
}

// _TurnAfter generate turn state after performed action
func (d *DominionAppData) _TurnAfter(at util.GeneralTypesOfActions) {
	switch at {
	case util.RngCommit:
		d.turn.performedAction = util.RngCommit
		d.turn.SetAllowed(util.RngTouch)
		d.turn.SetNextActor()
		break
	case util.RngTouch:
		d.turn.performedAction = util.RngTouch
		d.turn.SetAllowed(util.RngRelease)
		d.turn.SetNextActor()
		break
	case util.RngRelease:
		d.turn.performedAction = util.RngRelease
		d.turn.SetAllowed(util.DrawCard)
		break
	case util.DrawCard:
		d.turn.performedAction = util.DrawCard
		var allowedActions []util.GeneralTypesOfActions

		if !d.turn.MandatoryPartFulfilled && d.CardDecks[d.turn.nextActor].IsInitialHandDrawn() {
			d.turn.MandatoryPartFulfilled = true
		}

		if d.turn.MandatoryPartFulfilled {
			allowedActions = d._GetAllowedDeckActions()
			allowedActions = append(allowedActions, util.EndTurn)
		} else {
			allowedActions = []util.GeneralTypesOfActions{util.RngCommit}
		}

		d.turn.SetAllowed(allowedActions...)
		break
	case util.PlayCard:
		d.turn.performedAction = util.PlayCard
		allowedActions := d._GetAllowedDeckActions()
		allowedActions = append(allowedActions, util.EndTurn)
	case util.BuyCard:
		d.turn.performedAction = util.BuyCard
		allowedActions := d._GetAllowedDeckActions()
		allowedActions = append(allowedActions, util.EndTurn)
	case util.EndTurn:
		d.turn.performedAction = util.EndTurn
		d.turn.SetAllowed(util.RngCommit)
		d.turn.SetNextActor()
		break
	}
}

// _GetAllowedDeckActions helper function to avoid duplicated code
func (d *DominionAppData) _GetAllowedDeckActions() []util.GeneralTypesOfActions {

	var allowedActions []util.GeneralTypesOfActions
	currentDeck := d.CardDecks[d.turn.nextActor]
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
	return nil // TODO end game check
}

//------------------------ Decks ------------------------

// DrawCard draws one card to the hand pile.
// A rng need to be performed before.
func (d *DominionAppData) DrawCard(actorIdx channel.Index) error {
	errorInfo := util.ErrorInfo{FunctionName: "DrawCard", FileName: util.ErrorConstDATA}

	//------ Checks ------
	if d.turn.nextActor != uint8(actorIdx) {
		return errorInfo.ThrowError("Wrong actor")
	}
	if !d.turn.IsActionAllowed(util.DrawCard) {
		return errorInfo.ThrowError("DrawCard is not IsActionAllowed")
	}

	//------ Perform action ------
	value, err := d.rng.RNGValue()
	if err != nil {
		return errorInfo.ForwardError(err)
	}

	err = d.CardDecks[actorIdx].DrawCard(value)
	if err != nil {
		return errorInfo.ForwardError(err)
	}
	// Rng was used, therefore delete it
	d.rng = RNG{}

	//------ Update turn ------
	d._TurnAfter(util.DrawCard)

	return nil
}

// PlayCard plays one card of the hand pile.
func (d *DominionAppData) PlayCard(actorIdx channel.Index, index uint8) error {
	errorInfo := util.ErrorInfo{FunctionName: "PlayCard", FileName: util.ErrorConstDATA}

	//------ Checks ------
	if d.turn.nextActor != uint8(actorIdx) {
		return errorInfo.ThrowError("Wrong actor")
	}
	if !d.turn.IsActionAllowed(util.PlayCard) {
		return errorInfo.ThrowError("PlayCard is not IsActionAllowed")
	}

	//------ Perform action ------
	err := d.CardDecks[actorIdx].PlayCardWithIndex(uint(index))
	if err != nil {
		return errorInfo.ForwardError(err)
	}

	//------ Update turn ------
	d._TurnAfter(util.PlayCard)

	return nil
}

// BuyCard Buy one card for given CardType.
func (d *DominionAppData) BuyCard(actorIdx channel.Index, cardType util.CardType) error {
	errorInfo := util.ErrorInfo{FunctionName: "BuyCard", FileName: util.ErrorConstDATA}

	//------ Checks ------
	if d.turn.nextActor != uint8(actorIdx) {
		return errorInfo.ThrowError("Wrong actor")
	}
	if !d.turn.IsActionAllowed(util.BuyCard) {
		return errorInfo.ThrowError("BuyCard is not IsActionAllowed")
	}

	//------ Perform action ------
	card, err := d.stock.TakeOffCard(cardType)
	if err != nil {
		return errorInfo.ForwardError(err)
	}

	err = d.CardDecks[d.turn.nextActor].BoughtCard(card)
	if err != nil {
		return errorInfo.ForwardError(err)
	}

	//------ Update turn ------
	d._TurnAfter(util.BuyCard)

	return nil
}

//------------------------ Rng ------------------------

// RngCommit set an image for rng
// Players who want to draw a card need to start by committing to a preimage
func (d *DominionAppData) RngCommit(actorIdx channel.Index, image []byte) error {
	errorInfo := util.ErrorInfo{FunctionName: "RngCommit", FileName: util.ErrorConstDATA}

	//------ Checks ------
	if d.turn.nextActor != uint8(actorIdx) {
		return errorInfo.ThrowError("Wrong actor")
	}
	if !d.turn.IsActionAllowed(util.RngCommit) {
		return errorInfo.ThrowError("RngCommit is not IsActionAllowed")
	}

	//------ Perform action ------
	err := d.rng.Commit(image)
	if err != nil {
		return errorInfo.ForwardError(err)
	}

	//------ Update turn ------
	d._TurnAfter(util.RngCommit)

	return nil
}

// RngTouch set a second preimage for rng
// Players need accept the set image by selecting a second preimage
func (d *DominionAppData) RngTouch(actorIdx channel.Index) error {
	errorInfo := util.ErrorInfo{FunctionName: "RngTouch", FileName: util.ErrorConstDATA}

	//------ Checks ------
	if d.turn.nextActor != uint8(actorIdx) {
		return errorInfo.ThrowError("Actor can't touch for his own rng")
	}
	if !d.turn.IsActionAllowed(util.RngTouch) {
		return errorInfo.ThrowError("RngTouch is not IsActionAllowed")
	}

	//------ Perform action ------
	err := d.rng.Touch()
	if err != nil {
		return errorInfo.ForwardError(err)
	}

	//------ Update turn ------
	d._TurnAfter(util.RngTouch)

	return nil
}

// RngRelease set preimage of set image
// Players publish their preimage of the image, s.t. a shared random value can be calculated
func (d *DominionAppData) RngRelease(actorIdx channel.Index, image []byte) error {
	errorInfo := util.ErrorInfo{FunctionName: "RngRelease", FileName: util.ErrorConstDATA}

	//------ Checks ------
	if d.turn.nextActor != uint8(actorIdx) {
		return errorInfo.ThrowError("Wrong actor")
	}
	if !d.turn.IsActionAllowed(util.RngRelease) {
		return errorInfo.ThrowError("RngRelease is not IsActionAllowed")
	}

	//------ Perform action ------
	err := d.rng.Release(image)
	if err != nil {
		return errorInfo.ForwardError(err)
	}

	//------ Update turn ------
	d._TurnAfter(util.RngRelease)

	return nil
}
