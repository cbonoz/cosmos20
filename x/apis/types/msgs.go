package types

import (
	"errors"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	// TypeMsgPostResponse type of PostResponse msg
	TypeMsgPostResponse = "post_response"

	// MaxExpiry defines the max expiry time defined as UNIX time (9999-12-31 23:59:59 +0000 UTC)
	MaxExpiry = 253402300799
)

// ensure Msg interface compliance at compile time
var _ sdk.Msg = &MsgPostResponse{}

// MsgPostResponse struct representing a posted price message.
// Used by oracles to input prices to the apis
type MsgPostResponse struct {
	From     sdk.AccAddress `json:"from" yaml:"from"`           // client that sent in this address
	RequestID string         `json:"request_id" yaml:"request_id"` // asset code used by exchanges/api
	RawResponse    string        `json:"response" yaml:"response"`
}

// NewMsgPostResponse creates a new post price msg
func NewMsgPostResponse(
	from sdk.AccAddress,
	requestId string,
	response string) MsgPostResponse {
	return MsgPostResponse{
		From:     from,
		RequestID: requestId,
		RawResponse:    response,
	}
}

// Route Implements Msg.
func (msg MsgPostResponse) Route() string { return RouterKey }

// Type Implements Msg
func (msg MsgPostResponse) Type() string { return TypeMsgPostResponse }

// GetSignBytes Implements Msg.
func (msg MsgPostResponse) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgPostResponse) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

// ValidateBasic does a simple validation check that doesn't require access to any other information.
func (msg MsgPostResponse) ValidateBasic() error {
	if msg.From.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender address cannot be empty")
	}
	if strings.TrimSpace(msg.RequestID) == "" {
		return errors.New("request id cannot be blank")
	}
	return nil
}
