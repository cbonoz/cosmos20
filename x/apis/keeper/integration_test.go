package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cbonoz/cosmos20/app"
	"github.com/cbonoz/cosmos20/x/apis"
)

func NewApisGenStateMulti() app.GenesisState {
	pfGenesis := apis.GenesisState{
		Params: apis.Params{
			Requests: []apis.Request{
				{RequestID: "get_weather", BaseAsset: "btc", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
				{RequestID: "xrp:usd", BaseAsset: "xrp", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
			},
		},
	}
	return app.GenesisState{apis.ModuleName: apis.ModuleCdc.MustMarshalJSON(pfGenesis)}
}
