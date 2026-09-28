package sema

import (
	"strings"
	"testing"

	"github.com/starofrainnight/suckc/internal/frontend"
	"github.com/starofrainnight/suckc/internal/parser"
)

// analyzeSrc runs the full pipeline and returns Subs.
func analyzeSrc(t *testing.T, src string, opts Options) (Subs, error) {
	t.Helper()
	res, err := frontend.ParseSource(src)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if opts.FileName == "" {
		opts.FileName = "test.suckc"
	}
	return Analyze(res, opts)
}

// subForAuto returns the substitution of the first Auto token.
func subForAuto(t *testing.T, src string, subs Subs) string {
	t.Helper()
	res, err := frontend.ParseSource(src)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	for _, tok := range res.Tokens {
		if tok.GetTokenType() == parser.SuckCParserAuto {
			return subs[tok.GetTokenIndex()]
		}
	}
	t.Fatal("no Auto token in source")
	return ""
}

// allAutoSubs returns the substitution of every Auto token in src, in
// source order. Subs is keyed by token index, so the order has to come
// from the token stream rather than from the map.
func allAutoSubs(t *testing.T, src string, subs Subs) []string {
	t.Helper()
	res, err := frontend.ParseSource(src)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	var out []string
	for _, tok := range res.Tokens {
		if tok.GetTokenType() == parser.SuckCParserAuto {
			out = append(out, subs[tok.GetTokenIndex()])
		}
	}
	return out
}

func TestDeduceIdentifierForm(t *testing.T) {
	src := "void f() {\n\tauto i = 12;\n}\n"
	subs, err := analyzeSrc(t, src, Options{TargetBits: 32})
	if err != nil {
		t.Fatalf("analyze: %v", err)
	}
	if got := subForAuto(t, src, subs); got != "int" {
		t.Errorf("sub = %q, want %q", got, "int")
	}
}

func TestDeducePointerForm(t *testing.T) {
	src := "void f() {\n\tint x = 0;\n\tauto *p = &x;\n}\n"
	subs, err := analyzeSrc(t, src, Options{TargetBits: 32})
	if err != nil {
		t.Fatalf("analyze: %v", err)
	}
	if got := subForAuto(t, src, subs); got != "int" {
		t.Errorf("auto *p sub = %q, want %q", got, "int")
	}
}

func TestDeduceAddressFormKeepsStar(t *testing.T) {
	src := "void f() {\n\tint x = 0;\n\tauto p = &x;\n}\n"
	subs, err := analyzeSrc(t, src, Options{TargetBits: 32})
	if err != nil {
		t.Fatalf("analyze: %v", err)
	}
	if got := subForAuto(t, src, subs); got != "int *" {
		t.Errorf("auto p sub = %q, want %q", got, "int *")
	}
}

func TestDeduceArrayForm(t *testing.T) {
	src := "void f() {\n\tauto a[] = {1, 2, 3};\n}\n"
	subs, err := analyzeSrc(t, src, Options{TargetBits: 32})
	if err != nil {
		t.Fatalf("analyze: %v", err)
	}
	if got := subForAuto(t, src, subs); got != "int" {
		t.Errorf("array sub = %q, want %q", got, "int")
	}
}

func TestDeduceArraySizedForm(t *testing.T) {
	src := "void f() {\n\tauto a[10] = {1, 2, 3};\n}\n"
	subs, err := analyzeSrc(t, src, Options{TargetBits: 32})
	if err != nil {
		t.Fatalf("analyze: %v", err)
	}
	if got := subForAuto(t, src, subs); got != "int" {
		t.Errorf("sized array sub = %q, want %q", got, "int")
	}
}

func TestDeduceArrayStringInit(t *testing.T) {
	src := "void f() {\n\tauto s[] = \"abc\";\n}\n"
	subs, err := analyzeSrc(t, src, Options{TargetBits: 32})
	if err != nil {
		t.Fatalf("analyze: %v", err)
	}
	if got := subForAuto(t, src, subs); got != "char" {
		t.Errorf("string array sub = %q, want %q", got, "char")
	}
}

func TestDeduceForInit(t *testing.T) {
	src := "void f() {\n\tfor (auto i = 0; i < 10; i++) {\n\t\tint y = i;\n\t}\n}\n"
	subs, err := analyzeSrc(t, src, Options{TargetBits: 32})
	if err != nil {
		t.Fatalf("analyze: %v", err)
	}
	if got := subForAuto(t, src, subs); got != "int" {
		t.Errorf("for-init sub = %q, want %q", got, "int")
	}
}

func TestDeduceUsesEarlierDeclaration(t *testing.T) {
	src := "void f() {\n\tint a = 5;\n\tauto b = a;\n\tauto c = b + 1;\n}\n"
	subs, err := analyzeSrc(t, src, Options{TargetBits: 32})
	if err != nil {
		t.Fatalf("analyze: %v", err)
	}
	res, _ := frontend.ParseSource(src)
	count := 0
	for _, tok := range res.Tokens {
		if tok.GetTokenType() == parser.SuckCParserAuto {
			count++
			if subs[tok.GetTokenIndex()] != "int" {
				t.Errorf("auto #%d sub = %q, want int", count, subs[tok.GetTokenIndex()])
			}
		}
	}
	if count != 2 {
		t.Fatalf("expected 2 Auto tokens, got %d", count)
	}
}

