package global

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"perun.network/perun-examples/dominion-cli/app/util"
)

// ToImage generate image out of preimage
func ToImage(preImage [util.PreImageSize]byte) ([util.HashSize]byte, error) {
	errorInfo := util.ErrorInfo{FunctionName: "ToImage", FileName: util.ErrorConstCommitment}

	if uint16(len(preImage)) < util.HashSize {
		return [util.HashSize]byte{}, errorInfo.ThrowError("given preimage has not correct size")
	}

	image := [util.HashSize]byte{}
	copy(image[:], crypto.Keccak256(preImage[:]))
	return image, nil
}

// ValidatePreImage check if preImage can be used to generate image
func ValidatePreImage(image [util.HashSize]byte, preImage [util.PreImageSize]byte) error {
	errorInfo := util.ErrorInfo{FunctionName: "ValidatePreImage", FileName: util.ErrorConstCommitment}
	if !bytes.Equal(crypto.Keccak256(preImage[:]), image[:]) {
		return errorInfo.ThrowError(fmt.Sprintf("preimage: \n%v\nis not valid for image: \n%v\nas its \n%v\n", preImage, image, crypto.Keccak256(preImage[:])))
	}
	return nil
}

// RandomBytes of given size
func RandomBytes(size uint16) []byte { // TODO Add ERROR
	errorInfo := util.ErrorInfo{FunctionName: "RandomBytes", FileName: util.ErrorConstCommitment}

	buf := make([]byte, size)
	_, err := rand.Read(buf)
	if err != nil {
		errorInfo.ForwardError(err) // TODO ERROR handle
	}
	return buf
}

// Xor output xor of a and b
func Xor(a, b [util.PreImageSize]byte) ([]byte, error) {
	var c = make([]byte, util.PreImageSize)
	for i := uint16(0); i < util.PreImageSize; i++ {
		c[i] = a[i] ^ b[i]
	}
	return c, nil
}
