package main

import (
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
	"github.com/zhaolion/civitai-cli/cmd"
)

func init() {
	rootCmd.AddCommand(
		cmd.APICommand(),
		cmd.DownloadCommand(),
	)
}

func main() {
	flag.Parse()

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

var rootCmd = &cobra.Command{
	Use:   "civitai-cli",
	Short: "civitai-cli is a simple Go client to batch download models from CivitAI with CivitAI Api V1.",
}
