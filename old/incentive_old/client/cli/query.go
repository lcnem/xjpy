package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/lcnem/jpyx/x/incentive/types"
)

const (
	flagOwner = "owner"
	flagType  = "type"
)

// GetQueryCmd returns the cli query commands for the incentive module
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	incentiveQueryCmd := &cobra.Command{
		Use:   types.ModuleName,
		Short: "Querying commands for the incentive module",
	}

	incentiveQueryCmd.AddCommand(flags.GetCommands(
		queryParamsCmd(queryRoute, cdc),
		queryRewardsCmd(queryRoute, cdc),
	)...)

	return incentiveQueryCmd
}

func queryRewardsCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rewards",
		Short: "query claimable rewards",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query rewards with optional flags for owner and type

			Example:
			$ %s query %s rewards
			$ %s query %s rewards --owner jpyx15qdefkmwswysgg4qxgqpqr35k3m49pkx2jdfnw
			$ %s query %s rewards --type jpyx-minting
			`,
				version.ClientName, types.ModuleName, version.ClientName, types.ModuleName,
				version.ClientName, types.ModuleName)),
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			page := viper.GetInt(flags.FlagPage)
			limit := viper.GetInt(flags.FlagLimit)
			strOwner := viper.GetString(flagOwner)
			strType := viper.GetString(flagType)

			// Prepare params for querier
			owner, err := sdk.AccAddressFromBech32(strOwner)
			if err != nil {
				return err
			}

			switch strings.ToLower(strType) {
			case "jpyx-minting":
				params := types.NewQueryJPYXMintingRewardsParams(page, limit, owner)
				claims, err := executeJPYXMintingRewardsQuery(queryRoute, cdc, cliCtx, params)
				if err != nil {
					return err
				}
				return cliCtx.PrintOutput(claims)
			default:
				paramsJPYXMinting := types.NewQueryJPYXMintingRewardsParams(page, limit, owner)
				jpyxMintingClaims, err := executeJPYXMintingRewardsQuery(queryRoute, cdc, cliCtx, paramsJPYXMinting)
				if err != nil {
					return err
				}
				if len(jpyxMintingClaims) > 0 {
					cliCtx.PrintOutput(jpyxMintingClaims)
				}
			}
			return nil
		},
	}
	cmd.Flags().String(flagOwner, "", "(optional) filter by owner address")
	cmd.Flags().String(flagType, "", "(optional) filter by reward type")
	cmd.Flags().Int(flags.FlagPage, 1, "pagination page rewards of to to query for")
	cmd.Flags().Int(flags.FlagLimit, 100, "pagination limit of rewards to query for")
	return cmd
}

func queryParamsCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "params",
		Short: "get the incentive module parameters",
		Long:  "Get the current global incentive module parameters.",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			// Query
			route := fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryGetParams)
			res, height, err := cliCtx.QueryWithData(route, nil)
			if err != nil {
				return err
			}
			cliCtx = cliCtx.WithHeight(height)

			// Decode and print results
			var params types.Params
			if err := cdc.UnmarshalJSON(res, &params); err != nil {
				return fmt.Errorf("failed to unmarshal params: %w", err)
			}
			return cliCtx.PrintOutput(params)
		},
	}
}

func executeJPYXMintingRewardsQuery(queryRoute string, cdc *codec.Codec, cliCtx context.CLIContext,
	params types.QueryJPYXMintingRewardsParams) (types.JPYXMintingClaims, error) {
	bz, err := cdc.MarshalJSON(params)
	if err != nil {
		return types.JPYXMintingClaims{}, err
	}

	route := fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryGetJPYXMintingRewards)
	res, height, err := cliCtx.QueryWithData(route, bz)
	if err != nil {
		return types.JPYXMintingClaims{}, err
	}

	cliCtx = cliCtx.WithHeight(height)

	var claims types.JPYXMintingClaims
	if err := cdc.UnmarshalJSON(res, &claims); err != nil {
		return types.JPYXMintingClaims{}, fmt.Errorf("failed to unmarshal claims: %w", err)
	}

	return claims, nil
}