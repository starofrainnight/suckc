// Package frontend reads *.suckc sources and produces an AST (parse tree)
// from the ANTLR4-generated SuckC parser.
package frontend

import (
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"

	"github.com/starofrainnight/suckc/internal/parser"
)

// ParseResult is the outcome of parsing a *.suckc file.
type ParseResult struct {
	Tree   *parser.TranslationUnitContext
	Tokens []antlr.Token
	// FileName is the path the source was read from; ParseFile sets it and
	// ParseSource leaves it empty. It is the default file reported by
	// SourceAt for tokens no line marker applies to.
	FileName string
	// Regions holds every C-scope region of the tree, sorted by StartTok
	// and non-overlapping (see region.go).
	Regions []Region
	// LineMarkers holds the surviving gcc -E line markers, in token order.
	LineMarkers []LineMarker
	// LinkageSpecToks holds the wrapper tokens (`extern`, the string
	// literal and both braces) of every linkage specification in the tree.
	LinkageSpecToks []int
	// CFuncs and CTypedefs hold what the C-scope regions declare, in source
	// order (see cdecl.go). Both are empty when the translation unit has no
	// region.
	CFuncs    []CFuncSig
	CTypedefs []CTypedef
}

// ParseFile reads the source file at path, lexes and parses it, and returns
// the translation-unit parse tree. It reports an error if the file cannot be
// read or if the source contains syntax errors.
func ParseFile(path string) (*ParseResult, error) {
	src, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return parseSource(string(src), path)
}

// ParseSource lexes and parses the given SuckC source text.
func ParseSource(src string) (*ParseResult, error) {
	return parseSource(src, "")
}

// parseSource is the shared implementation; fileName is recorded for
// diagnostics (empty when the source has no associated path).
//
// A lexer error aborts the parse just like a parser error: the lexer drops
// the offending characters, and identity emission would copy that silently
// dropped text into the generated C — `foo$bar` would become `foobar`, a
// rename that still compiles. Inside a C-scope region nobody else can ever
// notice (sema treats the body as opaque), so the failure must be loud here.
func parseSource(src, fileName string) (*ParseResult, error) {
	lexer := parser.NewSuckCLexer(antlr.NewInputStream(src))
	lexer.RemoveErrorListeners()
	el := &syntaxErrorListener{}
	lexer.AddErrorListener(el)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	p := parser.NewSuckCParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(el)

	tree := p.TranslationUnit().(*parser.TranslationUnitContext)

	res := &ParseResult{Tree: tree, Tokens: stream.GetAllTokens(), FileName: fileName}
	res.collectRegions()
	res.collectLineMarkers()
	res.collectLinkageSpecs()
	res.collectCDecls()

	if len(el.errors) > 0 {
		return nil, res.locatedError(el.errors[0])
	}
	return res, nil
}

// parseError is one lexer or parser diagnostic with its source position.
type parseError struct {
	line   int
	column int
	msg    string
}

// locatedError formats the first recorded lexer/parser error. When the
// source carries line markers (i.e. it is gcc -E output) the position is
// mapped through SourceAt, so an error inside a region reports the header
// the text came from instead of the parsed file. Marker-free sources keep
// the plain "line N:C msg" form.
func (r *ParseResult) locatedError(e parseError) error {
	if len(r.LineMarkers) > 0 {
		if idx, ok := r.tokenIndexAt(e.line, e.column); ok {
			file, line := r.SourceAt(idx)
			if file != "" {
				// SourceAt maps the token's line; carry the same marker
				// offset over to the recorded error line. Columns are
				// unaffected: a marker occupies whole lines of its own.
				line = e.line + (line - r.Tokens[idx].GetLine())
				return fmt.Errorf("%s:%d:%d: %s", file, line, e.column, e.msg)
			}
		}
	}
	return fmt.Errorf("line %d:%d %s", e.line, e.column, e.msg)
}

// tokenIndexAt returns the index of the first token at or after the given
// source position, which anchors a diagnostic into the token stream. ok is
// false only for an empty stream.
func (r *ParseResult) tokenIndexAt(line, column int) (int, bool) {
	for i, tok := range r.Tokens {
		if tok.GetLine() > line || (tok.GetLine() == line && tok.GetColumn() >= column) {
			return i, true
		}
	}
	if n := len(r.Tokens); n > 0 {
		return n - 1, true
	}
	return 0, false
}

type syntaxErrorListener struct {
	antlr.DefaultErrorListener
	errors []parseError
}

func (l *syntaxErrorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {
	l.errors = append(l.errors, parseError{line: line, column: column, msg: msg})
}
