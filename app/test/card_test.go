package test

import (
	"github.com/stretchr/testify/assert"
	"perun.network/perun-examples/app-channel/app"
	"testing"
)

// Test_Card_Of
func Test_Card_Of(t *testing.T) {

	cardCheckOf(cardOfType(app.MoneyCopper), app.MoneyCopper, t)
	cardCheckOf(cardOfType(app.MoneySilver), app.MoneySilver, t)
	cardCheckOf(cardOfType(app.MoneyGold), app.MoneyGold, t)
	cardCheckOf(cardOfType(app.VictorySmall), app.VictorySmall, t)
	cardCheckOf(cardOfType(app.VictoryMid), app.VictoryMid, t)
	cardCheckOf(cardOfType(app.VictoryBig), app.VictoryBig, t)
}

// Test_Card_ToByte
func Test_Card_ToByte(t *testing.T) {
	cardCheckToByte(cardOfType(app.MoneyCopper), app.MoneyCopper, t)
	cardCheckToByte(cardOfType(app.MoneySilver), app.MoneySilver, t)
	cardCheckToByte(cardOfType(app.MoneyGold), app.MoneyGold, t)
	cardCheckToByte(cardOfType(app.VictorySmall), app.VictorySmall, t)
	cardCheckToByte(cardOfType(app.VictoryMid), app.VictoryMid, t)
	cardCheckToByte(cardOfType(app.VictoryBig), app.VictoryBig, t)
}

func cardCheckOf(expected app.Card, ct app.CardType, t *testing.T) {
	actual := app.Card{}
	actual.Of([]byte{byte(ct)})
	// TODO Empty List
	// TODO byte not in CardType
	assert.Equal(t, expected, actual, "Check card 'Of' for ct")
}

func cardCheckToByte(card app.Card, expected app.CardType, t *testing.T) {
	bytes := card.ToByte()
	assert.Len(t, bytes, 1, "byte length should be 1")
	actual := app.CardType(bytes[0])
	// TODO byte not in CardType
	assert.Equal(t, expected, actual, "Check card 'ToByte' for ct")
}
