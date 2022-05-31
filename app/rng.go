package app

import (
	"perun.network/perun-examples/app-channel/app/util"
)

type RNG struct {
	ImageA, PreImageB, PreImageA []byte
}

// Of create RNG out of a bytes
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

// ToByte create a byte representation of RNG
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

// Init sets up initial RNG state
func (r *RNG) Init() {
	r.PreImageB = make([]byte, util.HashSize)
	r.PreImageA = make([]byte, util.HashSize)
	r.ImageA = make([]byte, util.HashSize)
}
