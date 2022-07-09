package client

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"perun.network/go-perun/channel"
	"perun.network/go-perun/client"
	"perun.network/go-perun/log"
	"perun.network/perun-examples/dominion-cli/app"
	"perun.network/perun-examples/dominion-cli/app/util"
)

// DominionChannel is a wrapper for a Perun channel for the Dominion app use case.
type DominionChannel struct {
	*client.Channel
	log     log.Logger
	timeout time.Duration
}

func (g *DominionChannel) GetAppStateDataEncoded() []byte {
	w := strings.Builder{}
	g.State().Data.(*app.DominionAppData).Encode(&w)
	return []byte(w.String())
}

func (g *DominionChannel) GetAppStateData() *app.DominionAppData {
	return g.State().Data.(*app.DominionAppData)
}

// NewDominionChannel creates a new Dominion app channel.
func NewDominionChannel(ch *client.Channel, t time.Duration) *DominionChannel {
	return &DominionChannel{
		Channel: ch,
		log:     log.WithField("channel", ch.ID()),
		timeout: t,
	}
}

func (ch *DominionChannel) GetTimeout() time.Duration {
	return ch.timeout
}

// Settle settles the app channel and withdraws the funds.
func (g *DominionChannel) SettleAndClose() {
	// Channel should be finalized through last ("winning") move.
	// No need to set `isFinal` here.
	ctx, cancel := context.WithTimeout(context.Background(), g.timeout)
	defer cancel()
	err := g.Settle(ctx, false)
	if err != nil {
		panic(err)
	}

	// Cleanup.
	err = g.Close()
	if err != nil {
		return
	}
}

//------------------------ Funds and Balances ------------------------

func stateBals(state *channel.State) []channel.Bal {
	return state.Balances[0]
}

func (ch *DominionChannel) GetBalances() (our, other *big.Int) {
	bals := stateBals(ch.State())
	if len(bals) != 2 {
		return new(big.Int), new(big.Int)
	}
	return bals[ch.Idx()], bals[1-ch.Idx()]
}

//------------------------ RNG ------------------------

// func stateGameState(state *channel.State) (*app.DominionAppData, error) {
// 	dominionAppData, ok := state.Data.(*app.DominionAppData)
// 	if !ok {
// 		return nil, errors.New("state.Data could not be castet to *app.DominionAppData")
// 	}
// 	return dominionAppData, nil
// }

// func (ch *DominionChannel) GetGameState() (*app.DominionAppData, error) {
// 	return stateGameState(ch.State())
// }

//------------------------ RNG ------------------------

// RngCommit player who wants to DrawOneCard commit to an preimage by setting corresponding image
func (g *DominionChannel) RngCommit(preImage [util.PreImageSizeByte]byte) {
	errorInfo := util.ErrorInfo{FunctionName: "RngCommit", FileName: util.ErrorConstChannel}

	ctx, cancel := context.WithTimeout(context.Background(), g.timeout)
	defer cancel()
	err := g.UpdateBy(ctx, func(state *channel.State) error {
		dominionApp, ok := state.App.(*app.DominionApp)
		if !ok {
			return errorInfo.ThrowError(fmt.Sprintf("App is in an invalid data format %T", dominionApp))
		}

		return dominionApp.RngCommit(state, g.Idx(), preImage)
	})
	if err != nil {
		panic(err) // We panic on error to keep the code simple.
	}
}

// RngTouch the player how doesn't DrawOneCard choose an image
func (g *DominionChannel) RngTouch() {
	errorInfo := util.ErrorInfo{FunctionName: "RngTouch", FileName: util.ErrorConstChannel}

	ctx, cancel := context.WithTimeout(context.Background(), g.timeout)
	defer cancel()

	g.log.Debug("proposing RngTouch...")
	err := g.UpdateBy(ctx, func(state *channel.State) error {
		dominionApp, ok := state.App.(*app.DominionApp)
		if !ok {
			return errorInfo.ThrowError(fmt.Sprintf("App is in an invalid data format %T", dominionApp))
		}

		return dominionApp.RngTouch(state, g.Idx())
	})
	if err != nil {
		panic(err) // We panic on error to keep the code simple.
	}
}

// RngRelease player who wants to DrawOneCard publish preimage for published image
func (g *DominionChannel) RngRelease(preImage [util.PreImageSizeByte]byte) {
	errorInfo := util.ErrorInfo{FunctionName: "RngRelease", FileName: util.ErrorConstChannel}

	ctx, cancel := context.WithTimeout(context.Background(), g.timeout)
	defer cancel()

	g.log.Debug("proposing RngRelease...")
	err := g.UpdateBy(ctx, func(state *channel.State) error {
		dominionApp, ok := state.App.(*app.DominionApp)
		if !ok {
			return errorInfo.ThrowError(fmt.Sprintf("App is in an invalid data format %T", dominionApp))
		}

		return dominionApp.RngRelease(state, g.Idx(), preImage)
	})
	if err != nil {
		panic(err) // We panic on error to keep the code simple.
	}
}

