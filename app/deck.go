package app

import (
	"log"
)

type Deck struct {
	mainCardPile Pile
	handCards    Pile
}

func (d *Deck) ToByte() []byte {
	mainCardPileBytes := d.mainCardPile.ToByte()
	handCardsBytes := d.handCards.ToByte()

	var deckByteLength = len(mainCardPileBytes) + len(handCardsBytes) + 1 // +1 is for length
	var dataBytes = make([]byte, deckByteLength)

	dataBytes[0] = byte(deckByteLength)
	dataBytes = append(dataBytes, mainCardPileBytes...)
	dataBytes = append(dataBytes, handCardsBytes...)

	return dataBytes
}

func (d *Deck) Of(dataBytes []byte) {
	mainCardPileSize := dataBytes[0]
	log.Println(mainCardPileSize)
	log.Println(dataBytes[1:mainCardPileSize])
	d.mainCardPile.Of(dataBytes[1:mainCardPileSize])

	handCardSize := dataBytes[mainCardPileSize]
	log.Println(handCardSize)
	log.Println(dataBytes[handCardSize : +handCardSize+mainCardPileSize])
	d.handCards.Of(dataBytes[handCardSize : +handCardSize+mainCardPileSize])
}

func (d *Deck) draw(seed []byte) {
	card := d.mainCardPile.draw(seed)
	d.handCards.add(card)
}
