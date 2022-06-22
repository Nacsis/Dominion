pragma solidity ^0.7.0;

import "./constant.sol";

library Convert {
    function ByteToBool(byte b) internal pure returns (bool){
        return uint8(b) == 1;
    }

    function bytesToByteArray(bytes memory b) internal pure returns (byte[] memory){
        byte[] memory array = new byte[](b.length);
        for (uint32 i = 0; i < b.length; i++) {
            array[i] = b[i];
        }
        return array;
    }

    function byteArrayToBytes(byte[] memory array) internal pure returns (bytes memory){
        bytes memory b = new bytes(array.length);
        for (uint32 i = 0; i < array.length; i++) {
            b[i] = array[i];
        }
        return b;
    }

    function bytes32ToByteArray(bytes32 b) internal pure returns (byte[] memory){
        byte[] memory array = new byte[](b.length);
        for (uint32 i = 0; i < b.length; i++) {
            array[i] = b[i];
        }
        return array;
    }

    function Slice2Array(bytes1[] memory b, uint256 start, uint256 end) internal pure returns (byte[] memory){
        byte[] memory array = new byte[](end - start);

        for (uint8 i = 0; i < array.length; i++) {
            array[i] = b[start + i];
        }
        return array;
    }

    function bytesToUint(bytes memory b) public pure returns (uint256 value){
        assembly {
            value := mload(add(b, 0x20)) // Magic
        }
        return value;
    }
}