package app

import (
	"fmt"
	"log"
	"perun.network/go-perun/channel"
	"perun.network/perun-examples/app-channel/app/util"
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

// EndTurn ends current turn by switching actors
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
// A rng need to be performed before.
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

// RngCommit set an image for rng
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

// RngTouch set a second preimage for rng
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
