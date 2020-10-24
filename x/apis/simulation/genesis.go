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

var (
	// BaseAssets is a list of collateral asset denoms
	BaseAssets = [3]string{"bnb", "xrp", "btc"}
	QuoteAsset = "usd"
)

// RandomizedGenState generates a random GenesisState for apis
func RandomizedGenState(simState *module.SimulationState) {
	apisGenesis := loadPricefeedGenState(simState)
	fmt.Printf("Selected randomly generated %s parameters:\n%s\n", types.ModuleName, codec.MustMarshalJSONIndent(simState.Cdc, apisGenesis))
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(apisGenesis)
}

// loadPricefeedGenState loads a valid apis gen state
func loadPricefeedGenState(simState *module.SimulationState) apis.GenesisState {
	var requests []apis.Request
	var postedPrices []apis.PostedPrice
	for _, denom := range BaseAssets {
		// Select an account to be the oracle
		oracle, _ := simulation.RandomAcc(simState.Rand, simState.Accounts)

		requestID := fmt.Sprintf("%s:%s", denom, QuoteAsset)
		// Construct request for asset
		request := apis.Request{
			RequestID:   requestID,
			BaseAsset:  denom,
			QuoteAsset: QuoteAsset,
			Oracles:    []sdk.AccAddress{oracle.Address},
			Active:     true,
		}

		// Construct posted price for asset
		postedPrice := apis.PostedPrice{
			RequestID:      request.RequestID,
			OracleAddress: oracle.Address,
			Price:         getInitialPrice(requestID),
			Expiry:        simState.GenTimestamp.Add(time.Hour * 24),
		}
		requests = append(requests, request)
		postedPrices = append(postedPrices, postedPrice)
	}
	params := apis.NewParams(requests)
	return apis.NewGenesisState(params, postedPrices)
}

// getInitialPrice gets the starting price for each of the base assets
func getInitialPrice(requestID string) (price sdk.Dec) {
	switch requestID {
	case "btc:usd":
		return sdk.MustNewDecFromStr("7000")
	case "bnb:usd":
		return sdk.MustNewDecFromStr("14")
	case "xrp:usd":
		return sdk.MustNewDecFromStr("0.2")
	default:
		return sdk.MustNewDecFromStr("20") // Catch future additional assets
	}
}
