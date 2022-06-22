pragma solidity ^0.7.0;

import "./Util/constant.sol";
import "./turn.sol";

library StockLib {
    struct Stock {
        uint8[] CardAmounts;
    }

    function oof(byte[] memory data) internal pure returns (Stock memory){

        Stock memory stock = Stock(new uint8[](Constant.CardTypeCount));

        for (uint8 i = 0; i < Constant.CardTypeCount; i++) {
            stock.CardAmounts[i] = uint8(data[i]);
        }

        return stock;
    }

    function equalStock(Stock memory a, Stock memory b) internal pure{
        require(a.CardAmounts.length == b.CardAmounts.length,"Turn.stock length changed");

        for (uint i = 0; i < a.CardAmounts.length; i++) {
            require(a.CardAmounts[i] == b.CardAmounts[i], "Turn.stock difference");
        }
    }
}