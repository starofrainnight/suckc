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
}

// ParseFile reads the source file at path, lexes and parses it, and returns
// the translation-unit parse tree. It reports an error if the file cannot be
// read or if the source contains syntax errors.
func ParseFile(path string) (*ParseResult, error) {
	src, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return ParseSource(string(src))
}

// ParseSource lexes and parses the given SuckC source text.
func ParseSource(src string) (*ParseResult, error) {
	lexer := parser.NewSuckCLexer(antlr.NewInputStream(src))
	stream := antlr.NewCommonTokenStream(lexer, 0)

	p := parser.NewSuckCParser(stream)
	p.RemoveErrorListeners()
	el := &syntaxErrorListener{}
	p.AddErrorListener(el)

	tree := p.TranslationUnit().(*parser.TranslationUnitContext)

	if len(el.errors) > 0 {
		return nil, fmt.Errorf("%s", el.errors[0])
	}
	return &ParseResult{Tree: tree, Tokens: stream.GetAllTokens()}, nil
}

type syntaxErrorListener struct {
	antlr.DefaultErrorListener
	errors []string
}

func (l *syntaxErrorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {
	l.errors = append(l.errors, fmt.Sprintf("line %d:%d %s", line, column, msg))
}
