// Copyright 2021 PolyCrypt GmbH, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain ImageA copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package app

import (
	"fmt"
	"io"
	"perun.network/go-perun/channel"
	"perun.network/go-perun/wallet"
	"perun.network/perun-examples/app-channel/app/util"
	"perun.network/perun-examples/app-channel/global"
)

// DominionApp is ImageA channel app.
type DominionApp struct {
	Addr wallet.Address
}

var ErrorConst = util.ErrorConstAPP

func NewDominionApp(addr wallet.Address) *DominionApp {
	return &DominionApp{
		Addr: addr,
	}
}

// Def returns the app address.
// required for App - port
func (a *DominionApp) Def() wallet.Address {
	return a.Addr
}

// DecodeData decodes the data.
// required for App - interface
func (a *DominionApp) DecodeData(r io.Reader) (channel.Data, error) {
	errorDescription := util.ErrorInfo{FunctionName: "DecodeData", FileName: util.ErrorConstAPP}
	appData := DominionAppData{}

	// Next Actor
	err := util.ReadObject(r, &appData.turn)
	if err != nil {
		return nil, util.ForwardError(errorDescription, err)
	}

	// Deck data
	for deckIndex := 0; deckIndex < util.NumPlayers; deckIndex++ {
		err := util.ReadObject(r, &appData.CardDecks[deckIndex])
		if err != nil {
			return nil, util.ForwardError(errorDescription, err)
		}
	}

	// RNG data
	err = util.ReadObject(r, &appData.rng)
	if err != nil {
		return nil, util.ForwardError(errorDescription, err)
	}

	return &appData, nil
}

// ValidTransition is called whenever the channel state transitions
// perform required validation steps
// required for StateApp - interface
func (a *DominionApp) ValidTransition(params *channel.Params, from, to *channel.State, idx channel.Index) error {
	errorDescription := util.ErrorInfo{FunctionName: "ValidTransition", FileName: util.ErrorConstAPP}

	// TODO KP WAS DAS DAS CHECKT
	err := channel.AssetsAssertEqual(from.Assets, to.Assets)
	if err != nil {
		return util.ForwardError(errorDescription, err)
	}

	// correct fromData format
	fromData, ok := from.Data.(*DominionAppData)
	if !ok {
		return util.ThrowError(errorDescription, fmt.Sprintf("fromData is in an invalid data format %T", fromData))
	}

	// correct toData format
	toData, ok := from.Data.(*DominionAppData)
	if !ok {
		return util.ThrowError(errorDescription, fmt.Sprintf("toData is in an invalid data format %T", toData))
	}

	// correct count of player
	if len(params.Parts) != util.NumPlayers {
		return util.ThrowError(errorDescription, fmt.Sprintf("Player count is not correct. should be %v but is %v", util.NumPlayers, len(params.Parts)))
	}

	fromDataClone := &(*fromData)
	switch toData.turn.performedAction {
	case util.RngCommit:
		dumPreImage := global.RandomBytes(util.HashSize)
		err = fromDataClone.RngCommit(idx, dumPreImage)
		if err != nil {
			return util.ForwardError(errorDescription, err)
		}
		fromDataClone.RngCommit(idx, dumPreImage)

		// replace dummy
		fromDataClone.rng.ImageA = toData.rng.ImageA

		if fromDataClone == toData {
			return util.ThrowError(errorDescription, "State transition could not be replicated for RngCommit")
		}
	}

	// TODO  ACTOR Check. -> if nextActor != CalcNextActor(currentActor)...
	// TODO CHECK FINAL
	//TODO CHECK STATE

	return nil
}

// ValidInit perform validation checks for initial game set up.
// required for StateApp - interface
func (a *DominionApp) ValidInit(p *channel.Params, s *channel.State) error {
	errorDescription := util.ErrorInfo{FunctionName: "ValidInit", FileName: util.ErrorConstAPP}

	// correct count of player
	if len(p.Parts) != util.NumPlayers {
		return util.ThrowError(errorDescription, fmt.Sprintf("Player count is not correct. should be %v but is %v", util.NumPlayers, len(p.Parts)))
	}
	// correct app format
	d, ok := s.Data.(*DominionAppData)
	if !ok {
		return util.ThrowError(errorDescription, fmt.Sprintf("AppData is in an invalid data format %T", d))
	}

	// game not final
	if s.IsFinal {
		return util.ThrowError(errorDescription, "game state must not be final on init")
	}

	//valid next actor
	if d.turn.nextActor >= util.NumPlayers {
		return util.ThrowError(errorDescription, fmt.Sprintf("Next actor is not valid. Should be smaller than %v but is %v", util.NumPlayers, d.turn.nextActor))
	}

	// TODO CHECK if initData(Mit d.nextActor == d)
	return nil
}

