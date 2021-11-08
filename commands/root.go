package commands

import (
	"github.com/spf13/cobra"
)

var version = "0.0.0"

var RootCommand = &cobra.Command{
	Use:          "toolkit",
	SilenceUsage: true,
	Version:      version,
}

func init() {
	RootCommand.CompletionOptions.DisableDefaultCmd = true
	RootCommand.SetVersionTemplate("{{.Version}}\n")
	RootCommand.AddCommand(serveCommand)
}
