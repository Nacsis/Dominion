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
            imageA = Convert.Slice2Array(data,0, Constant.HashSize);
        }

        if (size >= 2 * Constant.HashSize) {
            preImageB = Convert.Slice2Array(data,Constant.HashSize, 2 * Constant.HashSize);
        }

        if (size >= 3 * Constant.HashSize) {
            preImageA = Convert.Slice2Array(data, 2 * Constant.HashSize, 3 * Constant.HashSize);
        }

        RNG  memory rng = RNG(imageA, preImageB, preImageA);
        return rng;
    }
}