package util

import (
	"io"
	"perun.network/perun-examples/app-channel/app/game"
)

func ReadUInt8(r io.Reader) (uint8, error) {
	buf := make([]byte, 1)
	_, err := io.ReadFull(r, buf)
	return buf[0], err
}

func ReadDeck(r io.Reader) (game.Deck, error) {
	var buf []byte
	_, err := io.ReadFull(r, buf)
	var d game.Deck
	d.Of(buf)
	return d, err
}
