package cli

import (
	"bufio"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"

	"github.com/cbonoz/cosmos20/x/apis/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	apisTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Apis transactions subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	apisTxCmd.AddCommand(flags.PostCommands(
		GetCmdPostResponse(cdc),
	)...)

	return apisTxCmd
}

// GetCmdPostResponse cli command for posting prices.
func GetCmdPostResponse(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "postresponse [requestID] [response]",
		Short: "post the response for a particular api call",
		Example: fmt.Sprintf("%s tx %s postresponse 123 {} --from validator",
			version.ClientName, types.ModuleName),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgPostResponse(cliCtx.GetFromAddress(), args[0], "{}")
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
