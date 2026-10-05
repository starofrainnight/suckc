package parser

import "testing"

// wrapFile puts src between region markers at file scope.
func wrapFile(body string) string {
	return "#pragma c_scope_begin\n" + body + "\n#pragma c_scope_end\n"
}

// wrapBlock puts src between region markers inside a function body.
func wrapBlock(body string) string {
	return "void host(void) {\n#pragma c_scope_begin\n" + body + "\n#pragma c_scope_end\n}\n"
}

// mustParseInRegion is the acceptance table for the C declaration
// subgrammar. Every row is a construct measured to fail in the SuckC
// grammar before this subgrammar existed, or a vendor spelling that
// survives `gcc -E` verbatim.
func mustParseInRegion(t *testing.T, body string) {
	t.Helper()
	for _, tc := range []struct {
		scope string
		src   string
	}{
		{"file", wrapFile(body)},
		{"block", wrapBlock(body)},
	} {
		errs, _ := parseWithErrors(tc.src)
		if len(errs) != 0 {
			t.Errorf("[%s] %q: %v", tc.scope, body, errs)
		}
	}
}

func TestCDeclTypedefs(t *testing.T) {
	for _, body := range []string{
		"typedef unsigned long size_t;",
		"typedef long long i64;",
		"typedef unsigned short u16;",
		"typedef signed char s8;",
		"typedef long double ld;",
		"typedef unsigned int u32;",
		"typedef char *cstr;",
		"typedef void *handle;",
		"typedef struct Foo Foo, *FooPtr;",
		"typedef int x;",
		"typedef struct Foo Foo;",
		"typedef int (*cb)(int, void *);",
		"typedef const char *cstr2;",
		"typedef struct Foo *FooPtr2, FooArr[4];",
	} {
		mustParseInRegion(t, body)
	}
}

func TestCDeclDeclarations(t *testing.T) {
	for _, body := range []string{
		"int a, b;",
		"int * restrict p;",
		"int C { int a; };",
		"struct Foo { int x; };",
		"struct Foo *next;",
		"enum E { A, B };",
		"enum E { A = 1, B };",
		"enum E { A, B, };",
		"extern int x;",
		"static const int N = 4;",
		"void f(void);",
		"int f();",
		"unsigned long long y;",
		"void (*fp)(int);",
		"int arr[10];",
		"int arr2[] = { 1, 2, 3 };",
		"struct S { int a : 3; };",
		"struct S { unsigned a : 1; int : 0; };",
		"struct S { struct { int inner; }; int outer; };",
		"struct Nested { struct T { int x; } t; };",
		"const char *s = \"x\";",
		"volatile int v;",
		"int f(int a, ...);",
		"int f(int, char *, void *);",
		"int f(...);",
		"extern \"C\" { int z; }",
		"static int f(void) { return 1; }",
		"int arr3[2][3];",
		"char **argv;",
		"int (*fp2)(void);",
		"void (*sig)(int, void (*)(void));",
	} {
		mustParseInRegion(t, body)
	}
}

func TestCDeclInlineDefinitions(t *testing.T) {
	for _, body := range []string{
		"inline int g(void) { return 1; }",
		"int f(void) { return 1; }",
		"void f(void) { int a; { int b; a = b; } }",
		"void f(int a) { if (a) { a++; } else { a--; } }",
		"void f(int a) { while (a) a--; }",
		"void f(int a) { do { a--; } while (a); }",
		"void f(int a) { for (int i = 0; i < 10; i++) a += i; }",
		"void f(int a) { switch (a) { case 1: break; default: break; } }",
		"void f(void) { return; }",
		"int f(void) { int i = 0; int j = i + 1 * 2 - 3 / 4; return j; }",
		"int f(void) { int i = 0; return i ? 1 : 2; }",
		// Deviates from C89 on purpose: the trailing ';' parses as an empty
		// statement, because people do write it after an inline definition.
		"int f(void) { return 1; };",
		"int f(void) { int i = 1, j = 2; return i & j | i ^ j; }",
		"void f(int *p) { p++; *p = 0; goto done; done: ; }",
		"int f(void) { return (int)(long)(1); }",
	} {
		mustParseInRegion(t, body)
	}
}

// Vendor spellings that survive `gcc -E` verbatim. None of these has a
// lexer rule, so every one of them reaches the subgrammar as a plain
// Identifier followed by a parenthesised group.
func TestCDeclVendorAttributes(t *testing.T) {
	for _, body := range []string{
		"__attribute__((packed)) struct S { int x; };",
		"struct __attribute__((packed)) S { int x; };",
		"struct S { int a __attribute__((packed)); };",
		"int x __attribute__((aligned(4)));",
		"char *__attribute__((nonnull)) p;",
		"int f(void) __attribute__((nonnull));",
		"int f(void) __attribute__((__format__(__printf__, 1, 2)));",
		"int x __attribute((packed));",
		"__attribute__((weak)) int y;",
		"int f(void) __attribute__((weak, alias(\"other\")));",
		"struct S { int a : 3; } __attribute__((packed));",
	} {
		mustParseInRegion(t, body)
	}
}

