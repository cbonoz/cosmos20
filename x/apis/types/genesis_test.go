package types

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	tmtypes "github.com/tendermint/tendermint/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestGenesisStateValidate(t *testing.T) {
	now := time.Now()
	mockPrivKey := tmtypes.NewMockPV()
	pubkey, err := mockPrivKey.GetPubKey()
	require.NoError(t, err)
	addr := sdk.AccAddress(pubkey.Address())

	testCases := []struct {
		msg          string
		genesisState GenesisState
		expPass      bool
	}{
		{
			msg:          "default",
			genesisState: DefaultGenesisState(),
			expPass:      true,
		},
		{
			msg: "valid genesis",
			genesisState: NewGenesisState(
				NewParams(Requests{
					{"request", "{}", true},
				}),
			),
			expPass: true,
		},
	}

	for _, tc := range testCases {
		err := tc.genesisState.Validate()
		if tc.expPass {
			require.NoError(t, err, tc.msg)
		} else {
			require.Error(t, err, tc.msg)
		}
	}
}
