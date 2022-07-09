package app

import (
	"math/big"

	"perun.network/go-perun/channel"
	"perun.network/perun-examples/dominion-cli/app/util"
)

// EndTurn ends current Turn by switching actors
func (d *DominionAppData) EndTurn(actorIdx channel.Index) error {
	errorInfo := util.ErrorInfo{FunctionName: "EndTurn", FileName: util.ErrorConstDATA}

	//------ Checks ------
	if d.Turn.NextActor != uint8(actorIdx) {
		return errorInfo.ThrowError("Wrong Actor")
	}
	if !d.Turn.IsActionAllowed(util.EndTurn) {
		return errorInfo.ThrowError("EndTurn is not an allowed action")
	}

	//------ Perform action ------
	err := d.CardDecks[actorIdx].DiscardHandCards()
	if err != nil {
		return errorInfo.ForwardError(err)
	}
	err = d.CardDecks[actorIdx].DiscardPlayedCards()
	if err != nil {
		return errorInfo.ForwardError(err)
	}
	err = d.CardDecks[actorIdx].ResetResources()
	if err != nil {
		return errorInfo.ForwardError(err)
	}

	//------ Update Turn ------

	d.TurnAfter(util.EndTurn, *EmptyParams())

	return nil
}

//ComputeFinalBalances calculate final balances for game end
func (d *DominionAppData) ComputeFinalBalances(b channel.Balances) channel.Balances {

	winner, err := d.GetWinnerId()
	if err != nil {
		return b.Clone()
	}
	WinnerBal := b.Sum()
	if winner == 1 {
		return channel.Balances{[]channel.Bal{big.NewInt(0)}, WinnerBal}
	} else {
		return channel.Balances{WinnerBal, []channel.Bal{big.NewInt(0)}}
	}
}

//GetWinnerId calculate winner id
func (d *DominionAppData) GetWinnerId() (int, error) {
	errorInfo := util.ErrorInfo{FunctionName: "GetWinnerId", FileName: util.ErrorConstDATA}

	cardsOne := d.CardDecks[0].VictoryPointInDeck()
	cardsTwo := d.CardDecks[1].VictoryPointInDeck()

	if cardsOne == cardsTwo {
		return -1, errorInfo.ThrowError("Player has same amount of VictoryPoints")
	} else if cardsOne > cardsTwo {
		return 0, nil
	} else {
		return 1, nil
	}
}

// TurnAfter generate Turn state after performed action
func (d *DominionAppData) TurnAfter(at util.GeneralTypesOfActions, params Params) {

	if d.isFinalGameStateReached() {
		d.Turn.Params = params
		d.Turn.PerformedAction = at
		d.Turn.SetAllowed(util.GameEnd)
		d.Turn.SetNextActor()
	}

	switch at {
	case util.RngCommit:
		d.Turn.Params = params
		d.Turn.PerformedAction = util.RngCommit
		d.Turn.SetAllowed(util.RngTouch)
		d.Turn.SetNextActor()
		break
	case util.RngTouch:
		d.Turn.Params = params
		d.Turn.PerformedAction = util.RngTouch
		d.Turn.SetAllowed(util.RngRelease)
		d.Turn.SetNextActor()
		break
	case util.RngRelease:
		d.Turn.Params = params
		d.Turn.PerformedAction = util.RngRelease
		d.Turn.SetAllowed(util.DrawCard)
		break
	case util.DrawCard:
		d.Turn.Params = params
		d.Turn.PerformedAction = util.DrawCard
		var allowedActions []util.GeneralTypesOfActions

		if !d.Turn.MandatoryPartFulfilled && d.CardDecks[d.Turn.NextActor].IsInitialHandDrawn() {
			d.Turn.MandatoryPartFulfilled = true
		}

		if d.Turn.MandatoryPartFulfilled {
			allowedActions = d._GetAllowedDeckActions()
			allowedActions = append(allowedActions, util.EndTurn)
		} else {
			allowedActions = []util.GeneralTypesOfActions{util.RngCommit}
		}

		d.Turn.SetAllowed(allowedActions...)
		break
	case util.PlayCard:
		d.Turn.Params = params
		d.Turn.PerformedAction = util.PlayCard
		allowedActions := d._GetAllowedDeckActions()
		allowedActions = append(allowedActions, util.EndTurn)
		d.Turn.SetAllowed(allowedActions...)
		break
	case util.BuyCard:
		d.Turn.Params = params
		d.Turn.PerformedAction = util.BuyCard
		allowedActions := d._GetAllowedDeckActions()
		allowedActions = append(allowedActions, util.EndTurn)
		d.Turn.SetAllowed(allowedActions...)
		break
	case util.EndTurn:
		d.Turn.Params = params
		d.Turn.PerformedAction = util.EndTurn
		d.Turn.MandatoryPartFulfilled = false
		d.Turn.SetAllowed(util.RngCommit)
		d.Turn.SetNextActor()
		break
	case util.GameEnd:
		d.Turn.Params = params
		d.Turn.PerformedAction = util.GameEnd
		d.Turn.MandatoryPartFulfilled = false
		d.Turn.SetAllowed()
		break
	}
}