func (a *DominionApp) InitData(firstActor channel.Index) *DominionAppData {
	// TODO Clean
	var ad DominionAppData
	ad.Init(firstActor) // TODO ERROR handle
	return &ad
}

//------------------------ RNG ------------------------

// RngCommit player who wants to DrawOneCard commit to an preimage by setting corresponding image
func (a *DominionApp) RngCommit(s *channel.State, actorIdx channel.Index, preImage []byte) error {
	errorDescription := util.ErrorInfo{FunctionName: "RngCommit", FileName: util.ErrorConstAPP}
	d, ok := s.Data.(*DominionAppData)
	if !ok {
		return util.ThrowError(errorDescription, fmt.Sprintf("AppData is in an invalid data format %T", d))
	}

	err := d.RngCommit(actorIdx, preImage)
	if err != nil {
		return util.ForwardError(errorDescription, err)
	}
	return nil
}

// RngTouch the player how doesn't DrawOneCard choose an image
func (a *DominionApp) RngTouch(s *channel.State, actorIdx channel.Index) error {
	errorDescription := util.ErrorInfo{FunctionName: "RngTouch", FileName: util.ErrorConstAPP}
	d, ok := s.Data.(*DominionAppData)
	if !ok {
		return util.ThrowError(errorDescription, fmt.Sprintf("AppData is in an invalid data format %T", d))
	}

	err := d.RngTouch(actorIdx)
	if err != nil {
		return util.ForwardError(errorDescription, err)
	}
	return nil
}

// RngRelease player who wants to DrawOneCard publish preimage for published image
func (a *DominionApp) RngRelease(s *channel.State, actorIdx channel.Index, image []byte) error {
	errorDescription := util.ErrorInfo{FunctionName: "RngRelease", FileName: util.ErrorConstAPP}
	d, ok := s.Data.(*DominionAppData)
	if !ok {
		return util.ThrowError(errorDescription, fmt.Sprintf("AppData is in an invalid data format %T", d))
	}

	err := d.RngRelease(actorIdx, image)
	if err != nil {
		return util.ForwardError(errorDescription, err)
	}
	return nil
}

//------------------------ Drawing ------------------------

// DrawOneCard draws one card to the hand pile. A full rng need to be performed before.
func (a *DominionApp) DrawOneCard(s *channel.State, actorIdx channel.Index) error {
	errorDescription := util.ErrorInfo{FunctionName: "DrawOneCard", FileName: util.ErrorConstAPP}

	d, ok := s.Data.(*DominionAppData)
	if !ok {
		return util.ThrowError(errorDescription, fmt.Sprintf("AppData is in an invalid data format %T", d))
	}

	err := d.DrawOneCard(actorIdx)
	if err != nil {
		return util.ForwardError(errorDescription, err)
	}
	return nil
}

// PlayCard draws one card to the hand pile. A full rng need to be performed before.
func (a *DominionApp) PlayCard(s *channel.State, actorIdx channel.Index, index uint8) error {
	errorDescription := util.ErrorInfo{FunctionName: "PlayCard", FileName: util.ErrorConstAPP}

	d, ok := s.Data.(*DominionAppData)
	if !ok {
		return util.ThrowError(errorDescription, fmt.Sprintf("AppData is in an invalid data format %T", d))
	}

	err := d.PlayCard(actorIdx, index)
	if err != nil {
		return util.ForwardError(errorDescription, err)
	}
	return nil
}

//------------------------ General turn mechanics ------------------------

// EndTurn ends the current turn
func (a *DominionApp) EndTurn(s *channel.State, actorIdx channel.Index) error {
	errorDescription := util.ErrorInfo{FunctionName: "EndTurn", FileName: util.ErrorConstAPP}
	d, ok := s.Data.(*DominionAppData)
	if !ok {
		return util.ThrowError(errorDescription, fmt.Sprintf("AppData is in an invalid data format %T", d))
	}

	err := d.EndTurn(actorIdx)
	if err != nil {
		return util.ForwardError(errorDescription, err)
	}

	return nil
}

// EndGame ends game
func (a *DominionApp) EndGame(s *channel.State, actorIdx channel.Index) error {
	errorDescription := util.ErrorInfo{FunctionName: "EndGame", FileName: util.ErrorConstAPP}
	d, ok := s.Data.(*DominionAppData)
	if !ok {
		return util.ThrowError(errorDescription, fmt.Sprintf("AppData is in an invalid data format %T", d))
	}

	err := d.EndGame(actorIdx)
	if err != nil {
		return util.ForwardError(errorDescription, err)
	}

	s.IsFinal = true
	s.Balances = ComputeFinalBalances(s.Balances)

	return nil
}
