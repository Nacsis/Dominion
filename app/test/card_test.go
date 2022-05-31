package test

import (
	"github.com/stretchr/testify/assert"
	"perun.network/perun-examples/app-channel/app"
	"perun.network/perun-examples/app-channel/app/util"
	"testing"
)

// Test_Card_Of
func Test_Card_Of(t *testing.T) {

	cardCheckOf(cardOfType(util.Copper), util.Copper, t)
	cardCheckOf(cardOfType(util.Silver), util.Silver, t)
	cardCheckOf(cardOfType(util.Gold), util.Gold, t)
	cardCheckOf(cardOfType(util.VictorySmall), util.VictorySmall, t)
	cardCheckOf(cardOfType(util.VictoryMid), util.VictoryMid, t)
	cardCheckOf(cardOfType(util.VictoryBig), util.VictoryBig, t)
}

// Test_Card_ToByte
func Test_Card_ToByte(t *testing.T) {
	cardCheckToByte(cardOfType(util.Copper), util.Copper, t)
	cardCheckToByte(cardOfType(util.Silver), util.Silver, t)
	cardCheckToByte(cardOfType(util.Gold), util.Gold, t)
	cardCheckToByte(cardOfType(util.VictorySmall), util.VictorySmall, t)
	cardCheckToByte(cardOfType(util.VictoryMid), util.VictoryMid, t)
	cardCheckToByte(cardOfType(util.VictoryBig), util.VictoryBig, t)
}

func cardCheckOf(expected app.Card, ct util.CardType, t *testing.T) {
	actual := app.Card{}
	actual.Of([]byte{byte(ct)})
	// TODO Empty List
	// TODO byte not in CardType
	assert.Equal(t, expected, actual, "Check card 'Of' for ct")
}

func cardCheckToByte(card app.Card, expected util.CardType, t *testing.T) {
	bytes := card.ToByte()
	assert.Len(t, bytes, 1, "byte length should be 1")
	actual := util.CardType(bytes[0])
	// TODO byte not in CardType
	assert.Equal(t, expected, actual, "Check card 'ToByte' for ct")
}
