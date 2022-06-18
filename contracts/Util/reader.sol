pragma solidity ^0.7.0;

import "./Util/constant.sol";
import "../turn.sol";

library Reader {
    struct Reader{
        byte[] data;
    }

    function ReadUInt8(Reader r) internal pure returns (uint8){
        return r.data[0];
    }

    function ReadX(Reader r, uint8 x) internal returns (byte[]){
        byte[x] buf = r.data[:x];
        r.data = r.data[x:];
        return buf;
    }

    function ReadTurn(Reader r,  Turn.turn o) internal{
        o = Turn.of(ReadX(r,ReadUint8(r)));
    }

    function ReadStock(Reader r,  Turn.turn o) internal{
        o = Stock.of(ReadX(r,ReadUint8(r)));
    }

    function ReadCardDeck(Reader r,  Turn.turn o) internal{
        o = Turn.of(ReadX(r,ReadUint8(r)));
    }

    function ReadRng(Reader r,  Turn.turn o) internal{
        o = Rng.of(ReadX(r,ReadUint8(r)));
    }

}