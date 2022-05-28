package app

import (
	"encoding/binary"
	"math/rand"
	"perun.network/perun-examples/app-channel/app/util"
)

type Pile struct {
	Cards []Card
}

func (p *Pile) ToByte() []byte {
	var pileSize = len(p.Cards)
	var dataBytes = make([]byte, 0)
	dataBytes = append(dataBytes, byte(pileSize))

	for i := 0; i < pileSize; i++ {
		dataBytes = append(dataBytes, p.Cards[i].ToByte()...)
	}
	return dataBytes
}

func (p *Pile) Of(dataBytes []byte) {
	p.Cards = make([]Card, len(dataBytes))
	for i := 0; i < len(dataBytes); i++ {
		p.Cards[i].Of([]byte{dataBytes[i]})
	}
}

// DrawCardBasedOnSeed remove and return one card based on seed
func (p *Pile) DrawCardBasedOnSeed(seed []byte) (Card, error) {

	index, err := p.SeedToIndex(seed)
	if err != nil {
		return Card{}, util.ForwardError(util.ErrorConstPILE, "DrawCardBasedOnSeed", err)
	}

	card := p.Cards[index]
	p.Cards[index] = Card{}

	err = p.ResizeCards()
	if err != nil {
		// TODO Rollback ?
		return Card{}, util.ForwardError(util.ErrorConstPILE, "DrawCardBasedOnSeed", err)
	}

	return card, nil
}

// SeedToIndex return index generated out of seed
func (p *Pile) SeedToIndex(seed []byte) (int, error) {

	if len(seed) < 1 {
		return 0, util.ThrowError(util.ErrorConstPILE, "SeedToIndex", "given seed has a size of 0")
	}

	seedAsInt := binary.BigEndian.Uint64(seed)
	rand.Seed(int64(seedAsInt))
	return rand.Intn(len(p.Cards) - 1), nil
}

// ResizeCards remove gaps in Cards array
func (p *Pile) ResizeCards() error {
	cards := make([]Card, 0)

	for i := 0; i < len(p.Cards); i++ {
		if (p.Cards[i] != Card{}) {
			cards = append(cards, []Card{p.Cards[i]}...)
		}
	}

	p.Cards = cards

	return nil
}

// AddCard append card to current Cards
func (p *Pile) AddCard(card Card) error {
	if (card == Card{}) {
		return util.ThrowError(util.ErrorConstPILE, "AddCard", "given card was empty")
	}

	p.Cards = append(p.Cards, []Card{card}...)
	return nil
}
