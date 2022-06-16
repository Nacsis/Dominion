pragma solidity ^0.7.0;

import "./Util/constant.sol";

library Card {
    struct Card {
        Constant.CardType CardType;
        uint8   Money;
        uint8   VictoryPoints;
        uint8   PlayCost;
        uint8   BuyCost;
    }
}