package sema

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/starofrainnight/suckc/internal/frontend"
	"github.com/starofrainnight/suckc/internal/parser"
)

// A call site is spelled in the grammar as `postfixExpression LeftParen
// expressionList? RightParen`, with no dedicated arguments rule, so a call has
// to be recognised from the child shape rather than from a context type. The
// shapes this must not match: a functional cast (`int(5)`, where the child
// before the paren is a type specifier), a subscript (`a[i]`, which ends in
// `]`), and the four `cast<type>(x)` forms.
type ccallSite struct {
	Name   string
	Args   int
	TokIdx int
}

// CallError is an arity diagnostic. It is separate from Error because that
// type's message is hardcoded to auto deduction.
type CallError struct {
	File string
	Line int
	Col  int // 1-based
	Name string
	Msg  string
}

func (e *CallError) Error() string {
	return fmt.Sprintf("%s:%d:%d: %s", e.File, e.Line, e.Col, e.Msg)
}

// checkCCalls verifies the arity of every SuckC call whose callee is declared
// by a C-scope region. It is arity only: argument types are out of scope until
// the type representation in sema.Type can name qualifiers and canonicalise
// spellings.
//
// Calls whose callee is also declared in SuckC are skipped. That keeps
// redeclaring a vendor prototype in a .suckc file working, and it is the
// conservative direction when scope resolution is not built yet: an unchecked
// call keeps compiling, a wrong error would not.
func checkCCalls(res *frontend.ParseResult, a *Analyzer) error {
	if len(res.CFuncs) == 0 {
		return nil
	}
	index, err := indexCFuncs(res, a)
	if err != nil {
		return err
	}
	if len(index) == 0 {
		return nil
	}
	scan := ccallScan{
		shadow: map[string]bool{},
		isCFunc: func(name string) bool {
			_, ok := index[name]
			return ok
		},
	}
	scanCallsAndNames(res.Tree, &scan)
	sites, shadow := scan.sites, scan.shadow
	for _, s := range sites {
		sig, ok := index[s.Name]
		if !ok || shadow[s.Name] {
			continue
		}
		// Unreachable while C region bodies use only the cC* expression rules,
		// but a call name is exactly what an InRegion guard is for, and a
		// wrong arity error pointing into a C header would be baffling.
		if res.InRegion(s.TokIdx) {
			continue
		}
		if sig.Unprototyped {
			continue
		}
		want := len(sig.Params)
		if s.Args >= want && (sig.Variadic || s.Args == want) {
			continue
		}
		where := ""
		if file, line := sourceAtToken(res, a, sig.TokIdx); file != "" {
			where = fmt.Sprintf(" (declared at %s:%d)", file, line)
		}
		verb := "too many arguments"
		if s.Args < want {
			verb = "too few arguments"
		}
		decl := fmt.Sprintf("%d", want)
		if sig.Variadic {
			decl = fmt.Sprintf("%d or more", want)
		}
		return &CallError{
			File: a.Opts.FileName,
			Line: res.Tokens[s.TokIdx].GetLine(),
			Col:  res.Tokens[s.TokIdx].GetColumn() + 1,
			Name: s.Name,
			Msg: fmt.Sprintf("%s in call to %s: %d given, C region declares %s%s",
				verb, quoteName(s.Name), s.Args, decl, where),
		}
	}
	return nil
}

// indexCFuncs keys the region signatures by name. Two headers may legitimately
// declare the same function, so duplicates are merged unless they impose
// different arities, which C would reject and which would make the arity check
// depend on include order.
func indexCFuncs(res *frontend.ParseResult, a *Analyzer) (map[string]frontend.CFuncSig, error) {
	index := make(map[string]frontend.CFuncSig, len(res.CFuncs))
	for _, sig := range res.CFuncs {
		if prev, ok := index[sig.Name]; ok {
			if !sigConflict(prev, sig) {
				// A parameter-less prototype constrains nothing, so the
				// function stays unchecked rather than keeping whichever
				// arity happened to be seen first.
				if prev.Unprototyped || sig.Unprototyped {
					prev.Unprototyped = true
					index[sig.Name] = prev
				}
				continue
			}
			file, line := sourceAtToken(res, a, sig.TokIdx)
			if file == "" {
				file = a.Opts.FileName
			}
			return nil, &CallError{
				File: file,
				Line: line,
				Col:  0,
				Name: sig.Name,
				Msg: fmt.Sprintf("conflicting declarations of %s in C regions: %d argument(s) versus %d",
					quoteName(sig.Name), len(prev.Params), len(sig.Params)),
			}
		}
		index[sig.Name] = sig
	}
	return index, nil
}

