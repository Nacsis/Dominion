pragma solidity ^0.7.0;

import "./Util/constant.sol";
import "./turn.sol";
import "./stock.sol";
import "./deck.sol";
import "./rng.sol";
import "./data.sol";
import "./Util/reader.sol";
import "./Util/constant.sol";
import "./decoder.sol";
import "./action.sol";

library DataValidityLib {

    function validDataChange(
        DataLib.DominionAppData memory from,
        DataLib.DominionAppData memory tmp,
        DataLib.DominionAppData memory to)
    internal pure {

        if (to.turn.performedAction == Constant.GeneralTypesOfActions.GameInit) {
            tmp = to;
            //byte[] memory initEnc = Convert.bytesToByteArray(bytes(hex"c000032d00"));
            //ReaderLib.Reader memory initReader = ReaderLib.Reader(initEnc);
            //tmp = DecoderLib.decodeData(initReader);
        }

        if (to.turn.performedAction == Constant.GeneralTypesOfActions.RngCommit) {
            require(to.rng.ImageA.length == Constant.HashSize);
            ActionLib.RngCommit(tmp);
            // overwriting with A's Image
            tmp.rng.ImageA = to.rng.ImageA;
        }

        if (to.turn.performedAction == Constant.GeneralTypesOfActions.RngTouch) {
            require(to.rng.PreImageB.length == Constant.HashSize);
            ActionLib.RngTouch(tmp);
            // overwriting with B's PreImage
            tmp.rng.PreImageB = to.rng.PreImageB;
        }

        if (to.turn.performedAction == Constant.GeneralTypesOfActions.RngRelease) {
            ActionLib.RngRelease(tmp, to.rng.PreImageA);
        }

        if (to.turn.performedAction == Constant.GeneralTypesOfActions.DrawCard) {
            ActionLib.DrawCard(tmp);
        }

        if (to.turn.performedAction == Constant.GeneralTypesOfActions.PlayCard) {
            //ActionLib.PlayCard(tmp);
            tmp = to;
        }

        if (to.turn.performedAction == Constant.GeneralTypesOfActions.BuyCard) {
            //ActionLib.BuyCard(tmp);
            tmp = to;
        }

        if (to.turn.performedAction == Constant.GeneralTypesOfActions.EndTurn) {
            //ActionLib.EndTurn(tmp);
            tmp = to;
        }

        if (to.turn.performedAction == Constant.GeneralTypesOfActions.GameEnd) {
            //ActionLib.EndGame(tmp);
            tmp = to;
        }

        DataLib.equalAppData(tmp, to);
    }
}