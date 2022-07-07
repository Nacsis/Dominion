package play

import (
	"perun.network/go-perun/channel"
	"perun.network/go-perun/wire"
)

func playerAlias(addr wire.Address) string {
	for alias, cfg := range config.Peers {
		if cfg.perunID.Equals(addr) {
			return alias
		}
	}
	return addr.String()
}

func (n *node) playerAlias(idx channel.Index) string {
	for _, peer := range n.peers {
		perunID := peer.ch.Params().Parts[idx]
		return playerAlias(perunID)
	}
	n.log.Panic("player not found")
	return ""
}
