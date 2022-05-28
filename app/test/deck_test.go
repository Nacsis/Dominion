package test

import (
	"github.com/stretchr/testify/assert"
	"perun.network/perun-examples/app-channel/app"
	"perun.network/perun-examples/app-channel/app/util"
	"perun.network/perun-examples/app-channel/global"
	"testing"
)

// Test_Deck_Serialization
func Test_Deck_Serialization(t *testing.T) {
	deck, _, _ := deckSetUp()

	var bytes = deck.ToByte()

	deckActual := app.Deck{}
	deckActual.Of(bytes[1:])

	assert.Equal(t, deck, deckActual)
}

// Test_Deck_Draw
func Test_Deck_Draw(t *testing.T) {
	deck, cardTypesMain, cardTypesHand := deckSetUp()

	err := deck.DrawOneCard(global.RandomBytes(util.HashSize))
	assert.Nil(t, err)
	assert.True(t, len(deck.MainCardPile.Cards) == len(cardTypesMain)-1)
	assert.True(t, len(deck.HandCards.Cards) == len(cardTypesHand)+1)

	var difMain []app.CardType
	for i := 0; i < len(cardTypesMain)-1; i++ {
		if deck.MainCardPile.Cards[i].CardType != cardTypesMain[i] {
			difMain = append(difMain, cardTypesMain[i])
		}
	}
	if len(difMain) == 0 {
		difMain = append(difMain, cardTypesMain[len(cardTypesMain)-1])
	}

	assert.Len(t, difMain, 1)
	assert.Equal(t, difMain[0], deck.HandCards.Cards[len(deck.HandCards.Cards)-1].CardType)
}
