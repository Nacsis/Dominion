package app

import (
	"perun.network/go-perun/channel"
	"perun.network/perun-examples/app-channel/app/util"
)

func ComputeFinalBalances(b channel.Balances) channel.Balances {
	return b.Clone()
}
func CalcNextActor(actor uint8) uint8 {
	return (actor + 1) % util.NumPlayers
}
