package app

import (
	"perun.network/go-perun/channel"
)

type ActionTypes uint8

const (
	GameInit ActionTypes = iota
	RngCommit
	RngTouch
	RngRelease
	DrawCard
	EndTurn
	GameEnd
)

func ComputeFinalBalances(b channel.Balances) channel.Balances {
	return b.Clone()
}
