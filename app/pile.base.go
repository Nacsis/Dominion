package app

import (
	"encoding/binary"
	"math/rand"
	"perun.network/perun-examples/app-channel/app/util"
	"sort"
)

// DrawCardBasedOnSeed remove and return one card based on seed
func (p *Pile) DrawCardBasedOnSeed(seed []byte) (Card, error) {
	errorInfo := util.ErrorInfo{FunctionName: "DrawCardBasedOnSeed", FileName: util.ErrorConstPILE}

	index, err := p._SeedToIndex(seed)
	if err != nil {
		return Card{}, errorInfo.ForwardError(err)
	}

	card := p.Cards[index]

	err = p._ResizeCardsWithOutIndex(index)
	if err != nil {
		// TODO Rollback ?
		return Card{}, errorInfo.ForwardError(err)
	}

	return card, nil
}

// DrawCardWithIndex draw card with given index
func (p *Pile) DrawCardWithIndex(index uint) (Card, error) {
	errorInfo := util.ErrorInfo{FunctionName: "DrawCardWithIndex", FileName: util.ErrorConstPILE}

	if int(index) >= p.Length() {
		return Card{}, errorInfo.ThrowError("index out of range")
	}

	card := p.Cards[index]

	err := p._ResizeCardsWithOutIndex(int(index))
	if err != nil {
		return Card{}, errorInfo.ForwardError(err)
	}

	return card, nil
}

// GetAndRemoveCardWithIndices draw cards with given indices
func (p *Pile) GetAndRemoveCardWithIndices(indices []uint8) ([]Card, error) {
	errorInfo := util.ErrorInfo{FunctionName: "GetAndRemoveCardWithIndices", FileName: util.ErrorConstPILE}
	sort.Slice(indices, func(i, j int) bool { return indices[i] > indices[j] })
	cards := make([]Card, 0)

	for _, index := range indices {
		if int(index) >= p.Length() {
			return nil, errorInfo.ThrowError("index out of range")
		}
		cards = append(cards, p.Cards[index])

		err := p._ResizeCardsWithOutIndex(int(index))
		if err != nil {
			return nil, errorInfo.ForwardError(err)
		}
	}

	return cards, nil
}

// ViewCardWithIndex view card with given index without removing it of deck
func (p *Pile) ViewCardWithIndex(index uint) (Card, error) {
	errorInfo := util.ErrorInfo{FunctionName: "ViewCardWithIndex", FileName: util.ErrorConstPILE}

	if int(index) >= p.Length() {
		return Card{}, errorInfo.ThrowError("index out of range")
	}
	return p.Cards[index], nil
}

// _SeedToIndex return index generated out of seed
func (p *Pile) _SeedToIndex(seed []byte) (int, error) {
	errorInfo := util.ErrorInfo{FunctionName: "_SeedToIndex", FileName: util.ErrorConstPILE}

	if len(seed) < 1 {
		return 0, errorInfo.ThrowError("given seed has a size of 0")
	}

	seedAsInt := binary.BigEndian.Uint64(seed)
	rand.Seed(int64(seedAsInt))
	return rand.Intn(len(p.Cards)), nil
}

// _ResizeCardsWithOutIndex remove gaps in Cards array
func (p *Pile) _ResizeCardsWithOutIndex(index int) error {
	cards := make([]Card, 0)

	for i := 0; i < len(p.Cards); i++ {
		if i != index {
			//This looks stupid but I think there is a problem with the local validation of a transaction.
			//I didn't really get it but looks like the fromdata and  todata use same shared memory.
			//Therefore create we need to create a new object  new object instead of the object it self
			card := Card{}
			card.Of([]byte{byte(p.Cards[i].CardType)})
			cards = append(cards, card)
		}
	}

	p.Cards = cards
	return nil
}

// Clear reset Cards
func (p *Pile) Clear() {
	p.Init()
}

// Length return length of Cards
func (p *Pile) Length() int {
	return len(p.Cards)
}

// AddCardToPile add card to current Cards
func (p *Pile) AddCardToPile(card Card) error {
	errorInfo := util.ErrorInfo{FunctionName: "AddCardToPile", FileName: util.ErrorConstPILE}

	if (card == Card{}) {
		return errorInfo.ThrowError("given card was empty")
	}

	p.Cards = append(p.Cards, card)
	return nil
}

// MinimalPlayCostInDeck return the play cost of the card with the minimal play cost
func (p *Pile) MinimalPlayCostInDeck() (uint8, error) {
	errorInfo := util.ErrorInfo{FunctionName: "MinimalPlayCostInDeck", FileName: util.ErrorConstPILE}

	if p.Length() <= 0 {
		return 0, errorInfo.ThrowError("No cards available")
	}

	var minValue = p.Cards[0].PlayCost
	for _, card := range p.Cards {
		if card.PlayCost < minValue {
			minValue = card.PlayCost
		}
	}
	return minValue, nil
}
