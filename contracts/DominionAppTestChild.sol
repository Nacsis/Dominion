pragma solidity ^0.7.0;
pragma experimental ABIEncoderV2;

import "./perun-eth-contracts/contracts/App.sol";
import "./data.sol";
import "./Util/constant.sol";
import "./Util/reader.sol";
import "./DominionApp.sol";

contract DominionAppTestChild is DominionApp {
    function decodeDataExt(
        ReaderLib.Reader calldata r,
        string calldata poi)
    external pure returns (bytes memory){
        DataLib.DominionAppData memory appData = decodeData(r);

        if (keccak256(bytes(poi)) == keccak256(bytes("turn.nextActor"))) {
            return new bytes(appData.turn.nextActor);
        }

        bytes memory none;
        return none;
    }
}