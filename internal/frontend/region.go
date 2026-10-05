package frontend

import (
	"regexp"
	"sort"
	"strconv"

	"github.com/antlr4-go/antlr/v4"

	"github.com/starofrainnight/suckc/internal/parser"
)

// Region is a C-scope region body: the inclusive token-index range between
// the two marker tokens, excluding the markers themselves.
type Region struct {
	StartTok int // index of the first token AFTER CScopeBegin
	EndTok   int // index of the last token BEFORE CScopeEnd
}

// LineMarker is a surviving `gcc -E` line marker: `# <linenum> "<file>" [flags]`.
type LineMarker struct {
	TokIdx int
	File   string
	Line   int
}

// lineMarkerPattern matches a gcc -E line marker such as
// `# 0 "<built-in>"`, `# 1 "vendor.h" 1` or `# 1 "x.h" 1 3 4`. The line
// number is capped at nine digits so it always fits an int.
var lineMarkerPattern = regexp.MustCompile(`^#[ \t]+(\d{1,9})[ \t]+"([^"]*)"(?:[ \t]+\d+)*[ \t]*\r?$`)

// collectRegions records every cImportRegion of the parse tree in source
// order. The grammar makes the result sorted and non-overlapping for free: a
// region body ends at the first CScopeEnd, so a marker inside a region is a
// syntax error and regions can neither nest nor interleave.
func (r *ParseResult) collectRegions() {
	if r.Tree == nil {
		return
	}
	var walk func(n antlr.Tree)
	walk = func(n antlr.Tree) {
		if ctx, ok := n.(*parser.CImportRegionContext); ok {
			begin := ctx.CScopeBegin()
			end := ctx.CScopeEnd()
			if begin != nil && end != nil {
				b := begin.GetSymbol().GetTokenIndex()
				e := end.GetSymbol().GetTokenIndex()
				if e >= b {
					r.Regions = append(r.Regions, Region{StartTok: b + 1, EndTok: e - 1})
				}
			}
			return
		}
		if pr, ok := n.(antlr.ParserRuleContext); ok {
			for _, ch := range pr.GetChildren() {
				walk(ch)
			}
		}
	}
	walk(r.Tree)
}

// collectLineMarkers scans the whole token stream — hidden channel
// included — for preprocessor line markers that survived gcc -E. The
// `Directive` lexer rule hides them, so only their text can identify them.
func (r *ParseResult) collectLineMarkers() {
	for _, tok := range r.Tokens {
		if tok.GetChannel() == antlr.TokenDefaultChannel {
			continue
		}
		m := lineMarkerPattern.FindStringSubmatch(tok.GetText())
		if m == nil {
			continue
		}
		line, err := strconv.Atoi(m[1])
		if err != nil {
			continue // unreachable: the pattern caps the digits at nine
		}
		r.LineMarkers = append(r.LineMarkers, LineMarker{
			TokIdx: tok.GetTokenIndex(),
			File:   m[2],
			Line:   line,
		})
	}
}

// InRegion reports whether tokIdx falls inside a C-scope region body. The
// two marker tokens themselves are outside, and so is everything beyond a
// region. Binary search over the sorted, non-overlapping Regions.
func (r *ParseResult) InRegion(tokIdx int) bool {
	i := sort.Search(len(r.Regions), func(i int) bool {
		return r.Regions[i].EndTok >= tokIdx
	})
	if i == len(r.Regions) {
		return false
	}
	rg := r.Regions[i]
	return tokIdx >= rg.StartTok && tokIdx <= rg.EndTok
}

// SourceAt maps a token index back to the file and line it originated
// from. Line markers are replayed: the last marker before tokIdx names the
// file, and its declared line applies to the output line following the
// marker, so lines shift by (markerLine - markerOwnLine - 1). Flags (1 =
// entering, 2 = returning, 3 = system header, 4 = extern "C") are not
// interpreted: every marker already carries the name of the file its text
// belongs to, which is both simpler and correct for flagless markers such
// as `# 1 "first.suckc"`. With no applicable marker the results default to
// FileName and the token's own line.
func (r *ParseResult) SourceAt(tokIdx int) (file string, line int) {
	file = r.FileName
	if tokIdx < 0 || tokIdx >= len(r.Tokens) {
		return file, 0
	}
	line = r.Tokens[tokIdx].GetLine()
	i := sort.Search(len(r.LineMarkers), func(i int) bool {
		return r.LineMarkers[i].TokIdx >= tokIdx
	})
	if i == 0 {
		return file, line
	}
	m := r.LineMarkers[i-1]
	line += m.Line - r.Tokens[m.TokIdx].GetLine() - 1
	return m.File, line
}

// linkageWrapper returns the four wrapper terminals of a linkage
// specification, or nil for any other node. SuckC and its C dialect have
// separate rules for this (linkageSpecification and
// cCLinkageSpecification) but identical accessor names, so both are listed.
func linkageWrapper(n antlr.Tree) []antlr.TerminalNode {
	switch ctx := n.(type) {
	case *parser.LinkageSpecificationContext:
		return []antlr.TerminalNode{ctx.Extern(), ctx.StringLiteral(), ctx.LeftBrace(), ctx.RightBrace()}
	case *parser.CCLinkageSpecificationContext:
		return []antlr.TerminalNode{ctx.Extern(), ctx.StringLiteral(), ctx.LeftBrace(), ctx.RightBrace()}
	}
	return nil
}

// collectLinkageSpecs records the wrapper tokens of every linkage
// specification in the tree, inside C regions or not.
//
// The wrapper is pure noise in the emitted C. The backend is an identity
// emitter, so without this the four tokens survive into the `.c` and C89
// rejects them: `extern "C" { int z; }` yields
// "error: expected identifier or '(' before string constant". The braces
// carry no meaning at file scope in C -- a declaration there is already
// external -- so dropping them leaves exactly the intended semantics.
func (r *ParseResult) collectLinkageSpecs() {
	if r.Tree == nil {
		return
	}
	var walk func(n antlr.Tree)
	walk = func(n antlr.Tree) {
		for _, t := range linkageWrapper(n) {
			if t != nil {
				r.LinkageSpecToks = append(r.LinkageSpecToks, t.GetSymbol().GetTokenIndex())
			}
		}
		if pr, ok := n.(antlr.ParserRuleContext); ok {
			for _, ch := range pr.GetChildren() {
				walk(ch)
			}
		}
	}
	walk(r.Tree)
}

// ToErase returns the token indexes the backend must drop from the emitted
// C: both markers of every region (visible `#pragma c_scope_*` tokens,
// which Keil/IAR reject as unknown pragmas), every line marker (hidden
// `# N "file"` directives, which are not valid C89 outside `#line`), and the
// wrapper tokens of every linkage specification. The result is sorted
// ascending.
func (r *ParseResult) ToErase() []int {
	out := make([]int, 0, 2*len(r.Regions)+len(r.LineMarkers)+len(r.LinkageSpecToks))
	for _, rg := range r.Regions {
		out = append(out, rg.StartTok-1, rg.EndTok+1)
	}
	for _, m := range r.LineMarkers {
		out = append(out, m.TokIdx)
	}
	out = append(out, r.LinkageSpecToks...)
	sort.Ints(out)
	return out
}
