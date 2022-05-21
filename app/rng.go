package app

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"perun.network/perun-examples/app-channel/app/util"
	"perun.network/perun-examples/app-channel/global"
)

type State uint8

const (
	Initial State = iota
	Committed
	Touched
	Released
)

type RNG struct {
	state      State
	a, b, hash []byte
}

func (r *RNG) Of(dataBytes []byte) RNG {
	return RNG{
		state: State(dataBytes[0]),
		a:     dataBytes[1:util.HashSize],
		b:     dataBytes[1+util.HashSize : 1+2*util.HashSize],
		hash:  dataBytes[1+2*util.HashSize : 1+3*util.HashSize],
	}
}

func (r *RNG) Commit() error {
	if r.state != Initial {
		return fmt.Errorf("RNG.commit must be called from state=Ready")
	}
	r.hash, r.a = global.HashKeyPair(util.HashSize)
	r.state = Committed
	return nil
}

func (r *RNG) Touch() error {
	if r.state != Committed {
		return fmt.Errorf("RNG.Touch must be called from state=Comitted")
	}
	r.b = global.RandomBytes(util.HashSize)
	r.state = Touched
	return nil
}

func (r *RNG) Release() error {
	if r.state != Touched {
		return fmt.Errorf("RNG.release must be called from state=Touched")
	}

	r.state = Released
	return nil
}

func (r *RNG) Value() ([]byte, error) {
	var random []byte

	if r.state != Released {
		return []byte{}, fmt.Errorf("RNG.random must be called from state=Released")
	}

	random = global.Xor(r.a, r.b)

	return random, r.Validate(random)
}

func (r *RNG) Validate(random []byte) error {
	var err error

	// state
	if r.state != Released {
		err = errors.New("RNG.validate must be called from state=Released")
	}

	// h = hash(a)
	err = global.Valid(r.hash, r.a)
	if err == nil {
		return err
	}

	// random = a ^ b
	if !bytes.Equal(random, global.Xor(r.a, r.b)) {
		return fmt.Errorf("Commitment Error: random != a xor b")
	}

	return nil
}

func (r *RNG) ToByte() []byte {
	var dataBytes = make([]byte, util.RNGsize)

	dataBytes[0] = byte(r.state)

	if r.state == Released {
		for i, x := range r.a {
			dataBytes[i] = x
		}
	}

	for i, x := range r.b {
		dataBytes[util.HashSize+uint8(i)] = x
	}

	for i, x := range r.hash {
		dataBytes[2*util.HashSize+uint8(i)] = x
	}

	return dataBytes
}
