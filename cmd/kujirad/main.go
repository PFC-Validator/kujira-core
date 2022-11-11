package main

import (
	"os"

	"github.com/Team-Kujira/core/app"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/ignite-hq/cli/ignite/pkg/cosmoscmd"
	tm "github.com/tendermint/tendermint/cmd/tendermint/commands"
)

func main() {
	rootCmd, _ := cosmoscmd.NewRootCmd(
		app.Name,
		app.AccountAddressPrefix,
		app.DefaultNodeHome,
		app.Name,
		app.ModuleBasics,
		app.NewIgniteApp,
		// this line is used by starport scaffolding # root/arguments
	)
	rootCmd.AddCommand(tm.ReIndexEventCmd)
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
