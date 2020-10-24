package simulation

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	appparams "github.com/cbonoz/cosmos20/app/params"
	"github.com/cbonoz/cosmos20/x/apis/keeper"
	"github.com/cbonoz/cosmos20/x/apis/types"
)

// Simulation operation weights constants
const (
	OpWeightMsgUpdatePrices = "op_weight_msg_update_prices"

	// Block time params are un-exported constants in cosmos-sdk/x/simulation.
	// Copy them here in lieu of importing them.
	minTimePerBlock time.Duration = (10000 / 2) * time.Second
	maxTimePerBlock time.Duration = 10000 * time.Second

	// Calculate the average block time
	AverageBlockTime time.Duration = (maxTimePerBlock - minTimePerBlock) / 2
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(
	appParams simulation.AppParams, cdc *codec.Codec, ak auth.AccountKeeper, k keeper.Keeper,
) simulation.WeightedOperations {
	var weightMsgUpdatePrices int
	// var numBlocks int

	appParams.GetOrGenerate(cdc, OpWeightMsgUpdatePrices, &weightMsgUpdatePrices, nil,
		func(_ *rand.Rand) {
			weightMsgUpdatePrices = appparams.DefaultWeightMsgUpdatePrices
		},
	)

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightMsgUpdatePrices,
			SimulateMsgUpdatePrices(ak, k, 10000),
		),
	}
}

// getExpiryTime gets a price expiry time by taking the current time and adding a delta to it
func getExpiryTime(ctx sdk.Context) (t time.Time) {
	// need to use the blocktime from the context as the context generates random start time when running simulations
	return ctx.BlockTime().Add(AverageBlockTime * 5000) // if blocks were 6 seconds, the expiry would be 8 hrs
}
