package global

import (
	"bytes"
	"crypto/rand"
	"crypto/sha1"
	"errors"
	"perun.network/perun-examples/app-channel/app/util"
)

const HashSize = 20

func HashKeyPair(size uint8) ([]byte, []byte) {
	k := RandomBytes(size)
	h := sha1.New()
	h.Write(k)
	return h.Sum(nil), k
}
func Valid(hToCheck []byte, kToCheck []byte) error {
	h := sha1.New()
	h.Write(kToCheck)
	if bytes.Compare(h.Sum(nil), hToCheck) != 0 {
		return errors.New("incorrect Preimage")
	}
	return nil
}

func RandomBytes(size uint8) []byte {
	buf := make([]byte, size)
	rand.Read(buf)
	return buf
}

func Xor(a, b []byte) []byte {
	for i := uint8(0); i < util.HashSize; i++ {
		a[i] = a[i] ^ b[i]
	}
	return a
}
