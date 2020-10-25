package apis

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis sets distribution information for genesis.
func InitGenesis(ctx sdk.Context, keeper Keeper, gs GenesisState) {
	// Set the requests and oracles from params
	keeper.SetParams(ctx, gs.Params)
	// params := keeper.GetParams(ctx)
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, keeper Keeper) GenesisState {

	// Get the params for requests and oracles
	params := keeper.GetParams(ctx)
	return NewGenesisState(params)
}
