package global

import (
	"bytes"
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/sha3"
	"perun.network/perun-examples/dominion-cli/app/util"
)

// ToImage generate image out of preimage
func ToImage(preImage [util.PreImageSize]byte) ([util.HashSize]byte, error) {
	return _Enc(preImage), nil
}
func _Enc(preImage [util.PreImageSize]byte) [util.HashSize]byte {
	h := sha3.New256()
	h.Write(preImage[:])
	return util.SliceToHashByte(h.Sum(nil))
}

// ValidatePreImage check if preImage can be used to generate image
func ValidatePreImage(image [util.HashSize]byte, preImage [util.PreImageSize]byte) error {
	errorInfo := util.ErrorInfo{FunctionName: "ValidatePreImage", FileName: util.ErrorConstCommitment}
	var tmp = _Enc(preImage)
	if !bytes.Equal(tmp[:], image[:]) {
		return errorInfo.ThrowError(fmt.Sprintf("preimage: %v is not valid for image: %v ", preImage, image))
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
