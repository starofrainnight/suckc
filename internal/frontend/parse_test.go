package frontend

import (
	"strings"
	"testing"
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
