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
	deck := app.Deck{}
	deck.Init()

	var bytes = deck.ToByte()

	deckActual := app.Deck{}
	deckActual.Of(bytes[1:])

	assert.Equal(t, deck, deckActual)
}

// Test_Deck_Draw
func Test_Deck_Draw(t *testing.T) {
	deck, cardTypesMain, cardTypesHand, discardTypesHand, playedTypesHand := deckSetUp()

	err := deck.DrawOneCard(global.RandomBytes(util.HashSize))
	assert.Nil(t, err)
	assert.True(t, deck.MainPile.Size() == len(cardTypesMain)-1)
	assert.True(t, deck.HandPile.Size() == len(cardTypesHand)+1)
	assert.True(t, deck.DiscardedPile.Size() == len(discardTypesHand))
	assert.True(t, deck.PlayedPile.Size() == len(playedTypesHand))

	var difMain []app.CardType
	indexDif := 0

	for i := 0; i < len(cardTypesMain)-1; i++ {
		if deck.MainPile.Cards[i].CardType != cardTypesMain[i+indexDif] {
			difMain = append(difMain, cardTypesMain[i])
			indexDif++
		}
	}
	if len(difMain) == 0 {
		difMain = append(difMain, cardTypesMain[len(cardTypesMain)-1])
	}

	assert.Len(t, difMain, 1)
	assert.Equal(t, difMain[0], deck.HandPile.Cards[len(deck.HandPile.Cards)-1].CardType)
}
