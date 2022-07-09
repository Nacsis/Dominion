package app

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"

	"perun.network/go-perun/channel"
	"perun.network/perun-examples/dominion-cli/app/util"
	"perun.network/perun-examples/dominion-cli/global"
)

// Init sets up initial game state
func (a *DominionApp) Init(firstActor channel.Index) *DominionAppData {
	var dominionAppData DominionAppData
	err := dominionAppData.Init(firstActor)
	if err != nil {
		log.Fatalf("Init DominionAppData failed with error: %T", err)
	}
	return &dominionAppData
}

// _ValidState validate the last performed action lead to current state
func (a *DominionApp) _ValidState(fromData, toData DominionAppData, idx channel.Index) error {
	errorInfo := util.ErrorInfo{FunctionName: "_ValidState", FileName: util.ErrorConstAPP}

	// Deep Copy. Seems like the other ways didn't work correctly ... idk
	origJSON, _ := json.Marshal(fromData)
	fromDataClone := DominionAppData{}
	json.Unmarshal(origJSON, &fromDataClone)

	var err error

	switch toData.Turn.PerformedAction {
	case util.RngCommit:
		if uint16(len(toData.Rng.ImageA)) != util.HashSizeByte {
			err = errorInfo.ThrowError(fmt.Sprintf("given Image has not correct size of %d", util.HashSizeByte))
			break
		}
		dummyPreImage := util.SliceToPreImageByte(global.RandomBytes(util.PreImageSizeByte))
		err = fromDataClone.RngCommit(idx, dummyPreImage)
		fromDataClone.Rng.ImageA = toData.Rng.ImageA
		break
	case util.RngTouch:
		if uint16(len(toData.Rng.PreImageB)) != util.PreImageSizeByte {
			err = errorInfo.ThrowError(fmt.Sprintf("given preImage has not correct size of %d", util.PreImageSizeByte))
			break
		}
		err = fromDataClone.RngTouch(idx)
		fromDataClone.Rng.PreImageB = toData.Rng.PreImageB
		break
	case util.RngRelease:
		if uint16(len(toData.Rng.PreImageA)) != util.PreImageSizeByte {
			err = errorInfo.ThrowError(fmt.Sprintf("given preImage has not correct size of %d", util.PreImageSizeByte))
			break
		}
		err = fromDataClone.RngRelease(idx, toData.Rng.PreImageA)
		break
	case util.DrawCard:
		err = fromDataClone.DrawCard(idx)
		break
	case util.PlayCard:
		playedIndex := 0
		toCards := toData.CardDecks[idx].HandPile.Cards
		fromCards := fromDataClone.CardDecks[idx].HandPile.Cards
		if len(fromCards) <= len(toCards) {
			err = errorInfo.ThrowError("toCards is to long. No card was played")
			break
		}
		for i := 0; i < len(fromCards); i++ {
			if i == len(toCards) || toCards[i].CardType != fromCards[i].CardType {
				playedIndex = i
				break
			}
		}
		if fromDataClone.CardDecks[idx].HandPile.Cards[playedIndex].CardType != toData.Turn.Params.MainTarget {
			err = errorInfo.ThrowError("Played card doesn't match card index")
			break
		}

		err = fromDataClone.PlayCard(idx, uint8(playedIndex), toData.Turn.Params.SecondLvlIndices, toData.Turn.Params.SecondLvlTarget)
		break
	case util.BuyCard:
		cardTypeToBuy := toData.CardDecks[idx].DiscardedPile.Cards[toData.CardDecks[idx].DiscardedPile.Length()-1].CardType
		err = fromDataClone.BuyCard(idx, cardTypeToBuy)
		break
	case util.EndTurn:
		err = fromDataClone.EndTurn(idx)
		break
	case util.GameEnd:
		err = fromDataClone.EndGame(idx)
		break
	}

	if err != nil {
		return errorInfo.ForwardError(err)
	}

	// fmt.Printf("fromDataClone: %+v\n", fromDataClone)
	// fmt.Printf("toData:        %+v\n", toData)
	// toDataClone := toData.Clone2DominionAppData()

	if !reflect.DeepEqual(fromDataClone.Turn, toData.Turn) {
		fmt.Println("Struct 'Turn' could not be replicated for action")
		fmt.Printf("Turn Prop: %+v\n", toData.Turn)
		fmt.Printf("Turn Repl: %+v\n", fromDataClone.Turn)
		if !reflect.DeepEqual(fromDataClone.Turn.MandatoryPartFulfilled, toData.Turn.MandatoryPartFulfilled) {
			fmt.Println("Struct 'Turn.MandatoryPartFulfilled' could not be replicated for action")
			fmt.Printf("Turn.MandatoryPartFulfilled Prop: %+v\n", toData.Turn.MandatoryPartFulfilled)
			fmt.Printf("Turn.MandatoryPartFulfilled Repl: %+v\n", fromDataClone.Turn.MandatoryPartFulfilled)
		}
		if !reflect.DeepEqual(fromDataClone.Turn.NextActor, toData.Turn.NextActor) {
			fmt.Println("Struct 'Turn.NextActor' could not be replicated for action")
			fmt.Printf("Turn.NextActor Prop: %+v\n", toData.Turn.NextActor)
			fmt.Printf("Turn.NextActor Repl: %+v\n", fromDataClone.Turn.NextActor)
		}
		if !reflect.DeepEqual(fromDataClone.Turn.Params, toData.Turn.Params) {
			fmt.Println("Struct 'Turn.Params' could not be replicated for action")
			fmt.Printf("Turn.Params Prop: %+v\n", toData.Turn.Params)
			fmt.Printf("Turn.Params Repl: %+v\n", fromDataClone.Turn.Params)
		}
		if !reflect.DeepEqual(fromDataClone.Turn.PerformedAction, toData.Turn.PerformedAction) {
			fmt.Println("Struct 'Turn.PerformedAction' could not be replicated for action")
			fmt.Printf("Turn.PerformedAction Prop: %+v\n", toData.Turn.PerformedAction)
			fmt.Printf("Turn.PerformedAction Repl: %+v\n", fromDataClone.Turn.PerformedAction)
		}
		if !reflect.DeepEqual(fromDataClone.Turn.PossibleActions, toData.Turn.PossibleActions) {
			fmt.Println("Struct 'Turn.PossibleActions' could not be replicated for action")
			fmt.Printf("Turn.PossibleActions Prop: %+v\n", toData.Turn.PossibleActions)
			fmt.Printf("Turn.PossibleActions Repl: %+v\n", fromDataClone.Turn.PossibleActions)
		}
		if !reflect.DeepEqual(fromDataClone.Turn.Params.MainTarget, toData.Turn.Params.MainTarget) {
			fmt.Println("Struct 'Turn.Params.MainTarget' could not be replicated for action")
			fmt.Printf("Turn.Params.MainTarget Prop: %+v\n", toData.Turn.Params.MainTarget)
			fmt.Printf("Turn.Params.MainTarget Repl: %+v\n", fromDataClone.Turn.Params.MainTarget)
		}
		if !reflect.DeepEqual(fromDataClone.Turn.Params.SecondLvlIndices, toData.Turn.Params.SecondLvlIndices) {
			fmt.Println("Struct 'Turn.Params.SecondLvlIndices' could not be replicated for action")
			fmt.Printf("Turn.Params.SecondLvlIndices Prop: %+v\n", toData.Turn.Params.SecondLvlIndices)
			fmt.Printf("Turn.Params.SecondLvlIndices Repl: %+v\n", fromDataClone.Turn.Params.SecondLvlIndices)
		}
		if !(fromDataClone.Turn.Params.SecondLvlIndices == nil) {
			fmt.Println("fromDataClone 'Turn.Params.SecondLvlIndices' nil")
		}
		if !(toData.Turn.Params.SecondLvlIndices == nil) {
			fmt.Println("toData 'Turn.Params.SecondLvlIndices' nil")
		}
		if !reflect.DeepEqual(fromDataClone.Turn.Params.SecondLvlTarget, toData.Turn.Params.SecondLvlTarget) {
			fmt.Println("Struct 'Turn.Params.SecondLvlTarget' could not be replicated for action")
			fmt.Printf("Turn.Params.SecondLvlTarget Prop: %+v\n", toData.Turn.Params.SecondLvlTarget)
			fmt.Printf("Turn.Params.SecondLvlTarget Repl: %+v\n", fromDataClone.Turn.Params.SecondLvlTarget)
		}
	}
	if !reflect.DeepEqual(fromDataClone.Stock, toData.Stock) {
		fmt.Println("Struct 'Stock' could not be replicated for action")
	}
	if !reflect.DeepEqual(fromDataClone.CardDecks, toData.CardDecks) {
		fmt.Println("Struct 'CardDecks' could not be replicated for action")
	}
	if !reflect.DeepEqual(fromDataClone.Rng, toData.Rng) {
		fmt.Println("Struct 'Rng' could not be replicated for action")
	}

	s := strings.Builder{}
	fromDataClone.Encode(&s)
	bytesRepl := s.String()
	s = strings.Builder{}
	toData.Encode(&s)
	bytesProp := s.String()

	if bytesRepl != bytesProp {
		fmt.Printf("bytecode unequal: \n1) %s\n2) %s\n", bytesProp, bytesRepl)
	}

	// if !reflect.DeepEqual(fromDataClone, toData) {
	// 	return errorInfo.ThrowError("State transition could not be replicated for action")
	// }

	jsonRepl, _ := json.Marshal(fromDataClone)
	jsonProp, _ := json.Marshal(toData)

	if !DeepEqual(jsonProp, jsonRepl) {
		return errorInfo.ThrowError("State transition could not be replicated for action")
	}

	return nil
}
func DeepEqual(v1, v2 interface{}) bool {
	if reflect.DeepEqual(v1, v2) {
		return true
	}
	var x1 interface{}
	bytesA, _ := json.Marshal(v1)
	_ = json.Unmarshal(bytesA, &x1)
	var x2 interface{}
	bytesB, _ := json.Marshal(v2)
	_ = json.Unmarshal(bytesB, &x2)
	if reflect.DeepEqual(x1, x2) {
		return true
	}
	return false
}

