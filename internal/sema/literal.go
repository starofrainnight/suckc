package sema

import (
	"errors"
	"strconv"
	"strings"
)

var (
	errWideChar    = errors.New("wide character literal not supported")
	errOutOfRange  = errors.New("integer literal out of range")
	errUnsupported = errors.New("unsupported expression")
)

// deduceInteger implements the spec section 6.1 integer table.
func deduceInteger(text string, o Options) (Type, error) {
	raw, suffix := splitIntSuffix(text)
	base := 10
	isDecimal := true
	switch {
	case strings.HasPrefix(raw, "0x"), strings.HasPrefix(raw, "0X"):
		base, isDecimal = 16, false
		raw = raw[2:]
	case strings.HasPrefix(raw, "0b"), strings.HasPrefix(raw, "0B"):
		base, isDecimal = 2, false
		raw = raw[2:]
	case len(raw) > 1 && raw[0] == '0':
		base, isDecimal = 8, false
		raw = raw[1:]
	}
	v, err := strconv.ParseUint(raw, base, 64)
	if err != nil {
		return Type{}, errOutOfRange
	}
	suffix = strings.ToLower(suffix)
	hasU := strings.Contains(suffix, "u")
	hasL := strings.Contains(suffix, "l")

	_, intMax, intUMax := o.intLimits()
	_, longMax, _ := o.longLimits()
	ulongMax := o.ulongMaxU64()

	switch {
	case hasU && hasL: // ul, lu, ull, llu
		if v <= ulongMax {
			return Type{"unsigned long", 0}, nil
		}
		return Type{}, errOutOfRange
	case hasU: // u only
		if v <= uint64(intUMax) {
			return Type{"unsigned int", 0}, nil
		}
		if v <= ulongMax {
			return Type{"unsigned long", 0}, nil
		}
		return Type{}, errOutOfRange
	case hasL: // l only
		if v <= uint64(longMax) {
			return Type{"long", 0}, nil
		}
		if v <= ulongMax {
			return Type{"unsigned long", 0}, nil
		}
		return Type{}, errOutOfRange
	case isDecimal:
		// Decimal unsuffixed never yields bare unsigned int (C89 rule).
		if v <= uint64(intMax) {
			return Type{"int", 0}, nil
		}
		if v <= uint64(longMax) {
			return Type{"long", 0}, nil
		}
		if v <= ulongMax {
			return Type{"unsigned long", 0}, nil
		}
		return Type{}, errOutOfRange
	default:
		// octal/hex/binary unsuffixed: int -> unsigned int -> unsigned long
		if v <= uint64(intMax) {
			return Type{"int", 0}, nil
		}
		if v <= uint64(intUMax) {
			return Type{"unsigned int", 0}, nil
		}
		if v <= ulongMax {
			return Type{"unsigned long", 0}, nil
		}
		return Type{}, errOutOfRange
	}
}

// splitIntSuffix peels C integer suffixes (u/U/l/L/ll/LL combinations)
// from the end of the token text. C allows both "lu" and "ul" orderings,
// so l and u are trimmed alternately until neither makes progress.
func splitIntSuffix(text string) (raw, suffix string) {
	i := len(text)
	trim := func(isSuffix func(rune) bool) {
		for i > 0 {
			r := rune(text[i-1])
			if !isSuffix(r) {
				return
			}
			i--
		}
	}
	isL := func(r rune) bool { return r == 'l' || r == 'L' }
	isU := func(r rune) bool { return r == 'u' || r == 'U' }
	for {
		before := i
		trim(isL)
		trim(isU)
		if i == before {
			break
		}
	}
	if i == 0 {
		return text, ""
	}
	return text[:i], text[i:]
}

// deduceLiteralText deduces a type from raw literal token text.
func deduceLiteralText(text string, o Options) (Type, error) {
	switch {
	case strings.HasPrefix(text, `"`):
		return Type{"const char", 1}, nil
	case strings.Contains(text, `"`):
		// prefixed (L"", u8"", ...) or raw string: deferred, not in spec 6.1
		return Type{}, errUnsupported
	case strings.HasPrefix(text, "'") && strings.HasSuffix(text, "'") && len(text) >= 3:
		return Type{"char", 0}, nil
	case strings.HasPrefix(text, "true"), strings.HasPrefix(text, "false"):
		return Type{"bool", 0}, nil
	case text == "nullptr":
		return Type{"void", 1}, nil
	case len(text) > 0 && strings.ContainsAny(text[:1], "uUL") && strings.Contains(text, "'"):
		// wide/prefixed character literal: deferred (spec 6.1)
		return Type{}, errWideChar
	}
	if isFloatLiteral(text) {
		if strings.HasSuffix(text, "l") || strings.HasSuffix(text, "L") {
			// long double deferred (spec §6.1)
			return Type{}, errUnsupported
		}
		// unsuffixed and f/F both deduce float (spec §6.1, user decision)
		return Type{"float", 0}, nil
	}
	return deduceInteger(text, o)
}

// isFloatLiteral reports whether text looks like a C floating constant.
func isFloatLiteral(s string) bool {
	if s == "" {
		return false
	}
	body := s
	if last := body[len(body)-1]; last == 'f' || last == 'F' || last == 'l' || last == 'L' {
		body = body[:len(body)-1]
	}
	if body == "" {
		return false
	}
	if strings.Contains(body, ".") {
		return true
	}
	if idx := strings.IndexAny(body, "eE"); idx > 0 && idx < len(body)-1 {
		return true
	}
	return false
}
