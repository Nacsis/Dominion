package app

import (
	"fmt"
	"perun.network/perun-examples/app-channel/app/util"
	"perun.network/perun-examples/app-channel/global"
)

// Commit calculate and set  image A
func (r *RNG) Commit(preImage []byte) error {
	errorInfo := util.ErrorInfo{FunctionName: "Commit", FileName: util.ErrorConstRNG}

	if uint8(len(preImage)) != util.HashSize {
		return errorInfo.ThrowError(fmt.Sprintf("given preImage has not correct size of %d", util.HashSize))
	}

	r.PreImageB = nil
	r.PreImageA = nil
	r.ImageA = nil

	var err error
	r.ImageA, err = global.ToImage(preImage)
	if err != nil {
		return errorInfo.ForwardError(err)
	}

	return nil
}

// Touch update preimage B
func (r *RNG) Touch() error {
	errorInfo := util.ErrorInfo{FunctionName: "Touch", FileName: util.ErrorConstRNG}

	if r.ImageA == nil {
		return errorInfo.ThrowError("ImageA is not set")
	}

	r.PreImageB = global.RandomBytes(util.HashSize)
	return nil
}

// Release update preimage A
func (r *RNG) Release(preImageA []byte) error {
	errorInfo := util.ErrorInfo{FunctionName: "Release", FileName: util.ErrorConstRNG}

	if uint8(len(preImageA)) != util.HashSize {
		return errorInfo.ThrowError(fmt.Sprintf("given preImage has not correct size of %d", util.HashSize))
	}

	if r.PreImageB == nil {
		return errorInfo.ThrowError("PreImageB is not set")
	}

	err := global.ValidatePreImage(r.ImageA, preImageA)
	if err != nil {
		return errorInfo.ForwardError(err)
	}

	r.PreImageA = append([]byte(nil), preImageA...)
	return nil
}

// RNGValue return joined random value
func (r *RNG) RNGValue() ([]byte, error) {
	errorInfo := util.ErrorInfo{FunctionName: "RNGValue", FileName: util.ErrorConstRNG}

	if r.PreImageB == nil {
		return nil, errorInfo.ThrowError("PreImageB is not set")
	}

	err := global.ValidatePreImage(r.ImageA, r.PreImageA)
	if err != nil {
		return nil, errorInfo.ForwardError(err)
	}

	result, err := global.Xor(r.PreImageA, r.PreImageB)
	if err != nil {
		return nil, errorInfo.ForwardError(err)
	}
	return result, nil
}
