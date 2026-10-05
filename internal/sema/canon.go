package sema

import (
	"fmt"
	"strings"

	"github.com/starofrainnight/suckc/internal/frontend"
)

// canonicalSpelling normalizes a C89 type-specifier list to one spelling per
// type, so that source order and redundant keywords stop creating types that
// are distinct here but identical in C: "long unsigned int", "unsigned long
// int" and "unsigned long" all become "unsigned long".
//
// Qualifiers are preserved and moved to the front, because Spelling is also the
// text emitted for a deduced "auto"; dropping const there would silently
// un-const the generated C. Value reads strip them separately, via stripCV.
//
// Anything that is not a plain builtin specifier list is returned unchanged: an
// unrecognised word is a name this compiler has no rule for, and inventing a
// type for it is worse than comparing it as written.
func canonicalSpelling(spelling string) string {
	words := strings.Fields(spelling)
	if len(words) == 0 {
		return spelling
	}
	var quals, rest []string
	for _, w := range words {
		if isQualifierWord(w) {
			quals = append(quals, w)
		} else {
			rest = append(rest, w)
		}
	}
	base, ok := canonicalBuiltin(rest)
	if !ok {
		return spelling
	}
	if len(quals) == 0 {
		return base
	}
	return strings.Join(quals, " ") + " " + base
}

func isQualifierWord(w string) bool {
	switch w {
	case "const", "volatile", "restrict", "__restrict", "__restrict__", "_Atomic":
		return true
	}
	return false
}

// canonicalBuiltin maps a builtin type-specifier list onto the single spelling
// C89 6.7.2 gives it.
func canonicalBuiltin(words []string) (string, bool) {
	var (
		signed, unsigned    bool
		void, boolT         bool
		charT, shortT, intT bool
		floatT, doubleT     bool
		longCount           int
		sawBuiltin          bool
	)
	for _, w := range words {
		switch w {
		case "signed":
			signed, sawBuiltin = true, true
		case "unsigned":
			unsigned, sawBuiltin = true, true
		case "void":
			void, sawBuiltin = true, true
		case "_Bool", "bool":
			boolT, sawBuiltin = true, true
		case "char":
			charT, sawBuiltin = true, true
		case "short":
			shortT, sawBuiltin = true, true
		case "int":
			intT, sawBuiltin = true, true
		case "long":
			sawBuiltin = true
			longCount++
		case "float":
			floatT, sawBuiltin = true, true
		case "double":
			doubleT, sawBuiltin = true, true
		default:
			return "", false
		}
	}
	if !sawBuiltin {
		return "", false
	}
	// Illegal combinations are rejected rather than guessed: an invented
	// canonical form would merge types C keeps apart, or invent one C lacks.
	switch {
	case void:
		if signed || unsigned || charT || shortT || intT || floatT || doubleT || boolT || longCount > 0 {
			return "", false
		}
		return "void", true
	case boolT:
		if signed || unsigned || charT || shortT || intT || floatT || doubleT || void || longCount > 0 {
			return "", false
		}
		return "bool", true
	case floatT:
		if doubleT || signed || unsigned || charT || shortT || intT || longCount > 0 {
			return "", false
		}
		return "float", true
	case doubleT:
		if signed || unsigned || charT || shortT || intT || longCount > 1 {
			return "", false
		}
		if longCount == 1 {
			return "long double", true
		}
		return "double", true
	}
	if floatT {
		return "", false
	}
	if charT {
		// char, signed char and unsigned char are three distinct types.
		if shortT || longCount > 0 || intT {
			return "", false
		}
		switch {
		case signed && unsigned:
			return "", false
		case signed:
			return "signed char", true
		case unsigned:
			return "unsigned char", true
		default:
			return "char", true
		}
	}
	if signed && unsigned {
		return "", false
	}
	prefix := ""
	if unsigned {
		prefix = "unsigned "
	}
	switch {
	case shortT:
		if longCount > 0 {
			return "", false
		}
		return prefix + "short", true
	case longCount == 1:
		return prefix + "long", true
	case longCount == 2:
		return prefix + "long long", true
	case longCount != 0:
		return "", false
	default:
		// signed and unsigned on their own mean int.
		return prefix + "int", true
	}
}

// canonType is a type reduced to what compatibility compares. Qualifiers are
// absent on purpose: C89 6.5.16.1 lets a value be passed to a more-qualified
// target, so no qualifier difference can make a valid call invalid, and
// keeping them would only produce false rejections.
type canonType struct {
	Base  string
	Stars int
}

func (c canonType) key() string {
	if c.Stars == 0 {
		return c.Base
	}
	return c.Base + strings.Repeat("*", c.Stars)
}

