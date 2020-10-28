package apis

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cbonoz/cosmos20/x/apis/types"
)

// EndBlocker submits configured API calls at the end of the current block
func EndBlocker(ctx sdk.Context, k Keeper) {
	// Iterate over configured events.
	for _, request := range k.GetRequests(ctx) {
		if !request.Active {
			continue
		}

		// Make programmatic http calls according to param configuration
		body, err := k.MakeRequest(ctx, request.RequestID)
		if err != nil {
			panic(err)
		}
		// Broadcast payload as post response event to any subscribers.
		response := types.NewMsgPostResponse(nil, request.RequestID, body)
		result, err := HandleMsgPostResponse(ctx, k, response)
		if (err != nil) {
			panic(err)
		}
		fmt.Printf("result: %s, response: %s", string(result.Data), response)
	}
}
