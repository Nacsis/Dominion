package global

import (
	"bytes"
	"crypto/rand"
	"crypto/sha1"
	"errors"
)

const HashSize = 20

func HashKeyPair() ([]byte, []byte) {
	k := RandomString()
	h := sha1.New()
	h.Write(k)
	return h.Sum(nil), k
}
func Valid(hToCheck []byte, kToCheck []byte) error {
	h := sha1.New()
	h.Write(kToCheck)
	if bytes.Compare(h.Sum(nil), hToCheck) == 0 {
		return nil
	} else {
		return errors.New("incorrect Preimage")
	}

}

func RandomString() []byte {
	token := make([]byte, HashSize)
	rand.Read(token)
	return token
}
