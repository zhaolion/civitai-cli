package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zhaolion/civitai-cli/civitai/api"
)

var (
	argTargetDir string
)

var downloadRootCmd = &cobra.Command{
	Use:   "download",
	Short: "Download files from CivitAI",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func DownloadCommand() *cobra.Command {
	downloadRootCmd.PersistentFlags().BoolVarP(&argDebug, "debug", "", false, "enable debug mode")
	downloadRootCmd.AddCommand(apiModelDownloadCmd())
	downloadRootCmd.AddCommand(apiModelVerDownloadCmd())
	return downloadRootCmd
}

func apiModelDownloadCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "model",
		Short: "download files in model from CivitAI",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			client := api.NewClient(api.GetAPIToken(),
				api.CivitaiClientOptionDebug(argDebug),
			)

			modelID := fmt.Sprintf("%d", argModelID)
			err := client.ModelDownloadByID(ctx, modelID, argTargetDir)
			if err != nil {
				panic(err)
			}
		},
	}
	cmd.PersistentFlags().Int64VarP(&argModelID, "mid", "", 0, "model id")
	cmd.PersistentFlags().StringVarP(&argTargetDir, "dir", "", ".", "target dir, default is current dir")
	_ = cmd.MarkFlagRequired("mid")
	return cmd
}

func apiModelVerDownloadCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "model_ver",
		Short: "download files in one model's version from CivitAI",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			client := api.NewClient(api.GetAPIToken(),
				api.CivitaiClientOptionDebug(argDebug),
			)

			modelID := fmt.Sprintf("%d", argModelID)
			verID := fmt.Sprintf("%d", argVerID)
			err := client.ModelVerDownloadByID(ctx, modelID, verID, argTargetDir)
			if err != nil {
				panic(err)
			}
		},
	}
	cmd.PersistentFlags().Int64VarP(&argModelID, "mid", "", 0, "model id")
	cmd.PersistentFlags().Int64VarP(&argVerID, "vid", "", 0, "model version id")
	cmd.PersistentFlags().StringVarP(&argTargetDir, "dir", "", ".", "target dir, default is current dir")
	_ = cmd.MarkFlagRequired("mid")
	_ = cmd.MarkFlagRequired("vid")

	return cmd
}
