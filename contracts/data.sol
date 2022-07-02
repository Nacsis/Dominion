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

    function equalAppData(DominionAppData memory a, DominionAppData memory b) internal pure {
        TurnLib.equalTurn(a.turn, b.turn);
        StockLib.equalStock(a.stock, b.stock);
        RNGLib.equalRNG(a.rng, b.rng);
        require(a.CardDecks.length == b.CardDecks.length, "Turn.CardDecks length changed");

        for (uint i = 0; i < a.CardDecks.length; i++) {
            DeckLib.equalDeck(a.CardDecks[i], b.CardDecks[i]);
        }
    }
}