package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/example/testnet_crypstarter/testutil/keeper"
	"github.com/example/testnet_crypstarter/x/testnetcrypstarter/keeper"
	"github.com/example/testnet_crypstarter/x/testnetcrypstarter/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.TestnetcrypstarterKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
