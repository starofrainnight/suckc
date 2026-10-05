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
// re-inserted, so generation is a round trip for any source accepted by
// the frontend, except for the tokens listed by ParseResult.ToErase.
//
// subs maps token indexes to replacement text (sema's auto deduction);
// nil or empty means pure identity emission. A substitution applies to
// the listed token whatever channel it is on — a visible terminal as well
// as a hidden token between terminals. Generate seeds the map with
// ParseResult.ToErase, so the C-scope region markers and the gcc -E line
// markers are dropped from the output while the region body between them
// passes through byte for byte.
func Generate(res *frontend.ParseResult, subs map[int]string) (string, error) {
	if res == nil || res.Tree == nil {
		return "", errors.New("backend: no parse tree")
	}

	// Effective substitutions: the caller's replacements plus the tokens
	// that must vanish from the emitted C.
	repl := make(map[int]string, len(subs)+2*len(res.Regions)+len(res.LineMarkers))
	for idx, text := range subs {
		repl[idx] = text
	}
	for _, idx := range res.ToErase() {
		repl[idx] = ""
	}

	e := &emitter{tokens: res.Tokens, lastIndex: -1, subs: repl}

	walker := antlr.NewParseTreeWalker()
	walker.Walk(e, res.Tree)

	// Flush any trailing hidden tokens after the last terminal node
	// (up to, but not including, the EOF token).
	for i := e.lastIndex + 1; i < len(e.tokens)-1; i++ {
		e.write(i)
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
	subs      map[int]string // token index -> replacement text; "" = erase
}

// write emits token i, applying its substitution when one exists. Both the
// hidden-token backfill loops and the trailing flush go through it, so a
// substitution replaces a token regardless of its channel.
func (e *emitter) write(i int) {
	if repl, ok := e.subs[i]; ok {
		e.sb.WriteString(repl)
		return
	}
	e.sb.WriteString(e.tokens[i].GetText())
}

func (e *emitter) VisitTerminal(node antlr.TerminalNode) {
	tok := node.GetSymbol()
	if tok.GetTokenType() == antlr.TokenEOF {
		return
	}
	idx := tok.GetTokenIndex()

	// Re-insert hidden tokens between the previous terminal and this one.
	for i := e.lastIndex + 1; i < idx; i++ {
		e.write(i)
	}
	if repl, ok := e.subs[idx]; ok {
		e.sb.WriteString(repl)
	} else {
		e.sb.WriteString(tok.GetText())
	}
	e.lastIndex = idx
}