// EndTurn ends current Turn by switching actors
func (a *DominionApp) EndTurn(s *channel.State, actorIdx channel.Index) error {
	errorInfo := util.ErrorInfo{FunctionName: "EndTurn", FileName: util.ErrorConstAPP}

	dominionAppData, ok := s.Data.(*DominionAppData)
	if !ok {
		return errorInfo.ThrowError(fmt.Sprintf("AppData is in an invalid data format %T", dominionAppData))
	}

	err := dominionAppData.EndTurn(actorIdx)
	if err != nil {
		return errorInfo.ForwardError(err)
	}

	return nil
}

// EndGame ends game by setting the isFinal value and Calculation final balance
func (a *DominionApp) EndGame(s *channel.State, actorIdx channel.Index) error {
	errorInfo := util.ErrorInfo{FunctionName: "EndGame", FileName: util.ErrorConstAPP}

	dominionAppData, ok := s.Data.(*DominionAppData)
	if !ok {
		return errorInfo.ThrowError(fmt.Sprintf("AppData is in an invalid data format %T", dominionAppData))
	}

	err := dominionAppData.EndGame(actorIdx)
	if err != nil {
		return errorInfo.ForwardError(err)
	}

	s.IsFinal = true
	s.Balances = dominionAppData.ComputeFinalBalances(s.Balances)

	return nil
}

