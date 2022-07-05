package util

import (
	"io"
)

func ReadUInt8(r io.Reader) (uint8, error) {
	buf := make([]byte, 1)
	_, err := io.ReadFull(r, buf)
	return buf[0], err
}

func ReadUInt16(r io.Reader) (uint16, error) {
	buf := make([]byte, 2)
	_, err := io.ReadFull(r, buf)
	return ByteArrayToUint16(*(*[2]byte)(buf)), err
}

func ReadUInt32(r io.Reader) (uint32, error) {
	buf := make([]byte, 4)
	_, err := io.ReadFull(r, buf)
	return ByteArrayToUint32(*(*[4]byte)(buf)), err
}

// ReadObject Read a Readable object from stream.
// It is necessary that every object stream starts with its size
func ReadObject(r io.Reader, o Readable) error {
	length, err := ReadUInt16(r)
	buf := make([]byte, length)
	io.ReadFull(r, buf) // TODO ERROR handle
	o.Of(buf)
	return err
}

func PopLength(s []byte) (length uint, data []byte) {
	return uint(ByteArrayToUint16(*(*[2]byte)(s[:2]))), s[2:]
}