// isFinalGameStateReached helper function to avoid duplicated code
func (d *DominionAppData) isFinalGameStateReached() bool {
	return d.Stock.EmptyCardSets() >= 3 || d.Stock.IsBigVictoryCardEmpty()
}

// _GetAllowedDeckActions helper function to avoid duplicated code
func (d *DominionAppData) _GetAllowedDeckActions() []util.GeneralTypesOfActions {

	var allowedActions []util.GeneralTypesOfActions
	currentDeck := d.CardDecks[d.Turn.NextActor]
	if currentDeck.IsPlayActionPossible() {
		allowedActions = append(allowedActions, util.PlayCard)
	}
	if currentDeck.IsDrawActionPossible() {
		allowedActions = append(allowedActions, util.RngCommit)
	}
	if currentDeck.IsBuyActionPossible() {
		allowedActions = append(allowedActions, util.BuyCard)
	}

	return allowedActions
}

// EndGame ends game
func (d *DominionAppData) EndGame(actorIdx channel.Index) error {
	errorInfo := util.ErrorInfo{FunctionName: "EndGame", FileName: util.ErrorConstDATA}

	//------ Checks ------
	if d.Turn.NextActor != uint8(actorIdx) {
		return errorInfo.ThrowError("Wrong Actor")
	}
	if !d.Turn.IsActionAllowed(util.GameEnd) {
		return errorInfo.ThrowError("GameEnd is not an allowed action")
	}

	d.TurnAfter(util.GameEnd, *EmptyParams())
	return nil
}

//------------------------ Decks ------------------------

// DrawCard draws one card to the hand pile.
// A Rng need to be performed before.
func (d *DominionAppData) DrawCard(actorIdx channel.Index) error {
	errorInfo := util.ErrorInfo{FunctionName: "DrawCard", FileName: util.ErrorConstDATA}

	//------ Checks ------
	if d.Turn.NextActor != uint8(actorIdx) {
		return errorInfo.ThrowError("Wrong Actor")
	}
	if !d.Turn.IsActionAllowed(util.DrawCard) {
		return errorInfo.ThrowError("DrawCard is not an allowed action")
	}

	//------ Perform action ------
	value, err := d.Rng.RNGValue()
	if err != nil {
		return errorInfo.ForwardError(err)
	}

	err = d.CardDecks[actorIdx].DrawCard(value)
	if err != nil {
		return errorInfo.ForwardError(err)
	}
	// Rng was used, therefore delete it
	d.Rng = RNG{}

	//------ Update Turn ------
	d.TurnAfter(util.DrawCard, *EmptyParams())

	return nil
}

