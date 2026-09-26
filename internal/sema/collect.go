package sema

import (
	"github.com/antlr4-go/antlr/v4"

	"github.com/starofrainnight/suckc/internal/frontend"
	"github.com/starofrainnight/suckc/internal/parser"
)

// Analyzer carries shared state across the two passes (spec section 4.1).
type Analyzer struct {
	Opts     Options
	Scopes   *scope
	Funcs    map[string]Type    // pass 1: function name -> return type
	FileVars map[string]binding // pass 1: file-scope name -> binding
	Subs     Subs               // pass 2: Auto token index -> replacement
}

func newAnalyzer(opts Options) *Analyzer {
	return &Analyzer{
		Opts:     opts,
		Scopes:   &scope{},
		Funcs:    map[string]Type{},
		FileVars: map[string]binding{},
		Subs:     Subs{},
	}
}

// lookupVar resolves a variable name: pass-2 scopes first, then file scope.
func (a *Analyzer) lookupVar(name string) (binding, bool) {
	if b, ok := a.Scopes.lookup(name); ok {
		return b, true
	}
	b, ok := a.FileVars[name]
	return b, ok
}

// pass1 walks top-level declarations only: function signatures (with or
// without body) and file-scope simple declarations. Function bodies and
// nested blocks are not entered.
func pass1(res *frontend.ParseResult, a *Analyzer) error {
	// zero-value Analyzer stays usable: tests build the struct literally
	if a.Funcs == nil {
		a.Funcs = map[string]Type{}
	}
	if a.FileVars == nil {
		a.FileVars = map[string]binding{}
	}
	return walkFileScope(res.Tree, a)
}

func walkFileScope(node antlr.Tree, a *Analyzer) error {
	switch ctx := node.(type) {
	case *parser.FunctionDefinitionContext:
		return collectFunction(ctx, a)
	case *parser.SimpleDeclarationContext:
		return collectFileDecl(ctx, a)
	case antlr.TerminalNode:
		return nil // EOF, ';', stray tokens contribute no bindings
	default:
		return walkChildren(node, a)
	}
}

func walkChildren(node antlr.Tree, a *Analyzer) error {
	pr, ok := node.(antlr.ParserRuleContext)
	if !ok {
		return nil
	}
	for _, ch := range pr.GetChildren() {
		if err := walkFileScope(ch, a); err != nil {
			return err
		}
	}
	return nil
}

// collectFunction records name -> return Type. Parameters are ignored
// (auto parameters are a hard error elsewhere; arg types are unchecked).
func collectFunction(ctx *parser.FunctionDefinitionContext, a *Analyzer) error {
	name, err := declaratorName(ctx.Declarator())
	if err != nil {
		return err
	}
	if name == "" {
		return nil // operator overloads etc.: out of scope, skip
	}
	ret := Type{"int", 0}
	if seq := ctx.DeclSpecifierSeq(); seq != nil {
		ret = Type{spellingOf(seq), 0}
	}
	shape, err := analyzeDeclarator(ctx.Declarator())
	if err != nil {
		return err
	}
	if shape.IsFuncPtr {
		return nil // function returning function pointer: deferred (spec 11)
	}
	ret.Stars = shape.Stars
	a.Funcs[name] = ret
	return nil
}

// collectFileDecl records a file-scope variable or function prototype.
// Function prototypes (declarator with parameters) go to Funcs.
func collectFileDecl(ctx *parser.SimpleDeclarationContext, a *Analyzer) error {
	seq := ctx.DeclSpecifierSeq()
	initDecl := ctx.InitDeclarator()
	if seq == nil || initDecl == nil || hasAuto(seq) {
		return nil // file-scope auto: rejected by the forbidden-position scan
	}
	d := initDecl.Declarator()
	if d == nil {
		return nil
	}
	name, err := declaratorName(d)
	if err != nil || name == "" {
		return err
	}
	shape, err := analyzeDeclarator(d)
	if err != nil {
		return err
	}
	if shape.IsFuncPtr {
		return nil // deferred (spec 11): function-pointer variables
	}
	if shape.HasParams {
		a.Funcs[name] = Type{spellingOf(seq), shape.Stars}
		return nil
	}
	a.FileVars[name] = binding{
		Type:    Type{spellingOf(seq), shape.Stars},
		IsArray: shape.IsArray,
	}
	return nil
}

