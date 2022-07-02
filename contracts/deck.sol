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
        PileLib.Pile memory mainPile = PileLib.oof(Convert.Slice2Array(data, 1, mainCardPileSize + 1));
        data = Convert.Slice2Array(data, mainCardPileSize + 1, data.length);

        uint8 handCardSize = uint8(data[0]);
        PileLib.Pile memory handPile = PileLib.oof(Convert.Slice2Array(data, 1, handCardSize + 1));
        data = Convert.Slice2Array(data, handCardSize + 1, data.length);

        uint8 discardPileSize = uint8(data[0]);
        PileLib.Pile memory discardPile = PileLib.oof(Convert.Slice2Array(data, 1, discardPileSize + 1));
        data = Convert.Slice2Array(data, discardPileSize + 1, data.length);

        uint8 playedPileSize = uint8(data[0]);
        PileLib.Pile memory playedPile = PileLib.oof(Convert.Slice2Array(data, 1, playedPileSize + 1));
        data = Convert.Slice2Array(data, discardPileSize + 1, data.length);

        uint8[] memory resources = new uint8[](Constant.DeckResourcesCount);
        uint8 actionValueSize = uint8(data[0]);
        for (uint8 i = 0; i < actionValueSize; i++) {
            resources[i] = uint8(data[i + 1]);
        }

        return Deck(mainPile, handPile, discardPile, playedPile, resources);
    }

    function victoryPoints(Deck memory deck) internal pure returns (uint8){
        uint8 vp = 0;
        PileLib.Pile[4] memory pileArray = [deck.MainPile, deck.DiscardedPile, deck.HandPile, deck.PlayedPile];

        for (uint pileIndex = 0; pileIndex < pileArray.length; pileIndex++) {
            for (uint cardIndex = 0; cardIndex < pileArray[pileIndex].Card.length; cardIndex++) {
                vp += pileArray[pileIndex].Card[cardIndex].VictoryPoints;
            }
        }

        return vp;
    }

    function equalDeck(Deck memory a, Deck memory b) internal pure {
        PileLib.Pile[4] memory pileArrayA = [a.MainPile, a.DiscardedPile, a.HandPile, a.PlayedPile];
        PileLib.Pile[4] memory pileArrayB = [b.MainPile, b.DiscardedPile, b.HandPile, b.PlayedPile];

        for (uint pileIndex = 0; pileIndex < pileArrayA.length; pileIndex++) {
            PileLib.equalPile(pileArrayA[pileIndex], pileArrayB[pileIndex]);
        }
    }

    function isInitialHandDrawn(Deck memory deck) internal pure returns (bool){
        return deck.Resources[uint(Constant.DeckResources.DrawableCards)] == 0 && deck.PlayedPile.Card.length == 0 && deck.HandPile.Card.length == Constant.InitialDrawResources;
    }

    // IsPlayActionPossible true if another play action is possible
    function IsPlayActionPossible(DeckLib.Deck memory d) internal pure returns (bool) {
        return d.Resources[uint(Constant.DeckResources.PlayableCards)] >= PileLib.MinimalPlayCost(d.HandPile);
    }

    // IsDrawActionPossible true if another draw action is possible
    function IsDrawActionPossible(DeckLib.Deck memory d) internal pure returns (bool) {
        return d.Resources[uint(Constant.DeckResources.DrawableCards)] > 0;
    }

    // IsBuyActionPossible true if another buy action is possible
    function IsBuyActionPossible(DeckLib.Deck memory d) internal pure returns (bool) {
        return d.Resources[uint(Constant.DeckResources.PurchasableCards)] > 0;
    }

    function DrawCard(DeckLib.Deck memory d, byte[] memory seed) internal pure {
        require(d.Resources[uint(Constant.DeckResources.DrawableCards)] > 0);
        if (d.MainPile.Card.length == 0) {
            MixAndReassignDiscardedPile(d);
        }
        CardLib.Card memory card = PileLib.DrawCardBasedOnSeed(d.MainPile, seed);
        PileLib.AddCardToPile(d.HandPile, card);

        d.Resources[uint(Constant.DeckResources.DrawableCards)]--;
    }

    function MixAndReassignDiscardedPile(DeckLib.Deck memory d) internal pure {
        PileLib.Pile memory discardPile = d.DiscardedPile;
        d.DiscardedPile.Card = new CardLib.Card[](0);
        d.MainPile = discardPile;
    }
}