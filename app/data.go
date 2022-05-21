package app

import (
	"github.com/pkg/errors"
	"io"
	"perun.network/go-perun/channel"
	"perun.network/perun-examples/app-channel/app/util"
)

type DominionAppData struct {
	NextActor uint8
	/*Commitments [][global.HashSize]byte
	PreImages   [][global.HashSize]byte*/
	cards     []Card                // static Card information
	CardDecks [util.NumPlayers]Deck // dynamic Card information
	//cards          map[uuid.UUID]uint8
}

// Encode encodes app data onto an io.Writer.
func (d *DominionAppData) Encode(w io.Writer) error {

	err := util.WriteUInt8(w, d.NextActor)
	if err != nil {
		return errors.WithMessage(err, "writing actor")
	}

	for i := 0; i < len(d.CardDecks); i++ {
		err := util.Write(w, &d.CardDecks[i])
		if err != nil {
			return errors.WithMessage(err, "writing card")
		}
	}

	for i := 0; i < len(d.cards); i++ {
		err := util.Write(w, &d.cards[i])
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

func (a *DominionAppData) NewInitialDeck() Deck {
	/*
		var cards = make([]uuid.UUID, util.InitialDeckSize)
		var c = NewMoneyCard()
		for i := 0; i < util.InitialMoneyCards; i++ {
			c = NewMoneyCard()
			a.cards[c.id] = c
			cards[i] = c.id
		}
		for i := util.InitialMoneyCards; i < util.InitialMoneyCards+util.InitialVictoryCards; i++ {
			c = NewVictoryCard()
			a.cards[c.id] = c
			cards[i] = c.id
		}
	*/
	return Deck{
		cardIds:  cards,
		deckSize: util.InitialDeckSize,
	}
}

func (a *DominionAppData) cardOf(ids []uint8) []Card {
	var cards []Card
	for i := 0; i < len(ids); i++ {
		cards[i] = a.cards[i]
	}
	return cards
}