// sigConflict reports whether two prototypes of the same function disagree.
// A prototype without a parameter list constrains nothing, so it never
// conflicts.
func sigConflict(a, b frontend.CFuncSig) bool {
	if a.Unprototyped || b.Unprototyped {
		return false
	}
	if len(a.Params) != len(b.Params) {
		return true
	}
	return a.Variadic != b.Variadic
}

func quoteName(s string) string { return "'" + s + "'" }

// sourceAtToken resolves a token index to the file and line a diagnostic
// should quote, preferring the originating header recorded by the preprocessor
// line markers over the .suckc path.
func sourceAtToken(res *frontend.ParseResult, a *Analyzer, tokIdx int) (string, int) {
	file, line := res.SourceAt(tokIdx)
	if file == "" {
		file = a.Opts.FileName
		if tokIdx < len(res.Tokens) {
			line = res.Tokens[tokIdx].GetLine()
		}
	}
	return file, line
}

type ccallScan struct {
	sites   []ccallSite
	shadow  map[string]bool
	isCFunc func(string) bool
}

// scanCallsAndNames walks once because a statement-level call and a declaration
// are the same subtree here: `f(a);` puts the callee in a declaratorid, so
// collecting names separately would let every call shadow itself out of the
// check.
func scanCallsAndNames(t antlr.Tree, s *ccallScan) {
	switch c := t.(type) {
	case *parser.PostfixExpressionContext:
		if c != nil {
			if site, ok := callOf(c); ok {
				s.sites = append(s.sites, site)
			}
		}
	case *parser.SimpleDeclarationContext:
		if c != nil {
			if site, ok := declarationShapedCall(c, s.isCFunc); ok {
				s.sites = append(s.sites, site)
				// The declarator itself holds no calls, but the initializer
				// does: `outer(inner(1));` is one declaration subtree.
				if in := c.InitDeclarator().Initializer(); in != nil {
					for i := 0; i < in.GetChildCount(); i++ {
						scanCallsAndNames(in.GetChild(i), s)
					}
				}
				return
			}
		}
	}
	if d, ok := t.(*parser.DeclaratoridContext); ok && d != nil {
		if idExpr := childOfIDExpression(d); idExpr != nil {
			if unq := childOfUnqualifiedID(idExpr); unq != nil {
				if ident := childOfIdentifier(unq); ident != nil {
					s.shadow[ident.GetText()] = true
				}
			}
		}
	}
	for i := 0; i < t.GetChildCount(); i++ {
		scanCallsAndNames(t.GetChild(i), s)
	}
}

