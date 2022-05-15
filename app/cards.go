package app

import "io"

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
	cost  uint
}

func (c CardTypeBase) Name() CardName {
	return c.name
}

func (c CardTypeBase) Group() CardGroup {
	return c.group
}

func (c CardTypeBase) Cost() uint {
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
	MonValueCopper int8 = 1
	MonValueSilver int8 = 2
	MonValueGold   int8 = 3
	CostsCopper    int8 = 0
	CostsSilver    int8 = 3
	CostsGold      int8 = 6
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

func NewCopper(d *AppData) {
	c := Card{
		owner:          Unowned,
		cost:           0,
		group:          TreasureCard,
		dActions:       0,
		dBuys:          0,
		dMoney:         1,
		dDraws:         0,
		dVictoryPoints: 0,
		id:             d.NumAllCards}

	d.AddCard(c)
}

//#### Card methods ####
func writeCards(w io.Writer, cards [256]Card) error {
	var err error
	for _, card := range cards {
		err = card.WriteCard(w)
		if err == nil {
			continue
		}
	}
	return err
}

func (c Card) WriteCard(w io.Writer) error {
	uint8_attributes := []uint8{uint8(c.owner), c.cost, uint8(c.group), c.dActions, c.dBuys, c.dMoney, c.dDraws, c.dVictoryPoints, c.id, uint8(c.name)}
	err := writeUInt8Array(w, uint8_attributes)
	return err
}

func ReadCard(r io.Reader) Card {
	buf := make([]byte, 10)
	io.ReadFull(r, buf)

	card := Card{
		owner:          Owner(buf[0]),
		cost:           buf[1],
		group:          CardType(buf[2]),
		dActions:       buf[3],
		dBuys:          buf[4],
		dMoney:         buf[5],
		dDraws:         buf[6],
		dVictoryPoints: buf[7],
		id:             buf[8],
		name:           CardName(buf[9]),
	}

	return card
}
