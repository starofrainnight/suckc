package sema

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"

	"github.com/starofrainnight/suckc/internal/parser"
)

// maxAutoDepth bounds the lazy deduction recursion; deeper input is a
// pathological cycle rather than a real program.
const maxAutoDepth = 64

// pass1b deduces every file-scope auto that pass 1 recorded.
func pass1b(a *Analyzer) error {
	if a.pendingByName == nil {
		a.pendingByName = make(map[string]*pendingAuto, len(a.pendingAutos))
	}
	for _, pa := range a.pendingAutos {
		if pa.Name == "" {
			continue
		}
		if _, seen := a.pendingByName[pa.Name]; !seen {
			a.pendingByName[pa.Name] = pa
		}
	}
	for _, pa := range a.pendingAutos {
		if err := deduceFileAuto(a, pa, 0); err != nil {
			return err
		}
	}
	return nil
}

func deduceFileAuto(a *Analyzer, pa *pendingAuto, depth int) error {
	switch pa.State {
	case autoDone:
		return nil
	case autoInProgress:
		return fmt.Errorf("circular dependency on '%s'", pa.Name)
	}
	if depth > maxAutoDepth {
		return fmt.Errorf("circular dependency on '%s'", pa.Name)
	}
	pa.State = autoInProgress
	if err := deduceDependencies(a, pa, depth); err != nil {
		return err
	}
	if err := deduceAuto(a, pa.Decl, true); err != nil {
		return err
	}
	pa.State = autoDone
	return nil
}

// deduceDependencies deduces the other file-scope autos that pa's
// initializer mentions, so any-order references work. Collecting every
// identifier rather than only the ones the expression rules require is a
// safe over-approximation: deducing an unrelated auto early is
// idempotent, because State makes each declaration happen once.
func deduceDependencies(a *Analyzer, pa *pendingAuto, depth int) error {
	initDecl := pa.Decl.InitDeclarator()
	if initDecl == nil || initDecl.Initializer() == nil {
		return nil
	}
	var names []string
	collectIdentifiers(initDecl.Initializer(), &names)
	for _, name := range names {
		dep, ok := a.pendingByName[name]
		if !ok || dep == pa {
			continue
		}
		if err := deduceFileAuto(a, dep, depth+1); err != nil {
			return err
		}
	}
	return nil
}

// collectIdentifiers appends the text of every identifier token in node.
func collectIdentifiers(node antlr.Tree, out *[]string) {
	if tok, ok := node.(antlr.TerminalNode); ok {
		if tok.GetSymbol().GetTokenType() == parser.SuckCParserIdentifier {
			*out = append(*out, tok.GetSymbol().GetText())
		}
		return
	}
	if ctx, ok := node.(antlr.ParserRuleContext); ok {
		for _, child := range ctx.GetChildren() {
			collectIdentifiers(child, out)
		}
	}
}
