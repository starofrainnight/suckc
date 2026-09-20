package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/starofrainnight/suckc/internal/frontend"
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
			debug, _ := cmd.Flags().GetBool("debug")

			result, err := frontend.ParseFile(file)
			if err != nil {
				return fmt.Errorf("parse %s: %w", file, err)
			}

			if debug {
				fmt.Println(result.Tree.ToStringTree(result.Tree.GetParser().GetRuleNames(), result.Tree.GetParser()))
				return nil
			}

			decls := 0
			if seq := result.Tree.DeclarationSeq(); seq != nil {
				decls = len(seq.AllDeclaration())
			}
			fmt.Printf("suckc: %s: parsed %d declarations\n", file, decls)
			return nil
		},
	}
	cmd.Flags().BoolP("debug", "d", false, "dump the parse tree")
	return cmd
}

func main() {
	if err := newRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