// spellingOf renders a declSpecifierSeq as type text: the terminal texts
// joined with single spaces, minus storage/function specifiers and
// attributes. Auto is never part of a recorded spelling (callers check).
func spellingOf(seq parser.IDeclSpecifierSeqContext) string {
	if seq == nil {
		return ""
	}
	var parts []string
	var walk func(n antlr.Tree)
	skip := map[int]bool{
		parser.SuckCParserStatic:       true,
		parser.SuckCParserExtern:       true,
		parser.SuckCParserRegister:     true,
		parser.SuckCParserThread_local: true,
		parser.SuckCParserInline:       true,
		parser.SuckCParserMutable:      true,
		parser.SuckCParserVirtual:      true,
		parser.SuckCParserExplicit:     true,
		parser.SuckCParserConstexpr:    true,
		parser.SuckCParserFriend:       true,
	}
	walk = func(n antlr.Tree) {
		if t, ok := n.(antlr.TerminalNode); ok {
			if skip[t.GetSymbol().GetTokenType()] {
				return
			}
			txt := t.GetText()
			if txt != "" {
				parts = append(parts, txt)
			}
			return
		}
		if pr, ok := n.(antlr.ParserRuleContext); ok {
			for _, ch := range pr.GetChildren() {
				walk(ch)
			}
		}
	}
	walk(seq)
	return joinSpaces(parts)
}

func joinSpaces(parts []string) string {
	out := ""
	for i, p := range parts {
		if i > 0 {
			out += " "
		}
		out += p
	}
	return out
}

// hasAuto reports whether seq contains an Auto token.
func hasAuto(seq parser.IDeclSpecifierSeqContext) bool {
	if seq == nil {
		return false
	}
	found := false
	var walk func(n antlr.Tree)
	walk = func(n antlr.Tree) {
		if t, ok := n.(antlr.TerminalNode); ok {
			if t.GetSymbol().GetTokenType() == parser.SuckCParserAuto {
				found = true
			}
			return
		}
		if pr, ok := n.(antlr.ParserRuleContext); ok {
			for _, ch := range pr.GetChildren() {
				walk(ch)
			}
		}
	}
	walk(seq)
	return found
}

// declaratorName extracts the declared identifier from a declarator.
func declaratorName(d parser.IDeclaratorContext) (string, error) {
	if d == nil {
		return "", nil
	}
	found := ""
	var walk func(n antlr.Tree)
	walk = func(n antlr.Tree) {
		if found != "" {
			return
		}
		if ctx, ok := n.(*parser.DeclaratoridContext); ok {
			if id := ctx.IdExpression(); id != nil {
				if u := id.UnqualifiedId(); u != nil {
					if ident := u.Identifier(); ident != nil {
						found = ident.GetText()
					}
				}
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
	return found, nil
}

type declShape struct {
	Stars     int  // pointer stars directly on this declarator
	IsFuncPtr bool // parenthesized pointer declarator, e.g. (*fp)(...)
	IsArray   bool // has array suffix, e.g. a[10] or a[]
	HasParams bool // has a parameter list, i.e. declares a function/prototype
}

// analyzeDeclarator walks a declarator once and classifies its shape.
// Star counting skips parameter lists and array bounds: the Star token is
// shared with the multiplicative operator, so `a[2*3]` must not count.
func analyzeDeclarator(d parser.IDeclaratorContext) (declShape, error) {
	var s declShape
	if d == nil {
		return s, nil
	}
	bracketDepth := 0
	var walk func(n antlr.Tree)
	walk = func(n antlr.Tree) {
		if t, ok := n.(antlr.TerminalNode); ok {
			switch t.GetSymbol().GetTokenType() {
			case parser.SuckCParserStar:
				if bracketDepth == 0 {
					s.Stars++
				}
			case parser.SuckCParserLeftBracket:
				s.IsArray = true
				bracketDepth++
			case parser.SuckCParserRightBracket:
				if bracketDepth > 0 {
					bracketDepth--
				}
			}
			return
		}
		switch ctx := n.(type) {
		case *parser.ParametersAndQualifiersContext:
			s.HasParams = true
			return // parameter stars are not ours
		case *parser.NoPointerDeclaratorContext:
			// parenthesized form: LeftParen pointerDeclarator RightParen
			if ctx.LeftParen() != nil && ctx.PointerDeclarator() != nil {
				s.IsFuncPtr = true
			}
		}
		if pr, ok := n.(antlr.ParserRuleContext); ok {
			for _, ch := range pr.GetChildren() {
				walk(ch)
			}
		}
	}
	walk(d)
	return s, nil
}
