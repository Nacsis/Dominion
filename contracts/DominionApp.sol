// Copyright 2021 - See NOTICE file for copyright holders.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.7.0;
pragma experimental ABIEncoderV2;

import "./perun-eth-contracts/contracts/App.sol";
import "./data.sol";
import "./Util/constant.sol";

contract DominionApp is App {


    /**
     * @notice ValidTransition checks if there was a valid transition between two states.
     * @param params The parameters of the channel.
     * @param from The current state.
     * @param to The potential next state.
     * @param signerIdx Index of the participant who signed this transition.
     */
    function validTransition(
        Channel.Params calldata params,
        Channel.State calldata from,
        Channel.State calldata to,
        uint256 signerIdx)
    external pure override
    {

        //require(params.participants.length == numParts, "number of participants");


        Util.Reader fromReader = Util.Reader(from);
        Data.DominionAppData fromAppData = decodeData(fromReader);
        Util.Reader fromReaderClone = Util.Reader(from);
        Data.DominionAppData fromAppDataClone = decodeData(fromReaderClone);
        Util.Reader toReader = Util.Reader(to);
        Data.DominionAppData toAppData = decodeData(toReader);

        if (toAppData.turn.performedAction == util.RngCommit) {
            //...
        }
    }

    function decodeData(
        Util.Reader r)
    internal pure returns (Data.DominionAppData){
        Data.DominionAppData appData = Data.DominionAppData();
        util.ReadTurn(r, appData.turn);
        util.ReadStock(r, appData.stock);

        for (uint deckIndex = 0; deckIndex < util.NumPlayers; deckIndex++) {
            util.ReadCardDeck(r, appData.CardDecks[decodeData]);
        }
        util.ReadRng(r, appData.rng);

        return appData;

    }
}