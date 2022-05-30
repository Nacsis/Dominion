package app

import (
	"github.com/pkg/errors"
	"io"
	"perun.network/go-perun/channel"
	"perun.network/perun-examples/app-channel/app/util"
)

type DominionAppData struct {
	turn      Turn
	CardDecks [util.NumPlayers]Deck // dynamic Card information
	rng       RNG
}

// Encode encodes app data onto an io.Writer.
func (d *DominionAppData) Encode(w io.Writer) error {

	// Write turn
	err := util.Write(w, &d.turn)
	if err != nil {
		return errors.WithMessage(err, "writing actor")
	}

	// Write decks
	for i := 0; i < len(d.CardDecks); i++ {
		err := util.Write(w, &d.CardDecks[i])
		if err != nil {
			return errors.WithMessage(err, "writing deck")
		}
	}

	err = util.Write(w, &d.rng)
	if err != nil {
		return errors.WithMessage(err, "writing rng")
	}
	return nil
}

// Clone returns ImageA deep copy of the app data.
func (d *DominionAppData) Clone() channel.Data {
	_d := *d
	return &_d
}

func (d *DominionAppData) Init(firstActor channel.Index) error {
	// Set first actor
	d.turn.Init(uint8(firstActor))

	// Set initial decks
	for deckNum := 0; deckNum < util.NumPlayers; deckNum++ {
		d.CardDecks[deckNum].Init()
	}
	return nil
}

//------------------------ RNG ------------------------

// RngCommit player who wants to DrawOneCard commit to an preimage by setting corresponding image
func (d *DominionAppData) RngCommit(actorIdx channel.Index, image []byte) error {
	errorDescription := util.ErrorInfo{FunctionName: "RngCommit", FileName: util.ErrorConstDATA}

	//------ Checks ------
	if d.turn.nextActor != uint8(actorIdx) {
		return util.ThrowError(errorDescription, "Wrong actor")
	}
	if !d.turn.allowed(util.RngCommit) {
		return util.ThrowError(errorDescription, "RngCommit is not allowed")
	}

	//------ Perform action ------
	err := d.rng.Commit(image)
	if err != nil {
		return util.ForwardError(errorDescription, err)
	}

	//------ Update turn ------
	d.turnAfter(util.RngCommit)

	return nil
}

// RngTouch the player how doesn't DrawOneCard choose an image
func (d *DominionAppData) RngTouch(actorIdx channel.Index) error {
	errorDescription := util.ErrorInfo{FunctionName: "RngTouch", FileName: util.ErrorConstDATA}
	//------ Checks ------
	if d.turn.nextActor != uint8(actorIdx) {
		return util.ThrowError(errorDescription, "Actor can't touch for his own rng")
	}
	if !d.turn.allowed(util.RngTouch) {
		return util.ThrowError(errorDescription, "RngTouch is not allowed")
	}

	//------ Perform action ------
	err := d.rng.Touch()
	if err != nil {
		return util.ForwardError(errorDescription, err)
	}

	//------ Update turn ------
	d.turnAfter(util.RngTouch)

	return nil
}

// RngRelease player who wants to DrawOneCard publish preimage for published image
func (d *DominionAppData) RngRelease(actorIdx channel.Index, image []byte) error {
	errorDescription := util.ErrorInfo{FunctionName: "RngRelease", FileName: util.ErrorConstDATA}
	//------ Checks ------
	if d.turn.nextActor != uint8(actorIdx) {
		return util.ThrowError(errorDescription, "Wrong actor")
	}
	if !d.turn.allowed(util.RngRelease) {
		return util.ThrowError(errorDescription, "RngTouch is not allowed")
	}

	//------ Perform action ------
	err := d.rng.Release(image)
	if err != nil {
		return util.ForwardError(errorDescription, err)
	}

	//------ Update turn ------
	d.turnAfter(util.RngRelease)

	return nil
}

//------------------------ Drawing ------------------------

