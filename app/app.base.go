package app

import (
	"encoding/json"
	"fmt"
	"log"
	"perun.network/go-perun/channel"
	"perun.network/perun-examples/app-channel/app/util"
	"perun.network/perun-examples/app-channel/global"
	"reflect"
)

// Init sets up initial game state
func (a *DominionApp) Init(firstActor channel.Index) *DominionAppData {
	var dominionAppData DominionAppData
	err := dominionAppData.Init(firstActor)
	if err != nil {
		log.Fatalf("Init DominionAppData failed with error: %T", err)
	}
	return &dominionAppData
}

// _ValidState validate the last performed action lead to current state
func (a *DominionApp) _ValidState(fromData, toData DominionAppData, idx channel.Index) error {
	errorInfo := util.ErrorInfo{FunctionName: "_ValidState", FileName: util.ErrorConstAPP}

	// Deep Copy. Seems like the other ways didn't work correctly ... idk
	origJSON, _ := json.Marshal(fromData)
	fromDataClone := DominionAppData{}
	json.Unmarshal(origJSON, &fromDataClone)

	var err error

	switch toData.Turn.PerformedAction {
	case util.RngCommit:
		if uint8(len(toData.Rng.ImageA)) != util.HashSize {
			err = errorInfo.ThrowError(fmt.Sprintf("given Image has not correct size of %d", util.HashSize))
			break
		}
		dummyPreImage := global.RandomBytes(util.HashSize)
		err = fromDataClone.RngCommit(idx, dummyPreImage)
		fromDataClone.Rng.ImageA = toData.Rng.ImageA
		break
	case util.RngTouch:
		if uint8(len(toData.Rng.PreImageB)) != util.HashSize {
			err = errorInfo.ThrowError(fmt.Sprintf("given preImage has not correct size of %d", util.HashSize))
			break
		}
		err = fromDataClone.RngTouch(idx)
		fromDataClone.Rng.PreImageB = toData.Rng.PreImageB
		break
	case util.RngRelease:
		if uint8(len(toData.Rng.PreImageA)) != util.HashSize {
			err = errorInfo.ThrowError(fmt.Sprintf("given preImage has not correct size of %d", util.HashSize))
			break
		}
		err = fromDataClone.RngRelease(idx, toData.Rng.PreImageA)
		break
	case util.DrawCard:
		err = fromDataClone.DrawCard(idx)
		break
	case util.PlayCard:
		playedIndex := 0
		toCards := toData.CardDecks[idx].HandPile.Cards
		fromCards := fromDataClone.CardDecks[idx].HandPile.Cards
		if len(fromCards) <= len(toCards) {
			err = errorInfo.ThrowError("toCards is to long. No card was played")
			break
		}
		for i := 0; i < len(fromCards); i++ {
			if i == len(toCards) || toCards[i].CardType != fromCards[i].CardType {
				playedIndex = i
				break
			}
		}
		err = fromDataClone.PlayCard(idx, uint8(playedIndex))
		break
	case util.BuyCard:
		cardTypeToBuy := toData.CardDecks[idx].DiscardedPile.Cards[toData.CardDecks[idx].DiscardedPile.Length()-1].CardType
		err = fromDataClone.BuyCard(idx, cardTypeToBuy)
		break
	case util.EndTurn:
		err = fromDataClone.EndTurn(idx)
		break
	case util.GameEnd:
		err = fromDataClone.EndGame(idx)
		break
	}

	if err != nil {
		return errorInfo.ForwardError(err)
	}

	if !reflect.DeepEqual(fromDataClone, toData) {
		return errorInfo.ThrowError("State transition could not be replicated for action")
	}
	return nil
}

// EndTurn ends current Turn by switching actors
func (a *DominionApp) EndTurn(s *channel.State, actorIdx channel.Index) error {
	errorInfo := util.ErrorInfo{FunctionName: "EndTurn", FileName: util.ErrorConstAPP}

	dominionAppData, ok := s.Data.(*DominionAppData)
	if !ok {
		return errorInfo.ThrowError(fmt.Sprintf("AppData is in an invalid data format %T", dominionAppData))
	}

	err := dominionAppData.EndTurn(actorIdx)
	if err != nil {
		return errorInfo.ForwardError(err)
	}

	return nil
}

