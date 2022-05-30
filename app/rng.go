package app

import (
	"bytes"
	"fmt"
	"perun.network/perun-examples/app-channel/app/util"
	"perun.network/perun-examples/app-channel/global"
)

type State uint8

type RNG struct {
	ImageA, PreImageB, PreImageA []byte
}

func (r *RNG) Of(dataBytes []byte) {
	var size = uint8(len(dataBytes))

	if size == 0 {
		return
	}
	if size >= util.HashSize {
		r.ImageA = dataBytes[:util.HashSize] // 1 - 20
	}
	if size >= 2*util.HashSize {
		r.PreImageB = dataBytes[util.HashSize : 2*util.HashSize] // 21 - 40
	}
	if size >= 3*util.HashSize {
		r.PreImageA = dataBytes[2*util.HashSize : 3*util.HashSize] //40 - 60
	}
}
func (r *RNG) ToByte() []byte {

	// if ImageA is not set end with rng length 0
	var dataBytes = make([]byte, 0)

	if r.ImageA != nil && uint8(len(r.ImageA)) == util.HashSize {
		dataBytes = append(dataBytes, r.ImageA...)
	}
	if r.PreImageB != nil && uint8(len(r.PreImageB)) == util.HashSize {
		dataBytes = append(dataBytes, r.PreImageB...)
	}
	if r.PreImageA != nil && uint8(len(r.PreImageA)) == util.HashSize {
		dataBytes = append(dataBytes, r.PreImageA...)
	}

	return append([]byte{byte(len(dataBytes))}, dataBytes...)
}

// Commit set image A
func (r *RNG) Commit(preImage []byte) error {
	errorDescription := util.ErrorInfo{FunctionName: "Commit", FileName: util.ErrorConstRNG}

	if uint8(len(preImage)) != util.HashSize {
		return util.ThrowError(errorDescription, fmt.Sprintf("given preImage has not correct size of %d", util.HashSize))
	}

	r.PreImageB = nil
	r.PreImageA = nil
	r.ImageA = nil

	r.ImageA = global.ToImage(preImage)

	return nil
}

// Touch update preimage B
func (r *RNG) Touch() error {
	errorDescription := util.ErrorInfo{FunctionName: "Touch", FileName: util.ErrorConstRNG}

	if r.ImageA == nil {
		return util.ThrowError(errorDescription, "ImageA is not set")
	}

	r.PreImageB = global.RandomBytes(util.HashSize)
	return nil
}

// Release update preimage A
func (r *RNG) Release(preImageA []byte) error {
	errorDescription := util.ErrorInfo{FunctionName: "Release", FileName: util.ErrorConstRNG}

	if uint8(len(preImageA)) != util.HashSize {
		return util.ThrowError(errorDescription, fmt.Sprintf("given preImage has not correct size of %d", util.HashSize))
	}

	if r.PreImageB == nil {
		return util.ThrowError(errorDescription, "PreImageB is not set")
	}

	err := global.ValidatePreImage(r.ImageA, preImageA)
	if err != nil {
		return util.ForwardError(errorDescription, err)
	}

	r.PreImageA = append([]byte(nil), preImageA...)
	return nil
}

// CalcCorrespondingValue return joined random value
func (r *RNG) CalcCorrespondingValue() ([]byte, error) {
	errorDescription := util.ErrorInfo{FunctionName: "CalcCorrespondingValue", FileName: util.ErrorConstRNG}

	if r.PreImageB == nil {
		return nil, util.ThrowError(errorDescription, "PreImageB is not set")
	}

	err := global.ValidatePreImage(r.ImageA, r.PreImageA)
	if err != nil {
		return nil, util.ForwardError(errorDescription, err)
	}

	result, err := global.Xor(r.PreImageA, r.PreImageB)
	if err != nil {
		return nil, util.ForwardError(errorDescription, err)
	}
	return result, r.Validate(result)
}

// Validate value is same as CalcCorrespondingValue()
func (r *RNG) Validate(value []byte) error {
	errorDescription := util.ErrorInfo{FunctionName: "Validate", FileName: util.ErrorConstRNG}

	err := global.ValidatePreImage(r.ImageA, r.PreImageA)
	if err != nil {
		return util.ForwardError(errorDescription, err)
	}

	v, err := global.Xor(r.PreImageA, r.PreImageB)
	if err != nil {
		return util.ForwardError(errorDescription, err)
	}

	if !bytes.Equal(value, v) {
		return util.ThrowError(errorDescription, fmt.Sprintf("given value %v doesn't match CalcCorrespondingValue() result", value))
	}

	return nil
}
