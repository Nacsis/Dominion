package test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"perun.network/perun-examples/dominion-cli/app"
	"perun.network/perun-examples/dominion-cli/app/util"
	"strings"
	"testing"
)

func _init() (app.DominionAppData, error) {
	data := app.DominionAppData{}
	err := data.Init(0)
	return data, err
}

// Clone returns a deep copy of the app data.
func _clone(data app.DominionAppData) app.DominionAppData {
	origJSON, _ := json.Marshal(data)
	fromDataClone := app.DominionAppData{}
	json.Unmarshal(origJSON, &fromDataClone)
	return fromDataClone
}
func setUpPlayAction(cardType []util.CardType) app.DominionAppData {
	data, _ := _init()
	for _, u := range cardType {
		data.CardDecks[0].HandPile.AddCardToPile(cardOfType(u))
	}
	data.Turn.SetAllowed(util.PlayCard)
	return data
}

func Test_Encode_Decode(t *testing.T) {
	before := app.DominionAppData{
		Turn: app.Turn{
			NextActor:              1,
			PerformedAction:        2,
			MandatoryPartFulfilled: false,
			PossibleActions:        [8]bool{true, false, true, false, true, false, true, false},
			Params: app.Params{
				MainTarget:       3,
				SecondLvlTarget:  4,
				SecondLvlIndices: []byte{5, 4, 3, 2, 1},
			},
		},
		Stock: app.Stock{
			CardAmounts: [16]uint8{10, 20, 30, 40, 30, 20, 10, 9, 8, 7, 6, 6, 5, 4, 12, 1},
			Trash:       [16]uint8{11, 21, 31, 41, 31, 21, 11, 1, 2, 3, 4, 5, 6, 7, 11, 2},
		},
		CardDecks: [2]app.Deck{{
			MainPile:      app.Pile{Cards: []app.Card{cardOfType(util.Copper), cardOfType(util.Mine)}},
			HandPile:      app.Pile{Cards: []app.Card{cardOfType(util.Copper), cardOfType(util.Mine)}},
			DiscardedPile: app.Pile{Cards: []app.Card{cardOfType(util.Copper), cardOfType(util.Mine)}},
			PlayedPile:    app.Pile{Cards: []app.Card{cardOfType(util.Copper), cardOfType(util.Mine)}},
			Resources:     [4]uint8{4, 3, 2, 1}}, {

			MainPile:      app.Pile{Cards: []app.Card{cardOfType(util.Copper), cardOfType(util.Mine)}},
			HandPile:      app.Pile{Cards: []app.Card{cardOfType(util.Copper), cardOfType(util.Mine)}},
			DiscardedPile: app.Pile{Cards: []app.Card{cardOfType(util.Copper), cardOfType(util.Mine)}},
			PlayedPile:    app.Pile{Cards: []app.Card{cardOfType(util.Copper), cardOfType(util.Mine)}},
			Resources:     [4]uint8{0, 0, 1, 0},
		},
		},
		Rng: app.RNG{
			ImageA:    [util.HashSizeByte]byte{},
			PreImageB: [util.PreImageSizeByte]byte{},
			PreImageA: [util.PreImageSizeByte]byte{},
		},
	}

	after := before.Clone2DominionAppData()

	s := strings.Builder{}
	after.Encode(&s)
	app2 := app.DominionApp{}
	after2, _ := app2.DecodeData(strings.NewReader(s.String()))
	after3 := after2.(*app.DominionAppData)
	after3.Clone()

	/*
		assert.Equal(t, true, true)
		assert.Equal(t, reflect.DeepEqual(before.Turn, after3.Turn), true)
		assert.Equal(t, reflect.DeepEqual(before.Stock, after3.Stock), true)
		assert.Equal(t, reflect.DeepEqual(before.CardDecks, after3.CardDecks), true)
		assert.Equal(t, reflect.DeepEqual(before.Rng, after3.Rng), true)
		assert.Equal(t, reflect.DeepEqual(before, after3), true)

	*/
	print("done")
}

