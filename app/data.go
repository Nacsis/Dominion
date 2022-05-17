package app

import (
	"github.com/pkg/errors"
	"io"
	"perun.network/go-perun/channel"
)

const (
	NumPlayers = 2
)

type DominionAppData struct {
	NextActor uint8
	/*Commitments [][global.HashSize]byte
	PreImages   [][global.HashSize]byte*/
	CardDecks [NumPlayers]Deck
}

// Encode encodes app data onto an io.Writer.
func (d *DominionAppData) Encode(w io.Writer) error {

	err := WriteUInt8(w, d.NextActor)
	if err != nil {
		return errors.WithMessage(err, "writing actor")
	}

	for i := 0; i < NumPlayers; i++ {
		err := Write(w, &d.CardDecks[i])
		if err != nil {
			return errors.WithMessage(err, "writing card")
		}
	}
	return nil
}

// Clone returns a deep copy of the app data.
func (d *DominionAppData) Clone() channel.Data {
	_d := *d
	return &_d
}

func (d *DominionAppData) switchActor(actorIdx channel.Index) {

	if d.NextActor != Uint8safe(uint16(actorIdx)) {
		panic("invalid actor")
	}
	d.NextActor += +1
}
