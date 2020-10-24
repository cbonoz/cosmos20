package incentive

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cbonoz/cosmos20/x/incentive/keeper"
)

// BeginBlocker runs at the start of every block
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	k.DeleteExpiredClaimsAndClaimPeriods(ctx)
	k.ApplyRewardsToCdps(ctx)
	k.CreateAndDeleteRewardPeriods(ctx)
}
