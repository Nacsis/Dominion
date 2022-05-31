package app

import (
	"fmt"
	"io"
	"perun.network/go-perun/channel"
	"perun.network/go-perun/wallet"
	"perun.network/perun-examples/app-channel/app/util"
)

type DominionApp struct {
	Addr wallet.Address
}

func NewDominionApp(addr wallet.Address) *DominionApp {
	return &DominionApp{
		Addr: addr,
	}
}

// Def returns the app address.
func (a *DominionApp) Def() wallet.Address {
	return a.Addr
}

// DecodeData decodes the DominionAppData.
func (a *DominionApp) DecodeData(r io.Reader) (channel.Data, error) {
	errorInfo := util.ErrorInfo{FunctionName: "DecodeData", FileName: util.ErrorConstAPP}

	appData := DominionAppData{}

	// Read turn
	err := util.ReadObject(r, &appData.turn)
	if err != nil {
		return nil, errorInfo.ForwardError(err)
	}

	// Read stock
	err = util.ReadObject(r, &appData.stock)
	if err != nil {
		return nil, errorInfo.ForwardError(err)
	}

	// Read decks
	for deckIndex := 0; deckIndex < util.NumPlayers; deckIndex++ {
		err := util.ReadObject(r, &appData.CardDecks[deckIndex])
		if err != nil {
			return nil, errorInfo.ForwardError(err)
		}
	}

	// Read RNG
	err = util.ReadObject(r, &appData.rng)
	if err != nil {
		return nil, errorInfo.ForwardError(err)
	}

	return &appData, nil
}

// ValidTransition validate if state is correct
func (a *DominionApp) ValidTransition(params *channel.Params, from, to *channel.State, idx channel.Index) error {
	errorInfo := util.ErrorInfo{FunctionName: "ValidTransition", FileName: util.ErrorConstAPP}

	// Values are correct
	err := channel.AssetsAssertEqual(from.Assets, to.Assets)
	if err != nil {
		return errorInfo.ForwardError(err)
	}

	// Format of fromData is valid
	fromData, ok := from.Data.(*DominionAppData)
	if !ok {
		return errorInfo.ThrowError(fmt.Sprintf("fromData is in an invalid data format %T", fromData))
	}

	// Format of toData is valid
	toData, ok := from.Data.(*DominionAppData)
	if !ok {
		return errorInfo.ThrowError(fmt.Sprintf("toData is in an invalid data format %T", toData))
	}

	// Correct count of players
	if len(params.Parts) != util.NumPlayers {
		return errorInfo.ThrowError(fmt.Sprintf("Player count is not correct. should be %v but is %v", util.NumPlayers, len(params.Parts)))
	}

	// TODO State Validation
	// TODO Check if final state is reached
	/*
		fromDataClone := &(*fromData)
		switch toData.turn.performedAction {
		case util.RngCommit:
			dumPreImage := global.RandomBytes(util.HashSize)
			err = fromDataClone.RngCommit(idx, dumPreImage)
			if err != nil {
				return errorInfo.ForwardError( err)
			}
			fromDataClone.RngCommit(idx, dumPreImage)

			// replace dummy
			fromDataClone.rng.ImageA = toData.rng.ImageA

			if fromDataClone == toData {
				return errorInfo.ThrowError( "State transition could not be replicated for RngCommit")
			}
		}*/

	return nil
}

// ValidInit validate if initial state was calculated correctly
func (a *DominionApp) ValidInit(p *channel.Params, s *channel.State) error {
	errorInfo := util.ErrorInfo{FunctionName: "ValidInit", FileName: util.ErrorConstAPP}

	// Correct count of players
	if len(p.Parts) != util.NumPlayers {
		return errorInfo.ThrowError(fmt.Sprintf("Player count is not correct. should be %v but is %v", util.NumPlayers, len(p.Parts)))
	}

	// Format of appdata is valid
	appdata, ok := s.Data.(*DominionAppData)
	if !ok {
		return errorInfo.ThrowError(fmt.Sprintf("AppData is in an invalid data format %T", appdata))
	}

	// Initial game state can't be final
	if s.IsFinal {
		return errorInfo.ThrowError("game state must not be final on init")
	}

	// Next player value is in range of player count
	if appdata.turn.nextActor >= util.NumPlayers {
		return errorInfo.ThrowError(fmt.Sprintf("Next actor is not valid. Should be smaller than %v but is %v", util.NumPlayers, appdata.turn.nextActor))
	}

	return nil
}
