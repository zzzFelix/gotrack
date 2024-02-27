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
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 || len(args) > 4 {
				return saveCmd.Help()
			} else {
				return saveCmd.Execute()
			}
		},
	}
)
