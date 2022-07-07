package play

import (
	"fmt"

	"perun.network/go-perun/channel"
)

// dominionClient "perun.network/perun-examples/dominion-cli/client"

// TODO implement
func (n *node) OnUpdate(from, to *channel.State) {
	fmt.Print("Hi! OnUpdate has been called! :)")
}
