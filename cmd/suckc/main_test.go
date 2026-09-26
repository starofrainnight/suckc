package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func runCommand(args ...string) (string, string, error) {
	cmd := newRootCmd()
	cmd.SetArgs(args)

	oldStdout := os.Stdout
	stdoutR, stdoutW, _ := os.Pipe()
	os.Stdout = stdoutW

	oldStderr := os.Stderr
	stderrR, stderrW, _ := os.Pipe()
	os.Stderr = stderrW

	err := cmd.Execute()

	stdoutW.Close()
	stderrW.Close()
	os.Stdout = oldStdout
	os.Stderr = oldStderr

	var stdoutBuf, stderrBuf bytes.Buffer
	stdoutBuf.ReadFrom(stdoutR)
	stderrBuf.ReadFrom(stderrR)

	return stdoutBuf.String(), stderrBuf.String(), err
}

func TestVersion(t *testing.T) {
	stdout, _, err := runCommand("--version")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !strings.Contains(stdout, "suckc version 0.1.0") {
		t.Errorf("expected version output, got: %s", stdout)
	}
}

func TestUnknownFlag(t *testing.T) {
	_, stderr, err := runCommand("--bogus")
	if err == nil {
		t.Fatal("expected error for unknown flag")
	}
	if !strings.Contains(stderr, "bogus") {
		t.Errorf("expected stderr to mention 'bogus', got: %s", stderr)
	}
}

func TestMissingFile(t *testing.T) {
	_, stderr, err := runCommand("nonexistent.suckc")
	if err == nil {
		t.Fatal("expected error for missing file")
	}
	if !strings.Contains(stderr, "file not found") {
		t.Errorf("expected 'file not found' in stderr, got: %s", stderr)
	}
}

func TestGenerateFile(t *testing.T) {
	dir := t.TempDir()
	src := "// test\nint a = 0;\n"
	in := filepath.Join(dir, "hello.suckc")
	if err := os.WriteFile(in, []byte(src), 0o644); err != nil {
		t.Fatalf("write input: %v", err)
	}

	stdout, _, err := runCommand(in)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	out := filepath.Join(dir, "hello.c")
	got, err := os.ReadFile(out)
	if err != nil {
		t.Fatalf("expected generated %s, got %v", out, err)
	}
	if string(got) != src {
		t.Errorf("expected output to equal input, got: %q", got)
	}
	if !strings.Contains(stdout, "generated "+out) {
		t.Errorf("expected generated message, got: %s", stdout)
	}
}

func TestGenerateFileKeepsDir(t *testing.T) {
	dir := t.TempDir()
	in := filepath.Join(dir, "sub", "deep.suckc")
	if err := os.MkdirAll(filepath.Dir(in), 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	if err := os.WriteFile(in, []byte("int x;\n"), 0o644); err != nil {
		t.Fatalf("write input: %v", err)
	}

	if _, _, err := runCommand(in); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if _, err := os.Stat(filepath.Join(dir, "sub", "deep.c")); err != nil {
		t.Errorf("expected deep.c next to deep.suckc, got %v", err)
	}
}

func TestDebugFlag(t *testing.T) {
	stdout, _, err := runCommand("-d", "../../SuckC.suckc")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !strings.Contains(stdout, "(translationUnit") {
		t.Errorf("expected parse tree dump, got: %s", stdout)
	}
}

func TestNoArgs(t *testing.T) {
	stdout, stderr, err := runCommand()
	if err == nil {
		t.Fatal("expected error for no args")
	}
	if !strings.Contains(stderr, "suckc") && !strings.Contains(stdout, "suckc") {
		t.Errorf("expected usage to be shown, stdout: %s, stderr: %s", stdout, stderr)
	}
}

func TestDirectoryAsFile(t *testing.T) {
	_, stderr, err := runCommand("src")
	if err == nil {
		t.Fatal("expected error for directory")
	}
	if !strings.Contains(stderr, "file not found") {
		t.Errorf("expected 'file not found' in stderr, got: %s", stderr)
	}
}

func TestTargetBitsValidation(t *testing.T) {
	dir := t.TempDir()
	in := filepath.Join(dir, "v.suckc")
	if err := os.WriteFile(in, []byte("void f() {}\n"), 0o644); err != nil {
		t.Fatalf("write input: %v", err)
	}
	_, stderr, err := runCommand("--target-bits", "24", in)
	if err == nil {
		t.Fatal("expected error for --target-bits 24")
	}
	if !strings.Contains(stderr, "invalid --target-bits 24") {
		t.Errorf("stderr = %s", stderr)
	}
}

func TestAutoEndToEnd(t *testing.T) {
	dir := t.TempDir()
	src := "int main(void) {\n\tauto i = 12;\n\treturn i;\n}\n"
	in := filepath.Join(dir, "e2e.suckc")
	if err := os.WriteFile(in, []byte(src), 0o644); err != nil {
		t.Fatalf("write input: %v", err)
	}
	if _, _, err := runCommand(in); err != nil {
		t.Fatalf("run: %v", err)
	}
	out, err := os.ReadFile(filepath.Join(dir, "e2e.c"))
	if err != nil {
		t.Fatalf("read output: %v", err)
	}
	if !strings.Contains(string(out), "int i = 12") {
		t.Errorf("output missing deduction:\n%s", out)
	}
	if strings.Contains(string(out), "auto") {
		t.Errorf("output still contains auto:\n%s", out)
	}
}

func TestAutoTargetBitsEndToEnd(t *testing.T) {
	dir := t.TempDir()
	src := "int main(void) {\n\tauto n = 40000;\n\treturn 0;\n}\n"
	in := filepath.Join(dir, "bits.suckc")
	if err := os.WriteFile(in, []byte(src), 0o644); err != nil {
		t.Fatalf("write input: %v", err)
	}
	if _, _, err := runCommand("--target-bits", "16", in); err != nil {
		t.Fatalf("run 16: %v", err)
	}
	out16, _ := os.ReadFile(filepath.Join(dir, "bits.c"))
	if !strings.Contains(string(out16), "long n = 40000") {
		t.Errorf("bits16 output:\n%s", out16)
	}
	if _, _, err := runCommand("--target-bits", "32", in); err != nil {
		t.Fatalf("run 32: %v", err)
	}
	out32, _ := os.ReadFile(filepath.Join(dir, "bits.c"))
	if !strings.Contains(string(out32), "int n = 40000") {
		t.Errorf("bits32 output:\n%s", out32)
	}
}
