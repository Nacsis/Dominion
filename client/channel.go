package client

import (
	"context"
	"perun.network/go-perun/client"
)

// DominionChannel is a wrapper for a Perun channel for the Tic-tac-toe app use case.
type DominionChannel struct {
	ch *client.Channel
}

// newDominionChannel creates a new tic-tac-toe app channel.
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
