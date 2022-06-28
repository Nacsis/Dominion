/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package play

import (
	"bufio"
	"fmt"
	"os"

	prompt "github.com/c-bata/go-prompt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "play dominion against a remote peer interactively",
	Long:  `play dominion in an interactive shell process against a remote peer via a go-perun app-channel.`,
	Run:   runPlay,
}

// CommandLineFlags contains the command line flags.
type CommandLineFlags struct {
	testAPIEnabled bool
	cfgFile        string
	cfgNetFile     string
	useStdIO       bool
}

var flags CommandLineFlags

func init() {
	playCmd.PersistentFlags().StringVar(&flags.cfgFile, "config", "config.yaml", "General config file")
	playCmd.PersistentFlags().StringVar(&flags.cfgNetFile, "network", "network.yaml", "Network config file")
	// playCmd.PersistentFlags().BoolVar(&flags.testAPIEnabled, "test-api", false, "Expose testing API at 8080")
	// playCmd.PersistentFlags().BoolVar(&GetConfig().Node.PersistenceEnabled, "persistence", false, "Enables the persistence")
	// playCmd.PersistentFlags().StringVar(&GetConfig().SecretKey, "sk", "", "ETH Secret Key")
	viper.BindPFlag("secretkey", playCmd.PersistentFlags().Lookup("sk"))
	playCmd.PersistentFlags().BoolVar(&flags.useStdIO, "stdio", false, "Read from stdin")
}

// GetPlayCmd exposes playCmd so that it can be used as a sub-command by another cobra command instance.
func GetPlayCmd() *cobra.Command {
	return playCmd
}

// runPlay is executed everytime the program is started with the `demo` sub-command.
func runPlay(c *cobra.Command, args []string) {
	Setup()
	// if flags.testAPIEnabled {
	// 	StartTestAPI()
	// }
	if flags.useStdIO {
		runWithStdIO(executor)
	} else {
		p := prompt.New(
			executor,
			completer,
			prompt.OptionPrefix("> "),
			prompt.OptionTitle("perun"),
		)
		p.Run()
	}
}

func runWithStdIO(executor func(string)) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("> ")
	for scanner.Scan() {
		executor(scanner.Text())
		fmt.Printf("> ")
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error scanning stdin: %v\n", err)
		os.Exit(1)
	}
}

func completer(prompt.Document) []prompt.Suggest {
	return []prompt.Suggest{}
}

// executor wraps the demo executor to print error messages.
func executor(in string) {
	AddInput(in)
}
