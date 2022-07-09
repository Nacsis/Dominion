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
import "./dataValidity.sol";
import "./decoder.sol";

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
        ReaderLib.Reader memory fromReader = ReaderLib.Reader(Convert.bytesToByteArray(from.appData));
        DataLib.DominionAppData memory fromAppData = DecoderLib.decodeData(fromReader);
        ReaderLib.Reader memory fromReaderClone = ReaderLib.Reader(Convert.bytesToByteArray(from.appData));
        DataLib.DominionAppData memory fromAppDataClone = DecoderLib.decodeData(fromReaderClone);
        ReaderLib.Reader memory toReader = ReaderLib.Reader(Convert.bytesToByteArray(to.appData));
        DataLib.DominionAppData memory toAppData = DecoderLib.decodeData(toReader);

        //Valid Signer
        require(toAppData.turn.nextActor == signerIdx, "Signer is not nextActor");


        //Valid Participants
        require(params.participants.length == Constant.NumPlayers, "Number of participants");


        //Valid Action
        require(fromAppData.turn.possibleActions[uint256(toAppData.turn.performedAction)]);

        /*
        // Valid Data Change
        DataValidityLib.validDataChange(fromAppData, fromAppDataClone, toAppData);
        */

        /*
        // Test Final State
        require((toAppData.turn.performedAction == Constant.GeneralTypesOfActions.GameEnd) == to.isFinal, "Channel set isFinal, without GameEnd Action performed");
        Array.requireEqualAddressArray(to.outcome.assets, from.outcome.assets);
        Channel.requireEqualSubAllocArray(to.outcome.locked, from.outcome.locked);
        uint256[][] memory expectedBalances = from.outcome.balances;

        (bool hasWinner, uint8 winner) = getWinner(toAppData);

        if (hasWinner) {
            uint8 loser = 1 - winner;
            expectedBalances = new uint256[][](expectedBalances.length);
            for (uint i = 0; i < expectedBalances.length; i++) {
                expectedBalances[i] = new uint256[](Constant.NumPlayers);
                expectedBalances[i][winner] = from.outcome.balances[i][0] + from.outcome.balances[i][1];
                expectedBalances[i][loser] = 0;
            }
        }
        requireEqualUint256ArrayArray(to.outcome.balances, expectedBalances);
        */
    }

    function requireEqualUint256ArrayArray(
        uint256[][] memory a,
        uint256[][] memory b
    )
    internal pure
    {
        require(a.length == b.length, "uint256[][]: unequal length");
        for (uint i = 0; i < a.length; i++) {
            Array.requireEqualUint256Array(a[i], b[i]);
        }
    }

    function getWinner(
        DataLib.DominionAppData memory appData)
    internal pure returns (bool, uint8){
        bool hasWinner = false;
        uint8 winner = 0;

        if (DeckLib.victoryPoints(appData.CardDecks[0]) != DeckLib.victoryPoints(appData.CardDecks[1])) {
            hasWinner = true;
            if (DeckLib.victoryPoints(appData.CardDecks[0]) > DeckLib.victoryPoints(appData.CardDecks[1])) {
                winner = 0;
            } else {
                winner = 1;
            }
        }

        return (hasWinner, winner);
    }
}