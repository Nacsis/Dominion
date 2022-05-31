package app

import (
	"io"
	"perun.network/go-perun/channel"
	"perun.network/perun-examples/app-channel/app/util"
)

type DominionAppData struct {
	turn      Turn
	CardDecks [util.NumPlayers]Deck
	rng       RNG
}

// Encode encodes the app data onto an io.Writer.
func (d *DominionAppData) Encode(w io.Writer) error {
	errorInfo := util.ErrorInfo{FunctionName: "Encode", FileName: util.ErrorConstDATA}

	// Encode turn
	err := util.Write(w, &d.turn)
	if err != nil {
		return errorInfo.ForwardError(err)
	}

	// Encode decks
	for i := 0; i < len(d.CardDecks); i++ {
		err = util.Write(w, &d.CardDecks[i])
		if err != nil {
			return errorInfo.ForwardError(err)
		}
	}

	// Encode rng
	err = util.Write(w, &d.rng)
	if err != nil {
		return errorInfo.ForwardError(err)
	}

	return nil
}

// Clone returns a deep copy of the app data.
func (d *DominionAppData) Clone() channel.Data {
	_d := *d
	return &_d
}

// Init sets up initial game state
func (d *DominionAppData) Init(firstActor channel.Index) error {
	errorInfo := util.ErrorInfo{FunctionName: "ValidatePreImage", FileName: util.ErrorConstDATA}

	// Init Turn
	d.turn.Init(uint8(firstActor))

	// Init decks
	for deckNum := 0; deckNum < util.NumPlayers; deckNum++ {
		err := d.CardDecks[deckNum].Init()
		if err != nil {
			return errorInfo.ForwardError(err)
		}
	}

	// Init rng
	//TODO init rng

	return nil
}
