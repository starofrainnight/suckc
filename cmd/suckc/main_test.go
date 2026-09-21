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
