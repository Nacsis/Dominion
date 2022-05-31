package app

type Pile struct {
	Cards []Card
}

// ToByte create a byte representation of Pile
func (p *Pile) ToByte() []byte {
	dataBytes := make([]byte, 0)
	for _, card := range p.Cards {
		dataBytes = append(dataBytes, card.ToByte()...)
	}
	return append([]byte{byte(len(dataBytes))}, dataBytes...)
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
