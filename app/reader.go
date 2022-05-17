package app

import (
	"io"
)

func ReadUInt8(r io.Reader) (uint8, error) {
	buf := make([]byte, 1)
	_, err := io.ReadFull(r, buf)
	return buf[0], err
}

func ReadDeck(r io.Reader) (Deck, error) {
	var buf []byte
	_, err := io.ReadFull(r, buf)
	var d Deck
	d.Of(buf)
	return d, err
}
