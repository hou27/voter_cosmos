package keeper

import (
	"github.com/hou27/voter/x/voter/types"
)

var _ types.QueryServer = Keeper{}
