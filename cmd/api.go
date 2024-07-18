package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zhaolion/civitai-cli/civitai/api"
)

var apiRootCmd = &cobra.Command{
	Use: "api",
	Run: func(cmd *cobra.Command, args []string) {},
}

func APICommand() *cobra.Command {
	apiRootCmd.AddCommand(apiTokenSetCmd())
	apiRootCmd.AddCommand(apiTokenShowCmd())
	apiRootCmd.AddCommand(apiModelInfoShowCmd())
	apiRootCmd.AddCommand(apiModelVersionShowCmd())
	return apiRootCmd
}

func apiTokenSetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "set_api_token",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("Please provide a token")
			}

			api.SetAPIToken(args[0])
		},
	}
	return cmd
}

func apiTokenShowCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "show_api_token",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(api.GetAPITokenMask())
		},
	}
	return cmd
}

func apiModelInfoShowCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "model",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("Please provide a model identifier")
			}
			ctx := context.Background()
			client := api.NewClient(api.GetAPIToken(),
				api.CivitaiClientOptionDebug(*flagDebug),
			)

			model, err := client.ModelInfoByID(args[0])
			if err != nil {
				panic(err)
			}

			_ = api.NewTerminal().PrintModelInfo(ctx, model, nil)
		},
	}

	return cmd
}

func apiModelVersionShowCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "model_ver",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("Please provide a model version identifier")
			}
			ctx := context.Background()
			client := api.NewClient(api.GetAPIToken(),
				api.CivitaiClientOptionDebug(*flagDebug),
			)

			ver, err := client.ModelVersionByID(args[0])
			if err != nil {
				panic(err)
			}

			_ = api.NewTerminal().PrintModelVersionByID(ctx, ver, nil)
		},
	}

	return cmd
}
