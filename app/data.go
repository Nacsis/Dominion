package app

import (
	"github.com/pkg/errors"
	"io"
	"perun.network/go-perun/channel"
	"perun.network/perun-examples/app-channel/app/util"
)

type DominionAppData struct {
	NextActor uint8
	CardDecks [util.NumPlayers]Deck // dynamic Card information
	rng       RNG
}

// Encode encodes app data onto an io.Writer.
func (d *DominionAppData) Encode(w io.Writer) error {

	// Write next actor
	err := util.WriteUInt8(w, d.NextActor)
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

// Clone returns imageA deep copy of the app data.
func (d *DominionAppData) Clone() channel.Data {
	_d := *d
	return &_d
}

func (d *DominionAppData) Init(firstActor channel.Index) error {
	// Set first actor
	d.NextActor = uint8(firstActor)

	// Set initial decks
	for deckNum := 0; deckNum < util.NumPlayers; deckNum++ {
		for i := 0; i < util.InitialMoneyCards; i++ {
			card, _ := CardOfType(MoneyCopper) // TODO HAndle error
			d.CardDecks[deckNum].mainCardPile.cards = append(d.CardDecks[deckNum].mainCardPile.cards, card)
		}
		for i := 0; i < util.InitialVictoryCards; i++ {
			card, _ := CardOfType(VictorySmall) // TODO HAndle error
			d.CardDecks[deckNum].mainCardPile.cards = append(d.CardDecks[deckNum].mainCardPile.cards, card)
		}
	}
	return nil
}

// TODO JUST FOR TEST; REMOVE LATER
func (d *DominionAppData) switchActor(actorIdx channel.Index) {

	// TODO Actor check

	d.NextActor += +1
}

//------ RNG ------

// RngCommit player who wants to DrawOneCard commit to an preimage by setting corresponding image
func (d *DominionAppData) RngCommit(actorIdx channel.Index, image []byte) error {

	// TODO Actor check
	err := d.rng.Commit(image)
	if err != nil {
		return util.ForwardError(util.ErrorConstDATA, "RngCommit", err)
	}
	return nil
}

// RngTouch the player how doesn't DrawOneCard choose an image
func (d *DominionAppData) RngTouch(actorIdx channel.Index) error {

	// TODO Actor check

	err := d.rng.Touch()
	if err != nil {
		return util.ForwardError(util.ErrorConstDATA, "RngTouch", err)
	}
	return nil
}

// RngRelease player who wants to DrawOneCard publish preimage for published image
func (d *DominionAppData) RngRelease(actorIdx channel.Index, image []byte) error {

	// TODO Actor check

	err := d.rng.Release(image)
	if err != nil {
		return util.ForwardError(util.ErrorConstDATA, "RngRelease", err)
	}
	return nil
}

// DrawOneCard draws one card to the hand pile. A full rng need to be performed before.
func (d *DominionAppData) DrawOneCard(actorIdx channel.Index) error {

	// TODO Actor check

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

	return nil
}
