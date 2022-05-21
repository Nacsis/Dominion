// Copyright 2021 PolyCrypt GmbH, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
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
	"log"
	"perun.network/go-perun/channel"
	"perun.network/go-perun/wallet"
	"perun.network/perun-examples/app-channel/app/util"
)

// DominionApp is a channel app.
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

// DecodeData decodes the channel data.
// required for App - interface
func (a *DominionApp) DecodeData(r io.Reader) (channel.Data, error) {
	dad := DominionAppData{}

	var err error
	dad.NextActor, err = util.ReadUInt8(r)

	for deckIndex := 0; deckIndex < util.NumPlayers; deckIndex++ {
		util.ReadObject(r, &dad.CardDecks[deckIndex])
	}

	return &dad, err
}

// ValidTransition is called whenever the channel state transitions.
// required for StateApp - interface
func (a *DominionApp) ValidTransition(params *channel.Params, from, to *channel.State, idx channel.Index) error {

	err := channel.AssetsAssertEqual(from.Assets, to.Assets)
	if err != nil {
		return fmt.Errorf("Invalid assets: %v", err)
	}

	fromData := ValidStateFormat(from)
	toData := ValidStateFormat(to)

	ValidActorInformation(fromData.NextActor, toData.NextActor, params.Parts, idx)

	return nil
}

// ValidInit should perform app-specific checks for a valid initial state.
// The framework guarantees to only pass initial states with version == 0,
// correct channel ID and valid initial allocation.
// required for StateApp - interface
func (a *DominionApp) ValidInit(p *channel.Params, s *channel.State) error {

	ValidWalletLen(p.Parts)

	appData := ValidStateFormat(s)
	log.Println(appData)

	if s.IsFinal {
		return fmt.Errorf("must not be final")
	}

	NextActorIsInRange(appData.NextActor)
	return nil
}

func (a *DominionApp) InitData(firstActor channel.Index) *DominionAppData {
	var ad DominionAppData
	ad.Init(firstActor)
	return &ad
}

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
