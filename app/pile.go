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
	dataBytes := make([]byte, 0)
	for _, card := range p.Cards {
		dataBytes = append(dataBytes, card.ToByte()...)
	}
	return append([]byte{byte(len(dataBytes))}, dataBytes...)
}

func (p *Pile) Of(dataBytes []byte) {
	p.Cards = make([]Card, len(dataBytes))
	for i := 0; i < len(dataBytes); i++ {
		p.Cards[i].Of([]byte{dataBytes[i]})
	}
}
func (p *Pile) Init() {
	p.Cards = make([]Card, 0)
}

// DrawCardBasedOnSeed remove and return one card based on seed
func (p *Pile) DrawCardBasedOnSeed(seed []byte) (Card, error) {
	errorDescription := util.ErrorInfo{FunctionName: "DrawCardBasedOnSeed", FileName: util.ErrorConstPILE}

	index, err := p.SeedToIndex(seed)
	if err != nil {
		return Card{}, util.ForwardError(errorDescription, err)
	}

	card := p.Cards[index]
	p.Cards[index] = Card{}

	err = p.ResizeCards()
	if err != nil {
		// TODO Rollback ?
		return Card{}, util.ForwardError(errorDescription, err)
	}

	return card, nil
}

// GetAndRemoveCard get card and remove it
func (p *Pile) GetAndRemoveCard(index uint) (Card, error) {
	errorDescription := util.ErrorInfo{FunctionName: "GetAndRemoveCard", FileName: util.ErrorConstPILE}

	if int(index) >= p.Size() {
		return Card{}, util.ThrowError(errorDescription, "index out of range")
	}

	card := p.Cards[index]
	p.Cards[index] = Card{}

	err := p.ResizeCards()
	if err != nil {
		// TODO Rollback ?
		return Card{}, util.ForwardError(errorDescription, err)
	}

	return card, nil
}

// ViewCard get card and remove it
func (p *Pile) ViewCard(index uint) (Card, error) {
	errorDescription := util.ErrorInfo{FunctionName: "ViewCard", FileName: util.ErrorConstPILE}

	if int(index) >= p.Size() {
		return Card{}, util.ThrowError(errorDescription, "index out of range")
	}
	return p.Cards[index], nil
}

// SeedToIndex return index generated out of seed
func (p *Pile) SeedToIndex(seed []byte) (int, error) {
	errorDescription := util.ErrorInfo{FunctionName: "SeedToIndex", FileName: util.ErrorConstPILE}

	if len(seed) < 1 {
		return 0, util.ThrowError(errorDescription, "given seed has a size of 0")
	}

	seedAsInt := binary.BigEndian.Uint64(seed)
	rand.Seed(int64(seedAsInt))
	return rand.Intn(len(p.Cards)), nil
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
func (p *Pile) Clear() {
	p.Cards = make([]Card, 0)
}
func (p *Pile) Size() int {
	return len(p.Cards)
}

// AddCard append card to current Cards
func (p *Pile) AddCard(card Card) error {
	errorDescription := util.ErrorInfo{FunctionName: "AddCard", FileName: util.ErrorConstPILE}

	if (card == Card{}) {
		return util.ThrowError(errorDescription, "given card was empty")
	}

	p.Cards = append(p.Cards, []Card{card}...)
	return nil
}
func (p *Pile) MinimalPlayCostInDeck() (uint8, error) {
	errorDescription := util.ErrorInfo{FunctionName: "MinimalPlayCostInDeck", FileName: util.ErrorConstPILE}

	if p.Size() == 0 {
		util.ThrowError(errorDescription, "No cards available")
	}
	var minValue = p.Cards[0].PlayCost
	for _, card := range p.Cards {
		if card.PlayCost < minValue {
			minValue = card.PlayCost
		}
	}
	return minValue, nil
}
