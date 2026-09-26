package sema

import (
	"testing"

	"github.com/starofrainnight/suckc/internal/frontend"
)

func TestAnalyzeNilTree(t *testing.T) {
	if _, err := Analyze(nil, Options{}); err == nil {
		t.Fatal("expected error for nil parse result")
	}
}

func TestAnalyzeRejectsBadTargetBits(t *testing.T) {
	res, err := frontend.ParseSource("int a = 0;\n")
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if _, err := Analyze(res, Options{TargetBits: 24, FileName: "x.suckc"}); err == nil {
		t.Fatal("expected error for TargetBits 24")
	}
}

func TestAnalyzeNoAutoReturnsEmptySubs(t *testing.T) {
	res, err := frontend.ParseSource("int a = 0;\nvoid f() { int b = 1; }\n")
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	subs, err := Analyze(res, Options{TargetBits: 32, FileName: "t.suckc"})
	if err != nil {
		t.Fatalf("analyze: %v", err)
	}
	if len(subs) != 0 {
		t.Fatalf("expected empty subs, got %v", subs)
	}
}

func TestAnalyzeDefaultTargetBitsIsHost(t *testing.T) {
	got := Options{}.resolved()
	if got.TargetBits != 32 && got.TargetBits != 64 {
		t.Fatalf("expected host width 32 or 64, got %d", got.TargetBits)
	}
}

func TestErrorFormat(t *testing.T) {
	e := &Error{File: "example.suckc", Line: 12, Col: 9, Msg: "unsupported expression"}
	want := "example.suckc:12:9: cannot deduce type for 'auto': unsupported expression"
	if e.Error() != want {
		t.Fatalf("got %q, want %q", e.Error(), want)
	}
}
