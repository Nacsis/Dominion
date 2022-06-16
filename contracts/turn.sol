pragma solidity ^0.7.0;

import "./Util/constant.sol";

library Turn {
    struct Turn {
        uint8                   nextActor;
        Constant.GeneralTypesOfActions   performedAction;
        bool                    MandatoryPartFulfilled;
        bool[util.GameEnd]      possibleActions;
}

}