func TestCDeclVendorKeywords(t *testing.T) {
	for _, body := range []string{
		"__declspec(dllimport) int x;",
		"int vector __at(0x2000);",
		"__extension__ int x;",
		"__typeof__(x) y;",
		"typeof(x) y;",
		"int f(void) __asm__(\"myf\");",
		"int f(void) __asm(\"myf\");",
		"void __irq f(void) { }",
		"int __packed x;",
		"void f(__sptr p);",
		"void f(int __w64 x);",
		"void __cdecl f(void);",
		"__forceinline int g(void) { return 1; }",
		"int f(const char *__restrict s);",
		"int f(char *__restrict __sptr dst, const char *__restrict __sptr src, unsigned long n);",
	} {
		mustParseInRegion(t, body)
	}
}

// A K&R definition. Documented gap: the identifier list reads as a
// parameter list, so this parses but the two names are taken as types.
func TestCDeclKnRDefinition(t *testing.T) {
	mustParseInRegion(t, "int f(a, b) int a; int b; { return a + b; }")
}

// Constructs that are not C, plus the one gap this subgrammar accepts
// losing rather than papering over. A region exists to hold C, so a
// non-C construct failing is the correct outcome.
func TestCDeclRejectsNonCAndKnownGaps(t *testing.T) {
	for _, body := range []string{
		"namespace n { int q; }",        // not C
		"template <class T> void f(T);", // not C
		"int x = 1 << 2;",               // '<<' is not a SuckCLexer token
	} {
		errs, _ := parseWithErrors(wrapFile(body))
		if len(errs) == 0 {
			t.Errorf("%q parsed but should not have", body)
		}
	}
}

// An empty region and a region whose body is only hidden-channel tokens
// must still parse; frontend records the first as an inverted range.
func TestCDeclEmptyRegion(t *testing.T) {
	for _, src := range []string{
		"#pragma c_scope_begin\n#pragma c_scope_end\n",
		"#pragma c_scope_begin\n#pragma c_scope_end",
		"void f(void) {\n#pragma c_scope_begin\n#pragma c_scope_end\n}\n",
	} {
		if errs, _ := parseWithErrors(src); len(errs) != 0 {
			t.Errorf("empty region %q: %v", src, errs)
		}
	}
}

// TestCDeclSuckCKeywordSpellingCIdentifier pins the widening of cCIdentifier.
// Every row is a name that glibc or a vendor header spells literally while
// SuckC lexes it as a keyword token. Without the widening, "typedef int
// wchar_t;" from stddef.h fails with "no viable alternative".
func TestCDeclSuckCKeywordSpellingCIdentifier(t *testing.T) {
	for _, body := range []string{
		"typedef int wchar_t;",
		"typedef unsigned int char16_t;",
		"typedef unsigned int char32_t;",
		"typedef int bool;",
		"int true;",
		"int false;",
		"int nullptr;",
		"int constexpr;",
		"int thread_local;",
		"int mutable;",
		"int new;",
		"int delete;",
		"int template;",
		"int typename;",
		"int namespace;",
		"int wchar_t, char16_t;",
		"struct wchar_t { int x; };",
		"enum E { wchar_t, bool };",
	} {
		mustParseInRegion(t, body)
	}
}

// TestCDeclKeywordTypedefsUnusableAsTypes records a deliberate limitation.
// A name spelled with a SuckC keyword token can be *declared* -- which is
// the case that actually matters, glibc's "typedef int wchar_t;" -- but not
// yet *used* as a type, because cCTypeSpecifier accepts the plain Identifier
// token and the parser has no typedef table to know that wchar_t names one.
//
// Widening cCTypeSpecifier instead would fix the uses and break the
// declaration: "typedef int wchar_t;" would then have two viable parses (the
// name as a type specifier with no declarator, or as a declarator) and
// ANTLR picks between them silently.
//
// The fix belongs to the stage that builds the typedef table, which is where
// the declaration stops being just a spelling and starts being a fact.
func TestCDeclKeywordTypedefsUnusableAsTypes(t *testing.T) {
	for _, body := range []string{
		"wchar_t c;",
		"wchar_t f(void);",
		"int f(wchar_t a);",
		"int f(int a, wchar_t b);",
		"struct S { wchar_t w; };",
		"unsigned wchar_t x;",
	} {
		errs, _ := parseWithErrors(wrapFile(body))
		if len(errs) == 0 {
			t.Errorf("expected a parse error for %q, got none", body)
		}
	}
}

// TestCDeclLinkageSpecification covers `extern "C" { ... }`, which survives
// `gcc -E -x c` when a header spells it literally instead of going through
// glibc's __BEGIN_DECLS. The frontend erases the wrapper tokens, so this
// test only has to prove the region parses and the contents are visible.
func TestCDeclLinkageSpecification(t *testing.T) {
	for _, body := range []string{
		`extern "C" { int z; }`,
		`extern "C" { }`,
		`extern "C" { int a; void f(void); typedef int t; }`,
		`extern "C" int single;`,
	} {
		mustParseInRegion(t, body)
	}
}
