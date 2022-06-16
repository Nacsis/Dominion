pragma solidity ^0.7.0;

import "./Util/constant.sol";
import "./turn.sol";
import "./stock.sol";
import "./deck.sol";
import "./rng.sol";

library Data {
    struct DominionAppData {
        Turn                    turn;
        Stock                   stock;
        Deck[util.NumPlayers]   CardDecks;
        RNG                     rng;
    }
}