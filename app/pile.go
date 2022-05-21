package app

type Pile struct {
	cards []Card
}

func (d *Pile) ToByte() []byte {
	var deckLength = len(d.cards)
	var dataBytes = make([]byte, deckLength+1)
	dataBytes[0] = byte(deckLength)
	for i := 0; i < deckLength; i++ {
		ct := uint8(d.cards[i].cardType)
		dataBytes = append(dataBytes, ct)
	}
	return dataBytes
}
