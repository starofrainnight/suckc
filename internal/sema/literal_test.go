package sema

import "testing"

func TestDeduceIntegerLiteral(t *testing.T) {
	o16 := Options{TargetBits: 16}
	o32 := Options{TargetBits: 32}
	o64 := Options{TargetBits: 64}
	cases := []struct {
		text string
		opts Options
		want Type
	}{
		// decimal unsuffixed: int -> long -> unsigned long (never bare uint)
		{"12", o32, Type{"int", 0}},
		{"12", o64, Type{"int", 0}},
		{"40000", o16, Type{"long", 0}}, // 16-bit int overflows immediately
		{"40000", o32, Type{"int", 0}},
		{"3000000000", o32, Type{"unsigned long", 0}}, // exceeds int and long (ILP32), fits unsigned long
		// non-decimal unsuffixed: int -> unsigned int -> unsigned long
		{"0x7fffffff", o32, Type{"int", 0}},
		{"0x80000000", o32, Type{"unsigned int", 0}},
		{"0xffffffff", o32, Type{"unsigned int", 0}},
		{"0x100000000", o64, Type{"unsigned long", 0}},
		{"0b1010", o32, Type{"int", 0}},
		{"017", o32, Type{"int", 0}},
		// u suffix
		{"40000u", o32, Type{"unsigned int", 0}},
		{"4000000000u", o32, Type{"unsigned int", 0}}, // fits 32-bit unsigned int -> uint (spec §6.1)
		// l suffix
		{"12l", o32, Type{"long", 0}},
		{"12ul", o32, Type{"unsigned long", 0}},
		{"12ll", o32, Type{"long", 0}}, // ll collapses to long (no long long in model)
		// 64-bit long target
		{"9223372036854775807", o64, Type{"long", 0}},
		{"9223372036854775808", o64, Type{"unsigned long", 0}},
	}
	for _, c := range cases {
		got, err := deduceInteger(c.text, c.opts)
		if err != nil {
			t.Errorf("deduceInteger(%q): %v", c.text, err)
			continue
		}
		if got != c.want {
			t.Errorf("deduceInteger(%q, bits=%d) = %+v, want %+v", c.text, c.opts.TargetBits, got, c.want)
		}
	}
}

func TestDeduceIntegerOutOfRange(t *testing.T) {
	o32 := Options{TargetBits: 32}
	if _, err := deduceInteger("0xfffffffffffffffff", o32); err == nil {
		t.Error("expected out-of-range error")
	}
	// exceeds uint AND 32-bit unsigned long: no candidate type on target
	if _, err := deduceInteger("0x100000000", o32); err == nil {
		t.Error("expected out-of-range error for 2^32 on 32-bit target")
	}
}

func TestDeduceOtherLiterals(t *testing.T) {
	o := Options{TargetBits: 64}
	cases := []struct {
		text string
		want Type
		err  bool
	}{
		{"1.5", Type{"float", 0}, false},
		{"1e3", Type{"float", 0}, false},
		{"1.5f", Type{"float", 0}, false},
		{"1.5F", Type{"float", 0}, false},
		{"1.5L", Type{}, true}, // long double deferred -> error (spec §6.1)
		{`"abc"`, Type{"const char", 1}, false},
		{"'a'", Type{"char", 0}, false},
		{"L'a'", Type{}, true}, // deferred wide char -> error
		{"u'a'", Type{}, true},
		{"U'a'", Type{}, true},
		{"true", Type{"bool", 0}, false},
		{"false", Type{"bool", 0}, false},
		{"nullptr", Type{"void", 1}, false},
	}
	for _, c := range cases {
		got, err := deduceLiteralText(c.text, o)
		if c.err {
			if err == nil {
				t.Errorf("deduceLiteralText(%q): expected error, got %+v", c.text, got)
			}
			continue
		}
		if err != nil {
			t.Errorf("deduceLiteralText(%q): %v", c.text, err)
			continue
		}
		if got != c.want {
			t.Errorf("deduceLiteralText(%q) = %+v, want %+v", c.text, got, c.want)
		}
	}
}
