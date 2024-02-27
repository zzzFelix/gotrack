package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var (
	rootCmd = &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 || len(args) > 4 {
				saveCmd.Help()
			} else {
				saveCmd.Execute()
			}
		},
	}
)
