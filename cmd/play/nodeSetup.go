// Copyright (c) 2020 Chair of Applied Cryptography, Technische Universität
// Darmstadt, Germany. All rights reserved. This file is part of
// perun-eth-demo. Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package play

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/pkg/errors"

	echannel "perun.network/go-perun/backend/ethereum/channel"
	cherrors "perun.network/go-perun/backend/ethereum/channel/errors"
	ewallet "perun.network/go-perun/backend/ethereum/wallet"
	phd "perun.network/go-perun/backend/ethereum/wallet/hd"
	"perun.network/go-perun/channel"
	"perun.network/go-perun/channel/persistence/keyvalue"
	"perun.network/go-perun/client"
	"perun.network/go-perun/log"
	"perun.network/go-perun/pkg/sortedkv/leveldb"
	"perun.network/go-perun/watcher/local"
	wirenet "perun.network/go-perun/wire/net"
	"perun.network/go-perun/wire/net/simple"
	"perun.network/perun-examples/dominion-cli/app"
	"perun.network/perun-examples/dominion-cli/contracts/generated/dominionApp"
)

var (
	backend         *node
	ethereumBackend *ethclient.Client
)

// Setup initializes the node, can not be done in init() since it needs the
// configuration from viper.
func Setup() {
	SetConfig()

	var err error
	if ethereumBackend, err = ethclient.Dial(config.Chain.URL); err != nil {
		log.WithError(err).Fatalln("Could not connect to ethereum node.")
	}
	if backend, err = newNode(); err != nil {
		log.WithError(err).Fatalln("Could not initialize node.")
	}
}

func newNode() (*node, error) {
	wallet, acc, err := setupWallet(config.Mnemonic, config.AccountIndex)
	if err != nil {
		return nil, errors.WithMessage(err, "importing mnemonic")
	}
	dialer := simple.NewTCPDialer(config.Node.DialTimeout)
	signer := types.NewEIP155Signer(big.NewInt(config.Chain.ID))

	n := &node{
		log:     log.Get(),
		onChain: acc,
		wallet:  wallet,
		dialer:  dialer,
		cb:      echannel.NewContractBackend(ethereumBackend, phd.NewTransactor(wallet.Wallet(), signer), config.Chain.TxFinalityDepth),
		peers:   make(map[string]*peer),
	}
	return n, n.setup()
}

// setup does:
//  - Create a new offChain account.
//  - Create a client with the node's dialer, funder, adjudicator and wallet.
//  - Setup a TCP listener for incoming connections.
//  - Load or create the database and setting up persistence with it.
//  - Set the OnNewChannel, Proposal and Update handler.
//  - Print the configuration.
func (n *node) setup() error {
	if err := n.setupContracts(); err != nil {
		return errors.WithMessage(err, "setting up contracts")
	}

	var err error

	n.offChain, err = n.wallet.NewAccount()
	if err != nil {
		return errors.WithMessage(err, "creating account")
	}

	n.log.WithField("off-chain", n.offChain.Address()).Info("Generating account")

	n.bus = wirenet.NewBus(n.onChain, n.dialer)

	// Setup dispute watcher.
	watcher, err := local.NewWatcher(n.adjudicator)
	if err != nil {
		return errors.WithMessage(err, "creating watcher")
	}
	// Setup Perun client.
	if n.client, err = client.New(n.onChain.Address(), n.bus, n.funder, n.adjudicator, n.wallet, watcher); err != nil {
		return errors.WithMessage(err, "creating client")
	}

	host := config.Node.IP + ":" + strconv.Itoa(int(config.Node.Port))
	n.log.WithField("host", host).Trace("Listening for connections")
	listener, err := simple.NewTCPListener(host)
	if err != nil {
		return errors.WithMessage(err, "could not start tcp listener")
	}

	n.client.OnNewChannel(n.setupChannel)
	if err := n.setupPersistence(); err != nil {
		return errors.WithMessage(err, "setting up persistence")
	}

	// Register DominionApp 2 channel
	channel.RegisterApp(n.app)

	go n.client.Handle(n, n)
	go n.bus.Listen(listener)
	n.PrintConfig()
	return nil
}

