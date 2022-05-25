package util

const (
	// Game Properties
	NumPlayers          = 2
	InitialDeckSize     = 8
	InitialMoneyCards   = InitialDeckSize / 2
	InitialVictoryCards = InitialDeckSize / 2
	NumCardTypes        = 16

	// Crypto / Framework
	HashSize uint8 = 20
	RNGsize        = 3*HashSize + 1

	// Treasure Values
	MonValueCopper uint8 = 1
	MonValueSilver uint8 = 2
	MonValueGold   uint8 = 3

	// Card costs
	CostsCopper uint8 = 0
	CostsSilver uint8 = 3
	CostsGold   uint8 = 6

	// Error constant used for logging
	ErrorConstRNG  = "rng"
	ErrorConstDATA = "data"
	ErrorConstAPP  = "app"
	ErrorConstPILE = "pile"
	ErrorConstDECK = "deck"
	ErrorConstCARD = "card"
)