func TestDeduceNestedBlockShadowing(t *testing.T) {
	src := "void f() {\n\tauto x = 1;\n\t{\n\t\tauto x = 1L;\n\t\tlong y = x;\n\t}\n}\n"
	subs, err := analyzeSrc(t, src, Options{TargetBits: 32})
	if err != nil {
		t.Fatalf("analyze: %v", err)
	}
	res, _ := frontend.ParseSource(src)
	var got []string
	for _, tok := range res.Tokens {
		if tok.GetTokenType() == parser.SuckCParserAuto {
			got = append(got, subs[tok.GetTokenIndex()])
		}
	}
	if len(got) != 2 || got[0] != "int" || got[1] != "long" {
		t.Errorf("shadowing subs = %v, want [int long]", got)
	}
}

func TestDeduceCallBeforeDefinition(t *testing.T) {
	src := "int g(void) {\n\tauto v = helper();\n\treturn v;\n}\nint helper(void) { return 7; }\n"
	subs, err := analyzeSrc(t, src, Options{TargetBits: 32})
	if err != nil {
		t.Fatalf("analyze: %v", err)
	}
	if got := subForAuto(t, src, subs); got != "int" {
		t.Errorf("call-before-definition sub = %q, want int", got)
	}
}

func TestForbiddenPositions(t *testing.T) {
	cases := []struct {
		name string
		src  string
	}{
		{"function return", "auto f() { return 1; }\n"},
		{"parameters", "void f(auto x) {}\n"},
		{"decltype", "void f() {\n\tauto x = 1;\n\t(void)x;\n}\ndecltype(auto) y = 1;\n"},
	}
	for _, c := range cases {
		_, err := analyzeSrc(t, c.src, Options{TargetBits: 32})
		if err == nil {
			t.Errorf("%s: expected error", c.name)
			continue
		}
		var se *Error
		if !asSemaError(err, &se) {
			t.Errorf("%s: error is %T, want *sema.Error", c.name, err)
			continue
		}
		if !strings.Contains(se.Msg, "'auto' not allowed here") {
			t.Errorf("%s: msg = %q", c.name, se.Msg)
		}
		if se.Line < 1 || se.Col < 1 {
			t.Errorf("%s: position line=%d col=%d", c.name, se.Line, se.Col)
		}
	}
}

func TestMissingInitializerError(t *testing.T) {
	src := "void f() {\n\tauto i;\n}\n"
	_, err := analyzeSrc(t, src, Options{TargetBits: 32})
	if err == nil {
		t.Fatal("expected error")
	}
	var se *Error
	if !asSemaError(err, &se) {
		t.Fatalf("error type %T", err)
	}
	if !strings.Contains(se.Error(), "cannot deduce type for 'auto': missing initializer") {
		t.Errorf("error = %q", se.Error())
	}
	if !strings.HasPrefix(se.Error(), "test.suckc:") {
		t.Errorf("error missing file prefix: %q", se.Error())
	}
}

func TestInsufficientStarsError(t *testing.T) {
	src := "void f() {\n\tauto *p = 12;\n}\n"
	_, err := analyzeSrc(t, src, Options{TargetBits: 32})
	var se *Error
	if !asSemaError(err, &se) {
		t.Fatalf("error type %T, err=%v", err, err)
	}
	if !strings.Contains(se.Msg, "insufficient pointer depth") {
		t.Errorf("msg = %q", se.Msg)
	}
}

func TestUnknownIdentifierError(t *testing.T) {
	src := "void f() {\n\tauto x = undeclared;\n}\n"
	_, err := analyzeSrc(t, src, Options{TargetBits: 32})
	var se *Error
	if !asSemaError(err, &se) {
		t.Fatalf("error type %T, err=%v", err, err)
	}
	if !strings.Contains(se.Msg, "unknown identifier 'undeclared'") {
		t.Errorf("msg = %q", se.Msg)
	}
}

func TestEmptyBraceInitError(t *testing.T) {
	src := "void f() {\n\tauto a[] = {};\n}\n"
	_, err := analyzeSrc(t, src, Options{TargetBits: 32})
	var se *Error
	if !asSemaError(err, &se) {
		t.Fatalf("error type %T, err=%v", err, err)
	}
	if !strings.Contains(se.Msg, "empty initializer list") {
		t.Errorf("msg = %q", se.Msg)
	}
}

func TestFuncPtrDeclaratorError(t *testing.T) {
	src := "int g(void) { return 0; }\nvoid f() {\n\tauto (*fp)(void) = g;\n}\n"
	_, err := analyzeSrc(t, src, Options{TargetBits: 32})
	var se *Error
	if !asSemaError(err, &se) {
		t.Fatalf("error type %T, err=%v", err, err)
	}
	if !strings.Contains(se.Msg, "function-pointer declarator not supported") {
		t.Errorf("msg = %q", se.Msg)
	}
}

