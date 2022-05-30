package test

import (
	"math/rand"
	"perun.network/perun-examples/app-channel/app"
	"perun.network/perun-examples/app-channel/app/util"
	"perun.network/perun-examples/app-channel/global"
	"time"
)

func pileSetUp() (app.Pile, []app.CardType) {
	var cards = make([]app.Card, 0)
	var cardTypes = make([]app.CardType, 0)

	for i := 0; i < 10; i++ {
		cardTypes = append(cardTypes, app.CardType(ranIntN(5)))
	}

	for _, cardType := range cardTypes {
		cards = append(cards, cardOfType(cardType))
	}

	return app.Pile{Cards: cards}, cardTypes
}

func deckSetUp() (app.Deck, []app.CardType, []app.CardType, []app.CardType) {

	main, cardTypesMain := pileSetUp()
	hand, cardTypesHand := pileSetUp()
	discard, discardTypesHand := pileSetUp()
	return app.Deck{MainCardPile: main, HandCards: hand, DiscardPile: discard}, cardTypesMain, cardTypesHand, discardTypesHand
}
func rngCommittedSetUp() app.RNG {
	preImage := global.RandomBytes(util.HashSize)
	return app.RNG{ImageA: global.ToImage(preImage)}
}

func rngTouchedSetUp() app.RNG {
	preImage := global.RandomBytes(util.HashSize)
	return app.RNG{ImageA: global.ToImage(preImage), PreImageB: global.RandomBytes(util.HashSize)}
}

func rngReleaseSetUp() app.RNG {
	preImage := global.RandomBytes(util.HashSize)
	return app.RNG{ImageA: global.ToImage(preImage), PreImageB: global.RandomBytes(util.HashSize), PreImageA: preImage}
}

func ranIntN(n int) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return r1.Intn(n)
}

func cardOfType(ct app.CardType) app.Card {
	switch ct {
	case app.MoneyCopper:
		return app.Card{Money: util.MonValueCopper, CardType: app.MoneyCopper}
	case app.MoneySilver:
		return app.Card{Money: util.MonValueSilver, CardType: app.MoneySilver}
	case app.MoneyGold:
		return app.Card{Money: util.MonValueGold, CardType: app.MoneyGold}
	case app.VictorySmall:
		return app.Card{VictoryPoints: 1, CardType: app.VictorySmall}
	case app.VictoryMid:
		return app.Card{VictoryPoints: 2, CardType: app.VictoryMid}
	default:
		return app.Card{VictoryPoints: 3, CardType: app.VictoryBig}
	}
}