// declarationShapedCall recognises the statement-level call `f(a, b);`, which
// this grammar parses as a declaration with no type specifier, the declarator
// `f`, and `(a, b)` landing in the initializer. C++ settles this in
// [dcl.ambiguity] by asking whether `f` is a typedef, which a parser cannot
// know, and the specifier sequence here is optional, so ANTLR never even tries.
//
// Valid C is not ambiguous: a declaration always spells a type, so a
// declarator without one can only be an expression.
func declarationShapedCall(sd *parser.SimpleDeclarationContext, isCFunc func(string) bool) (ccallSite, bool) {
	init := sd.InitDeclarator()
	if init == nil || init.Declarator() == nil {
		return ccallSite{}, false
	}
	if name, tokIdx, ok := typedefNameSpelling(sd.DeclSpecifierSeq()); ok &&
		isCFunc != nil && isCFunc(name) {
		return ccallSite{Name: name, TokIdx: tokIdx, Args: 1}, true
	}
	if hasTypeSpecifier(sd.DeclSpecifierSeq()) {
		return ccallSite{}, false
	}
	name, tokIdx := callTargetName(init.Declarator())
	if name == "" {
		return ccallSite{}, false
	}
	site := ccallSite{Name: name, TokIdx: tokIdx}
	if init.Initializer() == nil {
		// `f();` keeps the parens on the declarator, so a bare declarator
		// with no initializer here is an empty declaration, not a call.
		if !hasParameterList(init.Declarator()) {
			return ccallSite{}, false
		}
		return site, true
	}
	args, ok := parenthesizedArgCount(init.Initializer())
	if !ok {
		return ccallSite{}, false
	}
	site.Args = args
	return site, true
}

// hasTypeSpecifier reports whether the sequence names an actual type. A lone
// identifier is a typedef-name rather than a type definition, which is the
// reading C++ prefers where the two are ambiguous.
func hasTypeSpecifier(specs parser.IDeclSpecifierSeqContext) bool {
	if specs == nil {
		return false
	}
	if specs.GetChildCount() != 1 {
		return true
	}
	ds, ok := specs.GetChild(0).(*parser.DeclSpecifierContext)
	if !ok || ds == nil || ds.GetChildCount() != 1 {
		return true
	}
	// The typedef-name wrapping rule sits between declSpecifier and the token.
	return !firstTerminalIsIdentifier(ds.GetChild(0))
}

func firstTerminalIsIdentifier(t antlr.Tree) bool {
	_, ok := firstIdentifierToken(t)
	return ok
}

// typedefNameSpelling returns the name and token index of a specifier sequence
// holding nothing but a single identifier.
//
// `f(s);` with one identifier argument lands here: with no typedef table in the
// grammar, ANTLR reads `f` as a type name and `s` as the declared variable.
// C settles this in [dcl.ambiguity] by asking whether `f` is a typedef, which
// only the symbol table knows, so the caller decides: a name that is not a
// C-region function really is a declaration.
func typedefNameSpelling(specs parser.IDeclSpecifierSeqContext) (string, int, bool) {
	if specs == nil || specs.GetChildCount() != 1 {
		return "", 0, false
	}
	ds, ok := specs.GetChild(0).(*parser.DeclSpecifierContext)
	if !ok || ds == nil {
		return "", 0, false
	}
	term, ok := firstIdentifierToken(ds)
	if !ok {
		return "", 0, false
	}
	return term.GetText(), term.GetSymbol().GetTokenIndex(), true
}

// firstIdentifierToken returns the leftmost identifier token in the subtree, and
// whether there is one. The typedef-name wrapping rules sit between
// declSpecifier and the token, so the search has to descend.
func firstIdentifierToken(t antlr.Tree) (antlr.TerminalNode, bool) {
	if term, ok := t.(antlr.TerminalNode); ok && term != nil {
		return term, term.GetSymbol().GetTokenType() == parser.SuckCLexerIdentifier
	}
	for i := 0; i < t.GetChildCount(); i++ {
		if term, ok := firstIdentifierToken(t.GetChild(i)); ok {
			return term, true
		}
	}
	return nil, false
}

func hasParameterList(t antlr.Tree) bool {
	if _, ok := t.(*parser.ParametersAndQualifiersContext); ok {
		return true
	}
	for i := 0; i < t.GetChildCount(); i++ {
		if hasParameterList(t.GetChild(i)) {
			return true
		}
	}
	return false
}

func callTargetName(t antlr.Tree) (string, int) {
	if id, ok := t.(*parser.DeclaratoridContext); ok && id != nil {
		return id.GetText(), id.GetStart().GetTokenIndex()
	}
	for i := 0; i < t.GetChildCount(); i++ {
		if name, tokIdx := callTargetName(t.GetChild(i)); name != "" {
			return name, tokIdx
		}
	}
	return "", 0
}

