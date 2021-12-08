package testnetcrypstarter_test

import (
	"testing"

	keepertest "github.com/example/testnet_crypstarter/testutil/keeper"
	"github.com/example/testnet_crypstarter/x/testnetcrypstarter"
	"github.com/example/testnet_crypstarter/x/testnetcrypstarter/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TestnetcrypstarterKeeper(t)
	testnetcrypstarter.InitGenesis(ctx, *k, genesisState)
	got := testnetcrypstarter.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
