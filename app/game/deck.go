package game

const (
	InitialDeckSize     = 8
	InitialMoneyCards   = InitialDeckSize / 2
	InitialVictoryCards = InitialDeckSize / 2
)

type Deck struct {
	deckSize uint8
	cards    []Card
}

func NewInitialDeck() Deck {
	var cards = make([]Card, InitialDeckSize)

	for i := 0; i < InitialMoneyCards; i++ {
		cards[i] = NewMoneyCard()
	}
	for i := InitialMoneyCards; i < InitialMoneyCards+InitialVictoryCards; i++ {
		cards[i] = NewVictoryCard()
	}
	return Deck{
		cards:    cards,
		deckSize: InitialDeckSize,
	}
}

func (d *Deck) ToByte() []byte {
	dataBytes := []byte{d.deckSize}
	for i := 0; i < int(d.deckSize); i++ {
		dataBytes = append(dataBytes, d.cards[i].ToByte()...)
	}
	return dataBytes
}

func (d *Deck) Of(dataBytes []byte) Deck {
	var deckSize = dataBytes[0]
	var cards = make([]Card, deckSize)
	for i := 0; i < int(deckSize); i++ {
		var c Card
		cards[i] = c.Of(dataBytes[1+(i*CardSize) : i*CardSize])
	}
	return Deck{
		deckSize: dataBytes[0],
		cards:    cards,
	}
}
