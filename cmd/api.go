package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zhaolion/civitai-cli/civitai/api"
)

var apiRootCmd = &cobra.Command{
	Use:   "api",
	Short: "Interact with CivitAI.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func APICommand() *cobra.Command {
	apiRootCmd.PersistentFlags().BoolVarP(&argDebug, "debug", "", false, "enable debug mode")

	apiRootCmd.AddCommand(apiTokenSetCmd())
	apiRootCmd.AddCommand(apiTokenShowCmd())
	apiRootCmd.AddCommand(apiModelInfoShowCmd())
	apiRootCmd.AddCommand(apiModelVersionShowCmd())
	return apiRootCmd
}

func apiTokenSetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set_api_token",
		Short: "set api token for authentication.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("Please provide a token.")
			}
			token := strings.TrimSpace(args[0])
			if token == "" {
				fmt.Println("Please provide a none empty token.")
			}

			api.SetAPIToken(token)
		},
	}
	return cmd
}

func apiTokenShowCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "view_api_token",
		Short: "view which api token is currently set.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(api.GetAPITokenMask())
		},
	}
	return cmd
}

func apiModelInfoShowCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "model",
		Short: "view model info by model id",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			client := api.NewClient(api.GetAPIToken(),
				api.CivitaiClientOptionDebug(argDebug),
			)

			modelID := fmt.Sprintf("%d", argModelID)
			model, err := client.ModelInfoByID(modelID)
			if err != nil {
				panic(err)
			}

			_ = api.NewTerminal().PrintModelInfo(ctx, model, nil)
		},
	}
	cmd.PersistentFlags().Int64VarP(&argModelID, "mid", "", 0, "model id")
	_ = cmd.MarkFlagRequired("mid")

	return cmd
}

func apiModelVersionShowCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "model_ver",
		Short: "view model's version info by version id.",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			client := api.NewClient(api.GetAPIToken(),
				api.CivitaiClientOptionDebug(argDebug),
			)

			vID := fmt.Sprintf("%d", argVerID)
			ver, err := client.ModelVersionByID(vID)
			if err != nil {
				panic(err)
			}

			_ = api.NewTerminal().PrintModelVersionByID(ctx, ver, nil)
		},
	}
	cmd.PersistentFlags().Int64VarP(&argVerID, "vid", "", 0, "model version id")
	_ = cmd.MarkFlagRequired("vid")

	return cmd
}
