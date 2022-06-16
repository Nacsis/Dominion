pragma solidity ^0.7.0;

import "./Util/constant.sol";

library RNG {
    struct RNG {
        byte[] ImageA;
        byte[] PreImageB;
        byte[] PreImageA;
    }
}