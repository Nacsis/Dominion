package client

import (
	"context"
	"fmt"
	"perun.network/go-perun/channel"
	"perun.network/go-perun/client"
	"perun.network/perun-examples/app-channel/app"
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

// SwitchActor switch the current actor and ends the game
func (g *DominionChannel) SwitchActor() {
	err := g.ch.UpdateBy(context.TODO(), func(state *channel.State) error {
		dominionApp, ok := state.App.(*app.DominionApp)
		if !ok {
			return fmt.Errorf("invalid app type: %T", dominionApp)
		}

		return dominionApp.SwitchActor(state, g.ch.Idx())
	})
	if err != nil {
		panic(err) // We panic on error to keep the code simple.
	}
}
