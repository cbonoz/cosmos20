package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalid = sdkerrors.Register(ModuleName, 1, "invalid")
	ErrUknownError = sdkerrors.Register(ModuleName, 2, "unknown api error")
	ErrBadRequest = sdkerrors.Register(ModuleName, 3, "bad request")
)
