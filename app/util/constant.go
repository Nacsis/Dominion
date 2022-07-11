package util

type DeckResources uint8
type GeneralTypesOfActions uint8

const (
	PlayableCards DeckResources = iota
	DrawableCards
	BuyableCards
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

type CardType uint8

const (
	NONE CardType = iota
	Copper
	Silver
	Gold
	VictorySmall
	VictoryMid
	VictoryBig
	Cellar
	Market
	Oasis
	Mine
	Remodel
	Smithy
	Feast
	Chapel
	Workshop
	Village
)

const (

	//------------------------ Game Actions ------------------------
	GeneralTypesOfActionsCount = int(GameEnd) + 1

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
	DeckResourcesCount    = int(SpendableMoney) + 1

	//------------------------ Crypto / Framework ------------------------

	HashSizeByte     uint16 = 32
	PreImageSizeByte uint16 = 128

	//------------------------ Card ------------------------

	CardTypeCount                  = 16
	CopperInitialStock       uint8 = 60
	SilverInitialStock       uint8 = 40
	GoldInitialStock         uint8 = 30
	VictorySmallInitialStock uint8 = 24
	VictoryMidInitialStock   uint8 = 12
	VictoryBigInitialStock   uint8 = 12

	CopperMoneyValue uint8 = 1
	SilverMoneyValue uint8 = 2
	GoldMoneyValue   uint8 = 3

	CopperCost uint8 = 0
	SilverCost uint8 = 3
	GoldCost   uint8 = 6

	MoneyCardPlayCost uint8 = 0

	ActionCardPlayCost     uint8 = 1
	ActionCardVictoryPoint uint8 = 0

	VictorySmallVictoryValue uint8 = 1
	VictoryMidVictoryValue   uint8 = 3
	VictoryBigVictoryValue   uint8 = 6

	VictorySmallCost uint8 = 2
	VictoryMidCost   uint8 = 5
	VictoryBigCost   uint8 = 8

	VictoryCardPlayCost uint8 = 0

	//------------------------ Error const ------------------------

	ErrorConstRNG        = "rng"
	ErrorConstDATA       = "data"
	ErrorConstAPP        = "app"
	ErrorConstChannel    = "channel"
	ErrorConstPILE       = "pile"
	ErrorConstDECK       = "deck"
	ErrorConstCommitment = "commitment"
	ErrorConstStock      = "stock"
)
