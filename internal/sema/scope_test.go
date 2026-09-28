package sema

import "testing"

func TestScopeShadowing(t *testing.T) {
	s := &scope{}
	s.push()
	s.declare("x", binding{Type: Type{"int", 0}, IsArray: false})
	s.push()
	s.declare("x", binding{Type: Type{"long", 0}, IsArray: false})
	b, ok := s.lookup("x")
	if !ok || b.Type != (Type{"long", 0}) {
		t.Fatalf("inner lookup = %+v, %v", b, ok)
	}
	s.pop()
	b, ok = s.lookup("x")
	if !ok || b.Type != (Type{"int", 0}) {
		t.Fatalf("outer lookup = %+v, %v", b, ok)
	}
	s.pop()
	if _, ok := s.lookup("x"); ok {
		t.Fatal("expected miss after popping all frames")
	}
}

func TestScopeUnknownName(t *testing.T) {
	s := &scope{}
	if _, ok := s.lookup("nope"); ok {
		t.Fatal("expected miss on empty scope")
	}
}
