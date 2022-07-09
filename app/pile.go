package app

import (
	"strings"

	"perun.network/perun-examples/dominion-cli/app/util"
)

type Pile struct {
	Cards []Card
}

// ToByte create a byte representation of Pile
func (p *Pile) ToByte() []byte {
	dataBytes := make([]byte, 0)
	for _, card := range p.Cards {
		dataBytes = append(dataBytes, card.ToByte()...)
	}
	return util.AppendLength(dataBytes)
}

// Of create Pile out of a bytes
func (p *Pile) Of(dataBytes []byte) {
	p.Cards = make([]Card, len(dataBytes))
	for i := 0; i < len(dataBytes); i++ {
		p.Cards[i].Of([]byte{dataBytes[i]})
	}
}

// Init sets up initial Pile state
func (p *Pile) Init() {
	p.Cards = make([]Card, 0)
}

func (p *Pile) Pretty() string {
	cards := make([]string, p.Length())
	for i, card := range p.Cards {
		cards[i] = card.CardType.String()
	}
	return strings.Join(cards, ", ")
}
