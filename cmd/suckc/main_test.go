package main

import (
	"bytes"
	"os"
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

func TestExistingFile(t *testing.T) {
	stdout, _, err := runCommand("../../SuckC.suckc")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	expected := "suckc: ../../SuckC.suckc: transpilation not implemented yet"
	if !strings.Contains(stdout, expected) {
		t.Errorf("expected placeholder output, got: %s", stdout)
	}
}

func TestDebugFlag(t *testing.T) {
	stdout, _, err := runCommand("-d", "../../SuckC.suckc")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !strings.Contains(stdout, "transpilation not implemented yet") {
		t.Errorf("expected placeholder output, got: %s", stdout)
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
