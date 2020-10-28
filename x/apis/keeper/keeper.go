package keeper

import (
	"fmt"
	"net/http"
	"io/ioutil"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params/subspace"

	"github.com/cbonoz/cosmos20/x/apis/types"
)

// Keeper struct for apis module
type Keeper struct {
	// key used to access the stores from Context
	key sdk.StoreKey
	// Codec for binary encoding/decoding
	cdc *codec.Codec
	// The reference to the Paramstore to get and set apis specific params
	paramSubspace subspace.Subspace
}

// NewKeeper returns a new keeper for the apis module.
func NewKeeper(
	cdc *codec.Codec, key sdk.StoreKey, paramstore subspace.Subspace,
) Keeper {
	if !paramstore.HasKeyTable() {
		paramstore = paramstore.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:           cdc,
		key:           key,
		paramSubspace: paramstore,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}


// Makes the api calls configured in the param store broadcasting events for each completed result
func (k Keeper) MakeRequest(ctx sdk.Context, requestID string) (string, error) {
	requests := k.GetRequests(ctx)
	// TODO: create requests

	for i := range requests {
		if requests[i].RequestID == requestID {
			resp, err := http.Get(requests[i].URL)
			if (err != nil) {
				return "", err
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			return string(body), err
		}
	}


	return "", fmt.Errorf("No request with id found in config, ID: %s", requestID)
}