package backend

import (
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/starofrainnight/suckc/internal/frontend"
)

var gccLineMarkerRE = regexp.MustCompile(`(?m)^# \d+ "[^"]*"(?: \d+)*$`)

// eraseMarkers returns src with the region pragmas removed, the exact text
// Generate must produce (the newlines that ended the marker lines survive as
// blank lines). gcc -E line markers are expected to survive, so they are
// left alone here.
func eraseMarkers(src string) string {
	out := strings.Replace(src, "#pragma c_scope_begin", "", 1)
	out = strings.Replace(out, "#pragma c_scope_end", "", 1)
	return out
}

// The emitted C must lose both region markers while every byte of the
// body — vendor extensions, pragmas and dialect literals — stays.
func TestGenerateErasesRegionMarkersKeepsBody(t *testing.T) {
	body := "struct P { int a; } __attribute__((packed));\n" +
		"#pragma pack(1)\n" +
		"double d = 1.5;\n" +
		"long double ld = 1.5L;\n" +
		"#pragma pack()\n"
	src := "#pragma c_scope_begin\n" + body + "#pragma c_scope_end\nvoid f(void) {}\n"

	res, err := frontend.ParseSource(src)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	got, err := Generate(res, nil)
	if err != nil {
		t.Fatalf("generate: %v", err)
	}
	if strings.Contains(got, "c_scope_begin") || strings.Contains(got, "c_scope_end") {
		t.Errorf("region markers not erased:\n%s", got)
	}
	if !strings.Contains(got, body) {
		t.Errorf("region body not kept byte-identical:\ngot:\n%s\nwant substring:\n%s", got, body)
	}
	if want := eraseMarkers(src); got != want {
		t.Errorf("output differs from input modulo erasures:\ngot:\n%s\nwant:\n%s", got, want)
	}
}

// Real gcc -E output: the region markers disappear, the `# N "file"` line
// markers survive verbatim, everything else round trips. The markers are
// hidden tokens that still name the file and line each declaration came
// from, so they stay; the diagnostic side (SourceAt) reads the same table.
// Fixture command (run from /tmp/opencode/cscope):
//
//	gcc -E -x c first.suckc > internal/frontend/testdata/region_line_markers.c
func TestGenerateKeepsLineMarkersFromGccPreprocess(t *testing.T) {
	const fixture = "../frontend/testdata/region_line_markers.c"
	src, err := os.ReadFile(fixture)
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}
	res, err := frontend.ParseFile(fixture)
	if err != nil {
		t.Fatalf("parse fixture: %v", err)
	}
	got, err := Generate(res, nil)
	if err != nil {
		t.Fatalf("generate: %v", err)
	}

	kept := 0
	for _, line := range strings.Split(got, "\n") {
		if gccLineMarkerRE.MatchString(line) {
			kept++
		}
	}
	if want := len(gccLineMarkerRE.FindAllString(string(src), -1)); kept != want {
		t.Errorf("kept %d line markers, want %d:\n%s", kept, want, got)
	}
	if strings.Contains(got, "c_scope_begin") || strings.Contains(got, "c_scope_end") {
		t.Errorf("region markers not erased:\n%s", got)
	}
	for _, want := range []string{
		"typedef unsigned long vendor_size_t;",
		"__attribute__((nonnull))",
		"static double vendor_get_half(void) { return 1.5; }",
		"static long double vendor_get_ld(void) { return 1.5L; }",
		"#pragma pack(1)",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("region body text lost: %q\noutput:\n%s", want, got)
		}
	}
	if want := eraseMarkers(string(src)); got != want {
		t.Errorf("output differs from input modulo erasures:\ngot:\n%s\nwant:\n%s", got, want)
	}
}

// A linkage specification's four wrapper tokens must be dropped: the
// emitted file is C89, where `extern "C" {` is a hard error ("expected
// identifier or '(' before string constant"). The declarations inside it
// stay, because a file-scope declaration is already external in C.
func TestGenerateErasesLinkageSpecification(t *testing.T) {
	for _, tc := range []struct {
		name    string
		src     string
		kept    []string
		dropped []string
	}{
		{
			name: "suckc level",
			src:  "extern \"C\" {\nint z;\n}\nint w;\n",
			kept: []string{"int z;", "int w;"},
		},
		{
			name: "inside a c region",
			src: "#pragma c_scope_begin\n" +
				"extern \"C\" {\nint vendor_init(unsigned long flags);\n}\n" +
				"#pragma c_scope_end\n",
			kept: []string{"int vendor_init(unsigned long flags);"},
		},
		{
			name:    "single declaration form",
			src:     "extern \"C\" int single;\n",
			kept:    []string{"int single;"},
			dropped: []string{"extern", "\"C\""},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			res, err := frontend.ParseSource(tc.src)
			if err != nil {
				t.Fatalf("parse: %v", err)
			}
			got, err := Generate(res, nil)
			if err != nil {
				t.Fatalf("generate: %v", err)
			}
			if strings.Contains(got, `extern "C"`) {
				t.Errorf("linkage wrapper survived:\n%s", got)
			}
			if strings.Contains(got, "c_scope_") {
				t.Errorf("region marker survived:\n%s", got)
			}
			for _, k := range tc.kept {
				if !strings.Contains(got, k) {
					t.Errorf("declaration %q dropped:\n%s", k, got)
				}
			}
			for _, d := range tc.dropped {
				if strings.Contains(got, d) {
					t.Errorf("token %q survived:\n%s", d, got)
				}
			}
		})
	}
}
