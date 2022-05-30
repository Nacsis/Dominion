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
	d.turn = InitTurn(uint8(firstActor))

	// Set initial decks
	for deckNum := 0; deckNum < util.NumPlayers; deckNum++ {
		for i := 0; i < util.InitialMoneyCards; i++ {

			card := Card{}
			card.Of([]byte{byte(MoneyCopper)}) // TODO Handle error
			d.CardDecks[deckNum].MainCardPile.AddCard(card)
		}
		for i := 0; i < util.InitialVictoryCards; i++ {
			card := Card{}
			card.Of([]byte{byte(VictorySmall)}) // TODO Handle error
			d.CardDecks[deckNum].MainCardPile.AddCard(card)
		}
	}
	return nil
}

//------------------------ RNG ------------------------

// RngCommit player who wants to DrawOneCard commit to an preimage by setting corresponding image
func (d *DominionAppData) RngCommit(actorIdx channel.Index, image []byte) error {
	//------ Checks ------
	if d.turn.nextActor != uint8(actorIdx) {
		return util.ThrowError(util.ErrorConstDATA, "RngCommit", "Wrong actor")
	}
	if !d.turn.allowed(RngCommit) {
		return util.ThrowError(util.ErrorConstDATA, "RngCommit", "RngCommit is not allowed")
	}

	//------ Perform action ------
	err := d.rng.Commit(image)
	if err != nil {
		return util.ForwardError(util.ErrorConstDATA, "RngCommit", err)
	}

	//------ Update turn ------
	d.turnAfter(RngCommit)

	return nil
}

// RngTouch the player how doesn't DrawOneCard choose an image
func (d *DominionAppData) RngTouch(actorIdx channel.Index) error {

	//------ Checks ------
	if d.turn.nextActor != uint8(actorIdx) {
		return util.ThrowError(util.ErrorConstDATA, "RngTouch", "Actor can't touch for his own rng")
	}
	if !d.turn.allowed(RngTouch) {
		return util.ThrowError(util.ErrorConstDATA, "RngCommit", "RngTouch is not allowed")
	}

	//------ Perform action ------
	err := d.rng.Touch()
	if err != nil {
		return util.ForwardError(util.ErrorConstDATA, "RngTouch", err)
	}

	//------ Update turn ------
	d.turnAfter(RngTouch)

	return nil
}

// RngRelease player who wants to DrawOneCard publish preimage for published image
func (d *DominionAppData) RngRelease(actorIdx channel.Index, image []byte) error {

	//------ Checks ------
	if d.turn.nextActor != uint8(actorIdx) {
		return util.ThrowError(util.ErrorConstDATA, "RngRelease", "Wrong actor")
	}
	if !d.turn.allowed(RngRelease) {
		return util.ThrowError(util.ErrorConstDATA, "RngRelease", "RngTouch is not allowed")
	}

	//------ Perform action ------
	err := d.rng.Release(image)
	if err != nil {
		return util.ForwardError(util.ErrorConstDATA, "RngRelease", err)
	}

	//------ Update turn ------
	d.turnAfter(RngRelease)

	return nil
}

//------------------------ Drawing ------------------------

// DrawOneCard draws one card to the hand pile. A full rng need to be performed before.
func (d *DominionAppData) DrawOneCard(actorIdx channel.Index) error {
	//------ Checks ------
	if d.turn.nextActor != uint8(actorIdx) {
		return util.ThrowError(util.ErrorConstDATA, "DrawOneCard", "Wrong actor")
	}
	if !d.turn.allowed(DrawCard) {
		return util.ThrowError(util.ErrorConstDATA, "DrawOneCard", "DrawCard is not allowed")
	}

	//------ Perform action ------
	value, err := d.rng.CalcCorrespondingValue()
	if err != nil {
		return util.ForwardError(util.ErrorConstDATA, "DrawOneCard", err)
	}

	err = d.CardDecks[actorIdx].DrawOneCard(value)
	if err != nil {
		return util.ForwardError(util.ErrorConstDATA, "DrawOneCard", err)
	}
	// Rng was used, therefore delete it
	d.rng = RNG{}

	//------ Update turn ------
	d.turnAfter(DrawCard)

	return nil
}

//------------------------ General turn mechanics ------------------------

func (d *DominionAppData) EndTurn(actorIdx channel.Index) error {
	//------ Checks ------
	if d.turn.nextActor != uint8(actorIdx) {
		return util.ThrowError(util.ErrorConstDATA, "EndTurn", "Wrong actor")
	}
	if !d.turn.allowed(EndTurn) {
		return util.ThrowError(util.ErrorConstDATA, "EndTurn", "EndTurn is not allowed")
	}

	//------ Perform action ------
	err := d.CardDecks[actorIdx].DiscardHandCards()
	if err != nil {
		return util.ForwardError(util.ErrorConstDATA, "EndTurn", err)
	}

	//------ Update turn ------
	d.turnAfter(EndTurn)

	return nil
}

// turnAfter generate turn state after performed action
func (d *DominionAppData) turnAfter(at ActionTypes) {
	switch at {
	case RngCommit:
		d.turn.performedAction = RngCommit
		d.turn.SetAllowed(RngTouch)
		d.turn.NextActor()
		break
	case RngTouch:
		d.turn.performedAction = RngTouch
		d.turn.SetAllowed(RngRelease)
		d.turn.NextActor()
		break
	case RngRelease:
		d.turn.performedAction = RngRelease
		d.turn.SetAllowed(DrawCard)
		break
	case DrawCard:
		d.turn.performedAction = DrawCard
		var allowedActions []ActionTypes
		if d.CardDecks[d.turn.nextActor].IsInitialHandDrawn() {
			allowedActions = append(allowedActions, EndTurn)
		} else {
			allowedActions = append(allowedActions, RngCommit)
		}
		d.turn.SetAllowed(allowedActions...)
		break
	case EndTurn:
		d.turn.performedAction = EndTurn
		d.turn.SetAllowed(RngCommit)
		d.turn.NextActor()
		break
	}

}

func (d *DominionAppData) EndGame(idx channel.Index) error {
	return nil // TODO endGame check
}
