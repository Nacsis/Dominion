package app

import "io"

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
	owner          Player
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

func newCopper(id uint8) Card {
	c := Card{
		cost:           0,
		group:          TreasureCard,
		dActions:       0,
		dBuys:          0,
		dMoney:         1,
		dDraws:         0,
		dVictoryPoints: 0,
		id:             id}

	return c
}

//#### Card methods ####
func writeCards(w io.Writer, cards [256]Card) error {
	var err error
	for _, card := range cards {
		err = card.writeCard(w)
		if err == nil {
			continue
		}
	}
	return err
}

func (c Card) writeCard(w io.Writer) error {
	uint8_attributes := []uint8{c.owner.id, c.cost, uint8(c.group), c.dActions, c.dBuys, c.dMoney, c.dDraws, c.dVictoryPoints, c.id, uint8(c.name)}
	err := writeUInt8Array(w, uint8_attributes)
	return err
}

func readCard(r io.Reader) Card {
	buf := make([]byte, 10)
	io.ReadFull(r, buf)

	card := Card{
		owner:          getPlayer(buf[0]), //todo: getPlayer Method
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
