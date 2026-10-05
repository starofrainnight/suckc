package frontend

import (
	"reflect"
	"testing"
)

// parseCRegion wraps body in C-scope markers at file scope and parses it.
func parseCRegion(t *testing.T, body string) *ParseResult {
	t.Helper()
	src := "#pragma c_scope_begin\n" + body + "\n#pragma c_scope_end\n"
	res, err := ParseSource(src)
	if err != nil {
		t.Fatalf("parse region %q: %v", body, err)
	}
	return res
}

// TokIdx is a token-stream position, not part of what a declaration means, and
// it has its own test below. Blank it so the table tests can state expectations
// without hardcoding token indexes.
func withoutTokIdx(sigs []CFuncSig) []CFuncSig {
	out := append([]CFuncSig(nil), sigs...)
	for i := range out {
		out[i].TokIdx = 0
	}
	return out
}

func withoutTypedefTokIdx(tds []CTypedef) []CTypedef {
	out := append([]CTypedef(nil), tds...)
	for i := range out {
		out[i].TokIdx = 0
	}
	return out
}

func TestCDeclFuncSignatures(t *testing.T) {
	tests := []struct {
		name string
		body string
		want []CFuncSig
	}{
		{
			// "void" as the only parameter name means no parameters at all.
			name: "void parameter list is empty",
			body: "void f(void);",
			want: []CFuncSig{{Name: "f"}},
		},
		{
			// "int f()" declares a function whose parameters are unspecified,
			// which is not the same as declaring zero parameters.
			name: "empty parentheses are unprototyped",
			body: "int f();",
			want: []CFuncSig{{Name: "f", Unprototyped: true}},
		},
		{
			name: "named parameter and ellipsis",
			body: "int f(int a, ...);",
			want: []CFuncSig{{Name: "f", Params: []CParam{{TypeSpelling: "int", Name: "a"}}, Variadic: true}},
		},
		{
			// Unnamed parameters still contribute their type, including stars.
			name: "unnamed parameters keep their stars",
			body: "int f(int, char *, void *);",
			want: []CFuncSig{{Name: "f", Params: []CParam{
				{TypeSpelling: "int"},
				{TypeSpelling: "char *"},
				{TypeSpelling: "void *"},
			}}},
		},
		{
			name: "function pointer is a signature on its pointee name",
			body: "int (*fp)(int);",
			want: []CFuncSig{{Name: "fp", Params: []CParam{{TypeSpelling: "int"}}}},
		},
		{
			// fp is a function returning a pointer to int, so the parameter
			// list attached to fp is fp's own -- not the returned one's.
			name: "function returning pointer uses its own parameters",
			body: "int *fp(int a);",
			want: []CFuncSig{{Name: "fp", Params: []CParam{{TypeSpelling: "int", Name: "a"}}}},
		},
		{
			// An array of function pointers is not callable by name.
			name: "array of function pointers is not a signature",
			body: "int (*a[3])(void);",
			want: nil,
		},
		{
			name: "declarations without a function suffix are not signatures",
			body: "int a, b;",
			want: nil,
		},
		{
			// The body's own declarations are locals, not program-level
			// signatures; collecting them would invent callable names.
			name: "locals inside a definition body are skipped",
			body: "static int f(void) { int hidden_local; return hidden_local; }",
			want: []CFuncSig{{Name: "f"}},
		},
		{
			name: "multiple declarators on one specifier line",
			body: "int alpha(void), beta(int);",
			want: []CFuncSig{
				{Name: "alpha"},
				{Name: "beta", Params: []CParam{{TypeSpelling: "int"}}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseCRegion(t, tt.body).CFuncs
			if len(got) == 0 && len(tt.want) == 0 {
				return
			}
			if !reflect.DeepEqual(withoutTokIdx(got), tt.want) {
				t.Errorf("CFuncs =\n  %+v\nwant\n  %+v", got, tt.want)
			}
		})
	}
}

func TestCDeclTypedefs(t *testing.T) {
	tests := []struct {
		name string
		body string
		want []CTypedef
	}{
		{
			name: "multi-word base type",
			body: "typedef unsigned long size_t;",
			want: []CTypedef{{Name: "size_t", Base: "unsigned long"}},
		},
		{
			name: "pointer typedef keeps the star count",
			body: "typedef char *cstr;",
			want: []CTypedef{{Name: "cstr", Base: "char", Stars: 1}},
		},
		{
			name: "function pointer typedef carries its parameters",
			body: "typedef int (*cb)(int, void *);",
			want: []CTypedef{{
				Name: "cb", Base: "int", Stars: 1, FuncPtr: true,
				Params: []CParam{{TypeSpelling: "int"}, {TypeSpelling: "void *"}},
			}},
		},
		{
			name: "struct typedef",
			body: "typedef struct Foo Foo;",
			want: []CTypedef{{Name: "Foo", Base: "struct Foo"}},
		},
		{
			name: "comma separated typedef declarators",
			body: "typedef struct Foo Foo, *FooPtr;",
			want: []CTypedef{
				{Name: "Foo", Base: "struct Foo"},
				{Name: "FooPtr", Base: "struct Foo", Stars: 1},
			},
		},
		{
			name: "typedef is not a function signature",
			body: "typedef int plain_alias;",
			want: []CTypedef{{Name: "plain_alias", Base: "int"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := parseCRegion(t, tt.body)
			if got := res.CFuncs; len(got) != 0 {
				t.Errorf("CFuncs = %+v, want none for a typedef", got)
			}
			if !reflect.DeepEqual(withoutTypedefTokIdx(res.CTypedefs), tt.want) {
				t.Errorf("CTypedefs =\n  %+v\nwant\n  %+v", res.CTypedefs, tt.want)
			}
		})
	}
}

// A linkage specification is the wrapper vendor headers use for their whole
// API, so declarations inside one still have to be collected.
func TestCDeclCollectsThroughLinkageSpecification(t *testing.T) {
	res := parseCRegion(t, "extern \"C\" {\n  int vendor_init(int);\n  void vendor_shutdown(void);\n}\n")
	want := []CFuncSig{
		{Name: "vendor_init", Params: []CParam{{TypeSpelling: "int"}}},
		{Name: "vendor_shutdown"},
	}
	if !reflect.DeepEqual(withoutTokIdx(res.CFuncs), want) {
		t.Errorf("CFuncs =\n  %+v\nwant\n  %+v", res.CFuncs, want)
	}
}

func TestCDeclAttributesDoNotAffectExtraction(t *testing.T) {
	tests := []struct {
		name string
		body string
	}{
		{"attribute before the specifiers", "__attribute__((nonnull)) int f(int a);"},
		{"attribute after the declarator", "int f(int a) __attribute__((const));"},
		{"attribute between specifiers and declarator", "int __attribute__((nonnull)) f(int a);"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseCRegion(t, tt.body).CFuncs
			want := []CFuncSig{{Name: "f", Params: []CParam{{TypeSpelling: "int", Name: "a"}}}}
			if !reflect.DeepEqual(withoutTokIdx(got), want) {
				t.Errorf("CFuncs =\n  %+v\nwant\n  %+v", got, want)
			}
		})
	}
}

// A K&R definition's parameter names arrive as an identifier list, which the
// region grammar reads as unnamed parameters whose types are spelled with
// those same identifiers. The names come out right; the parameter types are
// meaningless. Only the name is pinned here so a future fix to that known gap
// does not have to rewrite this test.
func TestCDeclKnRDefinitionKeepsName(t *testing.T) {
	res := parseCRegion(t, "int knr(a, b) int a; int b; { return a + b; }")
	if len(res.CFuncs) != 1 {
		t.Fatalf("CFuncs = %+v, want exactly one signature", res.CFuncs)
	}
	if res.CFuncs[0].Name != "knr" {
		t.Errorf("name = %q, want %q", res.CFuncs[0].Name, "knr")
	}
}

// TokIdx must point at the declared name so SourceAt can report the line in
// the original header rather than in the translated file.
func TestCDeclTokIdxResolvesToOriginalHeaderLine(t *testing.T) {
	res := parseCRegion(t, "int before_marker_marker_name(void);\nint vendor_api_call(int a);")
	if len(res.CFuncs) != 2 {
		t.Fatalf("CFuncs = %+v, want two", res.CFuncs)
	}
	call := res.CFuncs[1]
	if call.TokIdx <= res.CFuncs[0].TokIdx {
		t.Fatalf("TokIdx not increasing: %d then %d", res.CFuncs[0].TokIdx, call.TokIdx)
	}
	tok := res.Tokens[call.TokIdx]
	if tok.GetText() != "vendor_api_call" {
		t.Errorf("TokIdx points at %q, want the declared name", tok.GetText())
	}
}