// DrawOneCard draws one card to the hand pile. A full rng need to be performed before.
func (d *DominionAppData) DrawOneCard(actorIdx channel.Index) error {
	errorDescription := util.ErrorInfo{FunctionName: "DrawOneCard", FileName: util.ErrorConstDATA}
	//------ Checks ------
	if d.turn.nextActor != uint8(actorIdx) {
		return util.ThrowError(errorDescription, "Wrong actor")
	}
	if !d.turn.allowed(util.DrawCard) {
		return util.ThrowError(errorDescription, "DrawCard is not allowed")
	}

	//------ Perform action ------
	value, err := d.rng.CalcCorrespondingValue()
	if err != nil {
		return util.ForwardError(errorDescription, err)
	}

	err = d.CardDecks[actorIdx].DrawOneCard(value)
	if err != nil {
		return util.ForwardError(errorDescription, err)
	}
	// Rng was used, therefore delete it
	d.rng = RNG{}

	//------ Update turn ------
	d.turnAfter(util.DrawCard)

	return nil
}

// PlayCard play card with index
func (d *DominionAppData) PlayCard(actorIdx channel.Index, index uint8) error {
	errorDescription := util.ErrorInfo{FunctionName: "PlayCard", FileName: util.ErrorConstDATA}
	//------ Checks ------
	if d.turn.nextActor != uint8(actorIdx) {
		return util.ThrowError(errorDescription, "Wrong actor")
	}
	if !d.turn.allowed(util.PlayCard) {
		return util.ThrowError(errorDescription, "PlayCard is not allowed")
	}

	//------ Perform action ------

	err := d.CardDecks[actorIdx].PlayCardWithIndex(uint(index))
	if err != nil {
		return util.ForwardError(errorDescription, err)
	}

	//------ Update turn ------
	d.turnAfter(util.PlayCard)

	return nil
}

//------------------------ General turn mechanics ------------------------

func (d *DominionAppData) EndTurn(actorIdx channel.Index) error {
	errorDescription := util.ErrorInfo{FunctionName: "EndTurn", FileName: util.ErrorConstDATA}
	//------ Checks ------
	if d.turn.nextActor != uint8(actorIdx) {
		return util.ThrowError(errorDescription, "Wrong actor")
	}
	if !d.turn.allowed(util.EndTurn) {
		return util.ThrowError(errorDescription, "EndTurn is not allowed")
	}

	//------ Perform action ------
	err := d.CardDecks[actorIdx].DiscardHandCards()
	if err != nil {
		return util.ForwardError(errorDescription, err)
	}

	//------ Update turn ------
	d.turnAfter(util.EndTurn)

	return nil
}

// turnAfter generate turn state after performed action
func (d *DominionAppData) turnAfter(at util.GeneralTypesOfActions) {
	switch at {
	case util.RngCommit:
		d.turn.performedAction = util.RngCommit
		d.turn.SetAllowed(util.RngTouch)
		d.turn.NextActor()
		break
	case util.RngTouch:
		d.turn.performedAction = util.RngTouch
		d.turn.SetAllowed(util.RngRelease)
		d.turn.NextActor()
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
			allowedActions = d.getAllowedDeckActions()
			allowedActions = append(allowedActions, util.EndTurn)
		} else {
			allowedActions = []util.GeneralTypesOfActions{util.RngCommit}
		}

		d.turn.SetAllowed(allowedActions...)
		break
	case util.PlayCard:
		d.turn.performedAction = util.PlayCard
		allowedActions := d.getAllowedDeckActions()
		allowedActions = append(allowedActions, util.EndTurn)
	case util.EndTurn:
		d.turn.performedAction = util.EndTurn
		d.turn.SetAllowed(util.RngCommit)
		d.turn.NextActor()
		break
	}

}

// getAllowedDeckActions helper function to avoid duplicated code
func (d *DominionAppData) getAllowedDeckActions() []util.GeneralTypesOfActions {

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

func (d *DominionAppData) EndGame(idx channel.Index) error {
	return nil // TODO endGame check
}
