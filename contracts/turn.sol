pragma solidity ^0.7.0;

import "./Util/constant.sol";
import "./Util/convert.sol";
import "./data.sol";


library TurnLib {
    struct Turn {
        uint8 nextActor;
        Constant.GeneralTypesOfActions performedAction;
        bool MandatoryPartFulfilled;
        bool[] possibleActions;
    }

    function oof(byte[] memory data) internal pure returns (Turn memory){
        bool[] memory possibleActions = new bool[](uint8(Constant.GeneralTypesOfActions.GameEnd));
        for (uint8 i = 0; i < data.length - 3; i++) {
            possibleActions[i] = Convert.ByteToBool(data[3 + i]);
        }

        return Turn(uint8(data[0]),
            Constant.GeneralTypesOfActions(uint8(data[1])),
            Convert.ByteToBool(Convert.Slice2Array(data, 2, 3)[0]),
            possibleActions
        );
    }

    function equalTurn(Turn memory a, Turn memory b) internal pure {
        require(a.nextActor == b.nextActor, "Turn.nextActor difference");
        require(a.performedAction == b.performedAction, "Turn.performedAction difference");
        require(a.MandatoryPartFulfilled == b.MandatoryPartFulfilled, "Turn.MandatoryPartFulfilled difference");
        require(a.possibleActions.length == a.possibleActions.length, "Turn.possibleActions length difference");
        for (uint i = 0; i < a.possibleActions.length; i++) {
            require(a.possibleActions[i] == b.possibleActions[i], "Turn.possibleActions difference");
        }
    }

    function setOneAllowed(Turn memory turn, Constant.GeneralTypesOfActions[] memory newPossibleActions) internal pure {
        for (uint i = 0; i < turn.possibleActions.length; i++) {
            turn.possibleActions[i] = false;
        }

        for (uint i = 0; i < newPossibleActions.length; i++) {
            turn.possibleActions[i] = true;
        }
    }


    function setNextActor(Turn memory turn) internal pure {
        turn.nextActor = (turn.nextActor + 1) % Constant.NumPlayers;
    }
}