func parenthesizedArgCount(in antlr.Tree) (int, bool) {
	kids := in.GetChildren()
	if len(kids) != 3 || tokenType(kids[0]) != parser.SuckCLexerLeftParen ||
		tokenType(kids[2]) != parser.SuckCLexerRightParen {
		return 0, false
	}
	list, ok := kids[1].(*parser.ExpressionListContext)
	if !ok || list == nil || list.InitializerList() == nil {
		return 0, false
	}
	return len(list.InitializerList().AllInitializerClause()), true
}

func callOf(pe *parser.PostfixExpressionContext) (ccallSite, bool) {
	kids := pe.GetChildren()
	n := len(kids)
	if n < 3 || tokenType(kids[n-1]) != parser.SuckCParserRightParen {
		return ccallSite{}, false
	}
	site := ccallSite{}
	switch head := kids[n-2].(type) {
	case *parser.ExpressionListContext:
		if head == nil || head.InitializerList() == nil {
			return ccallSite{}, false
		}
		site.Args = len(head.InitializerList().AllInitializerClause())
	case *antlr.TerminalNodeImpl:
		if head.GetSymbol().GetTokenType() != parser.SuckCParserLeftParen {
			return ccallSite{}, false
		}
	default:
		return ccallSite{}, false
	}
	// kids[0] is the callee in both shapes this rule accepts, and the offset
	// from the end is not: `f()` ends in [callee ( )] so the callee is at
	// n-3, while `f(a)` ends in [callee ( a )] with an extra ExpressionList
	// node, putting it at n-4. Reading n-3 finds LeftParen whenever the call
	// has arguments, so every argument-taking call silently dropped out.
	name, tokIdx := calleeName(kids[0])
	if name == "" {
		return ccallSite{}, false
	}
	site.Name, site.TokIdx = name, tokIdx
	return site, true
}

func calleeName(t antlr.Tree) (string, int) {
	switch c := t.(type) {
	case *parser.PostfixExpressionContext:
		if c == nil || c.GetChildCount() == 0 {
			return "", 0
		}
		return calleeName(c.GetChild(0))
	case *parser.PrimaryExpressionContext:
		idExpr := childOfIDExpression(c)
		if idExpr == nil {
			return "", 0
		}
		unq := childOfUnqualifiedID(idExpr)
		if unq == nil {
			return "", 0
		}
		ident := childOfIdentifier(unq)
		if ident == nil {
			return "", 0
		}
		return ident.GetText(), ident.GetSymbol().GetTokenIndex()
	}
	return "", 0
}

// collectDeclaratorNames gathers every name SuckC declares, at any scope and
// of any kind, so that arity checking can stand aside for names the .suckc file
// declares itself. Parameters count, which over-collects: standing aside is the
// safe direction.
func childOfIDExpression(t antlr.Tree) *parser.IdExpressionContext {
	for i := 0; i < t.GetChildCount(); i++ {
		if v, ok := t.GetChild(i).(*parser.IdExpressionContext); ok && v != nil {
			return v
		}
	}
	return nil
}

func childOfUnqualifiedID(t antlr.Tree) *parser.UnqualifiedIdContext {
	for i := 0; i < t.GetChildCount(); i++ {
		if v, ok := t.GetChild(i).(*parser.UnqualifiedIdContext); ok && v != nil {
			return v
		}
	}
	return nil
}

func childOfIdentifier(t antlr.Tree) *antlr.TerminalNodeImpl {
	for i := 0; i < t.GetChildCount(); i++ {
		if v, ok := t.GetChild(i).(*antlr.TerminalNodeImpl); ok && v != nil &&
			v.GetSymbol().GetTokenType() == parser.SuckCLexerIdentifier {
			return v
		}
	}
	return nil
}

func tokenType(t antlr.Tree) int {
	if term, ok := t.(*antlr.TerminalNodeImpl); ok && term != nil {
		return term.GetSymbol().GetTokenType()
	}
	return -1
}
