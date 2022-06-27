package test

import (
	"perun.network/perun-examples/dominion-cli/app"
	"perun.network/perun-examples/dominion-cli/app/util"
)

func cardOfType(ct util.CardType) app.Card {
	card := app.Card{}
	card.Of([]byte{byte(ct)})
	return card
}
