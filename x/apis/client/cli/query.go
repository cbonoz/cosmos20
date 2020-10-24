package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cbonoz/cosmos20/x/apis/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	// Group nameservice queries under a subcommand
	apisQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the apis module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	apisQueryCmd.AddCommand(flags.GetCommands(
		GetCmdRequests(queryRoute, cdc),
		GetCmdQueryParams(queryRoute, cdc),
	)...)

	return apisQueryCmd
}

// GetCmdRequests queries list of requests in the apis
func GetCmdRequests(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "requests",
		Short: "get the requests in the apis",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			route := fmt.Sprintf("custom/%s/requests", queryRoute)
			res, _, err := cliCtx.QueryWithData(route, nil)
			if err != nil {
				return err
			}
			var requests types.Requests
			cdc.MustUnmarshalJSON(res, &requests)
			return cliCtx.PrintOutput(requests)
		},
	}
}

// GetCmdQueryParams queries the apis module parameters
func GetCmdQueryParams(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "params",
		Short: "get the apis module parameters",
		Long:  "Get the current global apis module parameters.",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			// Query
			route := fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryGetParams)
			res, _, err := cliCtx.QueryWithData(route, nil)
			if err != nil {
				return err
			}

			// Decode and print results
			var out types.Params
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
