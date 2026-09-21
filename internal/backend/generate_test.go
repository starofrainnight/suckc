package backend

import (
	"testing"

	"github.com/starofrainnight/suckc/internal/frontend"
)

func TestGenerateRoundTrip(t *testing.T) {
	src := "typedef void (*func_type)(int a);\ntypedef int kknd;\nint a = 0;\n\nvoid func_a() { printf(\"func_a\\n\"); }\n"
	res, err := frontend.ParseSource(src)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	got, err := Generate(res)
	if err != nil {
		t.Fatalf("generate: %v", err)
	}
	if got != src {
		t.Errorf("expected output to equal input, got:\n%s", got)
	}
}

func TestGeneratePreservesCommentsAndWhitespace(t *testing.T) {
	src := "// leading comment\nint a = 0; // trailing\n\n/* block */\nvoid f() {}\n"
	res, err := frontend.ParseSource(src)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	got, err := Generate(res)
	if err != nil {
		t.Fatalf("generate: %v", err)
	}
	if got != src {
		t.Errorf("expected output to equal input, got:\n%s", got)
	}
}

func TestGenerateEmptySource(t *testing.T) {
	res, err := frontend.ParseSource("")
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	got, err := Generate(res)
	if err != nil {
		t.Fatalf("generate: %v", err)
	}
	if got != "" {
		t.Errorf("expected empty output, got: %q", got)
	}
}

func TestGenerateNoTree(t *testing.T) {
	if _, err := Generate(nil); err == nil {
		t.Fatal("expected error for nil result")
	}
}
