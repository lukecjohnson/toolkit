package main

import (
	"os"

	"github.com/lukecjohnson/toolkit/commands"
)

func main() {
	if err := commands.RootCommand.Execute(); err != nil {
		os.Exit(1)
	}
}
