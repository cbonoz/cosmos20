package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"

	abci "github.com/tendermint/tendermint/abci/types"
	tmtime "github.com/tendermint/tendermint/types/time"

	"github.com/cbonoz/cosmos20/app"
	"github.com/cbonoz/cosmos20/x/apis/keeper"
)

type KeeperTestSuite struct {
	suite.Suite

	keeper keeper.Keeper
	addrs  []sdk.AccAddress
	app    app.TestApp
	ctx    sdk.Context
}

func (suite *KeeperTestSuite) SetupTest() {
	tApp := app.NewTestApp()
	ctx := tApp.NewContext(true, abci.Header{Height: 1, Time: tmtime.Now()})
	_, addrs := app.GeneratePrivKeyAddressPairs(10)
	tApp.InitializeFromGenesisStates(
		NewApisGenStateMulti(),
	)
	suite.ctx = ctx
	suite.addrs = addrs
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
