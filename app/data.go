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
	rng RNG
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
			return errors.WithMessage(err, "writing card")
		}
	}
	return nil
}

// Clone returns a deep copy of the app data.
func (d *DominionAppData) Clone() channel.Data {
	_d := *d
	return &_d
}

func (d *DominionAppData) switchActor(actorIdx channel.Index) {

	if d.NextActor != util.Uint8safe(uint16(actorIdx)) {
		panic("invalid actor")
	}
	d.NextActor += +1
}

func (a *DominionAppData) Init(firstActor channel.Index) error {
	// Set first actor
	a.NextActor = uint8(firstActor)

	// Set initial decks
	for deckNum := 0; deckNum < util.NumPlayers; deckNum++ {
		for i := 0; i < util.InitialMoneyCards; i++ {
			a.CardDecks[deckNum].mainCardPile.cards = append(a.CardDecks[deckNum].mainCardPile.cards, CardOfType(MoneyCopper))
		}
		for i := 0; i < util.InitialVictoryCards; i++ {
			a.CardDecks[deckNum].mainCardPile.cards = append(a.CardDecks[deckNum].mainCardPile.cards, CardOfType(VictorySmall))
		}
	}
	return nil
}
