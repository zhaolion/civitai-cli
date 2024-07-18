package cmd

import (
	flag "github.com/spf13/pflag"
)

var (
	flagDebug = flag.Bool("debug", false, "enable debug mode")
)
