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
	"github.com/pkg/errors"
	"io"

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
func (a *DominionApp) Def() wallet.Address {
	return a.Addr
}

func (a *DominionApp) InitData(firstActor channel.Index) *Data {
	return &Data{
		NextActor: uint8(firstActor),
	}
}

// DecodeData decodes data specific to this application. This has to be
// defined on an application-level because every app can have completely
// different data; during decoding the application needs to be known to
// know how to decode the data.

// DecodeData decodes the channel data.
// Habe return type von (channel.Data, error)  abgeändert, zum testen, da ich nicht weiß wie das mit dem Interface geht
func (a *DominionApp) DecodeData(r io.Reader) (channel.Data, error) {
	d := Data{}

	var err error
	//Read next actor
	d.NextActor, err = readUInt8(r)
	if err != nil {
		return nil, errors.WithMessage(err, "reading actor")
	}
	//Read NumAllCards
	d.NumAllCards, err = readUInt8(r)
	if err != nil {
		return nil, errors.WithMessage(err, "reading NumAllCards")
	}

	//Read Cards
	for i := uint8(0); i < d.NumAllCards; i++ {
		d.AllCards[i] = ReadCard(r)
	}

	return &d, err
}

// ValidTransition is called whenever the channel state transitions.
func (a *DominionApp) ValidTransition(params *channel.Params, from, to *channel.State, idx channel.Index) error {
	//todo
	return nil
}

func GetPlayer(i uint8) Player {
	return Player{i} //todo
}
