package frontend

import (
	"strings"

	"github.com/antlr4-go/antlr/v4"

	"github.com/starofrainnight/suckc/internal/parser"
)

// CParam is one parameter of a C function signature. TypeSpelling is the
// source spelling of the declared type -- "unsigned long", "void *",
// "VendorReg" -- enough to compare against later, not a type system on its
// own. Name is empty for an unnamed parameter.
type CParam struct {
	TypeSpelling string
	Name         string
}

// CFuncSig is a function declaration or definition inside a C-scope region.
//
// Unprototyped marks "int f();". In C that form says nothing at all about
// the parameters, so an arity check must treat it as always compatible; it
// must not be read as a zero-parameter function, which would reject every
// call that passes arguments to a K&R-era header.
type CFuncSig struct {
	Name         string
	Params       []CParam
	Variadic     bool
	Unprototyped bool
	// TokIdx indexes the token that binds the name, so a later diagnostic
	// can be reported against the header the declaration came from.
	TokIdx int
}

// CTypedef is a typedef inside a C-scope region. Base is the source
// spelling of the type specifiers with the `typedef` keyword removed
// ("unsigned long"), Stars counts the pointer indirections after them, and
// FuncPtr says the typedef names a function pointer, in which case Params
// and Variadic describe the pointed-to function.
//
// Stars counts the pointer levels written around the name, so it includes the
// one belonging to the function pointer itself: `typedef int (*cb)(int)`
// gives {Base: "int", Stars: 1, FuncPtr: true}, the same shape as
// `typedef char *cstr`. Do not add Stars to Base again when FuncPtr is set.
type CTypedef struct {
	Name     string
	Base     string
	Stars    int
	FuncPtr  bool
	Params   []CParam
	Variadic bool
	TokIdx   int
}

// collectCDecls records the function signatures and typedefs declared inside
// C-scope regions.
//
// It stops at every external declaration instead of descending through it. A
// region may legitimately wrap a function definition -- vendor headers ship
// static inline helpers -- and that definition's body holds local
// declarations which are not program-level signatures. Not descending also
// skips the K&R parameter declarations of an old-style definition, which sit
// beside the body rather than inside it.
func (r *ParseResult) collectCDecls() {
	if r.Tree == nil {
		return
	}
	var walk func(antlr.Tree)
	walk = func(n antlr.Tree) {
		if decl, ok := n.(*parser.CCExternalDeclarationContext); ok {
			r.externalDecl(decl)
			return
		}
		if pr, ok := n.(antlr.ParserRuleContext); ok {
			for _, ch := range pr.GetChildren() {
				walk(ch)
			}
		}
	}
	walk(r.Tree)
}

func (r *ParseResult) externalDecl(decl *parser.CCExternalDeclarationContext) {
	specs := decl.CCDeclSpecifiers()
	if specs == nil {
		return
	}
	isTypedef := hasTypedefSpecifier(specs)
	if d := decl.CCDeclarator(); d != nil {
		r.declaratorDecl(d, specs, isTypedef)
	}
	if list := decl.CCInitDeclaratorList(); list != nil {
		for _, init := range list.AllCCInitDeclarator() {
			switch {
			case init.CCDeclarator() != nil:
				r.declaratorDecl(init.CCDeclarator(), specs, isTypedef)
			case init.CCPlainDeclarator() != nil:
				r.declaratorDecl(init.CCPlainDeclarator(), specs, isTypedef)
			}
		}
	}
}

func (r *ParseResult) declaratorDecl(d antlr.Tree, specs parser.ICCDeclSpecifiersContext, isTypedef bool) {
	name, fn, array := declaratorTarget(d)
	if name == "" {
		return
	}
	if isTypedef {
		td := CTypedef{
			Name:   name,
			Base:   typedefBase(specs),
			Stars:  pointerCount(d),
			TokIdx: nameTokenIndex(d),
		}
		if fn != nil {
			td.FuncPtr = true
			td.Params, td.Variadic, _ = paramList(fn)
		}
		r.CTypedefs = append(r.CTypedefs, td)
		return
	}
	// No function suffix means a plain variable; an array of function
	// pointers is not itself callable, so neither becomes a signature.
	if fn == nil || array {
		return
	}
	params, variadic, unprototyped := paramList(fn)
	r.CFuncs = append(r.CFuncs, CFuncSig{
		Name:         name,
		Params:       params,
		Variadic:     variadic,
		Unprototyped: unprototyped,
		TokIdx:       nameTokenIndex(d),
	})
}

