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
	deck, _, _, _ := deckSetUp()

	var bytes = deck.ToByte()

	deckActual := app.Deck{}
	deckActual.Of(bytes[1:])

	assert.Equal(t, deck, deckActual)
}

// Test_Deck_Draw
func Test_Deck_Draw(t *testing.T) {
	deck, cardTypesMain, cardTypesHand, discardTypesHand := deckSetUp()

	err := deck.DrawOneCard(global.RandomBytes(util.HashSize))
	assert.Nil(t, err)
	assert.True(t, deck.MainCardPile.Size() == len(cardTypesMain)-1)
	assert.True(t, deck.HandCards.Size() == len(cardTypesHand)+1)
	assert.True(t, deck.DiscardPile.Size() == len(discardTypesHand))
	var difMain []app.CardType
	indexDif := 0

	for i := 0; i < len(cardTypesMain)-1; i++ {
		if deck.MainCardPile.Cards[i].CardType != cardTypesMain[i+indexDif] {
			difMain = append(difMain, cardTypesMain[i])
			indexDif++
		}
	}
	if len(difMain) == 0 {
		difMain = append(difMain, cardTypesMain[len(cardTypesMain)-1])
	}

	assert.Len(t, difMain, 1)
	assert.Equal(t, difMain[0], deck.HandCards.Cards[len(deck.HandCards.Cards)-1].CardType)
}
