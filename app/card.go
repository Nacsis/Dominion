package app

import (
	"perun.network/perun-examples/dominion-cli/app/util"
)

type Card struct {
	CardType      util.CardType
	Money         uint8
	Play          uint8
	Buy           uint8
	Draw          uint8
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
	case util.Cellar:
		c.VictoryPoints = util.ActionCardVictoryPoint
		c.CardType = util.Cellar
		c.PlayCost = util.ActionCardPlayCost
		c.BuyCost = 2
		c.Play = 1
		break
	case util.Market:
		c.Money = 1
		c.VictoryPoints = util.ActionCardVictoryPoint
		c.CardType = util.Market
		c.PlayCost = util.ActionCardPlayCost
		c.BuyCost = 5
		c.Play = 1
		c.Buy = 1
		c.Play = 1
		break
	case util.Merchant:
		c.VictoryPoints = util.ActionCardVictoryPoint
		c.CardType = util.Merchant
		c.PlayCost = util.ActionCardPlayCost
		c.BuyCost = 3
		c.Play = 1
		c.Draw = 1
		break
	case util.Mine:
		c.VictoryPoints = util.ActionCardVictoryPoint
		c.CardType = util.Mine
		c.PlayCost = util.ActionCardPlayCost
		c.BuyCost = 5
		break
	case util.Remodel:
		c.VictoryPoints = util.ActionCardVictoryPoint
		c.CardType = util.Remodel
		c.PlayCost = util.ActionCardPlayCost
		c.BuyCost = 4
		break
	case util.Smithy:
		c.VictoryPoints = util.ActionCardVictoryPoint
		c.CardType = util.Smithy
		c.PlayCost = util.ActionCardPlayCost
		c.BuyCost = 4
		c.Draw = 3
		break
	case util.Chapel:
		c.VictoryPoints = util.ActionCardVictoryPoint
		c.CardType = util.Chapel
		c.PlayCost = util.ActionCardPlayCost
		c.BuyCost = 2
		break
	case util.Workshop:
		c.VictoryPoints = util.ActionCardVictoryPoint
		c.CardType = util.Workshop
		c.PlayCost = util.ActionCardPlayCost
		c.BuyCost = 3
		break
	case util.Feast:
		c.VictoryPoints = util.ActionCardVictoryPoint
		c.CardType = util.Feast
		c.PlayCost = util.ActionCardPlayCost
		c.BuyCost = 4
		break
	case util.Village:
		c.VictoryPoints = util.ActionCardVictoryPoint
		c.CardType = util.Village
		c.PlayCost = util.ActionCardPlayCost
		c.BuyCost = 3
		c.Play = 2
		c.Draw = 1
		break
	}
}
func (c *Card) IsMoneyCard() bool {
	return c.CardType == util.Copper || c.CardType == util.Silver || c.CardType == util.Gold
}
