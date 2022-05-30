package app

import (
	"perun.network/go-perun/channel"
)

func ComputeFinalBalances(b channel.Balances) channel.Balances {
	return b.Clone()
}
