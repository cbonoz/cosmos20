package simulation

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/cbonoz/cosmos20/x/apis/types"
	apis "github.com/cbonoz/cosmos20/x/apis/types"
)


// RandomizedGenState generates a random GenesisState for apis
func RandomizedGenState(simState *module.SimulationState) {
	apisGenesis := loadApisGenState(simState)
	fmt.Printf("Selected randomly generated %s parameters:\n%s\n", types.ModuleName, codec.MustMarshalJSONIndent(simState.Cdc, apisGenesis))
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(apisGenesis)
}

// loadApisGenState loads a valid apis gen state
func loadApisGenState(simState *module.SimulationState) apis.GenesisState {
	var requests []apis.Request
	params := apis.NewParams(requests)
	return apis.NewGenesisState(params, postedPrices)
}
