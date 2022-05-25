package util

import (
	"io"
)

func ReadUInt8(r io.Reader) (uint8, error) {
	buf := make([]byte, 1)
	_, err := io.ReadFull(r, buf)
	return buf[0], err
}

// ReadObject Read a Readable object from stream.
// It is necessary that every object stream starts with its size
func ReadObject(r io.Reader, o Readable) error {
	length, err := ReadUInt8(r)
	buf := make([]byte, length)
	io.ReadFull(r, buf) // TODO ERROR handle
	o.Of(buf)
	return err
}
