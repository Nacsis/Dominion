pragma solidity ^0.7.0;

import "./Util/constant.sol";
import "./pile.sol";


library Deck {
    struct Deck {
        Pile MainPile; // main deck
        Pile HandPile;// Hand cards
        Pile DiscardedPile; // deck of discarded cards
        Pile PlayedPile;// deck of played cards
        uint8[util.DeckResourcesCount] Resources;
    }
}