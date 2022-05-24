package app

import (
	"fmt"
	"perun.network/go-perun/channel"
	"perun.network/go-perun/wallet"
	"perun.network/perun-examples/app-channel/app/util"
)

func ValidStateFormat(s *channel.State) *DominionAppData {
	sData, ok := s.Data.(*DominionAppData)
	if !ok {
		panic(fmt.Sprintf("to state: invalid data type: %T", s.Data))
	}
	return sData
}

func ValidActorInformation(currentActor, nextActor uint8, parts []wallet.Address, idx channel.Index) {
	// Check actor.
	if currentActor != util.Uint8safe(uint16(idx)) {
		panic(fmt.Errorf("invalid actor: expected %v, got %v", currentActor, idx))
	}

	ValidWalletLen(parts)
	NextActorIsInRange(nextActor)

	expectedToNextActor := CalcNextActor(currentActor)
	if nextActor != expectedToNextActor {
		panic(fmt.Errorf("invalid next actor: expected %v, got %v", expectedToNextActor, nextActor))
	}
}

func ValidWalletLen(parts []wallet.Address) {
	// Check next actor.
	if len(parts) != util.NumPlayers {
		panic("invalid number of participants")
	}
}

func NextActorIsInRange(nextActor uint8) {
	// Check next actor.
	if nextActor >= util.NumPlayers {
		panic(fmt.Errorf("invalid next actor: got %d, expected < %d", nextActor, util.NumPlayers))
	}
}
