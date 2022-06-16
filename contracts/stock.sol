pragma solidity ^0.7.0;

import "./Util/constant.sol";
import "./turn.sol";

library Stock {
    struct Stock {
        uint8[util.CardTypeCount] CardAmounts;
    }
}