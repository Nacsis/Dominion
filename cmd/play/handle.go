package play

import (
	"fmt"

	"perun.network/go-perun/channel"
	"perun.network/perun-examples/dominion-cli/app"
	"perun.network/perun-examples/dominion-cli/app/util"
)

// Triggered on all peers after successfully accepted channel update proposals
// Summarizes done activityies and auto-triggers micro actions if useful
func (n *node) OnUpdate(from, to *channel.State) {
	fmt.Print("Hi! OnUpdate has been called! :)")
	toData, ok := to.Data.(*app.DominionAppData)
	if !ok {
		n.log.Errorf("toData is in an invalid data format %T", toData)
		return
	}

	// ###################################
	// --- Handle automatic actions ------
	// ###################################

	// >> initial card drawing
	ownTurn := n.ownTurn(toData)
	initHandDrawn := toData.Turn.MandatoryPartFulfilled
	toData.Turn.IsActionAllowed(util.RngCommit)
	if ownTurn && !initHandDrawn && toData.Turn.IsActionAllowed(util.RngCommit) {
		n.drawCard()
	}

}
