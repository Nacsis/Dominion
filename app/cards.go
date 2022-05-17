package app

import "fmt"

//### Card Type Constants ###
type CardGroup uint8

const (
	UnknownCardGroup CardGroup = iota
	ActionCard
	TreasureCard
	VictoryCard
)

//### Card Name Constants ###
type CardName uint8

const (
	UnknownCardName CardName = iota
	Copper
)

// TODO define when needed...
// type CardType interface {
// 	Name()
// 	Group()
// 	Cost()
// }

type CardTypeBase struct {
	name  CardName
	group CardGroup
	cost  uint8
}

func (c CardTypeBase) Name() CardName {
	return c.name
}

func (c CardTypeBase) Group() CardGroup {
	return c.group
}

func (c CardTypeBase) Cost() uint8 {
	return c.cost
}

type ActionCardType struct {
	// composition with CardBase
	CardTypeBase
	// specialized properties
	dActions, dBuys, dMoney, dDraws int8
}

type TreasureCardType struct {
	// composition with CardBase
	CardTypeBase
	// specialized properties
	dMoney int8
}

const (
	MonValueCopper int8  = 1
	MonValueSilver int8  = 2
	MonValueGold   int8  = 3
	CostsCopper    uint8 = 0
	CostsSilver    uint8 = 3
	CostsGold      uint8 = 6
)

type VictoryCardType struct {
	// composition with CardBase
	CardTypeBase
	// specialized properties
	dVictoryPoints int8
}

const VicValueProvince int8 = 6

// type CardType struct {
// 	cost           uint8
// 	group          CardGroup
// 	dActions       uint8
// 	dBuys          uint8
// 	dMoney         uint8
// 	dDraws         uint8
// 	dVictoryPoints uint8
// 	id             uint8
// 	name           CardName
// }

//#### Card constructors ####

// checks position validity
// returns (valid bool, isBase bool)
func CheckPositionValid(pos int8) (bool, bool) {
	// TODO
	if pos > int8(NumActionCardsInGame) || pos == 0 {
		return false, false
	}
	if pos < -int8(NumBaseCards) {
		return false, true
	}
	return true, pos < 0
}

// TODO replace return type with appropriate interface
func NewBaseCard(pos int8) (*TreasureCardType, error) {
	valid, isBase := CheckPositionValid(pos)
	if !valid || !isBase {
		return nil, fmt.Errorf("position not valid")
	}
	name := CardName(-pos)
	switch name {
	case Copper:
		{
			return &TreasureCardType{
				CardTypeBase: CardTypeBase{
					name:  Copper,
					group: TreasureCard,
					cost:  CostsCopper,
				},
				dMoney: MonValueCopper,
			}, nil
		}
	default:
		{
			return nil, fmt.Errorf("card not implemented")
		}
	}
}

//#### Card methods ####
