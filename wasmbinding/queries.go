package wasmbinding

import (
	denomkeeper "kujira/x/denom/keeper"

	oraclekeeper "kujira/x/oracle/keeper"

	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
)

type QueryPlugin struct {
	denomKeeper  denomkeeper.Keeper
	bankkeeper   bankkeeper.Keeper
	oraclekeeper oraclekeeper.Keeper
}

// NewQueryPlugin returns a reference to a new QueryPlugin.
func NewQueryPlugin(bk bankkeeper.Keeper, ok oraclekeeper.Keeper, dk denomkeeper.Keeper) *QueryPlugin {
	return &QueryPlugin{
		denomKeeper:  dk,
		bankkeeper:   bk,
		oraclekeeper: ok,
	}
}
