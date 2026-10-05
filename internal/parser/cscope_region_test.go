package parser

import (
	"strconv"
	"strings"
	"testing"

	"github.com/antlr4-go/antlr/v4"
)

// lexAll returns every token for src, including hidden-channel ones.
func lexAll(src string) []antlr.Token {
	stream := antlr.NewCommonTokenStream(NewSuckCLexer(antlr.NewInputStream(src)), 0)
	stream.Fill()
	return stream.GetAllTokens()
}

type errorRecorder struct {
	antlr.DefaultErrorListener
	msgs []string
}

func (r *errorRecorder) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {
	r.msgs = append(r.msgs, "line "+strconv.Itoa(line)+":"+strconv.Itoa(column)+" "+msg)
}

// parseWithErrors parses src as a translation unit and returns the syntax
// errors plus the parse-tree dump.
func parseWithErrors(src string) (errs []string, dump string) {
	lexer := NewSuckCLexer(antlr.NewInputStream(src))
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := NewSuckCParser(stream)
	rec := &errorRecorder{}
	p.RemoveErrorListeners()
	p.AddErrorListener(rec)
	tree := p.TranslationUnit()
	return rec.msgs, tree.ToStringTree(p.GetRuleNames(), p)
}

// The two C-scope markers must lex as VISIBLE tokens on the default channel.
// If a marker ever falls back to the hidden Directive rule the whole region
// feature silently disappears: ANTLR resolves equal-length matches by rule
// order and reports no warning, so this assertion is the regression guard.
func TestCScopeMarkersOnDefaultChannel(t *testing.T) {
	src := "#pragma c_scope_begin\nvoid f(void) {}\n#pragma c_scope_end\n"
	want := map[int]string{
		SuckCLexerCScopeBegin: "#pragma c_scope_begin",
		SuckCLexerCScopeEnd:   "#pragma c_scope_end",
	}
	seen := map[int]bool{}
	for _, tok := range lexAll(src) {
		tt := tok.GetTokenType()
		text, isMarker := want[tt]
		if !isMarker {
			continue
		}
		seen[tt] = true
		if ch := tok.GetChannel(); ch != antlr.TokenDefaultChannel {
			t.Errorf("token %q channel = %d, want default channel %d", tok.GetText(), ch, antlr.TokenDefaultChannel)
		}
		if got := tok.GetText(); got != text {
			t.Errorf("token text = %q, want %q", got, text)
		}
	}
	for tt, text := range want {
		if !seen[tt] {
			t.Errorf("marker %q (token type %d) not found on any channel as a visible token", text, tt)
		}
	}
}

// The markers tolerate trailing blanks and a CR, because Directive would
// otherwise match further and silently win the longest-match tie — which
// would leave the region unopened. Real trailing content must still fall
// through to Directive and stay a hidden preprocessor directive.
func TestCScopeMarkerTrailingBlanksAndCRLF(t *testing.T) {
	for _, src := range []string{
		"#pragma c_scope_begin\n",
		"#pragma c_scope_begin \n",
		"#pragma c_scope_begin   \t \n",
		"#pragma c_scope_end \n",
		"#pragma c_scope_begin\r\nvoid f(void) {}\r\n#pragma c_scope_end\r\n",
		"#pragma   c_scope_begin\t\n",
	} {
		found := false
		for _, tok := range lexAll(src) {
			if tok.GetTokenType() == SuckCLexerCScopeBegin || tok.GetTokenType() == SuckCLexerCScopeEnd {
				found = true
				if ch := tok.GetChannel(); ch != antlr.TokenDefaultChannel {
					t.Errorf("%q: marker %q on channel %d, want default", src, tok.GetText(), ch)
				}
			}
		}
		if !found {
			t.Errorf("%q: expected a marker token, got none (trailing blanks/CR must not break the marker)", src)
		}
	}
}

func TestCScopeMarkerIsStrict(t *testing.T) {
	for _, src := range []string{
		"#pragma c_scope_begin trailing\n",
		"#pragma c_scope_begin /* x */\n",
		"#pragma c_scope_end trailing\n",
		"#pragma c_scope_begin_ish\n",
	} {
		for _, tok := range lexAll(src) {
			tt := tok.GetTokenType()
			if tt == SuckCLexerCScopeBegin || tt == SuckCLexerCScopeEnd {
				t.Errorf("%q: marker token %q lexed despite trailing content", src, tok.GetText())
			}
		}
	}
}

// Regions must parse at file scope and inside a function body, and the tree
// must contain an actual cImportRegion node (ANTLR picks the first matching
// alternative silently, so "no error" alone would not prove the region rule
// matched).
func TestCScopeRegionParses(t *testing.T) {
	for _, src := range []string{
		"#pragma c_scope_begin\n#include <stdio.h>\nvoid hello(void) { printf(\"hi\\n\"); }\n#pragma c_scope_end\n",
		"void hello(void);\nvoid myMain(void) {\n#pragma c_scope_begin\nhello();\n#pragma c_scope_end\nhello();\n}\n",
		"#pragma c_scope_begin\nint x;\n#pragma c_scope_end\nvoid f(void) {\n#pragma c_scope_begin\nx++;\n#pragma c_scope_end\n}\n",
		// Body is all hidden-channel text: cRawRegion must accept zero visible tokens.
		"#pragma c_scope_begin\n#include <stdio.h>\n#pragma c_scope_end\nvoid f(void) {}\n",
		"#pragma c_scope_begin\n#pragma c_scope_end\nvoid f(void) {}\n",
	} {
		errs, dump := parseWithErrors(src)
		if len(errs) > 0 {
			t.Errorf("%q: unexpected syntax errors: %v", src, errs)
			continue
		}
		if !strings.Contains(dump, "cImportRegion") {
			t.Errorf("%q: tree has no cImportRegion node; dump: %s", src, dump)
		}
	}
}

// An unterminated region must fail with a syntax error (the raw-token loop
// excludes EOF so it stops at end of input instead of hanging), and nested
// regions are rejected because the first cScopeEnd closes the outer region.
func TestCScopeRegionErrors(t *testing.T) {
	for name, src := range map[string]string{
		"unterminated":       "#pragma c_scope_begin\nvoid f(void) {}\n",
		"unterminated-block": "void f(void) {\n#pragma c_scope_begin\ng();\n",
		"nested":             "#pragma c_scope_begin\nint x;\n#pragma c_scope_end\n#pragma c_scope_end\n",
		"dangling end":       "void f(void) {}\n#pragma c_scope_end\n",
	} {
		errs, _ := parseWithErrors(src)
		if len(errs) == 0 {
			t.Errorf("%s: expected syntax error, got none (src=%q)", name, src)
		}
	}
}
