package frontend

import (
	"strings"
	"testing"

	"github.com/starofrainnight/suckc/internal/parser"
)

// tokenIndexOf returns the index of the nth (0-based) token whose text is
// exactly text.
func tokenIndexOf(t *testing.T, res *ParseResult, text string, nth int) int {
	t.Helper()
	seen := 0
	for _, tok := range res.Tokens {
		if tok.GetText() == text {
			if seen == nth {
				return tok.GetTokenIndex()
			}
			seen++
		}
	}
	t.Fatalf("token %q (#%d) not found in stream", text, nth)
	return -1
}

func TestInRegionCoversBodyOnly(t *testing.T) {
	src := "int a = 0;\n#pragma c_scope_begin\nint rb = 1;\n#pragma c_scope_end\nvoid f(void) {}\n"
	res, err := ParseSource(src)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if len(res.Regions) != 1 {
		t.Fatalf("regions = %d, want 1", len(res.Regions))
	}
	begin := tokenIndexOf(t, res, "#pragma c_scope_begin", 0)
	end := tokenIndexOf(t, res, "#pragma c_scope_end", 0)
	rg := res.Regions[0]
	if rg.StartTok != begin+1 || rg.EndTok != end-1 {
		t.Errorf("region = %+v, want {StartTok:%d EndTok:%d}", rg, begin+1, end-1)
	}

	if res.InRegion(begin) {
		t.Error("CScopeBegin marker reported inside its own region")
	}
	if res.InRegion(end) {
		t.Error("CScopeEnd marker reported inside its own region")
	}
	if !res.InRegion(begin + 1) {
		t.Error("first body token not reported inside the region")
	}
	if !res.InRegion(end - 1) {
		t.Error("last body token not reported inside the region")
	}
	if res.InRegion(0) {
		t.Error("token before the region reported inside it")
	}
	if res.InRegion(end + 1) {
		t.Error("token after the region reported inside it")
	}
}

func TestRegionsSortedAndNonOverlapping(t *testing.T) {
	src := "#pragma c_scope_begin\nint r1;\n#pragma c_scope_end\n" +
		"void f(void) {\n#pragma c_scope_begin\nr1++;\n#pragma c_scope_end\n}\n"
	res, err := ParseSource(src)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if len(res.Regions) != 2 {
		t.Fatalf("regions = %d, want 2 (file scope and block scope)", len(res.Regions))
	}
	prev := res.Regions[0]
	for i, rg := range res.Regions[1:] {
		if rg.StartTok <= prev.StartTok {
			t.Errorf("region %d not sorted by StartTok: %+v after %+v", i+1, rg, prev)
		}
		if rg.StartTok <= prev.EndTok {
			t.Errorf("region %d overlaps previous: %+v vs %+v", i+1, rg, prev)
		}
		prev = rg
	}
	for i, rg := range res.Regions {
		if !res.InRegion(rg.StartTok) || !res.InRegion(rg.EndTok) {
			t.Errorf("region %d body not reported inside: %+v", i, rg)
		}
	}
}

// Regions cannot nest, and the parse says so instead of quietly coping.
// An earlier version of this file asserted the opposite: when the region
// body was a bare token loop, an inner CScopeBegin was swallowed as body
// text. That was not merely cosmetic. ToErase derives the marker token
// indexes to blank from the outer region's body range (StartTok-1 and
// EndTok+1), so an inner marker sitting inside a body was never erased and
// leaked a bare "#pragma c_scope_begin" into the generated C. A parsed
// body rejects the token outright, which is both correct and safer.
func TestNestedMarkerIsRejected(t *testing.T) {
	src := "#pragma c_scope_begin\n#pragma c_scope_begin\nint x;\n#pragma c_scope_end\nvoid f(void) {}\n"
	res, err := ParseSource(src)
	if err == nil {
		t.Fatalf("parse accepted a nested marker, regions = %d, want a syntax error", len(res.Regions))
	}
	if !strings.Contains(err.Error(), "c_scope_begin") {
		t.Errorf("error should name the offending marker, got: %v", err)
	}

	// A well-formed region still yields exactly one, and no marker token
	// may survive inside a body -- that is the property the old test was
	// accidentally relying on the token loop for.
	ok := "#pragma c_scope_begin\nint x;\n#pragma c_scope_end\nvoid f(void) {}\n"
	res, err = ParseSource(ok)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if len(res.Regions) != 1 {
		t.Fatalf("regions = %d, want 1", len(res.Regions))
	}
	for _, tok := range res.Tokens {
		switch tok.GetTokenType() {
		case parser.SuckCLexerCScopeBegin, parser.SuckCLexerCScopeEnd:
			idx := tok.GetTokenIndex()
			if res.InRegion(idx) {
				t.Errorf("token %d is a marker inside a region body; it would leak into the output", idx)
			}
		}
	}
}

