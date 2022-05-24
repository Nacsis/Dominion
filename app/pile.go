package app

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
