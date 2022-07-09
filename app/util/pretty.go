package util

import "fmt"

func (e GeneralTypesOfActions) String() string {
	switch e {
	case GameInit:
		return "GameInit"
	case RngCommit:
		return "RngCommit"
	case RngTouch:
		return "RngTouch"
	case RngRelease:
		return "RngRelease"
	case DrawCard:
		return "DrawCard"
	case PlayCard:
		return "PlayCard"
	case BuyCard:
		return "BuyCard"
	case EndTurn:
		return "EndTurn"
	case GameEnd:
		return "GameEnd"
	default:
		return fmt.Sprintf("%d", int(e))
	}
}

func PrettyPossibleActions(pa [GameEnd + 1]bool) []string {
	pas := make([]string, 0)
	for i, v := range pa {
		if v {
			pas = append(pas, GeneralTypesOfActions(i).String())
		}
	}
	return pas
}

func (e DeckResources) String() string {
	switch e {
	case PlayableCards:
		return "PlayableCards"
	case DrawableCards:
		return "DrawableCards"
	case BuyableCards:
		return "BuyableCards"
	case SpendableMoney:
		return "SpendableMoney"
	default:
		return fmt.Sprintf("%d", int(e))
	}
}

func PrettyResources(r [DeckResourcesCount]uint8) map[string]uint8 {
	m := make(map[string]uint8, 0)
	for i, v := range r {
		m[DeckResources(i).String()] = v
	}
	return m
}

func (e CardType) String() string {
	switch e {
	case NONE:
		return "NONE"
	case Copper:
		return "Copper"
	case Silver:
		return "Silver"
	case Gold:
		return "Gold"
	case VictorySmall:
		return "VictorySmall"
	case VictoryMid:
		return "VictoryMid"
	case VictoryBig:
		return "VictoryBig"
	case Cellar:
		return "Cellar"
	case Market:
		return "Market"
	case Oasis:
		return "Oasis"
	case Mine:
		return "Mine"
	case Remodel:
		return "Remodel"
	case Smithy:
		return "Smithy"
	case Feast:
		return "Feast"
	case Chapel:
		return "Chapel"
	case Workshop:
		return "Workshop"
	case Village:
		return "Village"
	default:
		return fmt.Sprintf("%d", int(e))
	}
}
