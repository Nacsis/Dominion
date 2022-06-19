pragma solidity ^0.7.0;

import "./Util/constant.sol";

library CardLib {
    struct Card {
        Constant.CardType CardType;
        uint8 Money;
        uint8 VictoryPoints;
        uint8 PlayCost;
        uint8 BuyCost;
    }

    function oof(byte[]  memory data) internal pure returns (Card memory){
        Constant.CardType cardType;
        uint8 cardTypeId = uint8(data[0]);
        uint8 money;
        uint8 playCost;
        uint8 buyCost;
        uint8 victoryPoints;

        if (uint8(cardTypeId) == uint8(Constant.CardType.Copper)) {
            cardType = Constant.CardType.Copper;
            money = Constant.CopperMoneyValue;
            playCost = Constant.MoneyCardPlayCost;
            buyCost = Constant.CopperCost;
        } else if (uint8(cardType) == uint8(Constant.CardType.Silver)) {
            cardType = Constant.CardType.Silver;
            money = Constant.SilverMoneyValue;
            playCost = Constant.MoneyCardPlayCost;
            buyCost = Constant.SilverCost;
        } else if (uint8(cardType) == uint8(Constant.CardType.Gold)) {
            cardType = Constant.CardType.Gold;
            money = Constant.GoldMoneyValue;
            playCost = Constant.MoneyCardPlayCost;
            buyCost = Constant.GoldCost;
        } else if (uint8(cardType) == uint8(Constant.CardType.VictorySmall)) {
            cardType = Constant.CardType.VictorySmall;
            victoryPoints = Constant.VictorySmallVictoryValue;
            playCost = Constant.VictoryCardPlayCost;
            buyCost = Constant.VictorySmallCost;
        } else if (uint8(cardType) == uint8(Constant.CardType.VictoryMid)) {
            cardType = Constant.CardType.VictoryMid;
            victoryPoints = Constant.VictoryMidVictoryValue;
            playCost = Constant.VictoryCardPlayCost;
            buyCost = Constant.VictoryMidCost;
        } else if (uint8(cardType) == uint8(Constant.CardType.VictoryBig)) {
            cardType = Constant.CardType.VictoryBig;
            victoryPoints = Constant.VictoryBigVictoryValue;
            playCost = Constant.VictoryCardPlayCost;
            buyCost = Constant.VictoryBigCost;
        }

        Card memory c = Card(cardType,money, victoryPoints, playCost, buyCost);
        return c;
    }

}