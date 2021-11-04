package keeper

import (
	"github.com/coreators/scavenge/x/scavenge/types"
)

var _ types.QueryServer = Keeper{}
