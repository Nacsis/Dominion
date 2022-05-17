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
// required for App - interface
func (a *DominionApp) Def() wallet.Address {
	return a.Addr
}

// DecodeData decodes the channel data.
// required for App - interface
func (a *DominionApp) DecodeData(r io.Reader) (channel.Data, error) {
	d := DominionAppData{}

	var err error
	d.NextActor, err = readUInt8(r)
	if err != nil {
		return nil, err
	}

	for i := uint8(0); i < (NumActionCardsInGame); i++ {
		d.ActionCardsInvolved[i], err = readActionCard(r)
	}
	if err != nil {
		return nil, err
	}

	for i := uint8(0); i < NumPlayers; i++ {
		d.LenCardDecks[i], err = readUInt8(r)
		if err != nil {
			return nil, err
		}
		d.LenCardHand[i], err = readUInt8(r)
		if err != nil {
			return nil, err
		}
		d.LenCardTrashs[i], err = readUInt8(r)
		if err != nil {
			return nil, err
		}
	}
	d.LenCardGrave, err = readUInt8(r)
	if err != nil {
		return nil, err
	}

	for i := uint16(0); i < NumMaxCirculation; i++ {
		d.CardsInCirculation[i], err = readInt8(r)
		if err != nil {
			return nil, err
		}
	}

	return &d, err
}

// ValidTransition is called whenever the channel state transitions.
// required for StateApp - interface
func (a *DominionApp) ValidTransition(params *channel.Params, from, to *channel.State, idx channel.Index) error {
	return nil
}

// ValidInit should perform app-specific checks for a valid initial state.
// The framework guarantees to only pass initial states with version == 0,
// correct channel ID and valid initial allocation.
// required for StateApp - interface
func (a *DominionApp) ValidInit(p *channel.Params, s *channel.State) error {
	appData, ok := s.Data.(*DominionAppData)
	if !ok {
		return fmt.Errorf("invalid data type: %T", s.Data)
	}
	log.Println(appData)
	return nil
}

func (a *DominionApp) InitData(firstActor channel.Index) *DominionAppData {
	return &DominionAppData{
		NextActor: uint8(firstActor),
		CardStock: [16]uint8{
			10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
			// 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, // TODO replace with above when Cards implemented
			NumSupplyProvince,
			NumSupplyDuchy,
			NumSupplyEstate,
			NumSupplyGold,
			NumSupplySilver,
			NumSupplyCopper,
		},
	}
}

func (a *DominionApp) SwitchActor(s *channel.State, actorIdx channel.Index) error {
	d, ok := s.Data.(*DominionAppData)
	if !ok {
		return fmt.Errorf("invalid data type: %T", d)
	}

	d.Set(actorIdx)

	s.IsFinal = true
	s.Balances = computeFinalBalances(s.Balances)

	return nil
}
