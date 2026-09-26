// Package sema performs semantic analysis on a parsed SuckC translation
// unit: it deduces the concrete type of block-scope auto declarations and
// returns a token-index substitution map for the backend emitter.
package sema

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/starofrainnight/suckc/internal/frontend"
)

// Options configures semantic analysis.
type Options struct {
	TargetBits int    // 16, 32 or 64; 0 means host width
	FileName   string // source file name used in diagnostics
}

// resolved returns opts with TargetBits defaulted to the host width.
func (o Options) resolved() Options {
	if o.TargetBits == 0 {
		o.TargetBits = strconv.IntSize
	}
	return o
}

// Subs maps a token index to the replacement text for that token.
type Subs map[int]string

// Error is a deduction diagnostic in file:line:col form.
type Error struct {
	File string
	Line int
	Col  int // 1-based
	Msg  string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s:%d:%d: cannot deduce type for 'auto': %s",
		e.File, e.Line, e.Col, e.Msg)
}

// Analyze deduces auto types in the parsed translation unit and returns the
// substitutions the backend must apply during emission.
func Analyze(res *frontend.ParseResult, opts Options) (Subs, error) {
	if res == nil || res.Tree == nil {
		return nil, errors.New("sema: no parse tree")
	}
	opts = opts.resolved()
	switch opts.TargetBits {
	case 16, 32, 64:
	default:
		return nil, fmt.Errorf("sema: invalid target bits %d (must be 16, 32, or 64)", opts.TargetBits)
	}
	a := newAnalyzer(opts)
	if err := pass1(res, a); err != nil {
		return nil, err
	}
	if err := pass2(res, a); err != nil {
		return nil, err
	}
	if err := scanForbiddenAutos(res, a); err != nil {
		return nil, err
	}
	return a.Subs, nil
}
