package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
	"github.com/zhaolion/civitai-cli/civitai/api"
)

var (
	flagModelID    = flag.Int64("mid", 0, "model id")
	flagModelVerID = flag.Int64("vid", 0, "model version id")
	flagTargetDir  = flag.String("dir", "", "target directory")
)

var downloadRootCmd = &cobra.Command{
	Use: "download",
	Run: func(cmd *cobra.Command, args []string) {},
}

func DownloadCommand() *cobra.Command {
	flag.Parse()

	downloadRootCmd.AddCommand(apiModelDownloadCmd())
	downloadRootCmd.AddCommand(apiModelVerDownloadCmd())
	return downloadRootCmd
}

func apiModelDownloadCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "model",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			client := api.NewClient(api.GetAPIToken(),
				api.CivitaiClientOptionDebug(*flagDebug),
			)

			modelID, targetDir, err := parseApiModelDownloadArgs(args)
			if err != nil {
				panic(err)
			}

			err = client.ModelDownloadByID(ctx, modelID, targetDir)
			if err != nil {
				panic(err)
			}
		},
	}

	return cmd
}

func parseApiModelDownloadArgs(args []string) (modelID, targetDir string, err error) {
	// 要么全部 flags 都是 0，要么全部 flags 都是非 0
	if *flagModelID != 0 {
		modelID = fmt.Sprintf("%d", *flagModelID)
	}

	if *flagTargetDir != "" {
		targetDir = *flagTargetDir
	} else {
		targetDir = "."
	}

	// 没有设置 flag，使用 args
	// 0: modelID, 1: verID
	if modelID == "" && len(args) != 2 {
		return "", "", fmt.Errorf("model id is required")
	}
	if len(args) == 2 {
		modelID = args[0]
	}
	return modelID, targetDir, nil
}

func apiModelVerDownloadCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "model_ver",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			client := api.NewClient(api.GetAPIToken(),
				api.CivitaiClientOptionDebug(*flagDebug),
			)

			modelID, verID, targetDir, err := parseApiModelVerDownloadArgs(args)
			if err != nil {
				panic(err)
			}

			err = client.ModelVerDownloadByID(ctx, modelID, verID, targetDir)
			if err != nil {
				panic(err)
			}
		},
	}

	return cmd
}

func parseApiModelVerDownloadArgs(args []string) (modelID, verID, targetDir string, err error) {
	// 要么全部 flags 都是 0，要么全部 flags 都是非 0
	if *flagModelID != 0 {
		modelID = fmt.Sprintf("%d", *flagModelID)
	}
	if *flagModelVerID != 0 {
		verID = fmt.Sprintf("%d", *flagModelVerID)
	}
	if *flagTargetDir != "" {
		targetDir = *flagTargetDir
	} else {
		targetDir = "."
	}

	// 没有设置 flag，使用 args
	// 0: modelID, 1: verID
	if modelID == "" && verID == "" && len(args) != 2 {
		return "", "", "", fmt.Errorf("model id and version id are required")
	}
	if len(args) == 2 {
		modelID = args[0]
		verID = args[1]
	}
	return modelID, verID, targetDir, nil
}
