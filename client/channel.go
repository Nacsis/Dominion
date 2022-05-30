package client

import (
	"context"
	"fmt"
	"perun.network/go-perun/channel"
	"perun.network/go-perun/client"
	"perun.network/perun-examples/app-channel/app"
	"perun.network/perun-examples/app-channel/app/util"
)

// DominionChannel is a wrapper for a Perun channel for the Dominion app use case.
type DominionChannel struct {
	ch *client.Channel
}

// newDominionChannel creates a new Dominion app channel.
func newDominionChannel(ch *client.Channel) *DominionChannel {
	return &DominionChannel{ch: ch}
}

// Settle settles the app channel and withdraws the funds.
func (g *DominionChannel) Settle() {
	// Channel should be finalized through last ("winning") move.
	// No need to set `isFinal` here.
	err := g.ch.Settle(context.TODO(), false)
	if err != nil {
		panic(err)
	}

	// Cleanup.
	err = g.ch.Close()
	if err != nil {
		return
	}
}

//------------------------ RNG ------------------------

// RngCommit player who wants to DrawOneCard commit to an preimage by setting corresponding image
func (g *DominionChannel) RngCommit(preImage []byte) {
	err := g.ch.UpdateBy(context.TODO(), func(state *channel.State) error {
		dominionApp, ok := state.App.(*app.DominionApp)
		if !ok {
			return util.ThrowError(util.ErrorConstChannel, "RngCommit", fmt.Sprintf("App is in an invalid data format %T", dominionApp))
		}

		return dominionApp.RngCommit(state, g.ch.Idx(), preImage)
	})
	if err != nil {
		panic(err) // We panic on error to keep the code simple.
	}
}

// RngTouch the player how doesn't DrawOneCard choose an image
func (g *DominionChannel) RngTouch() {
	err := g.ch.UpdateBy(context.TODO(), func(state *channel.State) error {
		dominionApp, ok := state.App.(*app.DominionApp)
		if !ok {
			return util.ThrowError(util.ErrorConstChannel, "RngTouch", fmt.Sprintf("App is in an invalid data format %T", dominionApp))
		}

		return dominionApp.RngTouch(state, g.ch.Idx())
	})
	if err != nil {
		panic(err) // We panic on error to keep the code simple.
	}
}

// RngRelease player who wants to DrawOneCard publish preimage for published image
func (g *DominionChannel) RngRelease(preImage []byte) {
	err := g.ch.UpdateBy(context.TODO(), func(state *channel.State) error {
		dominionApp, ok := state.App.(*app.DominionApp)
		if !ok {
			return util.ThrowError(util.ErrorConstChannel, "RngRelease", fmt.Sprintf("App is in an invalid data format %T", dominionApp))
		}

		return dominionApp.RngRelease(state, g.ch.Idx(), preImage)
	})
	if err != nil {
		panic(err) // We panic on error to keep the code simple.
	}
}

//------------------------ Drawing ------------------------

// DrawOneCard draws one card to the hand pile. A full rng need to be performed before.
func (g *DominionChannel) DrawOneCard() {
	err := g.ch.UpdateBy(context.TODO(), func(state *channel.State) error {
		dominionApp, ok := state.App.(*app.DominionApp)
		if !ok {
			return util.ThrowError(util.ErrorConstChannel, "DrawOneCard", fmt.Sprintf("App is in an invalid data format %T", dominionApp))
		}

		return dominionApp.DrawOneCard(state, g.ch.Idx())
	})
	if err != nil {
		panic(err) // We panic on error to keep the code simple.
	}
}

//------------------------ General turn mechanics ------------------------

// EndTurn switch the current actor and ends the game
func (g *DominionChannel) EndTurn() {
	err := g.ch.UpdateBy(context.TODO(), func(state *channel.State) error {
		dominionApp, ok := state.App.(*app.DominionApp)
		if !ok {
			return util.ThrowError(util.ErrorConstChannel, "EndTurn", fmt.Sprintf("App is in an invalid data format %T", dominionApp))
		}

		return dominionApp.EndTurn(state, g.ch.Idx())
	})
	if err != nil {
		panic(err) // We panic on error to keep the code simple.
	}
}

//------------------------ GameEnd ------------------------

// EndGame
func (g *DominionChannel) EndGame() {
	err := g.ch.UpdateBy(context.TODO(), func(state *channel.State) error {
		dominionApp, ok := state.App.(*app.DominionApp)
		if !ok {
			return util.ThrowError(util.ErrorConstChannel, "EndTurn", fmt.Sprintf("App is in an invalid data format %T", dominionApp))
		}

		return dominionApp.EndGame(state, g.ch.Idx())
	})
	if err != nil {
		panic(err) // We panic on error to keep the code simple.
	}
}
