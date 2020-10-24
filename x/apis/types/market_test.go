package types

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	tmtypes "github.com/tendermint/tendermint/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestRequestValidate(t *testing.T) {
	mockPrivKey := tmtypes.NewMockPV()
	pubkey, err := mockPrivKey.GetPubKey()
	require.NoError(t, err)
	addr := sdk.AccAddress(pubkey.Address())

	testCases := []struct {
		msg     string
		request  Request
		expPass bool
	}{
		{
			"valid request",
			Request{
				RequestID:   "request",
				BaseAsset:  "xrp",
				QuoteAsset: "bnb",
				Oracles:    []sdk.AccAddress{addr},
				Active:     true,
			},
			true,
		},
		{
			"invalid id",
			Request{
				RequestID: " ",
			},
			false,
		},
		{
			"invalid base asset",
			Request{
				RequestID:  "request",
				BaseAsset: "XRP",
			},
			false,
		},
		{
			"invalid request",
			Request{
				RequestID:   "request",
				BaseAsset:  "xrp",
				QuoteAsset: "BNB",
			},
			false,
		},
		{
			"empty oracle address ",
			Request{
				RequestID:   "request",
				BaseAsset:  "xrp",
				QuoteAsset: "bnb",
				Oracles:    []sdk.AccAddress{nil},
			},
			false,
		},
		{
			"empty oracle address ",
			Request{
				RequestID:   "request",
				BaseAsset:  "xrp",
				QuoteAsset: "bnb",
				Oracles:    []sdk.AccAddress{addr, addr},
			},
			false,
		},
	}

	for _, tc := range testCases {
		err := tc.request.Validate()
		if tc.expPass {
			require.NoError(t, err)
		} else {
			require.Error(t, err)
		}
	}
}

func TestPostedPriceValidate(t *testing.T) {
	now := time.Now()
	mockPrivKey := tmtypes.NewMockPV()
	pubkey, err := mockPrivKey.GetPubKey()
	require.NoError(t, err)
	addr := sdk.AccAddress(pubkey.Address())

	testCases := []struct {
		msg         string
		postedPrice PostedPrice
		expPass     bool
	}{
		{
			"valid posted price",
			PostedPrice{
				RequestID:      "request",
				OracleAddress: addr,
				Price:         sdk.OneDec(),
				Expiry:        now,
			},
			true,
		},
		{
			"invalid id",
			PostedPrice{
				RequestID: " ",
			},
			false,
		},
		{
			"invalid oracle",
			PostedPrice{
				RequestID:      "request",
				OracleAddress: nil,
			},
			false,
		},
		{
			"invalid price",
			PostedPrice{
				RequestID:      "request",
				OracleAddress: addr,
				Price:         sdk.NewDec(-1),
			},
			false,
		},
		{
			"zero expiry time ",
			PostedPrice{
				RequestID:      "request",
				OracleAddress: addr,
				Price:         sdk.OneDec(),
				Expiry:        time.Time{},
			},
			false,
		},
	}

	for _, tc := range testCases {
		err := tc.postedPrice.Validate()
		if tc.expPass {
			require.NoError(t, err)
		} else {
			require.Error(t, err)
		}
	}
}
