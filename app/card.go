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

func CardOfType(ct CardType) (Card, error) {
	switch ct {
	case MoneyCopper:
		return Card{money: util.MonValueCopper, cardType: MoneyCopper}, nil
	case MoneySilver:
		return Card{money: util.MonValueSilver, cardType: MoneySilver}, nil
	case MoneyGold:
		return Card{money: util.MonValueGold, cardType: MoneyGold}, nil
	case VictorySmall:
		return Card{victoryPoints: 1, cardType: VictorySmall}, nil
	case VictoryMid:
		return Card{victoryPoints: 2, cardType: VictoryMid}, nil
	case VictoryBig:
		return Card{victoryPoints: 3, cardType: VictoryBig}, nil
	default:
		return Card{}, util.ThrowError(util.ErrorConstCARD, "CardOfType", "no card for card type found")
	}
}
