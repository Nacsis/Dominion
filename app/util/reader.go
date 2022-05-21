package util

import (
	"io"
)

func ReadUInt8(r io.Reader) (uint8, error) {
	buf := make([]byte, 1)
	_, err := io.ReadFull(r, buf)
	return buf[0], err
}

/*
func ReadObject(r io.Reader, o Readable) (Readable, error) {
	var buf []byte
	_, err := io.ReadFull(r, buf)
	o.Of(buf)
	return o, err
}*/
func ReadBytes(r io.Reader, length uint8) ([]byte, error) {
	buf := make([]byte, length)
	_, err := io.ReadFull(r, buf)
	return buf, err
}
