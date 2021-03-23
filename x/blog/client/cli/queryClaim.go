package cli

import (
    "context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
    "github.com/windforcer/blog/x/blog/types"
)

func CmdListClaim() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-claim",
		Short: "list all claim",
		RunE: func(cmd *cobra.Command, args []string) error {
            clientCtx := client.GetClientContextFromCmd(cmd)

            pageReq, err := client.ReadPageRequest(cmd.Flags())
            if err != nil {
                return err
            }

            queryClient := types.NewQueryClient(clientCtx)

            params := &types.QueryAllClaimRequest{
                Pagination: pageReq,
            }

            res, err := queryClient.ClaimAll(context.Background(), params)
            if err != nil {
                return err
            }

            return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

    return cmd
}

func CmdShowClaim() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-claim [id]",
		Short: "shows a claim",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
            clientCtx := client.GetClientContextFromCmd(cmd)

            queryClient := types.NewQueryClient(clientCtx)

            params := &types.QueryGetClaimRequest{
                Id: args[0],
            }

            res, err := queryClient.Claim(context.Background(), params)
            if err != nil {
                return err
            }

            return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

    return cmd
}
