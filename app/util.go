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
	"io"

	"perun.network/go-perun/channel"
)

func uint8safe(a uint16) uint8 {
	b := uint8(a)
	if uint16(b) != a {
		panic("unsafe")
	}
	return b
}

func readUInt8(r io.Reader) (uint8, error) {
	buf := make([]byte, 1)
	_, err := io.ReadFull(r, buf)
	return buf[0], err
}

func writeUInt8(w io.Writer, v uint8) error {
	_, err := w.Write([]byte{v})
	return err
}

func readInt8(r io.Reader) (int8, error) {
	buf := make([]byte, 1)
	_, err := io.ReadFull(r, buf)
	return int8(buf[0]), err
}

func writeInt8(w io.Writer, v int8) error {
	_, err := w.Write([]byte{uint8(v)})
	return err
}

func writeActionCard(w io.Writer, v ActionCardType) error {
	_, err := w.Write([]byte{
		byte(v.Name()),
		byte(v.Group()),
		byte(v.Cost()),
		byte(v.dActions),
		byte(v.dBuys),
		byte(v.dDraws),
		byte(v.dMoney),
	})
	return err
}

func readActionCard(r io.Reader) (ActionCardType, error) {
	buf := make([]byte, 7)
	_, err := io.ReadFull(r, buf)
	return ActionCardType{
		CardTypeBase: CardTypeBase{CardName(buf[0]), CardGroup(buf[1]), buf[2]},
		dActions:     int8(buf[3]),
		dBuys:        int8(buf[4]),
		dDraws:       int8(buf[5]),
		dMoney:       int8(buf[6]),
	}, err
}

func computeFinalBalances(bals channel.Balances) channel.Balances {
	return bals.Clone()
}
