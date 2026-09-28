package sema

import (
	"testing"

	"github.com/starofrainnight/suckc/internal/frontend"
)

func TestPass1CollectsEnumConstants(t *testing.T) {
	res, err := frontend.ParseSource("enum Color { RED, GREEN = 3 };\n")
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	a := &Analyzer{Opts: Options{TargetBits: 32, FileName: "t.suckc"}}
	if err := pass1(res, a); err != nil {
		t.Fatalf("pass1: %v", err)
	}
	if !a.Enums["RED"] {
		t.Errorf("RED not collected: %v", a.Enums)
	}
	if !a.Enums["GREEN"] {
		t.Errorf("GREEN not collected: %v", a.Enums)
	}
	if a.Enums["g"] {
		t.Errorf("unexpected key g: %v", a.Enums)
	}
}

func TestPass1CollectsConstInits(t *testing.T) {
	src := "const int c = 5;\nconst int d = c;\nint e = 1;\n"
	res, err := frontend.ParseSource(src)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	a := &Analyzer{Opts: Options{TargetBits: 32, FileName: "t.suckc"}}
	if err := pass1(res, a); err != nil {
		t.Fatalf("pass1: %v", err)
	}
	for _, name := range []string{"c", "d"} {
		b, ok := a.FileVars[name]
		if !ok {
			t.Fatalf("%s missing from FileVars", name)
		}
		if !b.IsConst {
			t.Errorf("%s: IsConst = false, want true", name)
		}
		if b.Type.Text() != "const int" {
			t.Errorf("%s: type = %q, want %q", name, b.Type.Text(), "const int")
		}
		if _, ok := a.ConstInits[name]; !ok {
			t.Errorf("%s: no recorded initializer", name)
		}
	}
	if _, ok := a.ConstInits["e"]; ok {
		t.Errorf("non-const e must not have a recorded initializer")
	}
	if a.FileVars["e"].IsConst {
		t.Errorf("e: IsConst = true, want false")
	}
}

func TestPass1RecordsPendingFileAutos(t *testing.T) {
	src := "int a = 0;\nauto *p = &a;\nauto i = 12;\n"
	res, err := frontend.ParseSource(src)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	a := &Analyzer{Opts: Options{TargetBits: 32, FileName: "t.suckc"}}
	if err := pass1(res, a); err != nil {
		t.Fatalf("pass1: %v", err)
	}
	if len(a.pendingAutos) != 2 {
		t.Fatalf("pendingAutos = %d, want 2", len(a.pendingAutos))
	}
	for i, want := range []string{"p", "i"} {
		if a.pendingAutos[i].Name != want {
			t.Errorf("pendingAutos[%d].Name = %q, want %q", i, a.pendingAutos[i].Name, want)
		}
		if a.pendingAutos[i].State != autoPending {
			t.Errorf("pendingAutos[%d].State = %v, want autoPending", i, a.pendingAutos[i].State)
		}
	}
}
