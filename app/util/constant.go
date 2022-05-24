package util

const (
	// Game Properties
	NumPlayers          = 2
	CardSize            = 3
	InitialDeckSize     = 8
	InitialMoneyCards   = InitialDeckSize / 2
	InitialVictoryCards = InitialDeckSize / 2
	NumCardTypes        = 16

	// Treasure Values
	MonValueCopper uint8 = 1
	MonValueSilver uint8 = 2
	MonValueGold   uint8 = 3

	// Card costs
	CostsCopper uint8 = 0
	CostsSilver uint8 = 3
	CostsGold   uint8 = 6
)
