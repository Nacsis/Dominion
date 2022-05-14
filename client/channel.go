package client

import (
	"context"
	"perun.network/go-perun/client"
)

// TicTacToeChannel is a wrapper for a Perun channel for the Tic-tac-toe app use case.
type TicTacToeChannel struct {
	ch *client.Channel
}

// newTicTacToeChannel creates a new tic-tac-toe app channel.
func newTicTacToeChannel(ch *client.Channel) *TicTacToeChannel {
	return &TicTacToeChannel{ch: ch}
}

// Settle settles the app channel and withdraws the funds.
func (g *TicTacToeChannel) Settle() {
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
