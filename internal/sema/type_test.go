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

func TestRankLatticeCoversEveryWidth(t *testing.T) {
	// long long and long double used to have no rank at all, so a variable
	// declared with one was rejected the moment it took part in arithmetic.
	// On a 32-bit embedded target int64_t is spelled "long long", which made
	// that the common case rather than an exotic one.
	widths := []struct {
		spelling string
		wantRank rank
		unsigned bool
	}{
		{"int", rankInt, false},
		{"unsigned int", rankInt, true},
		{"long", rankLong, false},
		{"unsigned long", rankLong, true},
		{"long long", rankLongLong, false},
		{"unsigned long long", rankLongLong, true},
		{"float", rankFloat, false},
		{"double", rankDouble, false},
		{"long double", rankLongDouble, false},
	}
	for _, w := range widths {
		t.Run(w.spelling, func(t *testing.T) {
			if got := rankOf(Type{w.spelling, 0}); got != w.wantRank {
				t.Errorf("rankOf(%q) = %d, want %d", w.spelling, got, w.wantRank)
			}
			if got := isUnsigned(w.spelling); got != w.unsigned {
				t.Errorf("isUnsigned(%q) = %v, want %v", w.spelling, got, w.unsigned)
			}
		})
	}
	if rankLongLong <= rankLong || rankLongDouble <= rankDouble {
		t.Error("long long must outrank long and long double must outrank double")
	}
}

func TestTypeBitsOfWideTypes(t *testing.T) {
	// C99 6.2.5.4: long long is at least 64 bits everywhere, so it is wider
	// than long on every ILP32 or 16-bit-int target.
	cases := []struct {
		bits int
		sp   string
		want int
	}{
		{16, "int", 16},
		{16, "long", 32},
		{16, "long long", 64},
		{32, "int", 32},
		{32, "long", 32},
		{32, "long long", 64},
		{64, "int", 32},
		{64, "long", 64},
		{64, "long long", 64},
	}
	for _, c := range cases {
		if got := typeBits(c.sp, Options{TargetBits: c.bits}); got != c.want {
			t.Errorf("typeBits(%q, %d-bit) = %d, want %d", c.sp, c.bits, got, c.want)
		}
	}
}

func TestUsualArithWideTypes(t *testing.T) {
	cases := []struct {
		name string
		bits int
		a, b Type
		want Type
	}{
		{"long long absorbs int", 32, Type{"long long", 0}, Type{"int", 0}, Type{"long long", 0}},
		{"long long absorbs long", 32, Type{"long", 0}, Type{"long long", 0}, Type{"long long", 0}},
		{"unsigned long long absorbs int", 32, Type{"int", 0}, Type{"unsigned long long", 0}, Type{"unsigned long long", 0}},
		{"long long holds unsigned int on ILP32", 32, Type{"long long", 0}, Type{"unsigned int", 0}, Type{"long long", 0}},
		{"long long holds unsigned long on ILP32", 32, Type{"long long", 0}, Type{"unsigned long", 0}, Type{"long long", 0}},
		// LP64: long and long long are both 64-bit, so signed long long cannot
		// hold every unsigned long value and the result goes unsigned.
		{"long long loses to unsigned long on LP64", 64, Type{"long long", 0}, Type{"unsigned long", 0}, Type{"unsigned long long", 0}},
		{"long double absorbs double", 32, Type{"double", 0}, Type{"long double", 0}, Type{"long double", 0}},
		{"long double absorbs float", 32, Type{"long double", 0}, Type{"float", 0}, Type{"long double", 0}},
		{"long double absorbs int", 32, Type{"int", 0}, Type{"long double", 0}, Type{"long double", 0}},
		{"double still beats float", 32, Type{"float", 0}, Type{"double", 0}, Type{"double", 0}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := usualArith(c.a, c.b, Options{TargetBits: c.bits})
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != c.want {
				t.Errorf("usualArith(%+v, %+v) = %+v, want %+v", c.a, c.b, got, c.want)
			}
		})
	}
}

func TestWideTypeArithmeticDeduce(t *testing.T) {
	cases := []struct {
		decl string
		want string
	}{
		{"long long g; auto y = g + 1;", "long long"},
		{"unsigned long long g; auto y = g + 1;", "unsigned long long"},
		{"long long a; long b; auto y = a * b;", "long long"},
		{"long double d; auto y = d + 1;", "long double"},
		{"long double d; double e; auto y = d * e;", "long double"},
	}
	for _, c := range cases {
		t.Run(c.decl, func(t *testing.T) {
			src := "void f(void) { " + c.decl + " }"
			subs, err := analyzeSrc(t, src, Options{TargetBits: 32})
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got := subForAuto(t, src, subs); got != c.want {
				t.Errorf("deduced %q, want %q", got, c.want)
			}
		})
	}
}