func (n *node) setupContracts() error {
	var adjAddr common.Address
	var assAddr common.Address
	var appAddr common.Address
	var err error

	fmt.Println("💭 Validating contracts...")

	switch contractSetup := config.Chain.contractSetup; contractSetup {
	case contractSetupOptionValidate:
		if adjAddr, err = validateAdjudicator(n.cb); err == nil { // validate adjudicator
			if assAddr, err = validateAssetHolder(n.cb, adjAddr); err == nil { // validate asset holder
				appAddr, err = validateAppContract(n.cb) // validate app contract
			}
		}
	case contractSetupOptionDeploy:
		if adjAddr, err = deployAdjudicator(n.cb, n.onChain.Account); err == nil { // deploy adjudicator
			if assAddr, err = deployAssetHolder(n.cb, adjAddr, n.onChain.Account); err == nil { // deploy asset holder
				appAddr, err = deployAppContract(n.cb, n.onChain.Account) // deploy DominionApp.sol
			}
		}
	case contractSetupOptionValidateOrDeploy:
		if adjAddr, err = validateAdjudicator(n.cb); err != nil { // validate adjudicator
			fmt.Println("❌ Adjudicator invalid")
			adjAddr, err = deployAdjudicator(n.cb, n.onChain.Account) // deploy adjudicator
		}

		if err == nil {
			if assAddr, err = validateAssetHolder(n.cb, adjAddr); err != nil { // validate asset holder
				fmt.Println("❌ Asset holder invalid")
				assAddr, err = deployAssetHolder(n.cb, adjAddr, n.onChain.Account) // deploy asset holder
			}
		}

		if err == nil {
			if appAddr, err = validateAppContract(n.cb); err != nil { // validate app contract
				fmt.Println("❌ App contract invalid")
				appAddr, err = deployAppContract(n.cb, n.onChain.Account) // deploy DominionApp.sol
			}
		}
	default:
		// unsupported setup method
		err = errors.New(fmt.Sprintf("Unsupported contract setup method '%s'.", contractSetup))
	}

	fmt.Println("✅ Contracts validated.")

	if err != nil {
		return errors.WithMessage(err, "contract setup failed")
	}

	n.adjAddr = adjAddr
	n.assetAddr = assAddr
	n.appAddr = appAddr
	recvAddr := ewallet.AsEthAddr(n.onChain.Address())
	n.adjudicator = echannel.NewAdjudicator(n.cb, n.adjAddr, recvAddr, n.onChain.Account)
	n.asset = (*ewallet.Address)(&n.assetAddr)
	n.app = app.NewDominionApp(ewallet.AsWalletAddr(appAddr))
	n.log.WithField("Adj", n.adjAddr).WithField("Asset", n.assetAddr).WithField("App", n.appAddr).Debug("Set contracts")

	funder := echannel.NewFunder(n.cb)
	funder.RegisterAsset(ewallet.Address(n.assetAddr), new(echannel.ETHDepositor), n.onChain.Account)
	n.funder = funder

	return nil
}

func (n *node) setupPersistence() error {
	if config.Node.PersistenceEnabled {
		n.log.Info("Starting persistence")
		db, err := leveldb.LoadDatabase(config.Node.PersistencePath)
		if err != nil {
			return errors.WithMessage(err, "creating/loading database")
		}
		persister := keyvalue.NewPersistRestorer(db)
		n.client.EnablePersistence(persister)

		ctx, cancel := context.WithTimeout(context.Background(), config.Node.ReconnecTimeout)
		defer cancel()
		if err := n.client.Restore(ctx); err != nil {
			n.log.WithError(err).Warn("Could not restore client")
			// return the error.
		}
	} else {
		n.log.Info("Persistence disabled")
	}
	return nil
}

func validateAdjudicator(cb echannel.ContractBackend) (common.Address, error) {
	ctx, cancel := newTransactionContext()
	defer cancel()

	adjAddr := config.Chain.adjudicator
	return adjAddr, echannel.ValidateAdjudicator(ctx, cb, adjAddr)
}

func validateAssetHolder(cb echannel.ContractBackend, adjAddr common.Address) (common.Address, error) {
	ctx, cancel := newTransactionContext()
	defer cancel()

	assAddr := config.Chain.assetholder
	return assAddr, echannel.ValidateAssetHolderETH(ctx, cb, assAddr, adjAddr)
}

// validate contract bytecode on blockchain
// see echannel.asset.go
func validateContract(ctx context.Context,
	backend bind.ContractCaller, contract common.Address, bytecode string) error {
	code, err := backend.CodeAt(ctx, contract, nil)
	if err != nil {
		err = cherrors.CheckIsChainNotReachableError(err)
		return errors.WithMessage(err, "fetching contract code")
	}
	if hex.EncodeToString(code) != bytecode {
		return errors.Wrap(echannel.ErrInvalidContractCode, "incorrect contract code")
	}
	return nil
}

