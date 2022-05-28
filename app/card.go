package app

import (
	"perun.network/perun-examples/app-channel/app/util"
)

type Card struct {
	CardType      CardType
	Money         uint8
	VictoryPoints uint8
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

func (c *Card) ToByte() []byte {
	return []byte{byte(c.CardType)}
}

func (c *Card) Of(dataBytes []byte) {
	switch CardType(dataBytes[0]) {
	case MoneyCopper:
		c.Money = util.MonValueCopper
		c.CardType = MoneyCopper
	case MoneySilver:
		c.Money = util.MonValueSilver
		c.CardType = MoneySilver
	case MoneyGold:
		c.Money = util.MonValueGold
		c.CardType = MoneyGold
	case VictorySmall:
		c.VictoryPoints = 1
		c.CardType = VictorySmall
	case VictoryMid:
		c.VictoryPoints = 2
		c.CardType = VictoryMid
	case VictoryBig:
		c.VictoryPoints = 3
		c.CardType = VictoryBig
	}
}
