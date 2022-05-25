package app

import (
	"encoding/binary"
	"math/rand"
	"perun.network/perun-examples/app-channel/app/util"
)

type Pile struct {
	cards []Card
}

func (p *Pile) ToByte() []byte {
	var pileSize = len(p.cards)
	var dataBytes = make([]byte, pileSize+1)
	dataBytes[0] = byte(pileSize)
	for i := 0; i < pileSize; i++ {
		ct := uint8(p.cards[i].cardType)
		dataBytes = append(dataBytes, ct)
	}
	return dataBytes
}

func (p *Pile) Of(dataBytes []byte) {
	p.cards = make([]Card, len(dataBytes))
	for i := 0; i < len(dataBytes); i++ {
		p.cards[i], _ = CardOfType(CardType(dataBytes[i]))
		// TODO HANDLE ERROR
	}
}

// DrawCardBasedOnSeed remove and return one card based on seed
func (p *Pile) DrawCardBasedOnSeed(seed []byte) (Card, error) {

	index, err := p.SeedToIndex(seed)
	if err != nil {
		return Card{}, util.ForwardError(util.ErrorConstPILE, "DrawCardBasedOnSeed", err)
	}

	card := p.cards[index]
	p.cards[index] = Card{}

	err = p.resizeCards()
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
	return rand.Intn(len(p.cards) - 1), nil
}

// resizeCards remove gaps in cards array
func (p *Pile) resizeCards() error {
	cards := make([]Card, 0)

	for i := 0; i < len(p.cards); i++ {
		if (p.cards[i] != Card{}) {
			cards = append(cards, []Card{p.cards[i]}...)
		}
	}

	p.cards = cards

	return nil
}

// AddCard append card to current cards
func (p *Pile) AddCard(card Card) error {
	if (card == Card{}) {
		return util.ThrowError(util.ErrorConstPILE, "AddCard", "given card was empty")
	}

	p.cards = append(p.cards, []Card{card}...)
	return nil
}
