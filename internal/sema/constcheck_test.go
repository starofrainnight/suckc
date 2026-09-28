package sema

import (
	"strings"
	"testing"

	"github.com/starofrainnight/suckc/internal/frontend"
)

func TestFileConstantAccepted(t *testing.T) {
	cases := []string{
		"auto x = 12;",
		"enum Color { RED, GREEN = 3 };\nauto x = RED;",
		"const int c = 5;\nauto x = c;",
		"const int c = 5;\nauto x = c + 1;",
		"int a = 0;\nauto *p = &a;",
		"int a = 0;\nauto *p = &a + 1;",
		"int arr[2] = {0, 1};\nauto *p = &arr[0];",
		"auto x = (12 + 1);",
		"auto n = sizeof(int);",
		"auto list[] = {1, 2};",
		"auto str[] = \"abc\";",
		// A string literal may initialize a char array (C89 3.5.7) or a
		// char pointer, so a value position is legal too.
		"auto x = \"abc\";",
	}
	for _, src := range cases {
		res, err := frontend.ParseSource(src)
		if err != nil {
			t.Fatalf("parse %q: %v", src, err)
		}
		a := newAnalyzer(Options{TargetBits: 32, FileName: "t.suckc"})
		if err := pass1(res, a); err != nil {
			t.Fatalf("pass1 %q: %v", src, err)
		}
		if len(a.pendingAutos) != 1 {
			t.Fatalf("%q: pendingAutos = %d, want 1", src, len(a.pendingAutos))
		}
		initDecl := a.pendingAutos[0].Decl.InitDeclarator()
		if err := checkFileConstant(initDecl.Initializer(), a); err != nil {
			t.Errorf("%q: unexpected error: %v", src, err)
		}
	}
}

func TestFileConstantRejected(t *testing.T) {
	cases := []string{
		"int a = 0;\nauto x = a;",
		"auto x = f();",
		"auto x = (1, 2);",
		"int a = 0;\nauto x = a++;",
		"auto x = (int)1;",
		"int a = 0;\nauto x = (a = 1);",
		"const int c = f();\nauto x = c;",
	}
	for _, src := range cases {
		res, err := frontend.ParseSource(src)
		if err != nil {
			t.Fatalf("parse %q: %v", src, err)
		}
		a := newAnalyzer(Options{TargetBits: 32, FileName: "t.suckc"})
		if err := pass1(res, a); err != nil {
			t.Fatalf("pass1 %q: %v", src, err)
		}
		if len(a.pendingAutos) != 1 {
			t.Fatalf("%q: pendingAutos = %d, want 1", src, len(a.pendingAutos))
		}
		initDecl := a.pendingAutos[0].Decl.InitDeclarator()
		err = checkFileConstant(initDecl.Initializer(), a)
		if err == nil {
			t.Errorf("%q: expected an error", src)
			continue
		}
		if !strings.Contains(err.Error(), "not a constant expression") {
			t.Errorf("%q: error = %v, want not a constant expression", src, err)
		}
	}
}
