package game

import (
	"github.com/google/uuid"
)

const (
	CardSize = 18
)

type Card struct {
	money         uint8
	victoryPoints uint8
	id            uuid.UUID
}

func NewMoneyCard() Card {
	return NewCard(1, 0)
}
func NewVictoryCard() Card {
	return NewCard(0, 1)
}

func NewCard(m, v uint8) Card {
	return Card{
		money:         m,
		victoryPoints: v,
		id:            uuid.New(),
	}
}

func (c *Card) ToByte() []byte {
	var dataBytes = make([]byte, 16)
	dataBytes[0] = c.money
	dataBytes[1] = c.victoryPoints

	idBytes, err := c.id.MarshalBinary()
	if err != nil {
		// TODO
	}
	for i := 2; i < 16; i++ {
		dataBytes[i] = idBytes[i-2]
	}

	return dataBytes
}

func (c *Card) Of(dataBytes []byte) Card {
	var id uuid.UUID
	id.UnmarshalBinary(dataBytes[2:15])
	return Card{
		money:         dataBytes[0],
		victoryPoints: dataBytes[1],
		id:            id,
	}
}