//------------------------ Decks ------------------------

// DrawCard draws one card to the hand pile.
// A Rng need to be performed before.
func (a *DominionApp) DrawCard(s *channel.State, actorIdx channel.Index) error {
	errorInfo := util.ErrorInfo{FunctionName: "DrawCard", FileName: util.ErrorConstAPP}

	dominionAppData, ok := s.Data.(*DominionAppData)
	if !ok {
		return errorInfo.ThrowError(fmt.Sprintf("AppData is in an invalid data format %T", dominionAppData))
	}

	err := dominionAppData.DrawCard(actorIdx)
	if err != nil {
		return errorInfo.ForwardError(err)
	}
	return nil
}

// PlayCard plays one card of the hand pile.
func (a *DominionApp) PlayCard(s *channel.State, actorIdx channel.Index, index uint8, followUpIndices []uint8, followUpCardType util.CardType) error {
	errorInfo := util.ErrorInfo{FunctionName: "PlayCard", FileName: util.ErrorConstAPP}

	dominionAppData, ok := s.Data.(*DominionAppData)
	if !ok {
		return errorInfo.ThrowError(fmt.Sprintf("AppData is in an invalid data format %T", dominionAppData))
	}

	err := dominionAppData.PlayCard(actorIdx, index, followUpIndices, followUpCardType)
	if err != nil {
		return errorInfo.ForwardError(err)
	}
	return nil
}

