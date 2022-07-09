pragma solidity ^0.7.0;

import "./Util/constant.sol";
import "./turn.sol";

library StockLib {
    struct Stock {
        uint8[] CardAmounts;
        uint8[] Trash;
    }

    function oof(byte[] memory data) internal pure returns (Stock memory){

        uint8[] memory CardAmounts = new uint8[](Constant.CardTypeCount);
        uint8[] memory Trash = new uint8[](Constant.CardTypeCount);

        for (uint8 i = 0; i < Constant.CardTypeCount; i++) {
            CardAmounts[i] = uint8(data[i]);
        }

        for (uint i = 0; i < Constant.CardTypeCount; i++){
            Trash[i] = uint8(data[Constant.CardTypeCount+i]);
        }

        return Stock(CardAmounts, Trash);
    }

    function equalStock(Stock memory a, Stock memory b) internal pure{
        require(a.CardAmounts.length == b.CardAmounts.length,"Turn.stock.cardamounts length changed");

        for (uint i = 0; i < a.CardAmounts.length; i++) {
            require(a.CardAmounts[i] == b.CardAmounts[i], "Turn.stock.cardamounts difference");
        }

        require(a.Trash.length == b.Trash.length,"Turn.stock.trash length changed");

        for (uint i = 0; i < a.CardAmounts.length; i++) {
            require(a.Trash[i] == b.Trash[i], "Turn.stock.trash difference");
        }
    }
}