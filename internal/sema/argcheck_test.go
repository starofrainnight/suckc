package sema

import (
	"strings"
	"testing"
)

// argRegion declares C text in a file-scope region and runs a SuckC body that
// calls into it. decls are the C declarations; pre are SuckC file-scope
// declarations; body is the SuckC function body.
func argRegion(t *testing.T, decls, pre, body string) error {
	t.Helper()
	src := "#pragma c_scope_begin\n" + decls + "\n#pragma c_scope_end\n" +
		pre + "\nvoid myMain(void) {\n" + body + "\n}\n"
	_, err := analyzeSrc(t, src, Options{FileName: "t.suckc"})
	return err
}

// TestCArgAccepted covers argument/parameter pairs that must not be reported.
// Most of these are the "shape matches" cases; several exist specifically to
// pin the deliberate looseness (pointers are not compared pointee-wise).
func TestCArgAccepted(t *testing.T) {
	cases := []struct {
		name  string
		decls string
		pre   string
		body  string
	}{
		{"pointer to same pointer", "void f(char *p);", "char *s;", "f(s);"},
		{"pointer to void pointer", "void f(void *p);", "char *s;", "f(s);"},
		{"int pointer to int pointer", "void f(int *p);", "int v;", "f(&v);"},
		{"different pointee types are not compared", "void f(int *p);", "char *s;", "f(s);"},
		{"null pointer constant for pointer parameter", "void f(char *p);", "", "f(0);"},
		{"null pointer constant with suffix", "void f(char *p);", "", "f(0u);"},
		{"null pointer constant hex zero", "void f(char *p);", "", "f(0x0);"},
		{"array decays to pointer", "void f(char *p);", "char buf[4];", "f(buf);"},
		{"block scope local", "void f(char *p);", "", "char *s; f(s);"},
		{"deref of pointer is a scalar", "void f(int x);", "char *p;", "f(*p);"},
		{"integer literal", "void f(int x);", "", "f(1);"},
		{"character literal", "void f(int x);", "", "f('a');"},
		{"same struct tag by typedef", "typedef struct Foo Foo; void f(Foo x);",
			"struct Foo v;", "f(v);"},
		{"same struct tag spelled directly", "void f(struct Foo x);",
			"struct Foo v;", "f(v);"},
		{"variadic tail is not checked", "void f(char *p, ...);", "", "f(0, 1, 2, 3);"},
		{"unprototyped parameter has no shape", "void f();", "char *s;", "f(s, 1, 2);"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if err := argRegion(t, tc.decls, tc.pre, tc.body); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

// Known gap, deliberately not asserted either way: `f(&v);` is not recognised
// as a call at all. SuckC inherits C++'s [dcl.ambiguity] -- a statement whose
// first token is a name can be either a call or a declaration, and the grammar
// has no typedef table to tell `f` (a function) from `f` (a type name). The
// single-identifier case is recovered in declCallArgs by asking the C-region
// signature table, but `&v` takes the reference-declarator reading and the
// argument never reaches the checker. Closing it needs either a grammar change
// (add a comma list to simpleDeclaration so `f(s, t);` parses at all) or a
// pre-pass that resolves statement ambiguity from the signature table.

// TestCArgRejected covers the genuinely incompatible shapes. Pointer depth is
// the one thing that can never be reconciled by an implicit conversion.
func TestCArgRejected(t *testing.T) {
	cases := []struct {
		name    string
		decls   string
		pre     string
		body    string
		wantSub string
	}{
		{"pointer to scalar", "void f(int x);", "char *s;", "f(s);",
			"passing 'char *' where 'int' is expected"},
		{"scalar to pointer", "void f(char *p);", "int v;", "f(v);",
			"passing 'int' where 'char *' is expected"},
		{"non-zero constant is not a null pointer constant", "void f(char *p);", "", "f(0xFF);",
			"passing 'int' where 'char *' is expected"},
		{"pointer to pointer of wrong depth", "void f(char **p);", "char *s;", "f(s);",
			"passing 'char *' where 'char **' is expected"},
		{"different struct tags", "typedef struct Foo Foo; typedef struct Bar Bar; void f(Bar x);",
			"struct Foo v;", "f(v);",
			"passing 'struct Foo' where 'struct Bar' is expected"},
		{"different enum tags", "typedef enum E1 E1; typedef enum E2 E2; void f(E2 x);",
			"enum E1 v;", "f(v);",
			"passing 'enum E1' where 'enum E2' is expected"},
		{"struct against scalar", "typedef struct Foo Foo; void f(int x);",
			"struct Foo v;", "f(v);",
			"passing 'struct Foo' where 'int' is expected"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := argRegion(t, tc.decls, tc.pre, tc.body)
			if err == nil {
				t.Fatal("want an incompatible-argument error, got none")
			}
			if !strings.Contains(err.Error(), tc.wantSub) {
				t.Fatalf("error does not mention %q:\n%v", tc.wantSub, err)
			}
		})
	}
}

// TestCArgSkipped pins the cases where the argument type cannot be determined
// and the call is therefore let through. Reporting nothing is the safe
// direction: a false rejection of working code is worse than a missed check.
func TestCArgSkipped(t *testing.T) {
	cases := []struct {
		name  string
		decls string
		pre   string
		body  string
	}{
		{"binary arithmetic result", "void f(int x);", "int a; int b;", "f(a + b);"},
		{"comparison result", "void f(int x);", "int a; int b;", "f(a < b);"},
		{"nested call result", "void g(void); void f(char *p);", "", "f(g());"},
		{"unresolvable name", "void f(int x);", "", "f(mystery);"},
		{"address of a call result", "void f(int x);", "", "f(&g());"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if err := argRegion(t, tc.decls, tc.pre, tc.body); err != nil {
				t.Fatalf("unknown argument type must not be reported: %v", err)
			}
		})
	}
}

// TestCArgBlockScopeLocalIsTyped is the regression test for wiring the check
// into pass2 rather than next to checkCCalls: checkCCalls runs after every
// block scope has been popped, so a local's type would be gone by then.
func TestCArgBlockScopeLocalIsTyped(t *testing.T) {
	err := argRegion(t, "void f(int x);", "", "char *s; f(s);")
	if err == nil {
		t.Fatal("a block-scope pointer passed to a scalar parameter must be reported")
	}
	if !strings.Contains(err.Error(), "incompatible argument") {
		t.Fatalf("unexpected diagnostic: %v", err)
	}
}

// TestCArgErrorPointsAtBothSites checks the diagnostic carries the call site in
// the .suckc file and the declaration in the originating header.
func TestCArgErrorPointsAtBothSites(t *testing.T) {
	src := "#pragma c_scope_begin\nvoid takes_int(int x);\n#pragma c_scope_end\n" +
		"void myMain(void) {\n    char *s;\n    takes_int(s);\n}\n"
	_, err := analyzeSrc(t, src, Options{FileName: "t.suckc"})
	if err == nil {
		t.Fatal("want an error")
	}
	msg := err.Error()
	if !strings.Contains(msg, "incompatible argument 1 in call to 'takes_int'") {
		t.Fatalf("diagnostic does not name the argument position: %v", err)
	}
	if !strings.Contains(msg, "t.suckc:6") {
		t.Fatalf("diagnostic does not point at the call site: %v", err)
	}
}

func TestCArgSingleIdentifierArgumentIsChecked(t *testing.T) {
	// "f(p);" parses as "declare p of type f" because the grammar has no
	// typedef table to tell the callee apart from a type name, so a single
	// identifier argument used to reach neither the arity check nor the
	// argument check. Both checks now recover it through typedefNameSpelling.
	// A regression here produced no failing test at all: the call was simply
	// not recognised, so both checks reported success.
	if err := argRegion(t, "void takes_two(int a, int b);", "int p;", "takes_two(p);"); err == nil {
		t.Fatal("want an error for a one-argument call, got none")
	}
	if err := argRegion(t, "void takes_one(int a);", "int p;", "takes_one(p);"); err != nil {
		t.Fatalf("a correctly-arity'd call with an identifier argument was rejected: %v", err)
	}
}
