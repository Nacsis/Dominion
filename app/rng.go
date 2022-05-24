package app

import (
	"bytes"
	"fmt"
	"perun.network/perun-examples/app-channel/app/util"
	"perun.network/perun-examples/app-channel/global"
)

type State uint8

type RNG struct {
	imageA, preImageB, preImageA []byte
}

func (r *RNG) Of(dataBytes []byte) {
	var size = dataBytes[0]
	if size == 0 {
		return
	}
	if size >= util.HashSize {
		r.imageA = dataBytes[1 : 1+util.HashSize] // 1 - 20
	}
	if size >= 2*util.HashSize {
		r.imageA = dataBytes[1+util.HashSize : 1+2*util.HashSize] // 21 - 40
	}
	if size >= 3*util.HashSize {
		r.imageA = dataBytes[1+2*util.HashSize : 1+3*util.HashSize] //40 - 60
	}
}
func (r *RNG) ToByte() []byte {

	// if imageA is not set end with rng length 0
	var dataBytes = make([]byte, 1)

	// TODO Muss ich das setzen oder es es automatisch drin?
	dataBytes[0] = 0

	if r.imageA != nil && uint8(len(r.imageA)) == util.HashSize {
		dataBytes = append(dataBytes, r.imageA...)
	}
	if r.preImageB != nil && uint8(len(r.preImageB)) == util.HashSize {
		dataBytes = append(dataBytes, r.preImageB...)
	}
	if r.preImageA != nil && uint8(len(r.preImageA)) == util.HashSize {
		dataBytes = append(dataBytes, r.preImageA...)
	}

	// To define how lang the arrays a
	dataBytes[0] = byte(len(dataBytes) - 1)
	return dataBytes
}

// Commit is the first function to call when the rng should start.
// First Commit to a imageA
func (r *RNG) Commit(preImage []byte) error {

	// ensure clean state
	r.preImageB = nil
	r.preImageA = nil
	r.imageA = nil

	r.imageA = global.Hash(preImage)

	return nil
}

// Touch is the second step to execute for rng.
// A committed imageA is necessary.
func (r *RNG) Touch() error {
	if r.imageA == nil {
		return fmt.Errorf("RNG.Touch must be called from state=Comitted")
	}
	r.preImageB = global.RandomBytes(util.HashSize)
	return nil
}

func (r *RNG) Release(preImageA []byte) error {
	if r.preImageB != nil || !global.IsValid(r.imageA, preImageA) {
		return fmt.Errorf("RNG.Release must be called from state=Touched")
	}
	r.preImageA = preImageA
	return nil
}

// Value return joined random value
func (r *RNG) Value() ([]byte, error) {
	var random []byte

	if r.preImageB != nil || !global.IsValid(r.imageA, r.preImageA) {
		return []byte{}, fmt.Errorf("RNG.random must be called from state=Released")
	}

	random = global.Xor(r.preImageA, r.preImageB)

	return random, r.Validate(random)
}

func (r *RNG) Validate(random []byte) error {

	// h = hash(imageA)
	err := global.Valid(r.preImageA, r.preImageA)
	if err == nil {
		return err
	}

	// random = imageA ^ preImageB
	if !bytes.Equal(random, global.Xor(r.preImageA, r.preImageB)) {
		return fmt.Errorf("Commitment Error: random != imageA xor preImageB")
	}

	return nil
}
