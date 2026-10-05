package sema

import (
	"strings"
	"testing"
)

// cRegion declares the given C text inside a file-scope C-scope region and then
// runs a SuckC function body, so a call to the declared name is checked.
func cRegion(t *testing.T, cDecls, body string) error {
	t.Helper()
	src := "#pragma c_scope_begin\n" + cDecls + "\n#pragma c_scope_end\n" +
		"void myMain(void) {\n" + body + "\n}\n"
	_, err := analyzeSrc(t, src, Options{FileName: "t.suckc"})
	return err
}

func TestCCallArityAccepted(t *testing.T) {
	cases := []struct {
		name  string
		decls string
		body  string
	}{
		{"exact match", "void f(int a, int b);", "f(1, 2);"},
		{"no arguments for void prototype", "void f(void);", "f();"},
		{"extra arguments for variadic", "int f(int a, ...);", "f(1, 2, 3, 4);"},
		{"one argument for variadic", "int f(int a, ...);", "f(1);"},
		{"unprototyped takes anything", "int f();", "f(1, 2, 3, 4, 5);"},
		{"unprototyped takes none", "int f();", "f();"},
		{"name not declared in any region", "void other(int);", "undeclared(1, 2, 3);"},
		{"nested calls both checked below", "void inner(int a); void outer(int a);",
			"outer(inner(1));"},
		{"call in a loop", "void f(int a);", "for (int i = 0; i < 3; i++) { f(i); }"},
		{"call result discarded then another", "void f(void);", "f(); f();"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if err := cRegion(t, tc.decls, tc.body); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestCCallArityRejected(t *testing.T) {
	cases := []struct {
		name  string
		decls string
		body  string
		want  string
	}{
		{"too few", "int f(int a, int b);", "f(1);",
			"too few arguments in call to 'f': 1 given, C region declares 2"},
		{"none given", "int f(int a);", "f();",
			"too few arguments in call to 'f': 0 given, C region declares 1"},
		{"too many", "int f(int a);", "f(1, 2);",
			"too many arguments in call to 'f': 2 given, C region declares 1"},
		{"arguments to void prototype", "void f(void);", "f(1);",
			"too many arguments in call to 'f': 1 given, C region declares 0"},
		{"variadic still needs its fixed argument", "int f(int a, ...);", "f();",
			"too few arguments in call to 'f': 0 given, C region declares 1 or more"},
		{"variadic is not unlimited", "int f(int a);", "f(1, 2, 3);",
			"too many arguments in call to 'f': 3 given, C region declares 1"},
		{"inner call of a nested call", "void inner(int a, int b); void outer(int a);",
			"outer(inner(1));",
			"too few arguments in call to 'inner': 1 given, C region declares 2"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := cRegion(t, tc.decls, tc.body)
			if err == nil {
				t.Fatal("want an arity error, got none")
			}
			if !strings.Contains(err.Error(), tc.want) {
				t.Fatalf("error mismatch\n got: %v\nwant substring: %s", err, tc.want)
			}
		})
	}
}

// A name the .suckc file declares itself is skipped, so that redeclaring a
// vendor prototype keeps working.
func TestCCallSkipsSuckCDeclaredNames(t *testing.T) {
	t.Run("redeclaration is the accepted form", func(t *testing.T) {
		src := "#pragma c_scope_begin\nint f(int a, int b);\n#pragma c_scope_end\n" +
			"int f(int a);\nvoid myMain(void) {\n    f(1);\n}\n"
		if _, err := analyzeSrc(t, src, Options{FileName: "t.suckc"}); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
	t.Run("local declaration shadows too", func(t *testing.T) {
		src := "#pragma c_scope_begin\nint f(int a, int b);\n#pragma c_scope_end\n" +
			"void myMain(void) {\n    int f = 1;\n    f;\n}\n"
		if _, err := analyzeSrc(t, src, Options{FileName: "t.suckc"}); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

func TestCCallIgnoresNonCallSuffixes(t *testing.T) {
	cases := []struct {
		name  string
		decls string
		body  string
	}{
		{"subscript", "int a(int x);", "int arr[10]; arr[0] = 1;"},
		{"field access then call", "void f(void);", "int s; s;"},
		{"call on a parenthesised callee", "void f(void);", "(f)();"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if err := cRegion(t, tc.decls, tc.body); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestCCallDuplicateDeclarations(t *testing.T) {
	t.Run("identical prototypes merge", func(t *testing.T) {
		if err := cRegion(t, "void f(int a); void f(int a);", "f(1);"); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
	t.Run("unprototyped alongside prototyped merges", func(t *testing.T) {
		if err := cRegion(t, "void f(int a); void f();", "f(1, 2, 3);"); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
	t.Run("disagreeing arities are an error", func(t *testing.T) {
		err := cRegion(t, "void f(int a); void f(int a, int b);", "f(1);")
		if err == nil {
			t.Fatal("want a conflict error, got none")
		}
		if !strings.Contains(err.Error(), "conflicting declarations of 'f'") {
			t.Fatalf("error mismatch: %v", err)
		}
	})
	t.Run("disagreeing variadic flag is an error", func(t *testing.T) {
		err := cRegion(t, "int f(int a, ...); int f(int a);", "f(1);")
		if err == nil {
			t.Fatal("want a conflict error, got none")
		}
		if !strings.Contains(err.Error(), "conflicting declarations of 'f'") {
			t.Fatalf("error mismatch: %v", err)
		}
	})
}

// The diagnostic must point at the call site in the .suckc file and quote where
// the declaration came from. Without preprocessor line markers there is no
// header name to recover, so both fall back to the .suckc path.
func TestCCallErrorPointsAtCallSite(t *testing.T) {
	src := "#pragma c_scope_begin\nint vendor_open(int a, int b);\n#pragma c_scope_end\n" +
		"void myMain(void) {\n    vendor_open(1);\n}\n"
	_, err := analyzeSrc(t, src, Options{FileName: "t.suckc"})
	if err == nil {
		t.Fatal("want an arity error, got none")
	}
	msg := err.Error()
	if !strings.HasPrefix(msg, "t.suckc:") {
		t.Fatalf("error must name the .suckc call site, got: %v", err)
	}
	if !strings.Contains(msg, "declared at") {
		t.Fatalf("error must quote the declaration origin, got: %v", err)
	}
}

// Calls inside a C region are C, not SuckC: the region body is not walked for
// SuckC semantics, so an arity error must never be raised there.
func TestCCallInsideRegionBodyIsNotChecked(t *testing.T) {
	src := "#pragma c_scope_begin\n" +
		"void f(int a, int b);\n" +
		"void caller(void) { f(1); }\n" +
		"#pragma c_scope_end\n" +
		"void myMain(void) { }\n"
	if _, err := analyzeSrc(t, src, Options{FileName: "t.suckc"}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
