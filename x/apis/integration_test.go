package apis_test

import (
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cbonoz/cosmos20/app"
	"github.com/cbonoz/cosmos20/x/apis"
)

func NewApisGenStateMulti() app.GenesisState {
	pfGenesis := apis.GenesisState{
		Params: apis.Params{
			Requests: []apis.Request{
				{RequestID: "btc:usd",URL: "test", Active: true},
			},
		},
	}
	return app.GenesisState{apis.ModuleName: apis.ModuleCdc.MustMarshalJSON(pfGenesis)}
}

