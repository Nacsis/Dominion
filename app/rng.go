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

// Commit set image A
func (r *RNG) Commit(preImage []byte) error {

	if uint8(len(preImage)) != util.HashSize {
		return util.ThrowError(util.ErrorConstRNG, "Commit", fmt.Sprintf("given preImage has not correct size of %d", util.HashSize))
	}

	r.preImageB = nil
	r.preImageA = nil
	r.imageA = nil

	r.imageA = global.ToImage(preImage)

	return nil
}

// Touch update preimage B
func (r *RNG) Touch() error {

	if r.imageA == nil {
		return util.ThrowError(util.ErrorConstRNG, "Touch", "imageA is not set")
	}

	r.preImageB = global.RandomBytes(util.HashSize)
	return nil
}

// Release update preimage A
func (r *RNG) Release(preImageA []byte) error {
	if uint8(len(preImageA)) != util.HashSize {
		return util.ThrowError(util.ErrorConstRNG, "Release", fmt.Sprintf("given preImage has not correct size of %d", util.HashSize))
	}

	if r.preImageB == nil {
		return util.ThrowError(util.ErrorConstRNG, "Release", "preImageB is not set")
	}

	err := global.ValidatePreImage(r.imageA, preImageA)
	if err != nil {
		return util.ForwardError(util.ErrorConstRNG, "Release", err)
	}

	r.preImageA = append([]byte(nil), preImageA...)
	return nil
}

// CalcCorrespondingValue return joined random value
func (r *RNG) CalcCorrespondingValue() ([]byte, error) {
	if r.preImageB == nil {
		return nil, util.ThrowError(util.ErrorConstRNG, "CalcCorrespondingValue", "preImageB is not set")
	}

	err := global.ValidatePreImage(r.imageA, r.preImageA)
	if err != nil {
		return nil, util.ForwardError(util.ErrorConstRNG, "CalcCorrespondingValue", err)
	}

	result, err := global.Xor(r.preImageA, r.preImageB)
	if err != nil {
		return nil, util.ForwardError(util.ErrorConstRNG, "CalcCorrespondingValue", err)
	}
	return result, r.Validate(result)
}

// Validate value is same as CalcCorrespondingValue()
func (r *RNG) Validate(value []byte) error {
	err := global.ValidatePreImage(r.imageA, r.preImageA)
	if err != nil {
		return util.ForwardError(util.ErrorConstRNG, "Validate", err)
	}

	v, err := global.Xor(r.preImageA, r.preImageB)
	if err != nil {
		return util.ForwardError(util.ErrorConstRNG, "Validate", err)
	}

	if !bytes.Equal(value, v) {
		return util.ThrowError(util.ErrorConstRNG, "Commit", fmt.Sprintf("given value %v doesn't match CalcCorrespondingValue() result", value))
	}

	return nil
}
