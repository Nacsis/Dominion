package app

import (
	"perun.network/perun-examples/dominion-cli/app/util"
	"perun.network/perun-examples/dominion-cli/global"
)

// Commit calculate and set  image A
func (r *RNG) Commit(preImage [util.PreImageSizeByte]byte) error {
	errorInfo := util.ErrorInfo{FunctionName: "Commit", FileName: util.ErrorConstRNG}

	r.PreImageB = [util.PreImageSizeByte]byte{}
	r.PreImageA = [util.PreImageSizeByte]byte{}
	r.ImageA = [util.HashSizeByte]byte{}

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

	if len(r.ImageA) == 0 {
		return errorInfo.ThrowError("ImageA is not set")
	}

	r.PreImageB = util.SliceToPreImageByte(global.RandomBytes(util.PreImageSizeByte))
	return nil
}

// Release update preimage A
func (r *RNG) Release(preImageA [util.PreImageSizeByte]byte) error {
	errorInfo := util.ErrorInfo{FunctionName: "Release", FileName: util.ErrorConstRNG}

	if len(r.PreImageB) == 0 {
		return errorInfo.ThrowError("PreImageB is not set")
	}

	err := global.ValidatePreImage(r.ImageA, preImageA)
	if err != nil {
		return errorInfo.ForwardError(err)
	}

	r.PreImageA = preImageA
	return nil
}

// RNGValue return joined random value
func (r *RNG) RNGValue() ([]byte, error) {
	errorInfo := util.ErrorInfo{FunctionName: "RNGValue", FileName: util.ErrorConstRNG}

	if len(r.PreImageB) == 0 {
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
