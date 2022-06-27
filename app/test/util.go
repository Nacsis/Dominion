package test

import (
	"math/rand"
	"time"

	"perun.network/perun-examples/dominion-cli/app"
	"perun.network/perun-examples/dominion-cli/app/util"
	"perun.network/perun-examples/dominion-cli/global"
)

func pileSetUp() (app.Pile, []util.CardType) {
	var cards = make([]app.Card, 0)
	var cardTypes = make([]util.CardType, 0)

	for i := 0; i < 10; i++ {
		cardTypes = append(cardTypes, util.CardType(ranIntN(5)))
	}

	for _, cardType := range cardTypes {
		cards = append(cards, cardOfType(cardType))
	}

	return app.Pile{Cards: cards}, cardTypes
}

func deckSetUp() (app.Deck, []util.CardType, []util.CardType, []util.CardType, []util.CardType) {

	main, cardTypesMain := pileSetUp()
	hand, cardTypesHand := pileSetUp()
	discard, discardTypesHand := pileSetUp()
	played, playedTypesHand := pileSetUp()

	return app.Deck{MainPile: main, HandPile: hand, DiscardedPile: discard, PlayedPile: played}, cardTypesMain, cardTypesHand, discardTypesHand, playedTypesHand
}
func rngCommittedSetUp() app.RNG {
	preImage := global.RandomBytes(util.HashSize)
	image, _ := global.ToImage(preImage)
	return app.RNG{ImageA: image}
}

func rngTouchedSetUp() app.RNG {
	preImage := global.RandomBytes(util.HashSize)
	image, _ := global.ToImage(preImage)
	return app.RNG{ImageA: image, PreImageB: global.RandomBytes(util.HashSize)}
}

func rngReleaseSetUp() app.RNG {
	preImage := global.RandomBytes(util.HashSize)
	image, _ := global.ToImage(preImage)
	return app.RNG{ImageA: image, PreImageB: global.RandomBytes(util.HashSize), PreImageA: preImage}
}

func ranIntN(n int) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return r1.Intn(n)
}

func cardOfType(ct util.CardType) app.Card {
	switch ct {
	case util.Copper:
		return app.Card{Money: util.CopperMoneyValue, CardType: util.Copper}
	case util.Silver:
		return app.Card{Money: util.SilverMoneyValue, CardType: util.Silver}
	case util.Gold:
		return app.Card{Money: util.GoldMoneyValue, CardType: util.Gold}
	case util.VictorySmall:
		return app.Card{VictoryPoints: 1, CardType: util.VictorySmall}
	case util.VictoryMid:
		return app.Card{VictoryPoints: 2, CardType: util.VictoryMid}
	default:
		return app.Card{VictoryPoints: 3, CardType: util.VictoryBig}
	}
}
