package types

import (
	"testing"
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
				URL: "test.com/api",
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
