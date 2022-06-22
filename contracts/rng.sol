pragma solidity ^0.7.0;

import "./Util/constant.sol";
import "./Util/convert.sol";

library RNGLib {
    struct RNG {
        byte[] ImageA;
        byte[] PreImageB;
        byte[] PreImageA;
    }

    function oof(byte[]  memory data) internal pure returns (RNG memory){
        uint8 size = uint8(data.length);
        byte[] memory imageA;
        byte[] memory preImageB;
        byte[] memory preImageA;

        if (size >= Constant.HashSize) {
            imageA = Convert.Slice2Array(data, 0, Constant.HashSize);
        }

        if (size >= 2 * Constant.HashSize) {
            preImageB = Convert.Slice2Array(data, Constant.HashSize, 2 * Constant.HashSize);
        }

        if (size >= 3 * Constant.HashSize) {
            preImageA = Convert.Slice2Array(data, 2 * Constant.HashSize, 3 * Constant.HashSize);
        }

        RNG  memory rng = RNG(imageA, preImageB, preImageA);
        return rng;
    }

    function equalRNG(RNG memory a, RNG memory b) internal pure {
        //No check for preimageA as it can be secret here and will be replaced with dummy

        require(a.ImageA.length == b.ImageA.length);
        for (uint i = 0; i<a.ImageA.length; i++){
            require(a.ImageA[i]==b.ImageA[i]);
        }
        require(a.PreImageB.length == b.PreImageB.length);
        for (uint i = 0; i<a.PreImageA.length; i++){
            require(a.PreImageA[i]==b.PreImageA[i]);
        }
    }

    function RNGValue(RNG memory rng) internal pure returns(byte[] memory) {
        byte[] memory value = new byte[](Constant.HashSize);

        for (uint i = 0; i < rng.PreImageA.length; i++){
            value[i] = rng.PreImageA[i] ^ rng.PreImageB[i];
        }

        return value;
    }
}