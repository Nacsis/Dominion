package app

import (
	"perun.network/perun-examples/dominion-cli/app/util"
)

type RNG struct {
	ImageA               [util.HashSize]byte
	PreImageB, PreImageA [util.PreImageSize]byte
}

// Of create RNG out of a bytes
func (r *RNG) Of(dataBytes []byte) {
	var size = uint16(len(dataBytes))

	if size == util.HashSize {
		r.ImageA = util.SliceToHashByte(dataBytes[:util.HashSize])
	}
	if size == util.PreImageSize+util.HashSize {
		r.PreImageB = util.SliceToPreImageByte(dataBytes[util.HashSize : util.PreImageSize+util.HashSize])
	}
	if size == 2*util.PreImageSize+util.HashSize {
		r.PreImageA = util.SliceToPreImageByte(dataBytes[util.PreImageSize+util.HashSize : 2*util.PreImageSize+util.HashSize])
	}
}

// ToByte create a byte representation of RNG
func (r *RNG) ToByte() []byte {

	// if ImageA is not set end with Rng length 0
	var dataBytes = make([]byte, 0)

	if len(r.ImageA) != 0 {
		dataBytes = append(dataBytes, r.ImageA[:]...)
		if len(r.PreImageB) != 0 {
			dataBytes = append(dataBytes, r.PreImageB[:]...)
			if len(r.PreImageA) != 0 {
				dataBytes = append(dataBytes, r.PreImageA[:]...)
			}
		}
	}

	return append([]byte{byte(len(dataBytes))}, dataBytes...)
}

// Init sets up initial RNG state
func (r *RNG) Init() {
	r.PreImageB = [util.PreImageSize]byte{}
	r.PreImageA = [util.PreImageSize]byte{}
	r.ImageA = [util.HashSize]byte{}
}
