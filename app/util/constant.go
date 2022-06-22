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

	HashSize     uint16 = 256 / 8
	PreImageSize uint16 = 1028

	//------------------------ Card ------------------------

	CardTypeCount                  = 16
	CopperInitialStock       uint8 = 30
	SilverInitialStock       uint8 = 30
	GoldInitialStock         uint8 = 30
	VictorySmallInitialStock uint8 = 30
	VictoryMidInitialStock   uint8 = 30
	VictoryBigInitialStock   uint8 = 30

	CopperMoneyValue uint8 = 1
	SilverMoneyValue uint8 = 2
	GoldMoneyValue   uint8 = 3

	CopperCost uint8 = 0
	SilverCost uint8 = 2
	GoldCost   uint8 = 6

	MoneyCardPlayCost uint8 = 0

	ActionCardPlayCost     uint8 = 1
	ActionCardVictoryPoint uint8 = 0

	VictorySmallVictoryValue uint8 = 1
	VictoryMidVictoryValue   uint8 = 2
	VictoryBigVictoryValue   uint8 = 3

	VictorySmallCost uint8 = 1
	VictoryMidCost   uint8 = 2
	VictoryBigCost   uint8 = 6

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
