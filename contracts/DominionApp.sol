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

/**
 * @notice TicTacToeApp is a channel app for playing tic tac toe.
 * The data is encoded as follows:
 * - data[0]: The index of the next actor.
 * - data[i], i in [1,10]: The value of field i. 0 means no tick, 1 means tick by player 1, 2 means tick by player 2.
 */
contract DominionApp is App {
    uint8 constant actorDataIndex = 0;
    uint8 constant actorDataLength = 1;
    uint8 constant appDataLength = 1; // Actor index + grid.
    uint8 constant numParts = 2;
    uint8 constant firstPlayer = 1;
    uint8 constant secondPlayer = 2;

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
        /*
        require(params.participants.length == numParts, "number of participants");

        uint8 actorIndex = uint8(from.appData[actorDataIndex]);
        require(to.appData.length == appDataLength, "data length");
        require(actorIndex == signerIdx, "actor not signer");
        require((actorIndex + 1) % numParts == uint8(to.appData[actorDataIndex]), "next actor");*/
    }
}