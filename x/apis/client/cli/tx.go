package cli

import (
	"bufio"
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"

	tmtime "github.com/tendermint/tendermint/types/time"

	"github.com/cbonoz/cosmos20/x/apis/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	apisTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Pricefeed transactions subcommands",
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
		Use:   "postprice [requestID] [price] [expiry]",
		Short: "post the latest price for a particular request with a given expiry as a UNIX time",
		Example: fmt.Sprintf("%s tx %s postprice bnb:usd 25 9999999999 --from validator",
			version.ClientName, types.ModuleName),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			price, err := sdk.NewDecFromStr(args[1])
			if err != nil {
				return err
			}

			expiryInt, err := strconv.ParseInt(args[2], 10, 64)
			if err != nil {
				return fmt.Errorf("invalid expiry %s: %w", args[2], err)
			}

			if expiryInt > types.MaxExpiry {
				return fmt.Errorf("invalid expiry; got %d, max: %d", expiryInt, types.MaxExpiry)
			}

			expiry := tmtime.Canonical(time.Unix(expiryInt, 0))

			msg := types.NewMsgPostResponse(cliCtx.GetFromAddress(), args[0], price, expiry)
			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
