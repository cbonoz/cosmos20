package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cbonoz/cosmos20/app"
	"github.com/cbonoz/cosmos20/x/apis"
)

func NewPricefeedGenStateMulti() app.GenesisState {
	pfGenesis := apis.GenesisState{
		Params: apis.Params{
			Requests: []apis.Request{
				{RequestID: "btc:usd", BaseAsset: "btc", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
				{RequestID: "xrp:usd", BaseAsset: "xrp", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
			},
		},
		PostedPrices: []apis.PostedPrice{
			{
				RequestID:      "btc:usd",
				OracleAddress: sdk.AccAddress{},
				Price:         sdk.MustNewDecFromStr("8000.00"),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
			{
				RequestID:      "xrp:usd",
				OracleAddress: sdk.AccAddress{},
				Price:         sdk.MustNewDecFromStr("0.25"),
				Expiry:        time.Now().Add(1 * time.Hour),
			},
		},
	}
	return app.GenesisState{apis.ModuleName: apis.ModuleCdc.MustMarshalJSON(pfGenesis)}
}
