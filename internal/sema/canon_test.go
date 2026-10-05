package sema

import (
	"testing"

	"github.com/starofrainnight/suckc/internal/frontend"
)

func TestCanonicalSpelling(t *testing.T) {
	cases := []struct{ in, want string }{
		{"unsigned long int", "unsigned long"},
		{"long int", "long"},
		{"signed", "int"},
		{"long long int", "long long"},
		{"long unsigned", "unsigned long"},
		{"long unsigned int", "unsigned long"},
		{"unsigned long long int", "unsigned long long"},
		{"long double", "long double"},
		{"_Bool", "bool"},
		{"const char", "const char"},
		{"char const", "const char"},
		{"unsigned long const", "const unsigned long"},
		{"volatile int", "volatile int"},
		{"char", "char"},
		{"signed char", "signed char"},
		{"unsigned char", "unsigned char"},
		{"short int", "short"},
		{"unsigned short int", "unsigned short"},
		{"double", "double"},
		{"float", "float"},
		{"void", "void"},

		// Not plain builtin lists: returned as written rather than guessed at.
		{"struct Foo", "struct Foo"},
		{"union Bar", "union Bar"},
		{"enum E", "enum E"},
		{"int char", "int char"},
		{"long long long", "long long long"},
		{"unsigned signed", "unsigned signed"},
		{"void int", "void int"},
		{"static int", "static int"},
		{"size_t", "size_t"},
		{"unsigned long *", "unsigned long *"},
	}
	for _, c := range cases {
		if got := canonicalSpelling(c.in); got != c.want {
			t.Errorf("canonicalSpelling(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestCanonicalSpellingKeepsCharVariantsApart(t *testing.T) {
	// C89 6.2.5: char, signed char and unsigned char are three distinct types.
	// A normalizer that collapsed them would let an incompatible argument pass.
	seen := map[string]string{}
	for _, sp := range []string{"char", "signed char", "unsigned char"} {
		got := canonicalSpelling(sp)
		if prev, dup := seen[got]; dup {
			t.Errorf("canonicalSpelling(%q) = %q, collides with %q", sp, got, prev)
		}
		seen[got] = sp
	}
}

func TestSplitTrailingStars(t *testing.T) {
	cases := []struct {
		in    string
		base  string
		stars int
	}{
		{"char", "char", 0},
		{"char *", "char", 1},
		{"char **", "char", 2},
		{"unsigned long *", "unsigned long", 1},
		{"const char *", "const char", 1},
		{"struct Foo *", "struct Foo", 1},
	}
	for _, c := range cases {
		base, stars := splitTrailingStars(c.in)
		if base != c.base || stars != c.stars {
			t.Errorf("splitTrailingStars(%q) = (%q, %d), want (%q, %d)",
				c.in, base, stars, c.base, c.stars)
		}
	}
}

func TestCanonTypeKey(t *testing.T) {
	cases := []struct {
		ct   canonType
		want string
	}{
		{canonType{Base: "char"}, "char"},
		{canonType{Base: "char", Stars: 1}, "char*"},
		{canonType{Base: "char", Stars: 2}, "char**"},
		{canonType{Base: "unsigned long"}, "unsigned long"},
	}
	for _, c := range cases {
		if got := c.ct.key(); got != c.want {
			t.Errorf("canonType%+v key = %q, want %q", c.ct, got, c.want)
		}
	}
}

func TestResolveWithoutTypedefs(t *testing.T) {
	r := newCTypeResolver(nil)
	cases := []struct{ in, want string }{
		{"unsigned long int", "unsigned long"},
		{"long unsigned int", "unsigned long"},
		{"const char *", "char*"},
		{"struct VendorReg *", "struct VendorReg*"},
		{"enum Mode", "enum Mode"},
		// Not declared here, and not a builtin: compared as written.
		{"some_unknown_t", "some_unknown_t"},
		{"int char", "int char"},
	}
	for _, c := range cases {
		if got := r.resolve(c.in).key(); got != c.want {
			t.Errorf("resolve(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestResolveExpandsTypedefs(t *testing.T) {
	// The size_t spelling is what glibc actually produces: __SIZE_TYPE__
	// expands to "long unsigned int", not "unsigned long".
	r := newCTypeResolver([]frontend.CTypedef{
		{Name: "size_t", Base: "long unsigned int"},
		{Name: "vendor_cstr", Base: "char", Stars: 1},
		{Name: "vendor_handle", Base: "void", Stars: 1},
		{Name: "vendor_cb", Base: "int", Stars: 1, FuncPtr: true,
			Params: []frontend.CParam{{TypeSpelling: "int"}, {TypeSpelling: "void *"}}},
	})
	cases := []struct{ in, want string }{
		{"size_t", "unsigned long"},
		{"size_t *", "unsigned long*"},
		{"vendor_cstr", "char*"},
		{"vendor_cstr *", "char**"},
		{"vendor_handle", "void*"},
		// The signature is dropped, so two different callback types compare
		// equal. That misses checks; it never rejects a valid call.
		{"vendor_cb", "FUNCPTR(int)*"},
		{"vendor_cb *", "FUNCPTR(int)**"},
	}
	for _, c := range cases {
		if got := r.resolve(c.in).key(); got != c.want {
			t.Errorf("resolve(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestResolveTypedefChainsTerminate(t *testing.T) {
	r := newCTypeResolver([]frontend.CTypedef{
		{Name: "a", Base: "b"},
		{Name: "b", Base: "a"},
		{Name: "self", Base: "self *"},
	})
	done := make(chan string, 1)
	go func() { done <- r.resolve("a").key() + "|" + r.resolve("self").key() }()
	got := <-done
	if got != "a|self*" {
		t.Errorf("self-referential typedef chains: got %q, want %q", got, "a|self*")
	}
}

func TestResolveFirstTypedefOfADuplicateNameWins(t *testing.T) {
	r := newCTypeResolver([]frontend.CTypedef{
		{Name: "dup", Base: "int"},
		{Name: "dup", Base: "char *", Stars: 1},
	})
	if got := r.resolve("dup").key(); got != "int" {
		t.Errorf("resolve(dup) = %q, want %q (first declaration wins)", got, "int")
	}
}

// An untagged aggregate is a nominal type in C: two "struct { int a; }" are
// different types. The spelling cannot identify it, so each one gets its own
// name -- which is also why it cannot be compared by text.
func TestAnonymousAggregatesGetDistinctNames(t *testing.T) {
	r := newCTypeResolver(nil)
	if got := r.resolve("struct { int a; }").key(); got != "struct <anon 1>" {
		t.Errorf("anonymous struct = %q, want %q", got, "struct <anon 1>")
	}
	if got := r.resolve("union { int a; }").key(); got != "union <anon 2>" {
		t.Errorf("anonymous union = %q, want %q", got, "union <anon 2>")
	}
}

// Two anonymous aggregates with byte-identical bodies share a name, so they
// compare equal when C calls them distinct. That miss is deliberate: naming
// them per occurrence instead would give the same typedef a different name in
// two declarations and reject calls that are valid.
func TestIdenticalAnonymousBodiesShareAName(t *testing.T) {
	r := newCTypeResolver(nil)
	first := r.resolve("struct { int a; }").key()
	second := r.resolve("struct { int a; }").key()
	if first != second {
		t.Errorf("identical anonymous bodies got %q and %q, want the same name", first, second)
	}
}

func TestTaggedAggregatesShareOneName(t *testing.T) {
	r := newCTypeResolver(nil)
	if got := r.resolve("struct VendorReg *").key(); got != "struct VendorReg*" {
		t.Errorf("tagged struct = %q, want %q", got, "struct VendorReg *")
	}
}

func TestCanonicalSpellingIsWhatAutoEmits(t *testing.T) {
	// Redundant keywords used to survive into the type text, and rankOf only
	// knew the canonical spellings, so any variable declared with them was
	// rejected the moment it took part in arithmetic.
	cases := []struct {
		decl string
		want string
	}{
		{"unsigned long int g; auto y = g + 1;", "unsigned long"},
		{"long int g; auto y = g + 1;", "long"},
		{"unsigned g; auto y = g + 1;", "unsigned int"},
		{"long unsigned int g; auto y = g + 1;", "unsigned long"},
		{"signed char c; auto y = c + 1;", "int"},
		{"unsigned short s; auto y = s * 2;", "int"},
	}
	for _, c := range cases {
		src := "void f(void) { " + c.decl + " }"
		subs, err := analyzeSrc(t, src, Options{TargetBits: 32})
		if err != nil {
			t.Errorf("%s: unexpected error: %v", c.decl, err)
			continue
		}
		var got string
		for _, v := range subs {
			got = v
		}
		if got != c.want {
			t.Errorf("%s: deduced %q, want %q", c.decl, got, c.want)
		}
	}
}
