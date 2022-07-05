package app

import (
	"encoding/json"
	"io"

	"perun.network/go-perun/channel"
	"perun.network/perun-examples/dominion-cli/app/util"
)

type DominionAppData struct {
	Turn      Turn
	Stock     Stock
	CardDecks [util.NumPlayers]Deck
	Rng       RNG
}

// Encode encodes the app data onto an io.Writer.
func (d *DominionAppData) Encode(w io.Writer) error {
	errorInfo := util.ErrorInfo{FunctionName: "Encode", FileName: util.ErrorConstDATA}

	// Encode Turn
	err := util.Write(w, &d.Turn)
	if err != nil {
		return errorInfo.ForwardError(err)
	}

	// Encode Stock
	err = util.Write(w, &d.Stock)
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

	// Encode Rng
	err = util.Write(w, &d.Rng)
	if err != nil {
		return errorInfo.ForwardError(err)
	}

	return nil
}

// Clone returns a deep copy of the app data.
func (d *DominionAppData) Clone() channel.Data {
	origJSON, _ := json.Marshal(d)
	fromDataClone := DominionAppData{}
	json.Unmarshal(origJSON, &fromDataClone)
	return &fromDataClone
}

func (d *DominionAppData) Clone2DominionAppData() *DominionAppData {
	origJSON, _ := json.Marshal(d)
	fromDataClone := DominionAppData{}
	json.Unmarshal(origJSON, &fromDataClone)
	return &fromDataClone
}

// Init sets up initial game state
func (d *DominionAppData) Init(firstActor channel.Index) error {
	errorInfo := util.ErrorInfo{FunctionName: "ValidatePreImage", FileName: util.ErrorConstDATA}

	// Init Turn
	d.Turn.Init(uint8(firstActor))

	// Init Stock
	d.Stock.Init()

	// Init decks
	for deckNum := 0; deckNum < util.NumPlayers; deckNum++ {
		pile, err := d.Stock.TakeOffOneInitialDeck()
		if err != nil {
			return errorInfo.ForwardError(err)
		}

		err = d.CardDecks[deckNum].Init(pile)
		if err != nil {
			return errorInfo.ForwardError(err)
		}
	}
	d.Rng.Init()

	return nil
}