// PlayCard plays one card of the hand pile.
func (d *DominionAppData) PlayCard(actorIdx channel.Index, playCardIndex uint8, followUpIndices []uint8, followUpCardType util.CardType) error {
	errorInfo := util.ErrorInfo{FunctionName: "PlayCard", FileName: util.ErrorConstDATA}

	//------ Checks ------
	if d.Turn.NextActor != uint8(actorIdx) {
		return errorInfo.ThrowError("Wrong Actor")
	}
	if !d.Turn.IsActionAllowed(util.PlayCard) {
		return errorInfo.ThrowError("PlayCard is not an allowed action")
	}

	//------ Perform action ------
	card, err := d.CardDecks[actorIdx].GetHandCardWithIndex(uint(playCardIndex))
	if err != nil {
		return errorInfo.ForwardError(err)
	}

	err = d.CardDecks[actorIdx].UpdateResourcesAfterPlayedCard(card)
	if err != nil {
		return errorInfo.ForwardError(err)
	}

	params := Params{MainTarget: card.CardType, SecondLvlTarget: followUpCardType, SecondLvlIndices: followUpIndices}
	err = d._HandleAction(params)

	if err != nil {
		return errorInfo.ForwardError(err)
	}

	//------ Update Turn ------

	d.TurnAfter(util.PlayCard, params)

	return nil
}

