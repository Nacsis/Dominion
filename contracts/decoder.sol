pragma solidity ^0.7.0;

import "./perun-eth-contracts/contracts/App.sol";
import "./data.sol";
import "./Util/constant.sol";
import "./Util/reader.sol";
import "./dataValidity.sol";

library DecoderLib {
    function decodeData(
        ReaderLib.Reader memory r)
    internal pure returns (DataLib.DominionAppData  memory)
    {
        TurnLib.Turn memory turn = ReaderLib.ReadTurn(r);
        StockLib.Stock memory stock = ReaderLib.ReadStock(r);

        DeckLib.Deck[] memory decks = new DeckLib.Deck[](Constant.NumPlayers);
        for (uint deckIndex = 0; deckIndex < Constant.NumPlayers; deckIndex++) {
            ReaderLib.ReadCardDeck(r);
        }
        RNGLib.RNG memory rng = ReaderLib.ReadRng(r);

        DataLib.DominionAppData memory appData = DataLib.DominionAppData(turn, stock, decks, rng);
        return appData;
    }
}