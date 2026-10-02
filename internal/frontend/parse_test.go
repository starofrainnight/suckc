package frontend

import (
	"strings"
	"testing"

	"github.com/starofrainnight/suckc/internal/parser"
)

func TestParseSourceValid(t *testing.T) {
	src := "typedef int kknd;\nint a = 0;\nvoid func_a() {}\n"
	res, err := ParseSource(src)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if res.Tree == nil {
		t.Fatal("expected non-nil tree")
	}
	seq := res.Tree.DeclarationSeq()
	if seq == nil || len(seq.AllDeclaration()) != 3 {
		t.Fatalf("expected 3 declarations")
	}
}

func TestParseSourceSyntaxError(t *testing.T) {
	_, err := ParseSource("int a = ;\n")
	if err == nil {
		t.Fatal("expected syntax error")
	}
	if !strings.Contains(err.Error(), "line 1:") {
		t.Errorf("expected line info in error, got: %v", err)
	}
}

func TestParseSourceEmpty(t *testing.T) {
	res, err := ParseSource("")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if res.Tree == nil {
		t.Fatal("expected non-nil tree")
	}
}

func TestParseFileMissing(t *testing.T) {
	_, err := ParseFile("nonexistent.suckc")
	if err == nil {
		t.Fatal("expected error for missing file")
	}
}

func TestParseFileSample(t *testing.T) {
	res, err := ParseFile("../../SuckC.suckc")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if res.Tree == nil {
		t.Fatal("expected non-nil tree")
	}
}

// `class` is deliberately not a keyword: `struct` is the only type keyword and
// `typename` covers template parameters. So it must lex as a plain identifier.
func TestClassIsNotAKeyword(t *testing.T) {
	res, err := ParseSource("int class = 3;\n")
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	seen := false
	for _, tok := range res.Tokens {
		if tok.GetText() != "class" {
			continue
		}
		seen = true
		if got := tok.GetTokenType(); got != parser.SuckCParserIdentifier {
			t.Errorf("`class` token type = %d, want Identifier (%d)", got, parser.SuckCParserIdentifier)
		}
	}
	if !seen {
		t.Fatal("no `class` token in source")
	}
}

// A function definition must carry a parameter list. Without that requirement
// `int C { int a; };` parses as a definition of C returning `int`, which every
// C compiler rejects.
func TestFunctionDefinitionNeedsParameterList(t *testing.T) {
	for _, src := range []string{
		"int C { int a; };\n",
		"void f() {}\nint C { int a; };\n",
	} {
		if _, err := ParseSource(src); err == nil {
			t.Errorf("%q: expected syntax error, got none", src)
		}
	}
}

func TestParseSourceFunctionDefinition(t *testing.T) {
	for _, src := range []string{
		"void f() {}\n",
		"int g(int a) { return a; }\n",
		"int *h() { return 0; }\n",
		"int (*k)(int) { return 0; }\n",
	} {
		if _, err := ParseSource(src); err != nil {
			t.Errorf("%q: expected no error, got %v", src, err)
		}
	}
}

func TestParseSourceStruct(t *testing.T) {
	for _, src := range []string{
		"struct S { int a; };\n",
		"struct S;\n",
		"enum struct E { A };\n",
		"enum E : int { A };\n",
	} {
		if _, err := ParseSource(src); err != nil {
			t.Errorf("%q: expected no error, got %v", src, err)
		}
	}
}
