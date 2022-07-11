pragma solidity ^0.7.0;

library Constant {
    uint8 constant NumPlayers = 2;

    //------------------------ Deck resources ------------------------

    uint8 constant InitialMoneyCards = 5;
    uint8 constant InitialVictoryCards = 3;
    uint8 constant InitialDeckSize = InitialMoneyCards + InitialVictoryCards;

    uint8 constant InitialDrawResources = 5;
    uint8 constant InitialPlayResources = 1;
    uint8 constant InitialBuyResources = 1;
    uint8 constant InitialMoneyResources = 0;
    uint8 constant DeckResourcesCount = 4;

    //------------------------ Crypto / Framework ------------------------

    uint8 constant HashSize = 32;
    uint16 constant PreImageSize = 128;

    //------------------------ Card ------------------------

    uint8 constant CardTypeCount = 6;
    uint8 constant CopperInitialStock = 60;
    uint8 constant SilverInitialStock = 40;
    uint8 constant GoldInitialStock = 30;
    uint8 constant VictorySmallInitialStock = 24;
    uint8 constant VictoryMidInitialStock = 12;
    uint8 constant VictoryBigInitialStock = 12;

    uint8 constant CopperMoneyValue = 1;
    uint8 constant SilverMoneyValue = 2;
    uint8 constant GoldMoneyValue = 3;

    uint8 constant    CopperCost = 0;
    uint8 constant    SilverCost = 3;
    uint8 constant    GoldCost = 6;

    uint8 constant    MoneyCardPlayCost = 0;

    uint8 constant ActionCardPlayCost = 1;
    uint8 constant ActionCardVictoryPoint = 0;

    uint8 constant    VictorySmallVictoryValue = 1;
    uint8 constant    VictoryMidVictoryValue = 3;
    uint8 constant    VictoryBigVictoryValue = 6;

    uint8 constant    VictorySmallCost = 2;
    uint8 constant    VictoryMidCost = 5;
    uint8 constant    VictoryBigCost = 8;

    uint8 constant    VictoryCardPlayCost = 0;

    enum CardType {Copper, Silver, Gold, VictorySmall, VictoryMid, VictoryBig, Cellar, Market, Merchant, Mine, Remodel, Smithy, Feast, Chapel, Workshop, Village}
    enum GeneralTypesOfActions {GameInit, RngCommit, RngTouch, RngRelease, DrawCard, PlayCard, BuyCard, EndTurn, GameEnd}
    enum DeckResources {PlayableCards, DrawableCards, PurchasableCards, SpendableMoney}

}