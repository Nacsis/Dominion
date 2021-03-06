package app

import (
	"fmt"
	"io"
	"reflect"

	"github.com/ethereum/go-ethereum/log"
	"perun.network/go-perun/channel"
	"perun.network/go-perun/wallet"
	"perun.network/perun-examples/dominion-cli/app/util"
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

	// Read Turn
	err := util.ReadObject(r, &appData.Turn)
	if err != nil {
		return nil, errorInfo.ForwardError(err)
	}

	// Read Stock
	err = util.ReadObject(r, &appData.Stock)
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
	err = util.ReadObject(r, &appData.Rng)
	if err != nil {
		return nil, errorInfo.ForwardError(err)
	}

	return &appData, nil
}

// ValidTransition validate if state is correct
func (a *DominionApp) ValidTransition(params *channel.Params, from, to *channel.State, idx channel.Index) error {
	errorInfo := util.ErrorInfo{FunctionName: "ValidTransition", FileName: util.ErrorConstAPP}
	log.Info("ValidTransition called")

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
	toData, ok := to.Data.(*DominionAppData)
	if !ok {
		return errorInfo.ThrowError(fmt.Sprintf("toData is in an invalid data format %T", toData))
	}

	// Correct count of players
	if len(params.Parts) != util.NumPlayers {
		return errorInfo.ThrowError(fmt.Sprintf("Player count is not correct. should be %v but is %v", util.NumPlayers, len(params.Parts)))
	}

	err = a._ValidState(*fromData, *toData, idx)
	if err != nil {
		return errorInfo.ForwardError(err)
	}

	if toData.Turn.PerformedAction == util.GameEnd {
		expectedBalances := toData.ComputeFinalBalances(from.Balances)
		if !to.IsFinal {
			return errorInfo.ThrowError("Expected final flag to be true")
		}
		if !to.Allocation.Balances.Equal(expectedBalances) {
			return errorInfo.ThrowError("Balances was not calculated correctly")
		}

	}
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
	if appdata.Turn.NextActor >= util.NumPlayers {
		return errorInfo.ThrowError(fmt.Sprintf("Next Actor is not valid. Should be smaller than %v but is %v", util.NumPlayers, appdata.Turn.NextActor))
	}

	// Initial state is correct
	var newData DominionAppData
	err := newData.Init(channel.Index(appdata.Turn.NextActor))

	if err != nil {
		return errorInfo.ForwardError(err)
	}

	if reflect.DeepEqual(newData, appdata) {
		return errorInfo.ThrowError("State transition could not be replicated for action")
	}
	return nil
}
