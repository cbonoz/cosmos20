package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

const (
	RestRequestID = "request_id"
)

// PostResponseReq defines the properties of a PostResponse request's body.
type PostResponseReq struct {
	BaseReq  rest.BaseReq `json:"base_req"`
	RequestID string       `json:"request_id"`
	Price    string       `json:"price"`
	Expiry   string       `json:"expiry"`
}

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	registerQueryRoutes(cliCtx, r)
	registerTxRoutes(cliCtx, r)
}
