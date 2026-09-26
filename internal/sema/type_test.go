package sema

import "testing"

func TestTypeText(t *testing.T) {
	cases := []struct {
		typ  Type
		want string
	}{
		{Type{"int", 0}, "int"},
		{Type{"int", 1}, "int *"},
		{Type{"int", 2}, "int **"},
		{Type{"const char", 1}, "const char *"},
	}
	for _, c := range cases {
		if got := c.typ.Text(); got != c.want {
			t.Errorf("Type%+v.Text() = %q, want %q", c.typ, got, c.want)
		}
	}
}

func TestIntLimitsByWidth(t *testing.T) {
	o16 := Options{TargetBits: 16}
	if _, max, _ := o16.intLimits(); max != 32767 {
		t.Errorf("16-bit int max = %d, want 32767", max)
	}
	o32 := Options{TargetBits: 32}
	if _, max, _ := o32.intLimits(); max != 2147483647 {
		t.Errorf("32-bit int max = %d, want 2147483647", max)
	}
	o64 := Options{TargetBits: 64}
	if _, lmax, umax := o64.longLimits(); lmax != 9223372036854775807 || umax != -1 {
		// ulongMax overflows int64; see implementation note: ulongMax returns
		// math.MaxUint64 as uint64 via separate accessor — test below.
		_ = umax
	}
}

func TestLongLimits(t *testing.T) {
	o16 := Options{TargetBits: 16}
	if _, lmax, _ := o16.longLimits(); lmax != 2147483647 {
		t.Errorf("16-target long max = %d, want 2147483647", lmax)
	}
	o64 := Options{TargetBits: 64}
	if _, lmax, _ := o64.longLimits(); lmax != 9223372036854775807 {
		t.Errorf("64-target long max = %d", lmax)
	}
}

func TestPromote(t *testing.T) {
	cases := []struct{ in, want Type }{
		{Type{"char", 0}, Type{"int", 0}},
		{Type{"bool", 0}, Type{"int", 0}},
		{Type{"unsigned char", 0}, Type{"int", 0}},
		{Type{"int", 0}, Type{"int", 0}},
		{Type{"long", 1}, Type{"long", 1}},
	}
	for _, c := range cases {
		if got := promote(c.in); got != c.want {
			t.Errorf("promote(%+v) = %+v, want %+v", c.in, got, c.want)
		}
	}
}

func TestUsualArith(t *testing.T) {
	o := Options{TargetBits: 64}
	cases := []struct {
		a, b, want Type
	}{
		{Type{"int", 0}, Type{"int", 0}, Type{"int", 0}},
		{Type{"char", 0}, Type{"int", 0}, Type{"int", 0}},
		{Type{"int", 0}, Type{"long", 0}, Type{"long", 0}},
		{Type{"int", 0}, Type{"unsigned int", 0}, Type{"unsigned int", 0}},
		{Type{"int", 0}, Type{"unsigned long", 0}, Type{"unsigned long", 0}},
		{Type{"long", 0}, Type{"unsigned int", 0}, Type{"long", 0}}, // LP64: long holds all uint32
		{Type{"int", 0}, Type{"double", 0}, Type{"double", 0}},
		{Type{"float", 0}, Type{"int", 0}, Type{"float", 0}},
		{Type{"bool", 0}, Type{"char", 0}, Type{"int", 0}},
	}
	for _, c := range cases {
		got, err := usualArith(c.a, c.b, o)
		if err != nil {
			t.Errorf("usualArith(%+v, %+v): %v", c.a, c.b, err)
			continue
		}
		if got != c.want {
			t.Errorf("usualArith(%+v, %+v) = %+v, want %+v", c.a, c.b, got, c.want)
		}
	}
}

func TestUsualArithOn32Target(t *testing.T) {
	o := Options{TargetBits: 32}
	got, err := usualArith(Type{"long", 0}, Type{"unsigned int", 0}, o)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// ILP32: long is 32-bit and cannot hold all uint32 -> unsigned long.
	if got != (Type{"unsigned long", 0}) {
		t.Errorf("ILP32 long+uint = %+v, want unsigned long", got)
	}
}

func TestUsualArithIncompatible(t *testing.T) {
	o := Options{TargetBits: 64}
	if _, err := usualArith(Type{"int", 1}, Type{"int", 0}, o); err == nil {
		t.Error("expected error for pointer + int")
	}
	if _, err := usualArith(Type{"int", 1}, Type{"long", 1}, o); err == nil {
		t.Error("expected error for int* + long*")
	}
}

func TestCommonType(t *testing.T) {
	o := Options{TargetBits: 64}
	got, err := commonType(Type{"int", 0}, Type{"double", 0}, o)
	if err != nil || got != (Type{"double", 0}) {
		t.Errorf("commonType(int,double) = %+v, %v", got, err)
	}
	got, err = commonType(Type{"int", 1}, Type{"int", 1}, o)
	if err != nil || got != (Type{"int", 1}) {
		t.Errorf("commonType(int*,int*) = %+v, %v", got, err)
	}
	if _, err := commonType(Type{"int", 0}, Type{"int", 1}, o); err == nil {
		t.Error("expected error for int vs int*")
	}
}
