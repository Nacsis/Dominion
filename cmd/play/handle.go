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
	fmt.Println("Hi! OnUpdate has been called! :)")
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
		n.drawCardStart()
		return
	}

	fmt.Printf("Turn: %+v\n", toData.Turn)
	fmt.Printf("ownTurn?: %v\n", ownTurn)
	if toData.Turn.PerformedAction == util.RngCommit && ownTurn {
		fmt.Println("Should perform RngTouch")
		for _, peer := range n.peers {
			fmt.Println("Should REALLY perform RngTouch")
			peer.ch.RngTouch()
			return
		}
	}

	if toData.Turn.PerformedAction == util.RngTouch && ownTurn {
		for _, peer := range n.peers {
			peer.ch.RngRelease(n.preimage)
			return
		}
	}

	if toData.Turn.PerformedAction == util.RngRelease && ownTurn {
		for _, peer := range n.peers {
			peer.ch.DrawOneCard()
			return
		}
	}

}