// declaratorTarget reports what a C declarator binds: the name it declares,
// the function suffix that applies to that name, and whether an array
// derivation sits at or above the name level (which makes the name an array
// rather than something callable).
//
// Both declarator hierarchies are handled by one walk because they are
// parallel -- cCDirectDeclarator and cCPlainDirectDeclarator differ only in
// whether a function suffix is allowed, and attributes and pointer operators
// sit in the same place. A parenthesized group binds its own name, so the
// suffix that applies to that name is the group's own; a suffix further out
// describes the type derived from the group. Both "int (*fp)(int)" and
// "int (*get(void))(int)" therefore resolve the way C reads them.
func declaratorTarget(n antlr.Tree) (name string, fn parser.ICCFunctionSuffixContext, array bool) {
	switch d := n.(type) {
	case *parser.CCDeclaratorContext:
		if inner := d.CCDirectDeclarator(); inner != nil {
			return declaratorTarget(inner)
		}
	case *parser.CCPlainDeclaratorContext:
		if inner := d.CCPlainDirectDeclarator(); inner != nil {
			return declaratorTarget(inner)
		}
	case *parser.CCDirectDeclaratorContext:
		thisFn, thisArray := directSuffixes(len(d.AllCCDirectDeclaratorSuffix()), func(i int) antlr.Tree {
			return d.AllCCDirectDeclaratorSuffix()[i]
		})
		if id := d.CCIdentifier(); id != nil {
			return id.GetText(), thisFn, thisArray
		}
		group := d.CCDeclarator()
		if group == nil {
			return "", thisFn, thisArray
		}
		name, innerFn, innerArray := declaratorTarget(group)
		if innerFn == nil {
			innerFn = thisFn
		}
		return name, innerFn, thisArray || innerArray
	case *parser.CCPlainDirectDeclaratorContext:
		thisFn, thisArray := directSuffixes(len(d.AllCCPlainDirectDeclaratorSuffix()), func(i int) antlr.Tree {
			return d.AllCCPlainDirectDeclaratorSuffix()[i]
		})
		if id := d.CCIdentifier(); id != nil {
			return id.GetText(), thisFn, thisArray
		}
		group := d.CCPlainDeclarator()
		if group == nil {
			return "", thisFn, thisArray
		}
		name, innerFn, innerArray := declaratorTarget(group)
		if innerFn == nil {
			innerFn = thisFn
		}
		return name, innerFn, thisArray || innerArray
	}
	return
}

// directSuffixes classifies the suffixes bound to one declarator name. Each
// entry is a wrapper rule (cCDirectDeclaratorSuffix and friends) whose own
// alternative is the interesting one, so every branch looks one level in.
func directSuffixes(n int, at func(int) antlr.Tree) (fn parser.ICCFunctionSuffixContext, array bool) {
	pick := func(inner parser.ICCFunctionSuffixContext, isArray bool) {
		if inner != nil && fn == nil {
			fn = inner
		}
		if isArray {
			array = true
		}
	}
	for i := 0; i < n; i++ {
		switch s := at(i).(type) {
		case *parser.CCDirectDeclaratorSuffixContext:
			pick(s.CCFunctionSuffix(), s.CCArraySuffix() != nil)
		case *parser.CCAbstractDirectDeclaratorSuffixContext:
			pick(s.CCFunctionSuffix(), s.CCArraySuffix() != nil)
		case *parser.CCPlainDirectDeclaratorSuffixContext:
			pick(nil, s.CCArraySuffix() != nil)
		}
	}
	return
}

// pointerCount counts every pointer operator in a declarator, including those
// inside a parenthesized group or an abstract declarator.
func pointerCount(n antlr.Tree) int {
	switch d := n.(type) {
	case *parser.CCDeclaratorContext:
		total := len(d.AllCCPointerOperator())
		if inner := d.CCDirectDeclarator(); inner != nil {
			total += pointerCount(inner)
		}
		return total
	case *parser.CCPlainDeclaratorContext:
		total := len(d.AllCCPointerOperator())
		if inner := d.CCPlainDirectDeclarator(); inner != nil {
			total += pointerCount(inner)
		}
		return total
	case *parser.CCDirectDeclaratorContext:
		if inner := d.CCDeclarator(); inner != nil {
			return pointerCount(inner)
		}
	case *parser.CCPlainDirectDeclaratorContext:
		if inner := d.CCPlainDeclarator(); inner != nil {
			return pointerCount(inner)
		}
	case *parser.CCAbstractDeclaratorContext:
		total := len(d.AllCCPointerOperator())
		if inner := d.CCAbstractDirectDeclarator(); inner != nil {
			total += pointerCount(inner)
		}
		return total
	case *parser.CCAbstractDirectDeclaratorContext:
		if inner := d.CCAbstractDeclarator(); inner != nil {
			return pointerCount(inner)
		}
	}
	return 0
}

