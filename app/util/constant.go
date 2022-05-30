package util

type DeckResources uint8
type GeneralTypesOfActions uint8

const (
	PlayableCards DeckResources = iota
	DrawableCards
	PurchasableCards
	SpendableMoney
)

const (
	GameInit GeneralTypesOfActions = iota
	RngCommit
	RngTouch
	RngRelease
	DrawCard
	PlayCard
	BuyCard
	EndTurn
	GameEnd // must remain at last position
)

const (
	//------------------------ Game Properties ------------------------

	NumPlayers = 2

	//------------------------ Deck resources ------------------------

	InitialMoneyCards   = 5
	InitialVictoryCards = 3
	InitialDeckSize     = InitialMoneyCards + InitialVictoryCards

	InitialDrawResources  = 5
	InitialPlayResources  = 1
	InitialBuyResources   = 1
	InitialMoneyResources = 0
	DeckResourcesCount    = 4

	//------------------------ Crypto / Framework ------------------------

	HashSize uint8 = 20

	// Treasure Values
	MonValueCopper uint8 = 1
	MonValueSilver uint8 = 2
	MonValueGold   uint8 = 3

	// Card costs
	CostsCopper uint8 = 0
	CostsSilver uint8 = 3
	CostsGold   uint8 = 6

	//------------------------ Error const ------------------------

	ErrorConstRNG        = "rng"
	ErrorConstDATA       = "data"
	ErrorConstAPP        = "app"
	ErrorConstChannel    = "channel"
	ErrorConstPILE       = "pile"
	ErrorConstDECK       = "deck"
	ErrorConstCARD       = "card"
	ErrorConstCommitment = "commitment"
)
