pragma solidity ^0.7.0;
pragma abicoder v2;

import "truffle/Assert.sol";
import "./DominionApp.sol";
import "./DominionAppTestChild.sol";

contract TestDecoder is DominionApp {
    function testDecode() external {
        byte[] memory enc = Convert.bytesToByteArray(bytes(hex"c0001e9cc0"));
        DominionAppTestChild app = new DominionAppTestChild();
        ReaderLib.Reader memory reader = ReaderLib.Reader(enc);
        bytes memory data = app.decodeDataExt(reader, "turn.nextActor");
        bytes memory expected = new bytes(1);
        Assert.assert(data, expected, "Failed to read turn.nextActor");
    }
}