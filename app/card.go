package app

import (
	"perun.network/perun-examples/dominion-cli/app/util"
)

type Card struct {
	CardType      util.CardType
	Money         uint8
	VictoryPoints uint8
	PlayCost      uint8
	BuyCost       uint8
}

// ToByte create a byte representation of Card
func (c *Card) ToByte() []byte {
	return []byte{byte(c.CardType)}
}

// Of create Card out of a bytes
func (c *Card) Of(dataBytes []byte) {
	switch util.CardType(dataBytes[0]) {
	case util.Copper:
		c.Money = util.CopperMoneyValue
		c.CardType = util.Copper
		c.PlayCost = util.MoneyCardPlayCost
		c.BuyCost = util.CopperCost
		break
	case util.Silver:
		c.Money = util.SilverMoneyValue
		c.CardType = util.Silver
		c.PlayCost = util.MoneyCardPlayCost
		c.BuyCost = util.SilverCost
		break
	case util.Gold:
		c.Money = util.GoldMoneyValue
		c.CardType = util.Gold
		c.PlayCost = util.MoneyCardPlayCost
		c.BuyCost = util.GoldCost
		break
	case util.VictorySmall:
		c.VictoryPoints = util.VictorySmallVictoryValue
		c.CardType = util.VictorySmall
		c.PlayCost = util.VictoryCardPlayCost
		c.BuyCost = util.VictorySmallCost
		break
	case util.VictoryMid:
		c.VictoryPoints = util.VictoryMidVictoryValue
		c.CardType = util.VictoryMid
		c.PlayCost = util.VictoryCardPlayCost
		c.BuyCost = util.VictoryMidCost
		break
	case util.VictoryBig:
		c.VictoryPoints = util.VictoryBigVictoryValue
		c.CardType = util.VictoryBig
		c.PlayCost = util.VictoryCardPlayCost
		c.BuyCost = util.VictoryBigCost
		break
	}
}
