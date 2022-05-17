package app

import (
	"perun.network/go-perun/channel"
)

func ComputeFinalBalances(b channel.Balances) channel.Balances {
	return b.Clone()
}
func CalcNextActor(actor uint8) uint8 {
	return (actor + 1) % NumPlayers
}