// Line markers come from real `gcc -E -x c` output, not hand-written text.
// Fixture command (run from /tmp/opencode/cscope):
//
//	gcc -E -x c first.suckc > internal/frontend/testdata/region_line_markers.c
func TestLineMarkersFromRealGccPreprocess(t *testing.T) {
	res, err := ParseFile("testdata/region_line_markers.c")
	if err != nil {
		t.Fatalf("parse fixture: %v", err)
	}
	want := []LineMarker{
		{File: "first.suckc", Line: 0},
		{File: "<built-in>", Line: 0},
		{File: "<command-line>", Line: 0},
		{File: "/usr/include/stdc-predef.h", Line: 1},
		{File: "<command-line>", Line: 0},
		{File: "first.suckc", Line: 1},
		{File: "vendor.h", Line: 1},
		{File: "first.suckc", Line: 3},
	}
	if len(res.LineMarkers) != len(want) {
		t.Fatalf("line markers = %d, want %d: %+v", len(res.LineMarkers), len(want), res.LineMarkers)
	}
	prevIdx := -1
	for i, m := range res.LineMarkers {
		if m.File != want[i].File || m.Line != want[i].Line {
			t.Errorf("marker %d = (%d, %q, %d), want (_, %q, %d)",
				i, m.TokIdx, m.File, m.Line, want[i].File, want[i].Line)
		}
		if m.TokIdx <= prevIdx {
			t.Errorf("marker %d at index %d is out of token order (previous %d)", i, m.TokIdx, prevIdx)
		}
		prevIdx = m.TokIdx
	}
}

func TestSourceAtFollowsLineMarkers(t *testing.T) {
	res, err := ParseFile("testdata/region_line_markers.c")
	if err != nil {
		t.Fatalf("parse fixture: %v", err)
	}

	// Before the first marker the defaults apply: FileName and the token's
	// own line.
	file, line := res.SourceAt(0)
	if file != "testdata/region_line_markers.c" || line != 1 {
		t.Errorf("SourceAt(0) = (%q, %d), want (%q, 1)", file, line, "testdata/region_line_markers.c")
	}

	// vendor_len sits on line 6 of vendor.h (the header's own numbering).
	idx := tokenIndexOf(t, res, "vendor_len", 0)
	file, line = res.SourceAt(idx)
	if file != "vendor.h" || line != 6 {
		t.Errorf("SourceAt(vendor_len) = (%q, %d), want (vendor.h, 6)", file, line)
	}

	// The return marker puts the closing pragma back in first.suckc line 3.
	end := tokenIndexOf(t, res, "#pragma c_scope_end", 0)
	file, line = res.SourceAt(end)
	if file != "first.suckc" || line != 3 {
		t.Errorf("SourceAt(c_scope_end) = (%q, %d), want (first.suckc, 3)", file, line)
	}
}

// A lexer error is fatal: the dropped characters would otherwise reach the
// generated C through identity emission, silently renaming identifiers.
// Region bodies are the dangerous case — sema never looks inside them.
func TestLexerErrorAbortsParse(t *testing.T) {
	cases := []string{
		"int $x = 1;\n",
		"#pragma c_scope_begin\nint $x = 1;\n#pragma c_scope_end\nvoid f(void) {}\n",
	}
	for _, src := range cases {
		res, err := ParseSource(src)
		if err == nil {
			t.Errorf("%q: expected lexer error, got none", src)
			continue
		}
		if res != nil {
			t.Errorf("%q: want nil result when the parse fails", src)
		}
		if !strings.Contains(err.Error(), "token recognition error") {
			t.Errorf("%q: err = %v, want a token recognition error", src, err)
		}
	}
}

// The fixture is real gcc -E output whose region body contains `$` coming
// from bad_vendor.h, so the diagnostic must name the header and its line,
// not the parsed .suckc path.
// Fixture command (run from /tmp/opencode/cscope):
//
//	gcc -E -x c lexsrc.suckc > internal/frontend/testdata/lex_err_region.c
func TestLexerErrorReportsOriginatingFile(t *testing.T) {
	_, err := ParseFile("testdata/lex_err_region.c")
	if err == nil {
		t.Fatal("expected a lexer error for `$` in the region body")
	}
	if !strings.Contains(err.Error(), "bad_vendor.h:3:11: token recognition error") {
		t.Errorf("err = %q, want it to point at bad_vendor.h:3:11", err)
	}
}
