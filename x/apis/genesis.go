package apis

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis sets distribution information for genesis.
func InitGenesis(ctx sdk.Context, keeper Keeper, gs GenesisState) {
	// Set the requests and oracles from params
	keeper.SetParams(ctx, gs.Params)

	// Iterate through the posted prices and set them in the store if they are not expired
	for _, pp := range gs.PostedPrices {
		if pp.Expiry.After(ctx.BlockTime()) {
			_, err := keeper.SetPrice(ctx, pp.OracleAddress, pp.RequestID, pp.Price, pp.Expiry)
			if err != nil {
				panic(err)
			}
		}
	}
	params := keeper.GetParams(ctx)

	// Set the current price (if any) based on what's now in the store
	for _, request := range params.Requests {
		if !request.Active {
			continue
		}
		rps, err := keeper.GetRawPrices(ctx, request.RequestID)
		if err != nil {
			panic(err)
		}
		if len(rps) == 0 {
			continue
		}
		err = keeper.SetCurrentPrices(ctx, request.RequestID)
		if err != nil {
			panic(err)
		}
	}
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, keeper Keeper) GenesisState {

	// Get the params for requests and oracles
	params := keeper.GetParams(ctx)

	var postedPrices []PostedPrice
	for _, request := range keeper.GetRequests(ctx) {
		pp, err := keeper.GetRawPrices(ctx, request.RequestID)
		if err != nil {
			panic(err)
		}
		postedPrices = append(postedPrices, pp...)
	}

	return NewGenesisState(params, postedPrices)
}
