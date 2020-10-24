package apis

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cbonoz/cosmos20/x/apis/types"
)

// EndBlocker updates the current apis
func EndBlocker(ctx sdk.Context, k Keeper) {
	// Update the current price of each asset.
	for _, request := range k.GetRequests(ctx) {
		if !request.Active {
			continue
		}

		err := k.SetCurrentPrices(ctx, request.RequestID)
		if err != nil && !errors.Is(err, types.ErrNoValidPrice) {
			panic(err)
		}
	}
}
