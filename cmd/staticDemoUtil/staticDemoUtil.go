// Copyright 2022 PolyCrypt GmbH
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

package staticDemoUtil

import (
	"context"
	"log"
	"math/big"

	"perun.network/perun-examples/dominion-cli/app/util"
	"perun.network/perun-examples/dominion-cli/contracts/generated/dominionApp"
	"perun.network/perun-examples/dominion-cli/global"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethwallet "perun.network/go-perun/backend/ethereum/wallet"

	"perun.network/perun-examples/dominion-cli/app"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	ethchannel "perun.network/go-perun/backend/ethereum/channel"
	swallet "perun.network/go-perun/backend/ethereum/wallet/simple"
	"perun.network/go-perun/channel"
	"perun.network/go-perun/wire"
	"perun.network/perun-examples/dominion-cli/client"
)

// deployContracts deploys the contracts on the specified ledger.
func DeployContracts(nodeURL string, chainID uint64, privateKey string) (adj, ah, app common.Address) {
	k, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		panic(err)
	}
	w := swallet.NewWallet(k)
	cb, err := client.CreateContractBackend(nodeURL, chainID, w)
	if err != nil {
		panic(err)
	}
	acc := accounts.Account{Address: crypto.PubkeyToAddress(k.PublicKey)}

	// Deploy adjudicator.
	adj, err = ethchannel.DeployAdjudicator(context.TODO(), cb, acc)
	if err != nil {
		panic(err)
	}

	// Deploy asset holder.
	ah, err = ethchannel.DeployETHAssetholder(context.TODO(), cb, adj, acc)
	if err != nil {
		panic(err)
	}

	const gasLimit = 1100000 // Must be sufficient for deploying TicTacToe.sol.
	tops, err := cb.NewTransactor(context.TODO(), gasLimit, acc)
	if err != nil {
		panic(err)
	}
	// Deploy dominion App.
	appAddress, tx, _, err := dominionApp.DeployDominionApp(tops, cb)
	if err != nil {
		panic(err)
	}

	_, err = bind.WaitDeployed(context.TODO(), cb, tx)
	if err != nil {
		panic(err)
	}

	return adj, ah, appAddress
}

// setupGameClient sets up a new client with the given parameters.
func SetupGameClient(
	bus wire.Bus,
	nodeURL string,
	chainID uint64,
	adjudicator common.Address,
	asset ethwallet.Address,
	privateKey string,
	app *app.DominionApp,
	stake channel.Bal,
) *client.AppClient {
	// Create wallet and account.
	k, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		panic(err)
	}
	w := swallet.NewWallet(k)
	acc := crypto.PubkeyToAddress(k.PublicKey)

	// Create and start client.
	c, err := client.SetupAppClient(
		bus,
		w,
		acc,
		nodeURL,
		chainID,
		adjudicator,
		asset,
		app,
		stake,
	)
	if err != nil {
		panic(err)
	}

	return c
}

// balanceLogger is a utility for logging client balances.
type balanceLogger struct {
	ethClient *ethclient.Client
}

// newBalanceLogger creates a new balance errorHandling.go for the specified ledger.
func NewBalanceLogger(chainURL string) balanceLogger {
	c, err := ethclient.Dial(chainURL)
	if err != nil {
		panic(err)
	}
	return balanceLogger{ethClient: c}
}

// LogBalances prints the balances of the specified clients.
func (l balanceLogger) LogBalances(clients ...*client.AppClient) {
	bals := make([]*big.Float, len(clients))
	for i, c := range clients {
		bal, err := l.ethClient.BalanceAt(context.TODO(), c.WalletAddress(), nil)
		if err != nil {
			log.Fatal(err)
		}
		bals[i] = client.WeiToEth(bal)
	}
	log.Println("Client balances (ETH):", bals)
}

func DrawInitHand(drawer, other *client.DominionChannel) {
	for i := 0; i < 5; i++ {
		var alicePreimage = global.RandomBytes(util.HashSize)
		drawer.RngCommit(alicePreimage)
		other.RngTouch()
		drawer.RngRelease(alicePreimage)
		drawer.DrawOneCard()
	}
}