// EndGame ends game by setting the isFinal value and Calculation final balance
func (a *DominionApp) EndGame(s *channel.State, actorIdx channel.Index) error {
	errorInfo := util.ErrorInfo{FunctionName: "EndGame", FileName: util.ErrorConstAPP}

	dominionAppData, ok := s.Data.(*DominionAppData)
	if !ok {
		return errorInfo.ThrowError(fmt.Sprintf("AppData is in an invalid data format %T", dominionAppData))
	}

	err := dominionAppData.EndGame(actorIdx)
	if err != nil {
		return errorInfo.ForwardError(err)
	}

	s.IsFinal = true
	s.Balances = ComputeFinalBalances(s.Balances) // TODO a real balance need to be calculated when final state could be reached

	return nil
}

func ComputeFinalBalances(b channel.Balances) channel.Balances {
	return b.Clone()
}

//------------------------ Decks ------------------------

// DrawCard draws one card to the hand pile.
// A Rng need to be performed before.
func (a *DominionApp) DrawCard(s *channel.State, actorIdx channel.Index) error {
	errorInfo := util.ErrorInfo{FunctionName: "DrawCard", FileName: util.ErrorConstAPP}

	dominionAppData, ok := s.Data.(*DominionAppData)
	if !ok {
		return errorInfo.ThrowError(fmt.Sprintf("AppData is in an invalid data format %T", dominionAppData))
	}

	err := dominionAppData.DrawCard(actorIdx)
	if err != nil {
		return errorInfo.ForwardError(err)
	}
	return nil
}

// PlayCard plays one card of the hand pile.
func (a *DominionApp) PlayCard(s *channel.State, actorIdx channel.Index, index uint8) error {
	errorInfo := util.ErrorInfo{FunctionName: "PlayCard", FileName: util.ErrorConstAPP}

	dominionAppData, ok := s.Data.(*DominionAppData)
	if !ok {
		return errorInfo.ThrowError(fmt.Sprintf("AppData is in an invalid data format %T", dominionAppData))
	}

	err := dominionAppData.PlayCard(actorIdx, index)
	if err != nil {
		return errorInfo.ForwardError(err)
	}
	return nil
}

// BuyCard Buy one card for given CardType.
func (a *DominionApp) BuyCard(s *channel.State, actorIdx channel.Index, cardType util.CardType) error {
	errorInfo := util.ErrorInfo{FunctionName: "BuyCard", FileName: util.ErrorConstAPP}

	dominionAppData, ok := s.Data.(*DominionAppData)
	if !ok {
		return errorInfo.ThrowError(fmt.Sprintf("BuyCard is in an invalid data format %T", dominionAppData))
	}

	err := dominionAppData.BuyCard(actorIdx, cardType)
	if err != nil {
		return errorInfo.ForwardError(err)
	}
	return nil
}

//------------------------ Rng ------------------------

// RngCommit set an image for Rng
// Players who want to draw a card need to start by committing to a preimage
func (a *DominionApp) RngCommit(s *channel.State, actorIdx channel.Index, preImage []byte) error {
	errorInfo := util.ErrorInfo{FunctionName: "RngCommit", FileName: util.ErrorConstAPP}

	dominionAppData, ok := s.Data.(*DominionAppData)
	if !ok {
		return errorInfo.ThrowError(fmt.Sprintf("AppData is in an invalid data format %T", dominionAppData))
	}

	err := dominionAppData.RngCommit(actorIdx, preImage)
	if err != nil {
		return errorInfo.ForwardError(err)
	}
	return nil
}

// RngTouch set a second preimage for Rng
// Players need accept the set image by selecting a second preimage
func (a *DominionApp) RngTouch(s *channel.State, actorIdx channel.Index) error {
	errorInfo := util.ErrorInfo{FunctionName: "RngTouch", FileName: util.ErrorConstAPP}

	dominionAppData, ok := s.Data.(*DominionAppData)
	if !ok {
		return errorInfo.ThrowError(fmt.Sprintf("AppData is in an invalid data format %T", dominionAppData))
	}

	err := dominionAppData.RngTouch(actorIdx)
	if err != nil {
		return errorInfo.ForwardError(err)
	}
	return nil
}

// RngRelease set preimage of set image
// Players publish their preimage of the image, s.t. a shared random value can be calculated
func (a *DominionApp) RngRelease(s *channel.State, actorIdx channel.Index, image []byte) error {
	errorInfo := util.ErrorInfo{FunctionName: "RngRelease", FileName: util.ErrorConstAPP}

	dominionAppData, ok := s.Data.(*DominionAppData)
	if !ok {
		return errorInfo.ThrowError(fmt.Sprintf("AppData is in an invalid data format %T", dominionAppData))
	}

	err := dominionAppData.RngRelease(actorIdx, image)
	if err != nil {
		return errorInfo.ForwardError(err)
	}
	return nil
}
