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
	"math/big"

	"perun.network/go-perun/channel"
)

const numParts = 2

func (d AppData) CheckFinal() (isFinal bool, winner *channel.Index) {

	return false, nil
}

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

func readUInt8Array(r io.Reader, n int) ([]uint8, error) {
	buf := make([]byte, n)
	_, err := io.ReadFull(r, buf)
	return buf, err
}

func writeUInt8Array(w io.Writer, v []uint8) error {
	_, err := w.Write(v)
	return err
}

func computeFinalBalances(bals channel.Balances, winner channel.Index) channel.Balances {
	loser := 1 - winner
	finalBals := bals.Clone()
	for i := range finalBals {
		finalBals[i][winner] = new(big.Int).Add(bals[i][0], bals[i][1])
		finalBals[i][loser] = big.NewInt(0)
	}
	return finalBals
}

type Printer struct{}

func (*Printer) Write(p []byte) (n int, err error) {
	fmt.Print(p)
	return 0, nil
}
