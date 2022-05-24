package app

import (
	"encoding/binary"
	"log"
	"math/rand"
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
		p.cards[i] = CardOfType(CardType(dataBytes[i]))
	}
}
func (p *Pile) Size() int {
	return len(p.cards)
}
func (p *Pile) draw(seed []byte) Card {
	log.Println(seed)
	seedAsInt := binary.BigEndian.Uint64(seed)
	rand.Seed(int64(seedAsInt))
	index := rand.Intn(len(p.cards))
	card := p.cards[index]
	p.cards[index] = Card{}
	p.resizeCards()
	return card
}
func (p *Pile) resizeCards() {
	cards := make([]Card, 1)
	for i := 0; i < len(p.cards); i++ {
		if (p.cards[i] != Card{}) {
			cards = append(cards, []Card{p.cards[i]}...)
		}
	}
	p.cards = cards
}

func (p *Pile) add(card Card) {
	p.cards = append(p.cards, []Card{card}...)
}
