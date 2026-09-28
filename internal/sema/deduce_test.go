package sema

import (
	"strings"
	"testing"

	"github.com/starofrainnight/suckc/internal/frontend"
	"github.com/starofrainnight/suckc/internal/parser"
)

// fileScopeDecl returns an analyzer that has run pass 1 over src, plus the
// first file-scope auto declaration in it.
func fileScopeDecl(t *testing.T, src string) (*Analyzer, *parser.SimpleDeclarationContext) {
	t.Helper()
	res, err := frontend.ParseSource(src)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	a := newAnalyzer(Options{TargetBits: 32, FileName: "t.suckc"})
	if err := pass1(res, a); err != nil {
		t.Fatalf("pass1: %v", err)
	}
	if len(a.pendingAutos) != 1 {
		t.Fatalf("pendingAutos = %d, want 1", len(a.pendingAutos))
	}
	return a, a.pendingAutos[0].Decl
}

func TestDeduceAutoFileScope(t *testing.T) {
	src := "int i = 3;\nauto *p = &i;\n"
	a, decl := fileScopeDecl(t, src)
	if err := deduceAuto(a, decl, true); err != nil {
		t.Fatalf("deduceAuto: %v", err)
	}
	// The declarator star stays in the source, so the substituted text is
	// "int" while p itself is an int *.
	if got := allAutoSubs(t, src, a.Subs); len(got) != 1 || got[0] != "int" {
		t.Errorf("subs = %v, want [int]", got)
	}
	if b, ok := a.FileVars["p"]; !ok || b.Type.Text() != "int *" {
		t.Errorf("FileVars[p] = %v, ok = %v", b, ok)
	}
}

func TestDeduceAutoMissingInitializer(t *testing.T) {
	a, decl := fileScopeDecl(t, "auto i;\n")
	err := deduceAuto(a, decl, true)
	if err == nil {
		t.Fatal("expected an error")
	}
	if !strings.Contains(err.Error(), "missing initializer") {
		t.Errorf("error = %v, want missing initializer", err)
	}
}

func TestBlockScopeIsNotConstantGated(t *testing.T) {
	src := "int a = 0;\nvoid f() { auto x = a; }\n"
	if _, err := analyzeSrc(t, src, Options{TargetBits: 32, FileName: "t.suckc"}); err != nil {
		t.Fatalf("analyze: %v", err)
	}
}
