package sema

import "testing"

func TestStripCV(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"const int", "int"},
		{"int", "int"},
		{"const volatile int", "int"},
		{"unsigned int", "unsigned int"},
	}
	for _, c := range cases {
		if got := stripCV(c.in); got != c.want {
			t.Errorf("stripCV(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestValueReadStripsCVAddressOfKeepsIt(t *testing.T) {
	src := "const int c = 5;\nvoid f() { auto j = c; auto p = &c; }\n"
	subs, err := analyzeSrc(t, src, Options{TargetBits: 32, FileName: "t.suckc"})
	if err != nil {
		t.Fatalf("analyze: %v", err)
	}
	got := allAutoSubs(t, src, subs)
	want := []string{"int", "const int *"}
	if len(got) != len(want) {
		t.Fatalf("subs = %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("subs[%d] = %q, want %q", i, got[i], want[i])
		}
	}
}

func TestEnumConstantDeducesInt(t *testing.T) {
	src := "enum Color { RED, GREEN = 3 };\nvoid f() { auto x = RED; }\n"
	subs, err := analyzeSrc(t, src, Options{TargetBits: 32, FileName: "t.suckc"})
	if err != nil {
		t.Fatalf("analyze: %v", err)
	}
	got := allAutoSubs(t, src, subs)
	if len(got) != 1 || got[0] != "int" {
		t.Errorf("subs = %v, want [int]", got)
	}
}
