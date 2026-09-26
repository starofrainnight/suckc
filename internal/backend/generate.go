// Package backend emits C/C++ source from a parsed SuckC translation unit.
//
// Phase 1: the source grammar is the full C/C++ grammar, so code generation
// is identity emission — the parse tree is walked and the text of every
// terminal node is written back out, with whitespace and comments re-inserted
// from the token stream. Later phases transform the tree and emit trimmed C.
package backend

import (
	"errors"
	"strings"

	"github.com/antlr4-go/antlr/v4"

	"github.com/starofrainnight/suckc/internal/frontend"
)

// Generate walks the parse tree of a parsed SuckC translation unit and
// returns the corresponding C/C++ source text.
//
// The output is faithful to the input: hidden-channel tokens (whitespace,
// comments, preprocessor directives) between consecutive terminals are
// re-inserted, so generation is a round trip for any source accepted by the
// frontend.
//
// subs maps token indexes to replacement text (sema's auto deduction);
// nil or empty means pure identity emission. Only the listed tokens are
// replaced; hidden-channel tokens between terminals stay unchanged.
func Generate(res *frontend.ParseResult, subs map[int]string) (string, error) {
	if res == nil || res.Tree == nil {
		return "", errors.New("backend: no parse tree")
	}

	e := &emitter{tokens: res.Tokens, lastIndex: -1, subs: subs}

	walker := antlr.NewParseTreeWalker()
	walker.Walk(e, res.Tree)

	// Flush any trailing hidden tokens after the last terminal node
	// (up to, but not including, the EOF token).
	for i := e.lastIndex + 1; i < len(e.tokens)-1; i++ {
		e.sb.WriteString(e.tokens[i].GetText())
	}

	return e.sb.String(), nil
}

// emitter collects source text while walking the parse tree. Terminal nodes
// are visited in source order; hidden-channel tokens between consecutive
// terminals are re-inserted so the emitted source matches the input.
type emitter struct {
	antlr.BaseParseTreeListener
	sb        strings.Builder
	tokens    []antlr.Token  // full token stream, indexed by token index
	lastIndex int            // token index of the last terminal emitted; -1 = none yet
	subs      map[int]string // token index -> replacement text; nil = identity
}

func (e *emitter) VisitTerminal(node antlr.TerminalNode) {
	tok := node.GetSymbol()
	if tok.GetTokenType() == antlr.TokenEOF {
		return
	}
	idx := tok.GetTokenIndex()

	// Re-insert hidden tokens between the previous terminal and this one.
	for i := e.lastIndex + 1; i < idx; i++ {
		e.sb.WriteString(e.tokens[i].GetText())
	}
	if repl, ok := e.subs[idx]; ok {
		e.sb.WriteString(repl)
	} else {
		e.sb.WriteString(tok.GetText())
	}
	e.lastIndex = idx
}
