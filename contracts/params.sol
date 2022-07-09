pragma solidity ^0.7.0;

import "./Util/constant.sol";
import "./Util/convert.sol";
import "./data.sol";


library ParamsLib {
    struct Params {
        Constant.CardType MainTarget;
        Constant.CardType SecondLvlTarget;
        uint8[] SecondLvlIndices;
    }

    function oof(byte[] memory data) internal pure returns (Params memory){
        uint8 length = uint8(data[0]);
        Constant.CardType main = Constant.CardType(uint8(data[1]));
        Constant.CardType second = Constant.CardType(uint8(data[2]));
        uint8[] memory indices = new uint8[](length-3);
        for (uint8 i = 0; i < length - 3; i++) {
            indices[i] = uint8(data[3 + i]);
        }
        return Params(main, second, indices);
    }

    function equal(Params memory a, Params memory b) internal pure{
        require(a.MainTarget==b.MainTarget, "Turn.Params.MainTarget not equal");
        require(a.SecondLvlTarget==b.SecondLvlTarget, "Turn.Params.SecondlevelTarget not equal");
        require(a.SecondLvlIndices.length == b.SecondLvlIndices.length,"Turn.Params.SecondlevelIndices length not equal");
        for (uint i = 0; i < a.SecondLvlIndices.length; i++) {
            require(a.SecondLvlIndices[i] == b.SecondLvlIndices[i], "Turn.Params.SecondlevelIndices not equal");
        }
    }
}