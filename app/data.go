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
	Cards     []Card                // static Card information
	CardDecks [util.NumPlayers]Deck // dynamic Card information
	//Cards          map[uuid.UUID]uint8
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

	for i := 0; i < len(d.Cards); i++ {
		err := util.Write(w, &d.Cards[i])
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

func (a *DominionAppData) Init() error {
	a.CreateInitialCards()
	a.CreateInitialDecks()
	return nil
}

func (a *DominionAppData) CreateInitialDecks() error {

	for deckNum := 0; deckNum < util.NumPlayers; deckNum++ {
		for i := 0; i < util.InitialMoneyCards; i++ {
			a.CardDecks[deckNum].cards = append(a.CardDecks[deckNum].cards, a.Cards[MoneyCopper])
		}
		for i := 0; i < util.InitialVictoryCards; i++ {
			a.CardDecks[deckNum].cards = append(a.CardDecks[deckNum].cards, a.Cards[VictorySmall])
		}
	}
	return nil

}
func (a *DominionAppData) CreateInitialCards() error {
	a.Cards = make([]Card, util.NumCardTypes)
	a.Cards[MoneyCopper] = NewCardOfType(MoneyCopper)
	a.Cards[MoneySilver] = NewCardOfType(MoneySilver)
	a.Cards[MoneyGold] = NewCardOfType(MoneyGold)
	a.Cards[VictorySmall] = NewCardOfType(VictorySmall)
	a.Cards[VictoryMid] = NewCardOfType(VictoryMid)
	a.Cards[VictoryBig] = NewCardOfType(VictoryBig)
	return nil
}

func (a *DominionAppData) CardOf(ids []uint8) []Card {
	var cards []Card
	for i := 0; i < len(ids); i++ {
		cards[i] = a.Cards[i]
	}
	return cards
}