//------------------------ Drawing ------------------------

// DrawOneCard draws one card to the hand pile. A full rng need to be performed before.
func (g *DominionChannel) DrawOneCard() {
	errorInfo := util.ErrorInfo{FunctionName: "DrawCard", FileName: util.ErrorConstChannel}

	ctx, cancel := context.WithTimeout(context.Background(), g.timeout)
	defer cancel()

	g.log.Debug("proposing DrawCard...")
	err := g.UpdateBy(ctx, func(state *channel.State) error {
		dominionApp, ok := state.App.(*app.DominionApp)
		if !ok {
			return errorInfo.ThrowError(fmt.Sprintf("App is in an invalid data format %T", dominionApp))
		}

		return dominionApp.DrawCard(state, g.Idx())
	})
	if err != nil {
		panic(err) // We panic on error to keep the code simple.
	}
}

// PlayCard draws one card to the hand pile.
func (g *DominionChannel) PlayCard(index uint8, followUpIndices []uint8, followUpCardType util.CardType) {
	errorInfo := util.ErrorInfo{FunctionName: "PlayCard", FileName: util.ErrorConstChannel}

	ctx, cancel := context.WithTimeout(context.Background(), g.timeout)
	defer cancel()
	err := g.UpdateBy(ctx, func(state *channel.State) error {
		dominionApp, ok := state.App.(*app.DominionApp)
		if !ok {
			return errorInfo.ThrowError(fmt.Sprintf("App is in an invalid data format %T", dominionApp))
		}

		return dominionApp.PlayCard(state, g.Idx(), index, followUpIndices, followUpCardType)
	})
	if err != nil {
		panic(err) // We panic on error to keep the code simple.
	}
}

// BuyCard Buyables one card for given CardType.
func (g *DominionChannel) BuyCard(cardType util.CardType, force bool) error {
	errorInfo := util.ErrorInfo{FunctionName: "BuyCard", FileName: util.ErrorConstChannel}

	ctx, cancel := context.WithTimeout(context.Background(), g.timeout)
	defer cancel()
	updateCallback := func(state *channel.State) error {
		dominionApp, ok := state.App.(*app.DominionApp)
		if !ok {
			return errorInfo.ThrowError(fmt.Sprintf("App is in an invalid data format %T", dominionApp))
		}

		return dominionApp.BuyCard(state, g.Idx(), cardType)
	}

	if !force {
		return g.UpdateBy(ctx, updateCallback)
	} else {
		if err := g.ForceUpdate(ctx, func(state *channel.State) {
			updateCallback(state)
			// if err := updateCallback(state); err != nil {
			// 	panic(err) // We panic on error because this should have already failed in off-chain attempt.
			// }
		}); err != nil {
			g.log.Errorln("Channel ForceUpdate failed... :(")
			return err
		}
	}
	return nil
}

//------------------------ General turn mechanics ------------------------

// EndTurn switch the current actor and ends the game
func (g *DominionChannel) EndTurn() {
	errorInfo := util.ErrorInfo{FunctionName: "EndTurn", FileName: util.ErrorConstChannel}

	ctx, cancel := context.WithTimeout(context.Background(), g.timeout)
	defer cancel()
	err := g.UpdateBy(ctx, func(state *channel.State) error {
		dominionApp, ok := state.App.(*app.DominionApp)
		if !ok {
			return errorInfo.ThrowError(fmt.Sprintf("App is in an invalid data format %T", dominionApp))
		}

		return dominionApp.EndTurn(state, g.Idx())
	})
	if err != nil {
		panic(err) // We panic on error to keep the code simple.
	}
}

//------------------------ GameEnd ------------------------

// EndGame
func (g *DominionChannel) EndGame() {
	errorInfo := util.ErrorInfo{FunctionName: "EndTurn", FileName: util.ErrorConstChannel}
	ctx, cancel := context.WithTimeout(context.Background(), g.timeout)
	defer cancel()
	err := g.UpdateBy(ctx, func(state *channel.State) error {
		dominionApp, ok := state.App.(*app.DominionApp)
		if !ok {
			return errorInfo.ThrowError(fmt.Sprintf("App is in an invalid data format %T", dominionApp))
		}

		return dominionApp.EndGame(state, g.Idx())
	})
	if err != nil {
		panic(err) // We panic on error to keep the code simple.
	}
}