// _HandleAction helper function to avoid duplicated code
func (d *DominionAppData) _HandleAction(params Params) error {
	errorInfo := util.ErrorInfo{FunctionName: "_HandleAction", FileName: util.ErrorConstDATA}
	playedCard := Card{}
	playedCard.Of([]byte{byte(params.MainTarget)})
	switch params.MainTarget {
	case util.Feast:
		card := Card{}
		card.Of([]byte{byte(params.SecondLvlTarget)})
		if card.BuyCost > 5 {
			return errorInfo.ThrowError("Selected Card is too expensive")
		}

		// Buyables selected Card
		card, err := d.Stock.TakeOffCard(params.SecondLvlTarget)
		if err != nil {
			return errorInfo.ForwardError(err)
		}
		err = d.CardDecks[d.Turn.NextActor].MoveToDiscardPile(card)
		if err != nil {
			return errorInfo.ForwardError(err)
		}

		// Trash played cardToBuy
		err = d.Stock.TrashCard(params.MainTarget)
		if err != nil {
			return errorInfo.ForwardError(err)
		}
		break
	case util.Chapel:
		if len(params.SecondLvlIndices) > 4 {
			return errorInfo.ThrowError("Only 4 cardToBuy from the Hand Pile can be trashed")
		}

		// Remove Card with index from hand
		cards, err := d.CardDecks[d.Turn.NextActor].HandPile.GetAndRemoveCardWithIndices(params.SecondLvlIndices)
		if err != nil {
			return errorInfo.ForwardError(err)
		}
		for _, card := range cards {
			// Trash cardToBuy
			err = d.Stock.TrashCard(card.CardType)
			if err != nil {
				return errorInfo.ForwardError(err)
			}
		}

		// Move played cardToBuy to played pile
		err = d.CardDecks[d.Turn.NextActor].MoveToPlayedPile(playedCard)
		if err != nil {
			return errorInfo.ForwardError(err)
		}
		break
	case util.Workshop:
		card := Card{}
		card.Of([]byte{byte(params.SecondLvlTarget)})
		if card.BuyCost > 4 {
			return errorInfo.ThrowError("Selected Card is too expensive")
		}

		// Buyables selected Card
		card, err := d.Stock.TakeOffCard(params.SecondLvlTarget)
		if err != nil {
			return errorInfo.ForwardError(err)
		}
		err = d.CardDecks[d.Turn.NextActor].MoveToDiscardPile(card)
		if err != nil {
			return errorInfo.ForwardError(err)
		}

		// Move played cardToBuy to played pile
		err = d.CardDecks[d.Turn.NextActor].MoveToPlayedPile(playedCard)
		if err != nil {
			return errorInfo.ForwardError(err)
		}
		break
	case util.Remodel:
		if len(params.SecondLvlIndices) != 1 {
			return errorInfo.ThrowError("Only 1 cardToBuy from the Hand Pile can be trashed")
		}
		// Remove Card with index from hand
		removedCard, err := d.CardDecks[d.Turn.NextActor].HandPile.DrawCardWithIndex(uint(params.SecondLvlIndices[0]))
		if err != nil {
			return errorInfo.ForwardError(err)
		}
		// Trash cardToBuy
		err = d.Stock.TrashCard(removedCard.CardType)
		if err != nil {
			return errorInfo.ForwardError(err)
		}

		cardToBuy := Card{}
		cardToBuy.Of([]byte{byte(params.SecondLvlTarget)})
		if removedCard.BuyCost+2 < cardToBuy.BuyCost {
			return errorInfo.ThrowError("Selected Card is too expensive")
		}

		// Buyables selected Card
		cardToBuy, err = d.Stock.TakeOffCard(params.SecondLvlTarget)
		if err != nil {
			return errorInfo.ForwardError(err)
		}
		err = d.CardDecks[d.Turn.NextActor].MoveToDiscardPile(cardToBuy)
		if err != nil {
			return errorInfo.ForwardError(err)
		}

		// Move played cardToBuy to played pile
		err = d.CardDecks[d.Turn.NextActor].MoveToPlayedPile(playedCard)
		if err != nil {
			return errorInfo.ForwardError(err)
		}
		break
	case util.Cellar:
		// Remove Card with index from hand
		cards, err := d.CardDecks[d.Turn.NextActor].HandPile.GetAndRemoveCardWithIndices(params.SecondLvlIndices)
		if err != nil {
			return errorInfo.ForwardError(err)
		}
		for _, card := range cards {
			// Trash cardToBuy
			err = d.CardDecks[d.Turn.NextActor].MoveToDiscardPile(card)
			if err != nil {
				return errorInfo.ForwardError(err)
			}
		}

		// Move played cardToBuy to played pile
		err = d.CardDecks[d.Turn.NextActor].MoveToPlayedPile(playedCard)
		if err != nil {
			return errorInfo.ForwardError(err)
		}

		// Extent Drawable cards
		d.CardDecks[d.Turn.NextActor].Resources[util.DrawableCards] += uint8(len(params.SecondLvlIndices))

		break
	case util.Mine:
		if len(params.SecondLvlIndices) != 1 {
			return errorInfo.ThrowError("One money card to trash need to be selected")
		}
		// Remove Card with index from hand
		cardToTrash, err := d.CardDecks[d.Turn.NextActor].HandPile.DrawCardWithIndex(uint(params.SecondLvlIndices[0]))
		if err != nil {
			return errorInfo.ForwardError(err)
		}
		if !cardToTrash.IsMoneyCard() {
			return errorInfo.ThrowError("Card to trash needs to be a money card")
		}
		// Trash cardToBuy
		err = d.Stock.TrashCard(cardToTrash.CardType)
		if err != nil {
			return errorInfo.ForwardError(err)
		}

		card := Card{}
		card.Of([]byte{byte(params.SecondLvlTarget)})
		if !card.IsMoneyCard() {
			return errorInfo.ThrowError("Card to gain needs to be a money card")
		}

		if card.BuyCost > cardToTrash.BuyCost+3 {
			return errorInfo.ThrowError("Card to gain can only cost trashed card +3")
		}
		// Buyables selected Card
		cardToBuy := Card{}
		cardToBuy, err = d.Stock.TakeOffCard(params.SecondLvlTarget)
		if err != nil {
			return errorInfo.ForwardError(err)
		}
		err = d.CardDecks[d.Turn.NextActor].MoveToHandPile(cardToBuy)
		if err != nil {
			return errorInfo.ForwardError(err)
		}

		err = d.CardDecks[d.Turn.NextActor].MoveToPlayedPile(playedCard)
		if err != nil {
			return errorInfo.ForwardError(err)
		}
		break
	case util.Oasis:
		if len(params.SecondLvlIndices) != 1 {
			return errorInfo.ThrowError("One money card to discard need to be selected")
		}
		// Remove Card with index from hand
		cardToDiscard, err := d.CardDecks[d.Turn.NextActor].HandPile.DrawCardWithIndex(uint(params.SecondLvlIndices[0]))
		if err != nil {
			return errorInfo.ForwardError(err)
		}
		d.CardDecks[d.Turn.NextActor].DiscardedPile.AddCardToPile(cardToDiscard)

		// Move played cardToBuy to played pile
		err = d.CardDecks[d.Turn.NextActor].MoveToPlayedPile(playedCard)
		if err != nil {
			return errorInfo.ForwardError(err)
		}
		break
		break
	default:
		// Move Played cardToBuy to Pile
		err := d.CardDecks[d.Turn.NextActor].MoveToPlayedPile(playedCard)
		if err != nil {
			return errorInfo.ForwardError(err)
		}
	}
	return nil
}

