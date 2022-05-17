package app

import (
	"github.com/pkg/errors"
	"io"
	"perun.network/go-perun/channel"
	perunio "perun.network/go-perun/pkg/io"
)

type Data struct {
	NextActor   uint8
	NumAllCards uint8
	AllCards    [256]Card
	perunio.Encoder
}

// Encode encodes app data onto an io.Writer.
func (d *Data) Encode(w io.Writer) error {
	err := writeUInt8(w, d.NextActor)
	if err != nil {
		return errors.WithMessage(err, "writing actor")
	}
	err = writeUInt8(w, d.NumAllCards)
	if err != nil {
		return errors.WithMessage(err, "writing NumAllCards")
	}
	err = writeCards(w, d.AllCards)
	return errors.WithMessage(err, "writing grid")
}

// Clone returns a deep copy of the app data.
func (d *Data) Clone() channel.Data {
	_d := *d
	return &_d
}

func CalcNextActor(actor uint8) uint8 {
	return (actor + 1) % numParts
}

func (d *Data) AddCard(c Card) {
	d.AllCards[d.NumAllCards] = c
	d.NumAllCards += 1
}
