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

	toData, ok := to.Data.(*app.DominionAppData)
	if !ok {
		n.log.Errorf("toData is in an invalid data format %T", toData)
		return
	}

	// ###################################
	// --- Handle automatic actions ------
	// ###################################

	if n.ownTurn(toData) {

		// relvant info
		initHandDrawn := toData.Turn.MandatoryPartFulfilled

		// user info prompts

		if toData.Turn.PerformedAction == util.DrawCard {
			for _, peer := range n.peers {
				handCards := toData.CardDecks[peer.ch.Idx()].HandPile
				if initHandDrawn {
					fmt.Printf("Hand: %s\n", handCards.Pretty())
				} else {
					fmt.Printf("You drew 1 %s.\n", handCards.Cards[handCards.Length()-1].CardType)
				}
			}
		}

		// >> initial card drawing

		toData.Turn.IsActionAllowed(util.RngCommit)
		if !initHandDrawn && toData.Turn.IsActionAllowed(util.RngCommit) {
			go n.drawCardStart()
			return
		}

		if toData.Turn.PerformedAction == util.RngCommit {
			for _, peer := range n.peers {
				go peer.ch.RngTouch()
				return
			}
		}

		if toData.Turn.PerformedAction == util.RngTouch {
			for _, peer := range n.peers {
				go peer.ch.RngRelease(n.preimage)
				return
			}
		}

		if toData.Turn.PerformedAction == util.RngRelease {
			for _, peer := range n.peers {
				go peer.ch.DrawOneCard()
				return
			}
		}
	}

}