func Test_Play_Cellar(t *testing.T) {
	handCard := []util.CardType{util.Cellar, util.Market, util.Oasis, util.Mine, util.Remodel}
	data := setUpPlayAction(handCard)

	followUpIndices := []uint8{0, 1, 2}
	clone := _clone(data)
	clone.PlayCard(0, uint8(0), followUpIndices, util.NONE)

	deckClone := clone.CardDecks[0]

	data.CardDecks[0].Resources[util.PlayableCards] = 1
	data.CardDecks[0].Resources[util.DrawableCards] += uint8(len(followUpIndices))

	assert.Equal(t, data.CardDecks[0].MainPile, deckClone.MainPile)
	assert.Equal(t, data.CardDecks[0].Resources, deckClone.Resources)

	assert.Equal(t, len(deckClone.PlayedPile.Cards), 1)
	assert.Equal(t, deckClone.PlayedPile.Cards[0].CardType, util.Cellar)

	assert.Equal(t, len(deckClone.HandPile.Cards), 1)
	assert.Equal(t, deckClone.HandPile.Cards[0].CardType, util.Remodel)

	assert.Equal(t, len(deckClone.DiscardedPile.Cards), 3)
	assert.Equal(t, deckClone.DiscardedPile.Cards[0].CardType, util.Mine)
	assert.Equal(t, deckClone.DiscardedPile.Cards[1].CardType, util.Oasis)
	assert.Equal(t, deckClone.DiscardedPile.Cards[2].CardType, util.Market)
}

func Test_Play_Market(t *testing.T) {
	handCard := []util.CardType{util.Cellar, util.Market, util.Oasis, util.Mine, util.Remodel}
	data := setUpPlayAction(handCard)

	followUpIndices := []uint8{0, 1, 2}
	clone := _clone(data)
	clone.PlayCard(0, uint8(1), followUpIndices, util.NONE)

	deckClone := clone.CardDecks[0]

	data.CardDecks[0].Resources[util.PlayableCards] = 1
	data.CardDecks[0].Resources[util.DrawableCards] += 1
	data.CardDecks[0].Resources[util.SpendableMoney] += 1
	data.CardDecks[0].Resources[util.BuyableCards] += 1

	assert.Equal(t, data.CardDecks[0].MainPile, deckClone.MainPile)
	assert.Equal(t, data.CardDecks[0].Resources, deckClone.Resources)

	assert.Equal(t, len(deckClone.PlayedPile.Cards), 1)
	assert.Equal(t, deckClone.PlayedPile.Cards[0].CardType, util.Market)

	assert.Equal(t, len(deckClone.HandPile.Cards), 4)
	assert.Equal(t, deckClone.HandPile.Cards[0].CardType, util.Cellar)
	assert.Equal(t, deckClone.HandPile.Cards[1].CardType, util.Oasis)

	assert.Equal(t, len(deckClone.DiscardedPile.Cards), 0)
}

func Test_Play_Oasis(t *testing.T) {
	handCard := []util.CardType{util.Cellar, util.Market, util.Oasis, util.Mine, util.Remodel}
	data := setUpPlayAction(handCard)

	followUpIndices := []uint8{0}
	clone := _clone(data)
	clone.PlayCard(0, uint8(2), followUpIndices, util.NONE)

	deckClone := clone.CardDecks[0]

	data.CardDecks[0].Resources[util.PlayableCards] = 1
	data.CardDecks[0].Resources[util.DrawableCards] += 1
	data.CardDecks[0].Resources[util.SpendableMoney] += 1

	assert.Equal(t, data.CardDecks[0].MainPile, deckClone.MainPile)
	assert.Equal(t, data.CardDecks[0].Resources, deckClone.Resources)

	assert.Equal(t, len(deckClone.PlayedPile.Cards), 1)
	assert.Equal(t, deckClone.PlayedPile.Cards[0].CardType, util.Oasis)

	assert.Equal(t, len(deckClone.HandPile.Cards), 3)
	assert.Equal(t, deckClone.HandPile.Cards[0].CardType, util.Market)
	assert.Equal(t, deckClone.HandPile.Cards[1].CardType, util.Mine)
	assert.Equal(t, deckClone.HandPile.Cards[2].CardType, util.Remodel)

	assert.Equal(t, len(deckClone.DiscardedPile.Cards), 1)
	assert.Equal(t, deckClone.DiscardedPile.Cards[0].CardType, util.Cellar)
}

func Test_Play_Remodel(t *testing.T) {
	handCard := []util.CardType{util.Cellar, util.Market, util.Oasis, util.Mine, util.Remodel}
	data := setUpPlayAction(handCard)

	followUpIndices := []uint8{0}
	clone := _clone(data)
	clone.PlayCard(0, uint8(4), followUpIndices, util.Smithy)

	deckClone := clone.CardDecks[0]

	data.CardDecks[0].Resources[util.PlayableCards] = 0

	assert.Equal(t, data.CardDecks[0].MainPile, deckClone.MainPile)
	assert.Equal(t, data.CardDecks[0].Resources, deckClone.Resources)

	assert.Equal(t, len(deckClone.PlayedPile.Cards), 1)
	assert.Equal(t, deckClone.PlayedPile.Cards[0].CardType, util.Remodel)

	assert.Equal(t, len(deckClone.HandPile.Cards), 3)
	assert.Equal(t, deckClone.HandPile.Cards[0].CardType, util.Market)
	assert.Equal(t, deckClone.HandPile.Cards[1].CardType, util.Oasis)
	assert.Equal(t, deckClone.HandPile.Cards[2].CardType, util.Mine)

	assert.Equal(t, len(deckClone.DiscardedPile.Cards), 1)
	assert.Equal(t, clone.Stock.Trash[util.Cellar], uint8(1))
	assert.Equal(t, deckClone.DiscardedPile.Cards[0].CardType, util.Smithy)
}

