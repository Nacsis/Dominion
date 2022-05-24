package global

import (
	"bytes"
	"crypto/rand"
	"crypto/sha1"
	"errors"
	"perun.network/perun-examples/app-channel/app/util"
)

func HashKeyPair(size uint8) ([]byte, []byte) {
	k := RandomBytes(size)
	h := sha1.New()
	h.Write(k)
	return h.Sum(nil), k
}

func Hash(key []byte) []byte {
	if uint8(len(key)) < util.HashSize {
		//TODO Panic oder was auch immer
	}
	h := sha1.New()
	h.Write(key)
	return h.Sum(nil)
}
func Valid(hToCheck []byte, kToCheck []byte) error {
	if IsValid(hToCheck, kToCheck) {
		return errors.New("incorrect Preimage")
	}
	return nil
}
func IsValid(hToCheck []byte, kToCheck []byte) bool {
	h := sha1.New()
	h.Write(kToCheck)
	return bytes.Compare(h.Sum(nil), hToCheck) != 0
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
