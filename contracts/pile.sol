pragma solidity ^0.7.0;

import "./card.sol";
import "./Util/convert.sol";

library PileLib {
    struct Pile {
        CardLib.Card[] Card;
    }

    function oof(byte[]  memory data) internal pure returns (Pile memory){
        CardLib.Card[]  memory cards = new CardLib.Card[](data.length);
        for (uint8 i = 0; i < data.length; i++) {
            cards[i] = CardLib.oof(Convert.Slice2Array(data, i, i + 1));
        }

        Pile  memory pile = Pile(cards);
        return pile;
    }

    function equalPile(Pile memory a, Pile memory b) internal pure {
        require(a.Card.length != b.Card.length);
        for (uint cardIndex = 0; cardIndex < a.Card.length; cardIndex++) {
            CardLib.equalCard(a.Card[cardIndex], b.Card[cardIndex]);
        }
    }

    function MinimalPlayCost(Pile memory p) internal pure returns (uint8) {
        uint8 minValue = p.Card[0].PlayCost;
        for (uint i = 0; i < p.Card.length; i++) {
            minValue = minValue < p.Card[i].PlayCost ? minValue : p.Card[i].PlayCost;
        }
        return minValue;
    }

    function DrawCardBasedOnSeed(PileLib.Pile memory p, byte[] memory seed) internal pure returns (CardLib.Card memory) {
        uint index = uint64(Convert.bytesToUint(Convert.byteArrayToBytes(seed))) % uint64(p.Card.length);
        CardLib.Card memory card = p.Card[index];
        ResizeCardsWithoutIndex(p, index);
        return card;
    }

    function ResizeCardsWithoutIndex(Pile memory p, uint index) internal pure {
        CardLib.Card[] memory newCards = new CardLib.Card[](p.Card.length - 1);

        for (uint i = 0;
            i < p.Card.length;
            i++) {
            uint j = 0;
            if (i != uint(index)) {
                newCards[j] = p.Card[i];
            }
            j++;
        }
        p.Card = newCards;
    }

    // AddCardToPile add card to current Cards
    function AddCardToPile(Pile memory p, CardLib.Card memory card) internal pure {
        CardLib.Card[] memory newCards = new CardLib.Card[](p.Card.length + 1);
        for (uint i; i< p.Card.length; i++){
            newCards[i] = p.Card[i];
        }
        newCards[p.Card.length] = card;
        p.Card = newCards;
    }
}