func Test_Play_Mine(t *testing.T) {
	handCard := []util.CardType{util.Copper, util.Market, util.Oasis, util.Mine, util.Remodel}
	data := setUpPlayAction(handCard)

	followUpIndices := []uint8{0}
	clone := _clone(data)
	clone.PlayCard(0, uint8(3), followUpIndices, util.Silver)

	deckClone := clone.CardDecks[0]

	data.CardDecks[0].Resources[util.PlayableCards] = 0

	assert.Equal(t, data.CardDecks[0].MainPile, deckClone.MainPile)
	assert.Equal(t, data.CardDecks[0].Resources, deckClone.Resources)

	assert.Equal(t, len(deckClone.PlayedPile.Cards), 1)
	assert.Equal(t, deckClone.PlayedPile.Cards[0].CardType, util.Mine)

	assert.Equal(t, len(deckClone.HandPile.Cards), 4)
	assert.Equal(t, deckClone.HandPile.Cards[0].CardType, util.Market)
	assert.Equal(t, deckClone.HandPile.Cards[1].CardType, util.Oasis)
	assert.Equal(t, deckClone.HandPile.Cards[2].CardType, util.Remodel)
	assert.Equal(t, deckClone.HandPile.Cards[3].CardType, util.Silver)

	assert.Equal(t, len(deckClone.DiscardedPile.Cards), 0)

	assert.Equal(t, clone.Stock.Trash[util.Copper], uint8(1))
}
func Test_Play_Smithy(t *testing.T) {
	handCard := []util.CardType{util.Copper, util.Market, util.Oasis, util.Mine, util.Smithy}
	data := setUpPlayAction(handCard)

	followUpIndices := []uint8{}
	clone := _clone(data)
	clone.PlayCard(0, uint8(4), followUpIndices, util.NONE)

	deckClone := clone.CardDecks[0]

	data.CardDecks[0].Resources[util.PlayableCards] = 0
	data.CardDecks[0].Resources[util.DrawableCards] += 3

	assert.Equal(t, data.CardDecks[0].MainPile, deckClone.MainPile)
	assert.Equal(t, data.CardDecks[0].Resources, deckClone.Resources)

	assert.Equal(t, len(deckClone.PlayedPile.Cards), 1)
	assert.Equal(t, deckClone.PlayedPile.Cards[0].CardType, util.Smithy)

	assert.Equal(t, len(deckClone.HandPile.Cards), 4)
	assert.Equal(t, deckClone.HandPile.Cards[0].CardType, util.Copper)
	assert.Equal(t, deckClone.HandPile.Cards[1].CardType, util.Market)
	assert.Equal(t, deckClone.HandPile.Cards[2].CardType, util.Oasis)
	assert.Equal(t, deckClone.HandPile.Cards[3].CardType, util.Mine)

	assert.Equal(t, len(deckClone.DiscardedPile.Cards), 0)
}
func Test_Play_Village(t *testing.T) {
	handCard := []util.CardType{util.Copper, util.Market, util.Oasis, util.Mine, util.Village}
	data := setUpPlayAction(handCard)

	followUpIndices := []uint8{}
	clone := _clone(data)
	clone.PlayCard(0, uint8(4), followUpIndices, util.NONE)

	deckClone := clone.CardDecks[0]

	data.CardDecks[0].Resources[util.PlayableCards] = 2
	data.CardDecks[0].Resources[util.DrawableCards] += 1

	assert.Equal(t, data.CardDecks[0].MainPile, deckClone.MainPile)
	assert.Equal(t, data.CardDecks[0].Resources, deckClone.Resources)

	assert.Equal(t, len(deckClone.PlayedPile.Cards), 1)
	assert.Equal(t, deckClone.PlayedPile.Cards[0].CardType, util.Village)

	assert.Equal(t, len(deckClone.HandPile.Cards), 4)
	assert.Equal(t, deckClone.HandPile.Cards[0].CardType, util.Copper)
	assert.Equal(t, deckClone.HandPile.Cards[1].CardType, util.Market)
	assert.Equal(t, deckClone.HandPile.Cards[2].CardType, util.Oasis)
	assert.Equal(t, deckClone.HandPile.Cards[3].CardType, util.Mine)

	assert.Equal(t, len(deckClone.DiscardedPile.Cards), 0)
}
func Test_Play_Workshop(t *testing.T) {
	handCard := []util.CardType{util.Copper, util.Market, util.Oasis, util.Mine, util.Workshop}
	data := setUpPlayAction(handCard)

	followUpIndices := []uint8{}
	clone := _clone(data)
	clone.PlayCard(0, uint8(4), followUpIndices, util.Remodel)

	deckClone := clone.CardDecks[0]

	data.CardDecks[0].Resources[util.PlayableCards] = 0

	assert.Equal(t, data.CardDecks[0].MainPile, deckClone.MainPile)
	assert.Equal(t, data.CardDecks[0].Resources, deckClone.Resources)

	assert.Equal(t, len(deckClone.PlayedPile.Cards), 1)
	assert.Equal(t, deckClone.PlayedPile.Cards[0].CardType, util.Workshop)

	assert.Equal(t, len(deckClone.HandPile.Cards), 4)
	assert.Equal(t, deckClone.HandPile.Cards[0].CardType, util.Copper)
	assert.Equal(t, deckClone.HandPile.Cards[1].CardType, util.Market)
	assert.Equal(t, deckClone.HandPile.Cards[2].CardType, util.Oasis)
	assert.Equal(t, deckClone.HandPile.Cards[3].CardType, util.Mine)

	assert.Equal(t, len(deckClone.DiscardedPile.Cards), 1)
	assert.Equal(t, deckClone.DiscardedPile.Cards[0].CardType, util.Remodel)
}
func Test_Play_Chapel(t *testing.T) {
	handCard := []util.CardType{util.Copper, util.Market, util.Oasis, util.Mine, util.Chapel}
	data := setUpPlayAction(handCard)

	followUpIndices := []uint8{0, 1, 2, 3}
	clone := _clone(data)
	clone.PlayCard(0, uint8(4), followUpIndices, util.Remodel)

	deckClone := clone.CardDecks[0]

	data.CardDecks[0].Resources[util.PlayableCards] = 0

	assert.Equal(t, data.CardDecks[0].MainPile, deckClone.MainPile)
	assert.Equal(t, data.CardDecks[0].Resources, deckClone.Resources)

	assert.Equal(t, len(deckClone.PlayedPile.Cards), 1)
	assert.Equal(t, deckClone.PlayedPile.Cards[0].CardType, util.Chapel)

	assert.Equal(t, len(deckClone.HandPile.Cards), 0)

	assert.Equal(t, len(deckClone.DiscardedPile.Cards), 0)
	assert.Equal(t, clone.Stock.Trash[util.Market], uint8(1))
	assert.Equal(t, clone.Stock.Trash[util.Oasis], uint8(1))
	assert.Equal(t, clone.Stock.Trash[util.Mine], uint8(1))
	assert.Equal(t, clone.Stock.Trash[util.Copper], uint8(1))

}
func Test_Play_Feast(t *testing.T) {
	handCard := []util.CardType{util.Copper, util.Market, util.Oasis, util.Mine, util.Feast}
	data := setUpPlayAction(handCard)

	followUpIndices := []uint8{}
	clone := _clone(data)
	clone.PlayCard(0, uint8(4), followUpIndices, util.Mine)

	deckClone := clone.CardDecks[0]

	data.CardDecks[0].Resources[util.PlayableCards] = 0

	assert.Equal(t, data.CardDecks[0].MainPile, deckClone.MainPile)
	assert.Equal(t, data.CardDecks[0].Resources, deckClone.Resources)

	assert.Equal(t, len(deckClone.PlayedPile.Cards), 0)

	assert.Equal(t, len(deckClone.HandPile.Cards), 4)
	assert.Equal(t, deckClone.HandPile.Cards[0].CardType, util.Copper)
	assert.Equal(t, deckClone.HandPile.Cards[1].CardType, util.Market)
	assert.Equal(t, deckClone.HandPile.Cards[2].CardType, util.Oasis)
	assert.Equal(t, deckClone.HandPile.Cards[3].CardType, util.Mine)

	assert.Equal(t, len(deckClone.DiscardedPile.Cards), 1)
	assert.Equal(t, deckClone.DiscardedPile.Cards[0].CardType, util.Mine)

	assert.Equal(t, clone.Stock.Trash[util.Feast], uint8(1))
}
