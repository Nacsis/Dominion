package app

import (
	"perun.network/perun-examples/dominion-cli/app/util"
)

type RNG struct {
	ImageA               [util.HashSizeByte]byte
	PreImageB, PreImageA [util.PreImageSizeByte]byte
}

// Of create RNG out of a bytes
func (r *RNG) Of(dataBytes []byte) {

	r.ImageA = util.SliceToHashByte(dataBytes[:util.HashSizeByte])

	r.PreImageB = util.SliceToPreImageByte(dataBytes[util.HashSizeByte : util.PreImageSizeByte+util.HashSizeByte])

	r.PreImageA = util.SliceToPreImageByte(dataBytes[util.PreImageSizeByte+util.HashSizeByte : 2*util.PreImageSizeByte+util.HashSizeByte])

}

// ToByte create a byte representation of RNG
func (r *RNG) ToByte() []byte {

	var dataBytes = make([]byte, 0)

	dataBytes = append(dataBytes, r.ImageA[:]...)
	dataBytes = append(dataBytes, r.PreImageB[:]...)
	dataBytes = append(dataBytes, r.PreImageA[:]...)

	return util.AppendLength(dataBytes)
}

// Init sets up initial RNG state
func (r *RNG) Init() {
	r.PreImageB = [util.PreImageSizeByte]byte{}
	r.PreImageA = [util.PreImageSizeByte]byte{}
	r.ImageA = [util.HashSizeByte]byte{}
}
