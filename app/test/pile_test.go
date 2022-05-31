package test

import (
	"github.com/stretchr/testify/assert"
	"perun.network/perun-examples/app-channel/app"
	"perun.network/perun-examples/app-channel/app/util"
	"perun.network/perun-examples/app-channel/global"
	"testing"
)

// Test_Pile_Of
func Test_Pile_Of(t *testing.T) {
	expectedPile, cardTypes := pileSetUp()
	var bytes = make([]byte, 0)
	for _, cardType := range cardTypes {
		bytes = append(bytes, byte(cardType))
	}
	actualPile := app.Pile{}
	actualPile.Of(bytes)

	assert.Equal(t, expectedPile, actualPile)
}

// Test_Pile_ToByte
func Test_Pile_ToByte(t *testing.T) {
	pile, cardTypes := pileSetUp()

	var expectedBytes = make([]byte, 0)
	expectedBytes = append(expectedBytes, byte(len(cardTypes)))
	for _, cardType := range cardTypes {
		expectedBytes = append(expectedBytes, byte(cardType))
	}
	assert.Equal(t, expectedBytes, pile.ToByte())
}

/*
// Test_Pile_Resize
func Test_Pile_Resize(t *testing.T) {
	pile, cardTypes := pileSetUp()
	pile.Cards[ranIntN(len(cardTypes)-1)] = app.Card{}
	pile._ResizeCards()
	assert.Len(t, pile.Cards, len(cardTypes)-1)
}*/

// Test_Pile_Add
func Test_Pile_Add(t *testing.T) {
	pile, cardTypes := pileSetUp()
	pile.AddCardToPile(cardOfType(util.CardType(ranIntN(5))))
	assert.Len(t, pile.Cards, len(cardTypes)+1)
}

/*
// Test_Pile_SeedToIndex
func Test_Pile_SeedToIndex(t *testing.T) {
	pile, cardTypes := pileSetUp()
	index, err := pile._SeedToIndex(global.RandomBytes(util.HashSize))
	assert.Nil(t, err)
	assert.True(t, len(cardTypes) > index)
	assert.True(t, 0 <= index)
}*/

// Test_Pile_Serialization
func Test_Pile_Serialization(t *testing.T) {
	pile, _ := pileSetUp()

	var bytes = pile.ToByte()

	pileActual := app.Pile{}
	pileActual.Of(bytes[1:])

	assert.Equal(t, pile, pileActual)
}

// Test_Pile_DrawCard
func Test_Pile_DrawCard(t *testing.T) {
	pile, cardTypes := pileSetUp()
	card, err := pile.DrawCardBasedOnSeed(global.RandomBytes(util.HashSize))
	var counterCardTypeBefore = 0
	for _, cardType := range cardTypes {
		if cardType == card.CardType {
			counterCardTypeBefore++
		}
	}
	var counterCardTypeAfter = 0
	for _, c := range pile.Cards {
		if c.CardType == card.CardType {
			counterCardTypeAfter++
		}
	}

	assert.NotEqual(t, card, app.Card{})
	assert.Nil(t, err)
	assert.True(t, len(cardTypes) > len(pile.Cards))
	assert.Equal(t, counterCardTypeBefore-1, counterCardTypeAfter)
}
