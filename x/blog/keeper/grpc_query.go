package keeper

import (
	"github.com/windforcer/blog/x/blog/types"
)

var _ types.QueryServer = Keeper{}
