package play

import (
	"fmt"

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
	data := n.GetAppStateData()

	ownTurn := n.ownTurn(data)

	if ownTurn {
		fmt.Println("TODO: IMPLEMENT!")

	} else {
		firstActor := n.playerAlias(channel.Index(data.Turn.NextActor))
		fmt.Printf("Not your turn! %s goes next.\n", firstActor)
	}

	return nil
}

// Buy a card. usage: buy [card name]\n Card name is not case sensitive.
func (n *node) BuyCard(args []string) error {
	n.mtx.Lock()
	defer n.mtx.Unlock()
	n.log.Traceln("BuyCard...")
	data := n.GetAppStateData()

	ownTurn := n.ownTurn(data)

	if ownTurn {
		fmt.Println("TODO: IMPLEMENT!")

	} else {
		firstActor := n.playerAlias(channel.Index(data.Turn.NextActor))
		fmt.Printf("Not your turn! %s goes next.\n", firstActor)
	}

	return nil
}

// End your turn. If the game is final, EndGame is automatically called instead and settlement and payout are triggered.
func (n *node) EndTurnOrGame() error {
	n.mtx.Lock()
	defer n.mtx.Unlock()
	n.log.Traceln("EndTurnOrGame...")
	data := n.GetAppStateData()

	ownTurn := n.ownTurn(data)

	if ownTurn {
		fmt.Println("TODO: IMPLEMENT!")

	} else {
		firstActor := n.playerAlias(channel.Index(data.Turn.NextActor))
		fmt.Printf("Not your turn! %s goes next.\n", firstActor)
	}

	return nil
}
