package app

import "io"

//### Player Constants ###
type Owner uint8

const (
	Unowned Owner = iota
	Player_1
	Player_2
)

//### Card Type Constants ###
type CardType uint8

const (
	UnknownCardType CardType = iota
	ActionCard
	TreasureCard
	VictoryCard
)

//### Card Name Constants ###
type CardName uint8

const (
	UnknownCardName CardType = iota
	Copper
)

type Card struct {
	owner          Owner
	cost           uint8
	group          CardType
	dActions       uint8
	dBuys          uint8
	dMoney         uint8
	dDraws         uint8
	dVictoryPoints uint8
	id             uint8
	name           CardName
}

//#### Card constructors ####

func NewCopper(d *Data) {
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

func (c Card) equals(c2 Card) bool {
	equal :=
		c.owner == c2.owner &&
			c.cost == c2.cost &&
			c.group == c2.group &&
			c.dActions == c2.dActions &&
			c.dBuys == c2.dBuys &&
			c.dMoney == c2.dMoney &&
			c.dDraws == c2.dDraws &&
			c.dVictoryPoints == c2.dVictoryPoints &&
			c.id == c2.id &&
			c.name == c2.name

	return equal
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