// BuyCard Buyables one card for given CardType.
func (a *DominionApp) BuyCard(s *channel.State, actorIdx channel.Index, cardType util.CardType) error {
	errorInfo := util.ErrorInfo{FunctionName: "BuyCard", FileName: util.ErrorConstAPP}

	dominionAppData, ok := s.Data.(*DominionAppData)
	if !ok {
		return errorInfo.ThrowError(fmt.Sprintf("BuyCard is in an invalid data format %T", dominionAppData))
	}

	err := dominionAppData.BuyCard(actorIdx, cardType)
	if err != nil {
		return errorInfo.ForwardError(err)
	}
	return nil
}

//------------------------ Rng ------------------------

// RngCommit set an image for Rng
// Players who want to draw a card need to start by committing to a preimage
func (a *DominionApp) RngCommit(s *channel.State, actorIdx channel.Index, preImage [util.PreImageSizeByte]byte) error {
	errorInfo := util.ErrorInfo{FunctionName: "RngCommit", FileName: util.ErrorConstAPP}

	dominionAppData, ok := s.Data.(*DominionAppData)
	if !ok {
		return errorInfo.ThrowError(fmt.Sprintf("AppData is in an invalid data format %T", dominionAppData))
	}

	err := dominionAppData.RngCommit(actorIdx, preImage)
	if err != nil {
		return errorInfo.ForwardError(err)
	}
	return nil
}

// RngTouch set a second preimage for Rng
// Players need accept the set image by selecting a second preimage
func (a *DominionApp) RngTouch(s *channel.State, actorIdx channel.Index) error {
	errorInfo := util.ErrorInfo{FunctionName: "RngTouch", FileName: util.ErrorConstAPP}

	dominionAppData, ok := s.Data.(*DominionAppData)
	if !ok {
		return errorInfo.ThrowError(fmt.Sprintf("AppData is in an invalid data format %T", dominionAppData))
	}

	err := dominionAppData.RngTouch(actorIdx)
	if err != nil {
		return errorInfo.ForwardError(err)
	}
	return nil
}

// RngRelease set preimage of set image
// Players publish their preimage of the image, s.t. a shared random value can be calculated
func (a *DominionApp) RngRelease(s *channel.State, actorIdx channel.Index, preImage [util.PreImageSizeByte]byte) error {
	errorInfo := util.ErrorInfo{FunctionName: "RngRelease", FileName: util.ErrorConstAPP}

	dominionAppData, ok := s.Data.(*DominionAppData)
	if !ok {
		return errorInfo.ThrowError(fmt.Sprintf("AppData is in an invalid data format %T", dominionAppData))
	}

	err := dominionAppData.RngRelease(actorIdx, preImage)
	if err != nil {
		return errorInfo.ForwardError(err)
	}
	return nil
}
