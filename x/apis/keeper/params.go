package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cbonoz/cosmos20/x/apis/types"
)

// GetParams returns the params from the store
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	var p types.Params
	k.paramSubspace.GetParamSet(ctx, &p)
	return p
}

// SetParams sets params on the store
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSubspace.SetParamSet(ctx, &params)
}

// GetRequests returns the requests from params
func (k Keeper) GetRequests(ctx sdk.Context) types.Requests {
	return k.GetParams(ctx).Requests
}

// GetRequest returns the request if it is in the apis system
func (k Keeper) GetRequest(ctx sdk.Context, requestID string) (types.Request, bool) {
	requests := k.GetRequests(ctx)

	for i := range requests {
		if requests[i].RequestID == requestID {
			return requests[i], true
		}
	}
	return types.Request{}, false

}
