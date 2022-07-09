pragma solidity ^0.7.0;

import "./constant.sol";
import "../turn.sol";
import "../stock.sol";
import "../card.sol";
import "../card.sol";
import "../rng.sol";
import "../deck.sol";

library ReaderLib {
    struct Reader {
        byte[] data;
    }

    function ReadUInt8(Reader memory r) internal pure returns (uint8){
        byte buf = r.data[0];
        r.data = Convert.Slice2Array(r.data, 1, r.data.length);
        return uint8(buf);
    }

    function ReadUInt16(Reader memory r) internal pure returns (uint16){
        bytes memory buf = Convert.byteArrayToBytes(Convert.Slice2Array(r.data,0,2));
        r.data = Convert.Slice2Array(r.data, 2, r.data.length);
        return uint16(Convert.bytesToUint(buf));
    }

    function ReadX(Reader memory r, uint16 x) internal pure returns (byte[] memory){
        byte[] memory buf = Convert.Slice2Array(r.data,0, x);
        r.data = Convert.Slice2Array(r.data,x,r.data.length);
        return buf;
    }

    function ReadTurn(Reader memory r) internal pure returns (TurnLib.Turn memory) {
        TurnLib.Turn memory o = TurnLib.oof(ReadX(r, ReadUInt16(r)));
        return o;
    }

    function ReadStock(Reader memory r) internal pure returns (StockLib.Stock memory) {
        StockLib.Stock memory o = StockLib.oof(ReadX(r, ReadUInt16(r)));
        return o;
    }

    function ReadCardDeck(Reader memory r) internal pure returns (DeckLib.Deck memory) {
        DeckLib.Deck memory o = DeckLib.oof(ReadX(r, ReadUInt16(r)));
        return o;
    }

    function ReadRng(Reader memory r) internal pure returns (RNGLib.RNG memory){
        RNGLib.RNG memory o = RNGLib.oof(ReadX(r, ReadUInt16(r)));
        return o;
    }

}