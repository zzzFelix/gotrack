package cmd

import (
	"fmt"
	"os"
)

func Execute() {
	if err := saveCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
