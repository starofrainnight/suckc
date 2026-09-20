package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "suckc",
		Short:   "SuckC language transpiler",
		Version: "0.1.0",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			file := args[0]
			info, err := os.Stat(file)
			if err != nil {
				return fmt.Errorf("file not found: %s", file)
			}
			if info.IsDir() {
				return fmt.Errorf("file not found: %s", file)
			}
			fmt.Printf("suckc: %s: transpilation not implemented yet\n", file)
			return nil
		},
	}
	cmd.Flags().BoolP("debug", "d", false, "enable debug output")
	return cmd
}

func main() {
	if err := newRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
