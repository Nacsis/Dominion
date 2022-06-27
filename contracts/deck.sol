pragma solidity ^0.7.0;

import "./Util/constant.sol";
import "./pile.sol";
import "./Util/convert.sol";


library DeckLib {
    struct Deck {
        PileLib.Pile MainPile; // main deck
        PileLib.Pile HandPile;// Hand cards
        PileLib.Pile DiscardedPile; // deck of discarded cards
        PileLib.Pile PlayedPile;// deck of played cards
        uint8[] Resources;
    }

    function oof(byte[]  memory data) internal pure returns (Deck memory){
        uint8 mainCardPileSize = uint8(data[0]);
        PileLib.Pile memory mainPile = PileLib.oof(Convert.Slice2Array(data,1, mainCardPileSize + 1));
        data = Convert.Slice2Array(data,mainCardPileSize + 1, data.length);

        uint8 handCardSize = uint8(data[0]);
        PileLib.Pile memory handPile = PileLib.oof(Convert.Slice2Array(data,1, handCardSize + 1));
        data = Convert.Slice2Array(data,handCardSize + 1, data.length);

        uint8 discardPileSize = uint8(data[0]);
        PileLib.Pile memory discardPile = PileLib.oof(Convert.Slice2Array(data,1, discardPileSize + 1));
        data = Convert.Slice2Array(data,discardPileSize + 1, data.length);

        uint8 playedPileSize = uint8(data[0]);
        PileLib.Pile memory playedPile = PileLib.oof(Convert.Slice2Array(data,1, playedPileSize + 1));
        data = Convert.Slice2Array(data,discardPileSize + 1, data.length);

        uint8[] memory resources = new uint8[](Constant.DeckResourcesCount);
        uint8 actionValueSize = uint8(data[0]);
        for (uint8 i = 0; i < actionValueSize; i++) {
            resources[i] = uint8(data[i + 1]);
        }

        return Deck(mainPile, handPile, discardPile, playedPile, resources);
    }
}