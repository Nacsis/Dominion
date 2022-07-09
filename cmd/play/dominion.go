package play

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"perun.network/go-perun/channel"
	"perun.network/go-perun/wire"
	"perun.network/perun-examples/dominion-cli/app"
	"perun.network/perun-examples/dominion-cli/app/util"
	"perun.network/perun-examples/dominion-cli/global"
)

// needs to be called after Open
func (n *node) GetAppStateData() *app.DominionAppData {

	for _, peer := range n.peers { // TODO different access approach in >2 peers channel
		data := peer.ch.GetAppStateData()
		return data
	}

	n.log.Errorf("Cannot retrieve app data... no peers yet")
	return nil
}

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
		perunID := peer.ch.Peers()[idx]
		return playerAlias(perunID)
	}
	n.log.Debug("player not found")
	return ""
}

func (n *node) ownTurn(data *app.DominionAppData) bool {
	nextActorIdx := channel.Index(data.Turn.NextActor)

	for _, peer := range n.peers {
		return peer.ch.Idx() == nextActorIdx
	}

	n.log.Error("no peer yet")
	return false
}

func (n *node) drawCardStart() {

	fmt.Println("drawing card...")

	preimage := util.SliceToPreImageByte(global.RandomBytes(util.PreImageSizeByte))
	for _, peer := range n.peers {
		peer.ch.RngCommit(preimage)
	}

	n.preimage = preimage
}

// Start the dominion cli game or see if an other see who needs to start. An open channel is required.
func (n *node) Start() error {
	n.mtx.Lock()
	defer n.mtx.Unlock()
	n.log.Traceln("Start...")
	data := n.GetAppStateData()

	ownTurn := n.ownTurn(data)
	// initHandDrawn := toData.Turn.MandatoryPartFulfilled
	if ownTurn {
		fmt.Println("Starting game :)")
		n.drawCardStart()
	} else {
		firstActor := n.playerAlias(channel.Index(data.Turn.NextActor))
		fmt.Printf("%s needs to start.\n", firstActor)
	}

	return nil
}

// Play a card. usage: buy [card name]\n Card name is not case sensitive. The first card of matching type will be chosen.
func (n *node) PlayCard(args []string) error {
	n.mtx.Lock()
	defer n.mtx.Unlock()
	n.log.Traceln("PlayCard...")
	for _, peer := range n.peers {

		data := peer.ch.GetAppStateData()

		ownTurn := n.ownTurn(data)

		if ownTurn {

			hand := data.CardDecks[peer.ch.Idx()].HandPile
			cid, err := strconv.Atoi(args[0])
			if err != nil { // try name

				c, ok := app.NewCard(args[0])

				if !ok {
					return errors.New("No valid card name")
				}

				for i, card := range hand.Cards {
					if c.CardType == card.CardType {
						cid = i
					}
				}
			} else {
				if cid >= hand.Length() {
					return errors.Errorf("No valid position in hand. (wanted: <%v, given: %v)", hand.Length(), cid)
				}
			}

			followUpCard, _ := app.NewCard(args[2])
			followUpIndices, _ := strToUint8List(args[1])
			// TODO followUpIndices
			peer.ch.PlayCard(uint8(cid), followUpIndices, followUpCard.CardType)

		} else {
			firstActor := n.playerAlias(channel.Index(data.Turn.NextActor))
			fmt.Printf("Not your turn! %s goes next.\n", firstActor)
		}

		return nil
	}

	return errors.Errorf("No connected peer.")
}

// Buy a card. usage: buy [card name]\n Card name is not case sensitive.
func (n *node) BuyCard(args []string) error {
	n.mtx.Lock()
	defer n.mtx.Unlock()
	n.log.Traceln("BuyCard...")
	for _, peer := range n.peers {

		data := peer.ch.GetAppStateData()

		ownTurn := n.ownTurn(data)

		if ownTurn {

			c, _ := app.NewCard(args[0])
			if err := peer.ch.BuyCard(c.CardType, false); err != nil {

				// handle Off-Chain error
				msg := fmt.Sprint("BuyCard action failed or peer did not respond. Do you want to ForceUpdate on the blockchain? (y/n)")
				Prompt(msg, func(userInput string) {

					if userInput == "y" {
						if err := peer.ch.BuyCard(c.CardType, true); err != nil {
							fmt.Printf("❌ ForceUpdate failed for action 'buy' with args %v\n", args)
						} else {
							fmt.Printf("✅ ForceUpdate success!\n")
						}
					}

				})
			}

		} else {
			firstActor := n.playerAlias(channel.Index(data.Turn.NextActor))
			fmt.Printf("Not your turn! %s goes next.\n", firstActor)
		}

		return nil
	}

	return errors.Errorf("No connected peer.")
}

// End your turn. If the game is final, EndGame is automatically called instead and settlement and payout are triggered.
func (n *node) EndTurnOrGame() error {
	n.mtx.Lock()
	defer n.mtx.Unlock()
	n.log.Traceln("EndTurnOrGame...")
	for _, peer := range n.peers {

		data := peer.ch.GetAppStateData()

		ownTurn := n.ownTurn(data)

		if ownTurn {

			if data.Turn.IsActionAllowed(util.GameEnd) {
				peer.ch.EndGame()
			} else {
				peer.ch.EndTurn()
			}

		} else {
			firstActor := n.playerAlias(channel.Index(data.Turn.NextActor))
			fmt.Printf("Not your turn! %s goes next.\n", firstActor)
		}

		return nil
	}

	return errors.Errorf("No connected peer.")
}
