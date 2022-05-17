package app

import (
	"fmt"
	"io"

	"github.com/pkg/errors"

	"perun.network/go-perun/channel"
)

const (
	NumActionCardsInGame uint8  = 10
	NumBaseCards         uint8  = 6
	NumPlayers           uint8  = 2
	NumSupplyActionCard  uint8  = 10
	NumSupplyCopper      uint8  = 60
	NumSupplySilver      uint8  = 40
	NumSupplyGold        uint8  = 30
	NumSupplyEstate      uint8  = 24
	NumSupplyDuchy       uint8  = 12
	NumSupplyProvince    uint8  = 12
	NumMaxCirculation    uint16 = uint16(NumActionCardsInGame)*uint16(NumSupplyActionCard) +
		uint16(NumSupplyCopper) + uint16(NumSupplySilver) + uint16(NumSupplyGold) +
		uint16(NumSupplyEstate) + uint16(NumSupplyDuchy) + uint16(NumSupplyProvince)
)

type DominionAppData struct {
	NextActor uint8
	// properties of selected Action Cards (immutable)
	ActionCardsInvolved [NumActionCardsInGame]ActionCardType
	// num cards left in stock:
	//   [0 : NumActionCardsInGame]: action cards
	//   [NumActionCards : end]: base cards
	CardStock [NumActionCardsInGame + NumBaseCards]uint8
	// offsets for card assignment
	LenCardDecks  [NumPlayers]uint8
	LenCardHand   [NumPlayers]uint8
	LenCardTrashs [NumPlayers]uint8
	LenCardGrave  uint8 // burned cards
	// cards in circulation (owner + position derived from offsets/lengths above)
	CardsInCirculation [NumMaxCirculation]int8 // (>0): pos in ActionCardsInvolved+1, (0): unknown/empty, (<0): base cards
}

// get the deck for player by id (enumerated starting at player 1)
func (d *DominionAppData) getDeck(p uint8) ([]int8, error) {
	if p == 0 || p > NumPlayers {
		return nil, fmt.Errorf("invalid player p")
	}
	ppos := p - 1
	len := uint16(d.LenCardDecks[ppos])
	var offset uint16 = 0
	for i := uint8(0); i < ppos; i++ {
		offset += uint16(d.LenCardDecks[i])
	}

	return d.CardsInCirculation[offset : offset+len], nil
}

// Encode encodes app data onto an io.Writer.
func (d *DominionAppData) Encode(w io.Writer) error {

	err := writeUInt8(w, d.NextActor)
	if err != nil {
		return errors.WithMessage(err, "writing actor")
	}

	for _, v := range d.ActionCardsInvolved {
		err = writeActionCard(w, v)
		if err != nil {
			return errors.WithMessage(err, "writing action card type")
		}
	}

	for i := 0; i < int(NumPlayers); i++ {
		err = writeUInt8(w, d.LenCardDecks[i])
		if err != nil {
			return errors.WithMessage(err, "writing LenCardDecks")
		}
		err = writeUInt8(w, d.LenCardHand[i])
		if err != nil {
			return errors.WithMessage(err, "writing LenCardHand")
		}
		err = writeUInt8(w, d.LenCardTrashs[i])
		if err != nil {
			return errors.WithMessage(err, "writing LenCardTrashs")
		}
	}
	err = writeUInt8(w, d.LenCardGrave)
	if err != nil {
		return errors.WithMessage(err, "writing LenCardGrave")
	}

	for _, v := range d.CardsInCirculation {
		err = writeInt8(w, v)
		if err != nil {
			return errors.WithMessage(err, "writing CardsInCirculation")
		}
	}

	return nil
}

// Clone returns a deep copy of the app data.
func (d *DominionAppData) Clone() channel.Data {
	_d := *d
	return &_d
}

func (d *DominionAppData) Set(actorIdx channel.Index) {

	if d.NextActor != uint8safe(uint16(actorIdx)) {
		panic("invalid actor")
	}
	d.NextActor += +1
}

/*
func CalcNextActor(actor uint8) uint8 {
	return (actor + 1) % numParts
}*/
/*
func (d *DominionAppData) AddCard(c Card) {
	d.AllCards[d.NumAllCards] = c
	d.NumAllCards += 1
}*/