// display spells the type the way a C programmer writes it, for diagnostics.
// key() is the unspaced form used for comparison; reusing it in a message would
// print "char*" where the surrounding text says "passing 'x' where 'y'".
func (c canonType) display() string {
	if c.Stars == 0 {
		return c.Base
	}
	return c.Base + " " + strings.Repeat("*", c.Stars)
}

// cTypeResolver reduces a C-region type spelling to a canonType, expanding
// typedef names on the way. It feeds argument/parameter comparison only;
// nothing it produces reaches the emitted C.
type cTypeResolver struct {
	tds      map[string]frontend.CTypedef
	anonSeq  int
	visiting map[string]bool
	cache    map[string]canonType
}

func newCTypeResolver(tds []frontend.CTypedef) *cTypeResolver {
	m := make(map[string]frontend.CTypedef, len(tds))
	for _, td := range tds {
		if _, dup := m[td.Name]; !dup {
			m[td.Name] = td
		}
	}
	return &cTypeResolver{tds: m, visiting: map[string]bool{}, cache: map[string]canonType{}}
}

func (r *cTypeResolver) resolve(spelling string) canonType {
	if c, ok := r.cache[spelling]; ok {
		return c
	}
	c := r.resolveUncached(spelling)
	r.cache[spelling] = c
	return c
}

func (r *cTypeResolver) resolveUncached(spelling string) canonType {
	base, stars := splitTrailingStars(spelling)

	// canonType carries no qualifiers, so they are dropped here. Leaving them
	// in would make "const char *" and "char *" compare unequal even though no
	// qualifier difference can make a valid call invalid.
	var keep []string
	for _, w := range strings.Fields(base) {
		if !isQualifierWord(w) {
			keep = append(keep, w)
		}
	}
	words := keep

	if anon, ok := anonymousTag(words); ok {
		r.anonSeq++
		return canonType{Base: fmt.Sprintf("%s <anon %d>", anon, r.anonSeq), Stars: stars}
	}

	if len(words) == 1 {
		if inner, ok := r.resolveTypedef(words[0]); ok {
			inner.Stars += stars
			return inner
		}
		// An unknown single name may be a typedef from a header this region did
		// not include. Compare it as written.
		return canonType{Base: words[0], Stars: stars}
	}

	tag := strings.Join(words, " ")
	if isTagKeyword(words) {
		// A tag names one nominal type for the whole translation unit, so the
		// same tag reached through two headers is still the same type.
		return canonType{Base: tag, Stars: stars}
	}
	if c, ok := canonicalBuiltin(words); ok {
		return canonType{Base: c, Stars: stars}
	}
	return canonType{Base: tag, Stars: stars}
}

func (r *cTypeResolver) resolveTypedef(name string) (canonType, bool) {
	td, ok := r.tds[name]
	if !ok {
		return canonType{}, false
	}
	if r.visiting[name] {
		// A self-referential chain denotes no type; treat the name as opaque
		// instead of looping.
		return canonType{Base: name}, true
	}
	r.visiting[name] = true
	defer delete(r.visiting, name)

	if td.FuncPtr {
		// The signature is dropped on purpose: two distinct function-pointer
		// types then compare equal, which misses checks but never rejects a
		// valid call.
		inner := r.resolve(td.Base)
		return canonType{Base: "FUNCPTR(" + inner.Base + ")", Stars: td.Stars}, true
	}
	inner := r.resolve(td.Base)
	inner.Stars += td.Stars
	return inner, true
}

func isTagKeyword(words []string) bool {
	if len(words) == 0 {
		return false
	}
	switch words[0] {
	case "struct", "union", "enum", "__struct", "__union", "__enum":
		return true
	}
	return false
}

// anonymousTag reports an untagged aggregate specifier. C gives each such
// declaration its own nominal type, so the spelling alone cannot identify it.
//
// The check has to look past the first word: a typedef of an anonymous
// aggregate arrives spelled out in full, as "struct { int a; }", with no tag
// between the keyword and the opening brace.
func anonymousTag(words []string) (string, bool) {
	if len(words) == 0 {
		return "", false
	}
	var kind string
	switch words[0] {
	case "struct", "union", "__struct", "__union":
		kind = words[0]
	default:
		return "", false
	}
	if len(words) == 1 || words[1] == "{" {
		return kind, true
	}
	return "", false
}

// splitTrailingStars separates the stars that CParam.TypeSpelling appends as
// text from the type they belong to.
func splitTrailingStars(spelling string) (string, int) {
	s := strings.TrimSpace(spelling)
	n := 0
	for strings.HasSuffix(s, "*") {
		n++
		s = strings.TrimRight(s[:len(s)-1], " \t")
	}
	return s, n
}
