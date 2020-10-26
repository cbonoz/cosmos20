package keeper_test

import (
	// "time"
	// "github.com/cbonoz/cosmos20/x/apis/types"

	"github.com/cbonoz/cosmos20/app"
	"github.com/cbonoz/cosmos20/x/apis"
)

const (
	exampleAPICall = "https://www.metaweather.com/api/location/search/?query=boston"
)

func NewApisGenStateMulti() app.GenesisState {
	pfGenesis := apis.GenesisState{
		Params: apis.Params{
			Requests: []apis.Request{
				{RequestID: "get_weather", URL: exampleAPICall, Active: true},
			},
		},
	}
	return app.GenesisState{apis.ModuleName: apis.ModuleCdc.MustMarshalJSON(pfGenesis)}
}
