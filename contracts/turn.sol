pragma solidity ^0.7.0;

import "./Util/constant.sol";
import "./Util/convert.sol";
import "./data.sol";
import "./params.sol";
import "./Util/reader.sol";


library TurnLib {
    struct Turn {
        uint8 nextActor;
        Constant.GeneralTypesOfActions performedAction;
        bool MandatoryPartFulfilled;
        bool[] possibleActions;
        ParamsLib.Params params;
    }

    function oof(byte[] memory data) internal pure returns (Turn memory){
        bool[] memory possibleActions = new bool[](uint8(Constant.GeneralTypesOfActions.GameEnd));
        uint8 lengthAction = uint8(data[3]);
        for (uint8 i = 0; i < lengthAction; i++) {
            possibleActions[i] = Convert.ByteToBool(data[4 + i]);
        }

        uint8 lengthParams = uint8(data[3+lengthAction]);
        ParamsLib.Params memory params = ParamsLib.oof(Convert.Slice2Array(data, 3+lengthAction+1,data.length));


        return Turn(uint8(data[0]),
            Constant.GeneralTypesOfActions(uint8(data[1])),
            Convert.ByteToBool(Convert.Slice2Array(data, 2, 3)[0]),
            possibleActions,
            params
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
        ParamsLib.equal(a.params, b.params);
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