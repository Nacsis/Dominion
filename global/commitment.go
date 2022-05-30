package global

import (
	"bytes"
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"perun.network/perun-examples/app-channel/app/util"
)

// ToImage generate image out of preimage
func ToImage(preImage []byte) []byte {
	if uint8(len(preImage)) < util.HashSize {
		//TODO Panic oder was auch immer
	}
	h := sha1.New()
	h.Write(preImage)
	return h.Sum(nil)
}

// ValidatePreImage check if preImage can be used to generate image
func ValidatePreImage(image []byte, preImage []byte) error {
	errorDescription := util.ErrorInfo{FunctionName: "DrawOneCard", FileName: util.ErrorConstCommitment}

	h := sha1.New()
	h.Write(preImage)
	if !bytes.Equal(h.Sum(nil), image) {
		return util.ThrowError(errorDescription, fmt.Sprintf("preimage: %v is not valid for image: %v ", preImage, image))
	}
	return nil
}

// RandomBytes of given size
func RandomBytes(size uint8) []byte {
	buf := make([]byte, size)
	rand.Read(buf) // TODO ERROR handle
	return buf
}

// Xor output xor of a and b
func Xor(a, b []byte) ([]byte, error) {
	errorDescription := util.ErrorInfo{FunctionName: "Xor", FileName: util.ErrorConstCommitment}

	if len(a) != int(util.HashSize) || len(b) != int(util.HashSize) {
		return nil, util.ThrowError(errorDescription, fmt.Sprintf("a or b has not the correct size of %v", util.HashSize))
	}
	var c = make([]byte, util.HashSize)
	for i := uint8(0); i < util.HashSize; i++ {
		c[i] = a[i] ^ b[i]
	}
	return c, nil
}
