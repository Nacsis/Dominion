/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"math/big"

	ethwallet "perun.network/go-perun/backend/ethereum/wallet"
	"perun.network/go-perun/wire"
	"perun.network/perun-examples/dominion-cli/app"
	"perun.network/perun-examples/dominion-cli/app/util"
	"perun.network/perun-examples/dominion-cli/client"
	"perun.network/perun-examples/dominion-cli/cmd/staticDemoUtil"

	"github.com/spf13/cobra"
)

// staticDemoCmd represents the staticDemo command
var staticDemoCmd = &cobra.Command{
	Use:   "staticDemo",
	Short: "Execute a simple deterministic example run of one dominion game.",
	Long: `Execute a stitic example run of a simple dominion game.
	Main purpose is to show/test that everything is working`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("staticDemo called")

		const (
			chainURL = "ws://127.0.0.1:8545"
			chainID  = 1337

			// Private keys.
			keyDeployer = "79ea8f62d97bc0591a4224c1725fca6b00de5b2cea286fe2e0bb35c5e76be46e"
			keyAlice    = "1af2e950272dd403de7a5760d41c6e44d92b6d02797e51810795ff03cc2cda4f"
			keyBob      = "f63d7d8e930bccd74e93cf5662fde2c28fd8be95edb70c73f1bdd863d07f412e"
		)

		// Deploy contracts.
		log.Println("Deploying contracts.")
		adjudicator, assetHolder, appAddress := staticDemoUtil.DeployContracts(chainURL, chainID, keyDeployer)
		asset := *ethwallet.AsWalletAddr(assetHolder)
		dominionApp := app.NewDominionApp(ethwallet.AsWalletAddr(appAddress))

		// Setup clients.
		log.Println("Setting up clients.")
		bus := wire.NewLocalBus() // Message bus used for off-chain communication.
		stake := client.EthToWei(big.NewFloat(5))
		alice := staticDemoUtil.SetupGameClient(bus, chainURL, chainID, adjudicator, asset, keyAlice, dominionApp, stake)
		bob := staticDemoUtil.SetupGameClient(bus, chainURL, chainID, adjudicator, asset, keyBob, dominionApp, stake)

		// Print balances before transactions.
		l := staticDemoUtil.NewBalanceLogger(chainURL)
		l.LogBalances(alice, bob)

		// Open dominionApp channel and play.
		log.Println("Opening channel.")
		appAlice := alice.OpenAppChannel(bob.WireAddress())
		appBob := bob.AcceptedChannel()
		log.Println("Channel Open")

		staticDemoUtil.DrawInitHand(appAlice, appBob)
		log.Println("Alice drawn init hand")
		appAlice.PlayCard(0)
		log.Println("Alice played a card")
		appAlice.BuyCard(util.Copper)
		log.Println("Alice Bought a card ", util.Copper)

		appAlice.EndTurn()
		log.Println("Alice end turn")

		staticDemoUtil.DrawInitHand(appBob, appAlice)
		log.Println("bob drawn init hand")
		appBob.EndTurn()
		log.Println("bob end turn")

		appAlice.EndGame()
		log.Println("Alice end game")

		// Payout.
		appAlice.Settle()
		appBob.Settle()
		log.Println("Settled")

		// Print balances after transactions.
		l.LogBalances(alice, bob)

		// Cleanup.
		alice.Shutdown()
		bob.Shutdown()
	},
}

func init() {
	rootCmd.AddCommand(staticDemoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// staticDemoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// staticDemoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
