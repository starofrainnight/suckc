package sema

import (
	"strings"
	"testing"
)

// A C-scope region is raw C to SuckC: C23 `auto` and `const` are legal
// there, so neither the forbidden-auto scan nor the constant gate may look
// at the body. The same constructs outside a region must still report.
func TestRegionHidesAutoAndConstFromSema(t *testing.T) {
	inRegion := "#pragma c_scope_begin\n" +
		"auto int ra = 1;\n" +
		"const int rc = f();\n" +
		"auto x = rc;\n" +
		"#pragma c_scope_end\n" +
		"void f(void) {}\n"
	if _, err := analyzeSrc(t, inRegion, Options{TargetBits: 32}); err != nil {
		t.Errorf("region body produced a diagnostic: %v", err)
	}

	outside := []struct {
		name string
		src  string
		want string
	}{
		{"forbidden auto", "auto f() { return 1; }\n", "'auto' not allowed here"},
		{"const gate", "const int rc = f();\nauto x = rc;\n", "not a constant expression"},
	}
	for _, c := range outside {
		_, err := analyzeSrc(t, c.src, Options{TargetBits: 32})
		if err == nil {
			t.Errorf("%s: expected an error outside a region", c.name)
			continue
		}
		if !strings.Contains(err.Error(), c.want) {
			t.Errorf("%s: err = %v, want substring %q", c.name, err, c.want)
		}
	}
}

// The forbidden-auto diagnostic reports SourceAt: with line markers the
// originating file wins over Options.FileName, without them the option is
// the fallback.
func TestForbiddenAutoUsesSourceAt(t *testing.T) {
	_, err := analyzeSrc(t, "auto f() { return 1; }\n", Options{TargetBits: 32})
	if err == nil {
		t.Fatal("expected an error")
	}
	var se *Error
	if !asSemaError(err, &se) {
		t.Fatalf("error type %T", err)
	}
	if se.File != "test.suckc" {
		t.Errorf("file = %q, want Options.FileName fallback", se.File)
	}

	// `# 1 "vendor.h" 1` puts the following line at vendor.h line 1.
	marked := "# 1 \"vendor.h\" 1\nauto f() { return 1; }\n"
	_, err = analyzeSrc(t, marked, Options{TargetBits: 32})
	if err == nil {
		t.Fatal("expected an error")
	}
	if !asSemaError(err, &se) {
		t.Fatalf("error type %T", err)
	}
	if se.File != "vendor.h" {
		t.Errorf("file = %q, want vendor.h from the line marker", se.File)
	}
	if se.Line != 1 {
		t.Errorf("line = %d, want 1 (vendor.h numbering)", se.Line)
	}
}
