package apis

import (

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// EndBlocker submits configured API calls at the end of the current block
func EndBlocker(ctx sdk.Context, k Keeper) {
	for _, request := range k.GetRequests(ctx) {
		if !request.Active {
			continue
		}

		err := k.MakeRequest(ctx, request.RequestID)
		if err != nil { // && !errors.Is(err, types.ErrNoValidPrice) {
			panic(err)
		}
	}
}
