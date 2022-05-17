package app

import (
	"fmt"
	"io"

	"github.com/pkg/errors"
	"perun.network/go-perun/channel"

	perunio "perun.network/go-perun/pkg/io"
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

type Data struct {
	NextActor           uint8
	ActionCardsInvolved [NumActionCardsInGame]ActionCardType
	CardStock           [NumActionCardsInGame + NumBaseCards]uint8
	LenCardDecks        [NumPlayers]uint8
	LenCardHand         [NumPlayers]uint8
	LenCardTrashs       [NumPlayers]uint8
	LenCardGrave        uint8
	CardsInCirculation  [NumMaxCirculation]CardName // alternative approach: pos in ActionCardsInvolved
	// NumAllCards         uint8
	// AllCards    [256]Card
	perunio.Encoder
}

// TODO design an interface instead of CardName?
func (d *Data) getDeck(p uint8) ([]CardName, error) {
	if p == 0 || p > NumPlayers {
		return nil, fmt.Errorf("Invalid player p")
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
func (d *Data) Encode(w io.Writer) error {
	err := writeUInt8(w, d.NextActor)
	if err != nil {
		return errors.WithMessage(err, "writing actor")
	}
	err = writeUInt8(w, d.NumAllCards)
	if err != nil {
		return errors.WithMessage(err, "writing NumAllCards")
	}
	err = writeCards(w, d.AllCards)
	return errors.WithMessage(err, "writing grid")
}

// Clone returns a deep copy of the app data.
func (d *Data) Clone() channel.Data {
	_d := *d
	return &_d
}

func CalcNextActor(actor uint8) uint8 {
	return (actor + 1) % numParts
}

func (d *Data) AddCard(c Card) {
	d.AllCards[d.NumAllCards] = c
	d.NumAllCards += 1
}
