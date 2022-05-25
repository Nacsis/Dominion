// Copyright 2021 PolyCrypt GmbH, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain imageA copy of the License at
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
)

// DominionApp is imageA channel app.
type DominionApp struct {
	Addr wallet.Address
}

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

	appData := DominionAppData{}
	var err error

	// Next Actor
	appData.NextActor, err = util.ReadUInt8(r)

	if err != nil {
		return nil, util.ForwardError(util.ErrorConstAPP, "DecodeData", err)
	}

	// Deck data
	for deckIndex := 0; deckIndex < util.NumPlayers; deckIndex++ {
		err := util.ReadObject(r, &appData.CardDecks[deckIndex])
		if err != nil {
			return nil, util.ForwardError(util.ErrorConstAPP, "DecodeData", err)
		}
	}

	// RNG data
	err = util.ReadObject(r, &appData.rng)
	if err != nil {
		return nil, util.ForwardError(util.ErrorConstAPP, "DecodeData", err)
	}

	return &appData, nil
}

// ValidTransition is called whenever the channel state transitions
// perform required validation steps
// required for StateApp - interface
func (a *DominionApp) ValidTransition(params *channel.Params, from, to *channel.State, idx channel.Index) error {

	// TODO KP WAS DAS DAS CHECKT
	err := channel.AssetsAssertEqual(from.Assets, to.Assets)
	if err != nil {
		return util.ForwardError(util.ErrorConstAPP, "ValidTransition", err)
	}

	// correct fromData format
	fromData, ok := from.Data.(*DominionAppData)
	if !ok {
		return util.ThrowError(util.ErrorConstAPP, "RngCommit", fmt.Sprintf("fromData is in an invalid data format %T", fromData))
	}

	// correct toData format
	toData, ok := from.Data.(*DominionAppData)
	if !ok {
		return util.ThrowError(util.ErrorConstAPP, "RngCommit", fmt.Sprintf("toData is in an invalid data format %T", toData))
	}

	// correct count of player
	if len(params.Parts) != util.NumPlayers {
		return util.ThrowError(util.ErrorConstAPP, "ValidInit", fmt.Sprintf("Player count is not correct. should be %v but is %v", util.NumPlayers, len(params.Parts)))
	}

	// TODO  ACTOR Check. -> if nextActor != CalcNextActor(currentActor)...
	// TODO CHECK FINAL
	//TODO CHECK STATE

	return nil
}

// ValidInit perform validation checks for initial game set up.
// required for StateApp - interface
func (a *DominionApp) ValidInit(p *channel.Params, s *channel.State) error {

	// correct count of player
	if len(p.Parts) != util.NumPlayers {
		return util.ThrowError(util.ErrorConstAPP, "ValidInit", fmt.Sprintf("Player count is not correct. should be %v but is %v", util.NumPlayers, len(p.Parts)))
	}
	// correct app format
	d, ok := s.Data.(*DominionAppData)
	if !ok {
		return util.ThrowError(util.ErrorConstAPP, "RngCommit", fmt.Sprintf("AppData is in an invalid data format %T", d))
	}

	// game not final
	if s.IsFinal {
		return util.ThrowError(util.ErrorConstAPP, "ValidInit", "game state must not be final on init")
	}

	//valid next actor
	if d.NextActor >= util.NumPlayers {
		return util.ThrowError(util.ErrorConstAPP, "ValidInit", fmt.Sprintf("Next actor is not valid. Should be smaller than %v but is %v", util.NumPlayers, d.NextActor))
	}

	// TODO CHECK if initData(Mit d.NextActor == d)
	return nil
}

func (a *DominionApp) InitData(firstActor channel.Index) *DominionAppData {
	// TODO Clean
	var ad DominionAppData
	ad.Init(firstActor) // TODO ERROR handle
	return &ad
}

// TODO Just a test remove later
func (a *DominionApp) SwitchActor(s *channel.State, actorIdx channel.Index) error {
	d, ok := s.Data.(*DominionAppData)
	if !ok {
		return fmt.Errorf("invalid data type: %T", d)
	}

	d.switchActor(actorIdx)

	s.IsFinal = true
	s.Balances = ComputeFinalBalances(s.Balances)

	return nil
}

//------ RNG ------

// RngCommit player who wants to DrawOneCard commit to an preimage by setting corresponding image
func (a *DominionApp) RngCommit(s *channel.State, actorIdx channel.Index, image []byte) error {

	d, ok := s.Data.(*DominionAppData)
	if !ok {
		return util.ThrowError(util.ErrorConstAPP, "RngCommit", fmt.Sprintf("AppData is in an invalid data format %T", d))
	}

	err := d.RngCommit(actorIdx, image)
	if err != nil {
		return util.ForwardError(util.ErrorConstAPP, "RngCommit", err)
	}
	return nil
}

// RngTouch the player how doesn't DrawOneCard choose an image
func (a *DominionApp) RngTouch(s *channel.State, actorIdx channel.Index) error {

	d, ok := s.Data.(*DominionAppData)
	if !ok {
		return util.ThrowError(util.ErrorConstAPP, "RngTouch", fmt.Sprintf("AppData is in an invalid data format %T", d))
	}

	err := d.RngTouch(actorIdx)
	if err != nil {
		return util.ForwardError(util.ErrorConstAPP, "RngTouch", err)
	}
	return nil
}

// RngRelease player who wants to DrawOneCard publish preimage for published image
func (a *DominionApp) RngRelease(s *channel.State, actorIdx channel.Index, image []byte) error {

	d, ok := s.Data.(*DominionAppData)
	if !ok {
		return util.ThrowError(util.ErrorConstAPP, "RngRelease", fmt.Sprintf("AppData is in an invalid data format %T", d))
	}

	err := d.RngRelease(actorIdx, image)
	if err != nil {
		return util.ForwardError(util.ErrorConstAPP, "RngRelease", err)
	}
	return nil
}

// DrawOneCard draws one card to the hand pile. A full rng need to be performed before.
func (a *DominionApp) DrawOneCard(s *channel.State, actorIdx channel.Index) error {
	d, ok := s.Data.(*DominionAppData)
	if !ok {
		return util.ThrowError(util.ErrorConstAPP, "DrawOneCard", fmt.Sprintf("AppData is in an invalid data format %T", d))
	}

	err := d.DrawOneCard(actorIdx)
	if err != nil {
		return util.ForwardError(util.ErrorConstAPP, "RngRelease", err)
	}
	return nil
}
