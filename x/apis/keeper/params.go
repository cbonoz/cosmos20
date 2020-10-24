package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

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

// GetOracles returns the oracles in the apis store
func (k Keeper) GetOracles(ctx sdk.Context, requestID string) ([]sdk.AccAddress, error) {
	for _, m := range k.GetRequests(ctx) {
		if requestID == m.RequestID {
			return m.Oracles, nil
		}
	}
	return []sdk.AccAddress{}, sdkerrors.Wrap(types.ErrInvalidRequest, requestID)
}

// GetOracle returns the oracle from the store or an error if not found
func (k Keeper) GetOracle(ctx sdk.Context, requestID string, address sdk.AccAddress) (sdk.AccAddress, error) {
	oracles, err := k.GetOracles(ctx, requestID)
	if err != nil {
		return sdk.AccAddress{}, sdkerrors.Wrap(types.ErrInvalidRequest, requestID)
	}
	for _, addr := range oracles {
		if address.Equals(addr) {
			return addr, nil
		}
	}
	return sdk.AccAddress{}, sdkerrors.Wrap(types.ErrInvalidOracle, address.String())
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
