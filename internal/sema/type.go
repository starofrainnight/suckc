package sema

import (
	"errors"
	"math"
	"strings"
)

// Type is a source-spelling type with a pointer depth (spec section 5).
// Spelling is kept verbatim ("int", "unsigned long", "kknd"); no typedef
// resolution is performed. Qualifiers stay inside Spelling, and only
// value reads strip them (see stripCV).
type Type struct {
	Spelling string
	Stars    int
}

// Text renders the type for emission, e.g. "int" or "int **".
func (t Type) Text() string {
	if t.Stars == 0 {
		return t.Spelling
	}
	return t.Spelling + " " + repeat("*", t.Stars)
}

func repeat(s string, n int) string {
	out := ""
	for i := 0; i < n; i++ {
		out += s
	}
	return out
}

// stripCV removes const and volatile from a type spelling. A value read
// yields the unqualified type (C89 6.3.2.1), while taking an address must
// keep the qualifier, so only value reads call this.
func stripCV(spelling string) string {
	fields := strings.Fields(spelling)
	kept := fields[:0]
	for _, f := range fields {
		if f == "const" || f == "volatile" {
			continue
		}
		kept = append(kept, f)
	}
	return strings.Join(kept, " ")
}

// errIncompatible is wrapped by callers into the stable diagnostic
// "incompatible operand types 'A' and 'B'".
var errIncompatible = errors.New("incompatible operand types")

// intLimits returns (min, max, unsignedMax) for the target int width.
// TargetBits 16 => 16-bit int; 32/64 => 32-bit int (spec section 6.1).
func (o Options) intLimits() (int64, int64, int64) {
	if o.TargetBits == 16 {
		return math.MinInt16, math.MaxInt16, math.MaxUint16
	}
	return math.MinInt32, math.MaxInt32, math.MaxUint32
}

// longLimits returns (min, max, unsignedMax) for the target long width.
// 16 => 32-bit long, 32 => 32-bit long, 64 => 64-bit long (LP64).
// The unsigned long max for 64-bit targets overflows int64; it is returned
// as -1 sentinel and ulongFits must use ulongMaxU64 instead.
func (o Options) longLimits() (int64, int64, int64) {
	if o.TargetBits == 64 {
		return math.MinInt64, math.MaxInt64, -1
	}
	return math.MinInt32, math.MaxInt32, math.MaxUint32
}

// ulongMaxU64 returns the target unsigned long max as uint64.
func (o Options) ulongMaxU64() uint64 {
	if o.TargetBits == 64 {
		return math.MaxUint64
	}
	return math.MaxUint32
}

// promote applies the C89 integer promotions used by this deduction subset:
// types narrower than int become int; everything else is unchanged
// (stars are preserved).
func promote(t Type) Type {
	if t.Stars != 0 {
		return t
	}
	switch t.Spelling {
	case "char", "signed char", "unsigned char", "short", "unsigned short", "bool":
		return Type{"int", 0}
	}
	return t
}

type rank int

const (
	rankNone rank = iota
	rankInt
	rankLong
	rankFloat
	rankDouble
)

func rankOf(t Type) rank {
	switch t.Spelling {
	case "int", "unsigned int":
		return rankInt
	case "long", "unsigned long":
		return rankLong
	case "float":
		return rankFloat
	case "double":
		return rankDouble
	}
	return rankNone
}

func isUnsigned(sp string) bool {
	return sp == "unsigned int" || sp == "unsigned long"
}

// usualArith implements the C89 6.3.1.8 subset over the deduction table
// (spec section 6.2): integer promotions first, then rank rules.
func usualArith(a, b Type, o Options) (Type, error) {
	if a.Stars != 0 || b.Stars != 0 {
		return Type{}, errIncompatible
	}
	a, b = promote(a), promote(b)
	if a == b {
		return a, nil
	}
	ra, rb := rankOf(a), rankOf(b)
	if ra == rankNone || rb == rankNone {
		return Type{}, errIncompatible
	}
	// Float dominance: double > float > integer ranks.
	if ra >= rankFloat || rb >= rankFloat {
		if ra == rankDouble || rb == rankDouble {
			return Type{"double", 0}, nil
		}
		if ra == rankFloat || rb == rankFloat {
			return Type{"float", 0}, nil
		}
	}
	// Both integers at this point.
	if isUnsigned(a.Spelling) == isUnsigned(b.Spelling) {
		// Same signedness: wider rank wins.
		if ra >= rb {
			return a, nil
		}
		return b, nil
	}
	unsigned, signed := a, b
	if isUnsigned(b.Spelling) {
		unsigned, signed = b, a
	}
	ru, rs := rankOf(unsigned), rankOf(signed)
	if ru >= rs {
		// Unsigned rank >= signed rank -> unsigned type of that rank.
		if ru == rankLong {
			return Type{"unsigned long", 0}, nil
		}
		return Type{"unsigned int", 0}, nil
	}
	// Signed rank > unsigned rank: signed wins only if it can hold every
	// value of the unsigned operand (C89 6.3.1.8 case 4).
	if signedCanHoldUnsigned(signed.Spelling, unsigned.Spelling, o) {
		return signed, nil
	}
	if rs == rankLong {
		return Type{"unsigned long", 0}, nil
	}
	return Type{"unsigned int", 0}, nil
}

// signedCanHoldUnsigned reports whether the signed target type can represent
// all values of the unsigned type.
func signedCanHoldUnsigned(signed, unsigned string, o Options) bool {
	sBits := typeBits(signed, o)
	uBits := typeBits(unsigned, o)
	// Signed needs one sign bit: representable iff sBits > uBits.
	return sBits > uBits
}

func typeBits(sp string, o Options) int {
	switch sp {
	case "int", "unsigned int":
		if o.TargetBits == 16 {
			return 16
		}
		return 32
	case "long", "unsigned long":
		if o.TargetBits == 64 {
			return 64
		}
		return 32
	}
	return 0
}

// commonType returns the ternary ?: common type (spec section 6.2):
// usual arithmetic conversions if both arithmetic, identical type
// otherwise, error if incompatible.
func commonType(a, b Type, o Options) (Type, error) {
	if a == b {
		return a, nil
	}
	if a.Stars == 0 && b.Stars == 0 {
		if rankOf(promote(a)) != rankNone && rankOf(promote(b)) != rankNone {
			return usualArith(a, b, o)
		}
	}
	return Type{}, errIncompatible
}
