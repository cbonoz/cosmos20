package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// DONTCOVER

var (
	ErrEmptyInput = sdkerrors.Register(ModuleName, 2, "input must not be empty")
	ErrEndpointNotFound = sdkerrors.Register(ModuleName, 7, "endpoint not found")
	ErrUknownError = sdkerrors.Register(ModuleName, 2, "unknown api error")
	ErrBadRequest = sdkerrors.Register(ModuleName, 3, "bad request")
)