func TestWideCharError(t *testing.T) {
	src := "void f() {\n\tauto c = L'a';\n}\n"
	_, err := analyzeSrc(t, src, Options{TargetBits: 32})
	var se *Error
	if !asSemaError(err, &se) {
		t.Fatalf("error type %T, err=%v", err, err)
	}
	if !strings.Contains(se.Msg, "wide character literal not supported") {
		t.Errorf("msg = %q", se.Msg)
	}
}

func TestTargetBits16ChangesLiteral(t *testing.T) {
	src := "void f() {\n\tauto n = 40000;\n}\n"
	subs32, err := analyzeSrc(t, src, Options{TargetBits: 32})
	if err != nil {
		t.Fatalf("bits32: %v", err)
	}
	subs16, err := analyzeSrc(t, src, Options{TargetBits: 16})
	if err != nil {
		t.Fatalf("bits16: %v", err)
	}
	if subForAuto(t, src, subs32) != "int" {
		t.Errorf("bits32 = %q, want int", subForAuto(t, src, subs32))
	}
	if subForAuto(t, src, subs16) != "long" {
		t.Errorf("bits16 = %q, want long", subForAuto(t, src, subs16))
	}
}

// asSemaError reports whether err is a *Error.
func asSemaError(err error, dst **Error) bool {
	e, ok := err.(*Error)
	if ok {
		*dst = e
	}
	return ok
}

func TestFileScopeAutoDeducesFromEarlierDecl(t *testing.T) {
	src := "int a = 0;\nauto *ff = &a;\n"
	subs, err := analyzeSrc(t, src, Options{TargetBits: 32})
	if err != nil {
		t.Fatalf("analyze: %v", err)
	}
	// Only the auto token is replaced, so the declarator's own star stays:
	// the emitted line reads "int *ff = &a;".
	if got := allAutoSubs(t, src, subs); len(got) != 1 || got[0] != "int" {
		t.Errorf("subs = %v, want [int]", got)
	}
}

func TestFileScopeAutoAnyOrder(t *testing.T) {
	src := "auto *p = &q;\nauto *q = &i;\nint i = 3;\n"
	subs, err := analyzeSrc(t, src, Options{TargetBits: 32})
	if err != nil {
		t.Fatalf("analyze: %v", err)
	}
	if got := allAutoSubs(t, src, subs); len(got) != 2 || got[0] != "int *" || got[1] != "int" {
		t.Errorf("subs = %v, want [int * int]", got)
	}
}

func TestFileScopeAutoCycle(t *testing.T) {
	src := "auto p = &q;\nauto *q = &p;\n"
	_, err := analyzeSrc(t, src, Options{TargetBits: 32})
	if err == nil {
		t.Fatal("expected a circular dependency error")
	}
	// Deduction starts at the first declaration, so the re-entered name is
	// 'p': the declaration whose type transitively depends on itself.
	if !strings.Contains(err.Error(), "circular dependency on 'p'") {
		t.Errorf("error = %v, want circular dependency on 'p'", err)
	}
}

func TestFileScopeAutoVisibleInBlockScope(t *testing.T) {
	src := "int a = 0;\nauto *ff = &a;\nvoid f() { auto q = ff; }\n"
	subs, err := analyzeSrc(t, src, Options{TargetBits: 32})
	if err != nil {
		t.Fatalf("analyze: %v", err)
	}
	if got := allAutoSubs(t, src, subs); len(got) != 2 || got[0] != "int" || got[1] != "int *" {
		t.Errorf("subs = %v, want [int int *]", got)
	}
}

func TestFileScopeAutoFromConstValue(t *testing.T) {
	src := "const int c = 5;\nauto j = c;\n"
	subs, err := analyzeSrc(t, src, Options{TargetBits: 32})
	if err != nil {
		t.Fatalf("analyze: %v", err)
	}
	if got := allAutoSubs(t, src, subs); len(got) != 1 || got[0] != "int" {
		t.Errorf("subs = %v, want [int]", got)
	}
}

func TestFileScopeAutoMissingInitializer(t *testing.T) {
	src := "auto i;\n"
	_, err := analyzeSrc(t, src, Options{TargetBits: 32})
	if err == nil {
		t.Fatal("expected error")
	}
	if !strings.Contains(err.Error(), "missing initializer") {
		t.Errorf("error = %v, want missing initializer", err)
	}
}

func TestFileScopeAutoNonConstantRejected(t *testing.T) {
	src := "int a = 0;\nauto x = a;\n"
	_, err := analyzeSrc(t, src, Options{TargetBits: 32})
	if err == nil {
		t.Fatal("expected error")
	}
	if !strings.Contains(err.Error(), "not a constant expression") {
		t.Errorf("error = %v, want not a constant expression", err)
	}
}

func TestFileScopeAutoTargetBits(t *testing.T) {
	src := "auto n = 40000;\n"
	subs, err := analyzeSrc(t, src, Options{TargetBits: 16})
	if err != nil {
		t.Fatalf("analyze: %v", err)
	}
	if got := allAutoSubs(t, src, subs); len(got) != 1 || got[0] != "long" {
		t.Errorf("subs = %v, want [long]", got)
	}
}
