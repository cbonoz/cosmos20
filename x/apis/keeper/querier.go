package keeper

import (
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cbonoz/cosmos20/x/apis/types"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {
		case types.QueryRequests:
			return queryRequests(ctx, req, keeper)
		case types.QueryGetParams:
			return queryGetParams(ctx, req, keeper)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown %s query endpoint", types.ModuleName)
		}
	}

}

func queryPrice(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) (res []byte, sdkErr error) {
	var requestParams types.QueryWithRequestIDParams
	err := types.ModuleCdc.UnmarshalJSON(req.Data, &requestParams)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	_, found := keeper.GetRequest(ctx, requestParams.RequestID)
	if !found {
		return []byte{}, sdkerrors.Wrap(types.ErrAssetNotFound, requestParams.RequestID)
	}
	currentPrice, sdkErr := keeper.GetCurrentPrice(ctx, requestParams.RequestID)
	if sdkErr != nil {
		return nil, sdkErr
	}
	bz, err := codec.MarshalJSONIndent(types.ModuleCdc, currentPrice)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func queryRawPrices(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) (res []byte, sdkErr error) {
	var requestParams types.QueryWithRequestIDParams
	err := types.ModuleCdc.UnmarshalJSON(req.Data, &requestParams)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	_, found := keeper.GetRequest(ctx, requestParams.RequestID)
	if !found {
		return []byte{}, sdkerrors.Wrap(types.ErrAssetNotFound, requestParams.RequestID)
	}

	rawPrices, err := keeper.GetRawPrices(ctx, requestParams.RequestID)
	if err != nil {
		return nil, err
	}

	bz, err := codec.MarshalJSONIndent(types.ModuleCdc, rawPrices)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func queryOracles(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) (res []byte, sdkErr error) {
	var requestParams types.QueryWithRequestIDParams
	err := types.ModuleCdc.UnmarshalJSON(req.Data, &requestParams)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	oracles, err := keeper.GetOracles(ctx, requestParams.RequestID)
	if err != nil {
		return []byte{}, sdkerrors.Wrap(types.ErrAssetNotFound, requestParams.RequestID)
	}

	bz, err := codec.MarshalJSONIndent(types.ModuleCdc, oracles)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func queryRequests(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) (res []byte, sdkErr error) {
	requests := keeper.GetRequests(ctx)

	bz, err := codec.MarshalJSONIndent(types.ModuleCdc, requests)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

// query params in the apis store
func queryGetParams(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	params := keeper.GetParams(ctx)

	// Encode results
	bz, err := codec.MarshalJSONIndent(types.ModuleCdc, params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}