// BuyCard Buyables one card for given CardType.
func (d *DominionAppData) BuyCard(actorIdx channel.Index, cardType util.CardType) error {
	errorInfo := util.ErrorInfo{FunctionName: "BuyCard", FileName: util.ErrorConstDATA}

	//------ Checks ------
	if d.Turn.NextActor != uint8(actorIdx) {
		return errorInfo.ThrowError("Wrong Actor")
	}
	if !d.Turn.IsActionAllowed(util.BuyCard) {
		return errorInfo.ThrowError("BuyCard is not an allowed action")
	}

	//------ Perform action ------
	card, err := d.Stock.TakeOffCard(cardType)
	if err != nil {
		return errorInfo.ForwardError(err)
	}

	err = d.CardDecks[d.Turn.NextActor].BoughtCard(card)
	if err != nil {
		return errorInfo.ForwardError(err)
	}

	//------ Update Turn ------
	d.TurnAfter(util.BuyCard, Params{MainTarget: cardType})

	return nil
}

//------------------------ Rng ------------------------

// RngCommit set an image for Rng
// Players who want to draw a card need to start by committing to a preimage
func (d *DominionAppData) RngCommit(actorIdx channel.Index, image [util.PreImageSizeByte]byte) error {
	errorInfo := util.ErrorInfo{FunctionName: "RngCommit", FileName: util.ErrorConstDATA}

	//------ Checks ------
	if d.Turn.NextActor != uint8(actorIdx) {
		return errorInfo.ThrowError("Wrong Actor")
	}
	if !d.Turn.IsActionAllowed(util.RngCommit) {
		return errorInfo.ThrowError("RngCommit is not an allowed action")
	}

	//------ Perform action ------
	err := d.Rng.Commit(image)
	if err != nil {
		return errorInfo.ForwardError(err)
	}

	//------ Update Turn ------
	d.TurnAfter(util.RngCommit, *EmptyParams())

	return nil
}

// RngTouch set a second preimage for Rng
// Players need accept the set image by selecting a second preimage
func (d *DominionAppData) RngTouch(actorIdx channel.Index) error {
	errorInfo := util.ErrorInfo{FunctionName: "RngTouch", FileName: util.ErrorConstDATA}

	//------ Checks ------
	if d.Turn.NextActor != uint8(actorIdx) {
		return errorInfo.ThrowError("Actor can't touch for his own Rng")
	}
	if !d.Turn.IsActionAllowed(util.RngTouch) {
		return errorInfo.ThrowError("RngTouch is not an allowed action")
	}

	//------ Perform action ------
	err := d.Rng.Touch()
	if err != nil {
		return errorInfo.ForwardError(err)
	}

	//------ Update Turn ------
	d.TurnAfter(util.RngTouch, *EmptyParams())

	return nil
}

// RngRelease set preimage of set image
// Players publish their preimage of the image, s.t. a shared random value can be calculated
func (d *DominionAppData) RngRelease(actorIdx channel.Index, preImage [util.PreImageSizeByte]byte) error {
	errorInfo := util.ErrorInfo{FunctionName: "RngRelease", FileName: util.ErrorConstDATA}

	//------ Checks ------
	if d.Turn.NextActor != uint8(actorIdx) {
		return errorInfo.ThrowError("Wrong Actor")
	}
	if !d.Turn.IsActionAllowed(util.RngRelease) {
		return errorInfo.ThrowError("RngRelease is not an allowed action")
	}

	//------ Perform action ------
	err := d.Rng.Release(preImage)
	if err != nil {
		return errorInfo.ForwardError(err)
	}

	//------ Update Turn ------
	d.TurnAfter(util.RngRelease, *EmptyParams())

	return nil
}
