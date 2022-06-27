pragma solidity ^0.7.0;

import "./Util/constant.sol";
import "./Util/convert.sol";

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
            Convert.ByteToBool(Convert.Slice2Array(data,2,3)[0]),
            possibleActions
        );
    }
}