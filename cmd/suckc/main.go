package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/starofrainnight/suckc/internal/backend"
	"github.com/starofrainnight/suckc/internal/frontend"
	"github.com/starofrainnight/suckc/internal/sema"
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

			bits, err := cmd.Flags().GetInt("target-bits")
			if err != nil {
				return err
			}
			if bits != 16 && bits != 32 && bits != 64 {
				return fmt.Errorf("invalid --target-bits %d (must be 16, 32, or 64)", bits)
			}

			subs, err := sema.Analyze(result, sema.Options{TargetBits: bits, FileName: file})
			if err != nil {
				return err
			}

			src, err := backend.Generate(result, subs)
			if err != nil {
				return fmt.Errorf("generate %s: %w", file, err)
			}

			out := strings.TrimSuffix(file, filepath.Ext(file)) + ".c"
			if err := os.WriteFile(out, []byte(src), 0o644); err != nil {
				return fmt.Errorf("write %s: %w", out, err)
			}

			fmt.Printf("suckc: %s: generated %s\n", file, out)
			return nil
		},
	}
	cmd.Flags().BoolP("debug", "d", false, "dump the parse tree")
	cmd.Flags().Int("target-bits", strconv.IntSize,
		"target integer width in bits: 16, 32, or 64 (default: host)")
	return cmd
}

func main() {
	if err := newRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
