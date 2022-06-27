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
import "./Util/reader.sol";

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


        ReaderLib.Reader memory fromReader = ReaderLib.Reader(Convert.bytesToByteArray(from.appData));
        DataLib.DominionAppData memory fromAppData = decodeData(fromReader);
        ReaderLib.Reader memory fromReaderClone = ReaderLib.Reader(Convert.bytesToByteArray(from.appData));
        DataLib.DominionAppData memory fromAppDataClone = decodeData(fromReaderClone);
        ReaderLib.Reader memory toReader = ReaderLib.Reader(Convert.bytesToByteArray(to.appData));
        DataLib.DominionAppData memory toAppData = decodeData(toReader);

        if (toAppData.turn.performedAction == Constant.GeneralTypesOfActions.RngCommit) {
            //...
        }
    }

    function decodeData(
        ReaderLib.Reader memory r)
    internal pure returns (DataLib.DominionAppData  memory){
        TurnLib.Turn memory turn = ReaderLib.ReadTurn(r);
        StockLib.Stock memory stock = ReaderLib.ReadStock(r);

        DeckLib.Deck[] memory decks = new DeckLib.Deck[](Constant.NumPlayers);
        for (uint deckIndex = 0; deckIndex < Constant.NumPlayers; deckIndex++) {
            ReaderLib.ReadCardDeck(r);
        }
        RNGLib.RNG memory rng = ReaderLib.ReadRng(r);

        DataLib.DominionAppData memory appData = DataLib.DominionAppData(turn,stock,decks,rng);
        return appData;

    }
}