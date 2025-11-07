package main

import (
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var RootCmd = &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}

	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
