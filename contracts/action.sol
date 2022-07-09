pragma solidity ^0.7.0;

import "./data.sol";
import "./Util/constant.sol";
import "./rng.sol";

library ActionLib {// Performing actions and state and do the turnafter stuff
    function RngCommit(DataLib.DominionAppData memory data) internal pure {
        data.turn.performedAction = Constant.GeneralTypesOfActions.RngCommit;
        data.rng.PreImageA = Convert.bytesToByteArray(bytes(""));
        data.rng.ImageA = Convert.bytesToByteArray(bytes(""));
        data.rng.PreImageB = Convert.bytesToByteArray(bytes(""));
        Constant.GeneralTypesOfActions[] memory allowed = new Constant.GeneralTypesOfActions[](1);
        allowed[0] = Constant.GeneralTypesOfActions.RngTouch;
        TurnLib.setOneAllowed(data.turn, allowed);
        TurnLib.setNextActor(data.turn);
    }

    function RngTouch(DataLib.DominionAppData memory data) internal pure {
        data.turn.performedAction = Constant.GeneralTypesOfActions.RngTouch;
        Constant.GeneralTypesOfActions[] memory allowed = new Constant.GeneralTypesOfActions[](1);
        allowed[0] = Constant.GeneralTypesOfActions.RngRelease;
        TurnLib.setOneAllowed(data.turn, allowed);
        TurnLib.setNextActor(data.turn);
    }

    function RngRelease(DataLib.DominionAppData memory data, byte[] memory preImageA) internal pure {
        require(preImageA.length == Constant.HashSize);
        byte[] memory ImageOfPreImageA = Convert.Slice2Array(Convert.bytes32ToByteArray(keccak256(Convert.byteArrayToBytes(preImageA))), 0, Constant.HashSize);
        for (uint i = 0; i < ImageOfPreImageA.length; i++) {
            require(ImageOfPreImageA[i] == data.rng.ImageA[i], "RngRelease Incorrect preImage released");
        }
        data.rng.PreImageA = preImageA;
        data.turn.performedAction = Constant.GeneralTypesOfActions.RngRelease;
        Constant.GeneralTypesOfActions[] memory allowed = new Constant.GeneralTypesOfActions[](1);
        allowed[0] = Constant.GeneralTypesOfActions.DrawCard;
        TurnLib.setOneAllowed(data.turn, allowed);
    }


    function DrawCard(DataLib.DominionAppData memory data) internal pure {
        data.turn.performedAction = Constant.GeneralTypesOfActions.DrawCard;
        byte[] memory rngValue = RNGLib.RNGValue(data.rng);
        DeckLib.DrawCard(data.CardDecks[data.turn.nextActor], rngValue);
        data.rng = RNGLib.RNG(Convert.bytesToByteArray(bytes("")),Convert.bytesToByteArray(bytes("")),Convert.bytesToByteArray(bytes("")));

        DeckLib.Deck memory deck = data.CardDecks[data.turn.nextActor];
        if (!data.turn.MandatoryPartFulfilled && DeckLib.isInitialHandDrawn(deck)) {
            data.turn.MandatoryPartFulfilled = true;
        }

        if (data.turn.MandatoryPartFulfilled) {
            ActionLib.setAllowedActions(data);
        } else {
            Constant.GeneralTypesOfActions[] memory allowed = new Constant.GeneralTypesOfActions[](1);
            allowed[0] = Constant.GeneralTypesOfActions.RngCommit;
            TurnLib.setOneAllowed(data.turn, allowed);
        }
    }

    function setAllowedActions(DataLib.DominionAppData memory data) internal pure {
        TurnLib.Turn memory t = data.turn;
        DeckLib.Deck memory currentDeck = data.CardDecks[t.nextActor];

        bool[] memory currentlyAllowed = t.possibleActions;

        for (uint i = 0; i < currentlyAllowed.length; i++) {
            currentlyAllowed[i] = false;
        }

        if (DeckLib.IsPlayActionPossible(currentDeck)) {
            currentlyAllowed[uint(Constant.GeneralTypesOfActions.PlayCard)] = true;
        }
        if (DeckLib.IsDrawActionPossible(currentDeck)) {
            currentlyAllowed[uint(Constant.GeneralTypesOfActions.RngCommit)] = true;
        }
        if (DeckLib.IsBuyActionPossible(currentDeck)) {
            currentlyAllowed[uint(Constant.GeneralTypesOfActions.BuyCard)] = true;
        }
    }



    /*
    function PlayCard(DataLib.DominionAppData data) internal pure {
        data.turn.PerformedAction = Constant.GeneralTypesOfActions.PlayCard;

    }
    */

}