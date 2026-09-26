package sema

import (
	"testing"

	"github.com/starofrainnight/suckc/internal/frontend"
)

func TestPass1CollectsFunctionsAndFileVars(t *testing.T) {
	src := `int g = 1;
int arr[4] = {1,2,3,4};
int add(int a, int b);
int add(int a, int b) { return a + b; }
void nothing(void) { int local = 3; }
`
	res, err := frontend.ParseSource(src)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	a := &Analyzer{Opts: Options{TargetBits: 32, FileName: "t.suckc"}}
	if err := pass1(res, a); err != nil {
		t.Fatalf("pass1: %v", err)
	}
	if got, ok := a.Funcs["add"]; !ok || got != (Type{"int", 0}) {
		t.Errorf("Funcs[add] = %+v, %v; want int", got, ok)
	}
	if got, ok := a.Funcs["nothing"]; !ok || got != (Type{"void", 0}) {
		t.Errorf("Funcs[nothing] = %+v, %v; want void", got, ok)
	}
	if got, ok := a.FileVars["g"]; !ok || got.Type != (Type{"int", 0}) || got.IsArray {
		t.Errorf("FileVars[g] = %+v, %v", got, ok)
	}
	if got, ok := a.FileVars["arr"]; !ok || !got.IsArray || got.Type != (Type{"int", 0}) {
		t.Errorf("FileVars[arr] = %+v, %v; want element int, IsArray", got, ok)
	}
	if _, ok := a.FileVars["local"]; ok {
		t.Error("block-scope local must not be a file var")
	}
}

func TestPass1FunctionPointerReturnNotCollected(t *testing.T) {
	src := `int (*fp)(void);
`
	res, err := frontend.ParseSource(src)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	a := &Analyzer{Opts: Options{TargetBits: 32, FileName: "t.suckc"}}
	if err := pass1(res, a); err != nil {
		t.Fatalf("pass1: %v", err)
	}
	if _, ok := a.Funcs["fp"]; ok {
		t.Error("function-pointer variable must not enter Funcs")
	}
}