// paramList reads a function suffix. unprototyped is set for the empty
// parameter list "()", which C defines to mean "unspecified" rather than
// "none".
func paramList(fn parser.ICCFunctionSuffixContext) (params []CParam, variadic, unprototyped bool) {
	list := fn.CCParameterTypeList()
	if list == nil {
		return nil, false, true
	}
	variadic = list.Ellipsis() != nil
	for _, pd := range list.AllCCParameterDeclaration() {
		params = append(params, paramOf(pd))
	}
	// "void f(void)" is C's spelling of an empty parameter list, not a
	// single parameter of type void. Only the lone, unnamed, pointerless
	// void means that.
	if !variadic && len(params) == 1 && params[0].TypeSpelling == "void" && params[0].Name == "" {
		return nil, false, false
	}
	return params, variadic, false
}

// paramOf reads one parameter declaration. The abstract-declarator branch is
// what keeps the star on "void *".
func paramOf(pd parser.ICCParameterDeclarationContext) CParam {
	p := CParam{}
	if specs := pd.CCDeclSpecifiers(); specs != nil {
		p.TypeSpelling = spelling(specs)
	}
	switch {
	case pd.CCDeclarator() != nil:
		p.TypeSpelling += stars(pointerCount(pd.CCDeclarator()))
		p.Name, _, _ = declaratorTarget(pd.CCDeclarator())
	case pd.CCAbstractDeclarator() != nil:
		p.TypeSpelling += stars(pointerCount(pd.CCAbstractDeclarator()))
	}
	return p
}

// specifierParts walks the specifier list's children, skipping the trailing
// typedef name: that name is the declarator, not part of the type.
func specifierParts(specs parser.ICCDeclSpecifiersContext) []antlr.Tree {
	var out []antlr.Tree
	for _, ch := range specs.GetChildren() {
		if _, isName := ch.(*parser.CCTypedefNameContext); isName {
			continue
		}
		out = append(out, ch)
	}
	return out
}

func hasTypedefSpecifier(specs parser.ICCDeclSpecifiersContext) bool {
	for _, part := range specifierParts(specs) {
		if spelling(part) == "typedef" {
			return true
		}
	}
	return false
}

func typedefBase(specs parser.ICCDeclSpecifiersContext) string {
	var parts []string
	for _, part := range specifierParts(specs) {
		if s := spelling(part); s != "" && s != "typedef" {
			parts = append(parts, s)
		}
	}
	return strings.Join(parts, " ")
}

func nameTokenIndex(d antlr.Tree) int {
	found := -1
	var walk func(antlr.Tree)
	walk = func(n antlr.Tree) {
		if id, ok := n.(*parser.CCIdentifierContext); ok {
			if found < 0 {
				found = id.GetStart().GetTokenIndex()
			}
			return
		}
		if pr, ok := n.(antlr.ParserRuleContext); ok {
			for _, ch := range pr.GetChildren() {
				walk(ch)
			}
		}
	}
	walk(d)
	if found >= 0 {
		return found
	}
	if pr, ok := d.(antlr.ParserRuleContext); ok && pr.GetStart() != nil {
		return pr.GetStart().GetTokenIndex()
	}
	return 0
}

// spelling returns the source text of a context with single spaces between
// tokens. ANTLR's own GetText concatenates token text with no separator,
// which would turn "unsigned long" into "unsignedlong" and "void *" into
// "void*"; these strings are recorded as source spellings, so the spaces
// matter.
func spelling(ctx antlr.Tree) string {
	var parts []string
	var walk func(antlr.Tree)
	walk = func(n antlr.Tree) {
		if t, ok := n.(antlr.TerminalNode); ok {
			parts = append(parts, t.GetText())
			return
		}
		if pr, ok := n.(antlr.ParserRuleContext); ok {
			for _, ch := range pr.GetChildren() {
				walk(ch)
			}
		}
	}
	walk(ctx)
	return strings.Join(parts, " ")
}

func stars(n int) string {
	if n == 0 {
		return ""
	}
	return " " + strings.Repeat("*", n)
}
