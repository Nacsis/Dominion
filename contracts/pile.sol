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
            cards[i] = CardLib.oof(Convert.Slice2Array(data,i,i+1));
        }

        Pile  memory pile = Pile(cards);
        return pile;
    }

}