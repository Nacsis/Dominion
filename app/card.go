package app

import (
	"perun.network/perun-examples/app-channel/app/util"
)

type Card struct {
	cardType      CardType
	money         uint8
	victoryPoints uint8
}

type CardType uint8

const (
	MoneyCopper CardType = iota
	MoneySilver
	MoneyGold
	VictorySmall
	VictoryMid
	VictoryBig
)

func NewCardOfType(ct CardType) Card {
	// Money
	if 0 <= int(ct) && int(ct) < 3 {
		return NewMoneyCard(ct)
	} else {
		// victory
		return NewVictoryCard(ct)
	}
}

func NewMoneyCard(ct CardType) Card {
	switch ct {
	case MoneyCopper:
		return NewCard(util.MonValueCopper, 0, MoneyCopper)
	case MoneySilver:
		return NewCard(util.MonValueSilver, 0, MoneySilver)
	case MoneyGold:
		return NewCard(util.MonValueGold, 0, MoneyGold)
	default:
		panic("Not Found")
	}
}
func NewVictoryCard(ct CardType) Card {
	switch ct {
	case VictorySmall:
		return NewCard(0, 1, VictorySmall)
	case VictoryMid:
		return NewCard(0, 2, VictoryMid)
	case VictoryBig:
		return NewCard(0, 6, VictoryBig)
	default:
		panic("Not Found")
	}
}

func NewCard(m, v uint8, ct CardType) Card {
	return Card{
		money:         m,
		victoryPoints: v,
		cardType:      ct,
	}
}

func (c *Card) ToByte() []byte {
	var dataBytes = make([]byte, util.CardSize)
	dataBytes[0] = byte(c.cardType)
	dataBytes[1] = c.money
	dataBytes[2] = c.victoryPoints
	return dataBytes
}

func (c *Card) Of(dataBytes []byte) Card {
	return Card{
		cardType:      CardType(dataBytes[0]),
		money:         dataBytes[1],
		victoryPoints: dataBytes[2],
	}
}
