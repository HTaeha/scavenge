package cli

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/coreators/scavenge/x/scavenge/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
)

var _ = strconv.Itoa(0)

func CmdSubmitScavenge() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-scavenge [solution-hash] [description] [reward]",
		Short: "Broadcast message submit-scavenge",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// find a hash of the solution
			solutionHash := sha256.Sum256([]byte(args[0]))
			// convert the hash to string
			solutionHashString := hex.EncodeToString(solutionHash[:])

			argDescription := string(args[1])
			argReward := string(args[2])

			// create a new message
			msg := types.NewMsgSubmitScavenge(
				clientCtx.GetFromAddress().String(),
				string(solutionHashString),
				string(argDescription),
				string(argReward),
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			// broadcast the transaction with the message to the blockchain
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
