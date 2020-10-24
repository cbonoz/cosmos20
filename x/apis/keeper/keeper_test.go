package keeper_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cbonoz/cosmos20/app"
	"github.com/cbonoz/cosmos20/x/apis/types"
)

// TestKeeper_SetGetRequest tests adding requests to the api set, getting requests from the store
func TestKeeper_SetGetRequest(t *testing.T) {
	tApp := app.NewTestApp()
	ctx := tApp.NewContext(true, abci.Header{})
	keeper := tApp.GetPriceFeedKeeper()

	mp := types.Params{
		Requests: types.Requests{
			types.Request{RequestID: "tstusd", BaseAsset: "tst", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
		},
	}
	keeper.SetParams(ctx, mp)
	requests := keeper.GetRequests(ctx)
	require.Equal(t, len(requests), 1)
	require.Equal(t, requests[0].RequestID, "tstusd")

	_, found := keeper.GetRequest(ctx, "tstusd")
	require.Equal(t, found, true)

	mp = types.Params{
		Requests: types.Requests{
			types.Request{RequestID: "tstusd", BaseAsset: "tst", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
			types.Request{RequestID: "tst2usd", BaseAsset: "tst2", QuoteAsset: "usd", Oracles: []sdk.AccAddress{}, Active: true},
		},
	}
	keeper.SetParams(ctx, mp)
	requests = keeper.GetRequests(ctx)
	require.Equal(t, len(requests), 2)
	require.Equal(t, requests[0].RequestID, "tstusd")
	require.Equal(t, requests[1].RequestID, "tst2usd")

	_, found = keeper.GetRequest(ctx, "nan")
	require.Equal(t, found, false)
}
