package keeper

import (
	"github.com/example/testnet_crypstarter/x/testnetcrypstarter/types"
)

var _ types.QueryServer = Keeper{}
