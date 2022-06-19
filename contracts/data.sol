pragma solidity ^0.7.0;

import "./Util/constant.sol";
import "./turn.sol";
import "./stock.sol";
import "./deck.sol";
import "./rng.sol";

library DataLib {
    struct DominionAppData {
        TurnLib.Turn turn;
        StockLib.Stock stock;
        DeckLib.Deck[] CardDecks;
        RNGLib.RNG rng;
    }
}