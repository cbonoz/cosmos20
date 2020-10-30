package keeper_test

import (
	"testing"
	// // "time"
	"github.com/stretchr/testify/require"
	"github.com/cbonoz/cosmos20/x/apis/types"
	// sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/cbonoz/cosmos20/app"
)

// TestKeeper_BasicGetRequest tests adding requests to the api set, getting requests from the store
func TestKeeper_BasicGetRequest(t *testing.T) {
	tApp := app.NewTestApp()
	ctx := tApp.NewContext(true, abci.Header{})
	keeper := tApp.GetApisKeeper()
	requestId := "123"

	mp := types.Params{
		Requests: types.Requests{
			types.NewRequest(requestId, exampleAPICall, "GET", "", true),
		},
	}
	keeper.SetParams(ctx, mp)
	body, err := keeper.MakeRequest(ctx, requestId)
	t.Log("response", body)

	// Perform basic validation.
	require.Nil(t, err)
	require.Contains(t, body, "Boston") // Response should contain Boston from weather API.
}