func validateAppContractETH(ctx context.Context,
	backend bind.ContractCaller, appAddr common.Address) error {
	return validateContract(ctx, backend, appAddr, dominionApp.DominionAppBinRuntime)
}

// TODO IMPLEMENT!
func validateAppContract(cb echannel.ContractBackend) (common.Address, error) {
	ctx, cancel := newTransactionContext()
	defer cancel()

	appAddr := config.Chain.appContract
	return appAddr, validateAppContractETH(ctx, cb, appAddr)
}

// deployAdjudicator deploys the Adjudicator to the blockchain and returns its address
// or an error.
func deployAdjudicator(cb echannel.ContractBackend, acc accounts.Account) (common.Address, error) {
	fmt.Println("🌐 Deploying adjudicator")
	ctx, cancel := context.WithTimeout(context.Background(), config.Chain.TxTimeout)
	defer cancel()
	adjAddr, err := echannel.DeployAdjudicator(ctx, cb, acc)
	return adjAddr, errors.WithMessage(err, "deploying eth adjudicator")
}

// deployAssetHolder deploys the Assetholder to the blockchain and returns its address
// or an error. Needs an Adjudicator address as second argument.
func deployAssetHolder(cb echannel.ContractBackend, adjudicator common.Address, acc accounts.Account) (common.Address, error) {
	fmt.Println("🌐 Deploying asset holder")
	ctx, cancel := context.WithTimeout(context.Background(), config.Chain.TxTimeout)
	defer cancel()
	asset, err := echannel.DeployETHAssetholder(ctx, cb, adjudicator, acc)
	return asset, errors.WithMessage(err, "deploying eth assetholder")
}

// deployApp deploys the DominionApp contract to the blockchain and returns its address
// or an error.
func deployAppContract(cb echannel.ContractBackend, acc accounts.Account) (common.Address, error) {

	const gasLimit = 1100000 // Must be sufficient for deploying DominionApp.sol.
	var appAddress common.Address

	fmt.Println("🌐 Deploying App contract 'DominionApp.sol'")
	ctx, cancel := context.WithTimeout(context.Background(), config.Chain.TxTimeout)
	defer cancel()

	// Deploy dominion App.
	var err error
	var tops *bind.TransactOpts
	var tx *types.Transaction
	if tops, err = cb.NewTransactor(ctx, gasLimit, acc); err == nil {
		if appAddress, tx, _, err = dominionApp.DeployDominionApp(tops, cb); err == nil {
			_, err = bind.WaitDeployed(ctx, cb, tx)
		}
	}

	return appAddress, errors.WithMessage(err, "deploying eth app contract")
}

// setupWallet imports the mnemonic and returns a corresponding wallet and
// the derived account at the given account index.
func setupWallet(mnemonic string, accountIndex uint) (*phd.Wallet, *phd.Account, error) {
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return nil, nil, errors.WithMessage(err, "creating hdwallet")
	}

	perunWallet, err := phd.NewWallet(wallet, accounts.DefaultBaseDerivationPath.String(), accountIndex)
	if err != nil {
		return nil, nil, errors.WithMessage(err, "creating perun wallet")
	}
	acc, err := perunWallet.NewAccount()
	if err != nil {
		return nil, nil, errors.WithMessage(err, "creating account")
	}

	return perunWallet, acc, nil
}

func newTransactionContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), config.Chain.TxTimeout)
}

func (n *node) PrintConfig() error {
	fmt.Printf(
		"Alias: %s\n"+
			"Listening: %s:%d\n"+
			"ETH RPC URL: %s\n"+
			"Perun ID: %s\n"+
			"OffChain: %s\n"+
			"ETHAssetHolder: %s\n"+
			"Adjudicator: %s\n"+
			"AppContract: %s\n"+
			"", config.Alias, config.Node.IP, config.Node.Port, config.Chain.URL, n.onChain.Address().String(), n.offChain.Address().String(), n.assetAddr.String(), n.adjAddr.String(), n.appAddr.String())

	fmt.Println("Known peers:")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.TabIndent)
	for alias, peer := range config.Peers {
		fmt.Fprintf(w, "%s\t%v\t%s:%d\n", alias, peer.PerunID, peer.Hostname, peer.Port)
	}
	return w.Flush()
}
