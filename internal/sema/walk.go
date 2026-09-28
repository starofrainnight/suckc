package sema

import (
	"errors"
	"strings"

	"github.com/antlr4-go/antlr/v4"

	"github.com/starofrainnight/suckc/internal/frontend"
	"github.com/starofrainnight/suckc/internal/parser"
)

// walker is a pass-2 listener over function bodies.
type walker struct {
	antlr.BaseParseTreeListener
	a      *Analyzer
	inFunc int   // function nesting depth (>0 => inside a function body)
	err    error // first error wins
}

// pass2 walks function bodies in source order, deducing auto declarations.
func pass2(res *frontend.ParseResult, a *Analyzer) error {
	w := &walker{a: a}
	l := antlr.NewParseTreeWalker()
	l.Walk(w, res.Tree)
	return w.err
}

func (w *walker) fail(err error) {
	if w.err == nil {
		w.err = err
	}
}

// errAt wraps a position-less deduction error with the Auto token position.
func (w *walker) errAt(tok antlr.Token, err error) error {
	return &Error{
		File: w.a.Opts.FileName,
		Line: tok.GetLine(),
		Col:  tok.GetColumn() + 1,
		Msg:  err.Error(),
	}
}

func (w *walker) EnterEveryRule(node antlr.ParserRuleContext) {
	if w.err != nil {
		return
	}
	switch ctx := node.(type) {
	case *parser.FunctionDefinitionContext:
		w.inFunc++
		w.a.Scopes.push()
		w.declareParams(ctx)
	case *parser.CompoundStatementContext:
		w.a.Scopes.push()
	case *parser.IterationStatementContext:
		w.a.Scopes.push()
		w.declareCondition(ctx)
	case *parser.SelectionStatementContext:
		// C++-style condition declaration: if (int x = ...) — record it.
		w.declareCondition(ctx)
	case *parser.SimpleDeclarationContext:
		if w.inFunc > 0 {
			w.handleDecl(ctx)
		}
	}
}

func (w *walker) ExitEveryRule(node antlr.ParserRuleContext) {
	if w.err != nil {
		return
	}
	switch node.(type) {
	case *parser.FunctionDefinitionContext:
		w.a.Scopes.pop()
		w.inFunc--
	case *parser.CompoundStatementContext:
		w.a.Scopes.pop()
	case *parser.IterationStatementContext:
		w.a.Scopes.pop()
	}
}

// declareParams records named non-auto parameters of a function.
func (w *walker) declareParams(fd *parser.FunctionDefinitionContext) {
	d := fd.Declarator()
	if d == nil {
		return
	}
	var walk func(n antlr.Tree)
	walk = func(n antlr.Tree) {
		pd, ok := n.(*parser.ParameterDeclarationContext)
		if !ok {
			if pr, ok := n.(antlr.ParserRuleContext); ok {
				for _, ch := range pr.GetChildren() {
					walk(ch)
				}
			}
			return
		}
		seq := pd.DeclSpecifierSeq()
		dcl := pd.Declarator()
		if seq == nil || dcl == nil || hasAuto(seq) {
			return // auto params: forbidden scan rejects; skip recording
		}
		name, err := declaratorName(dcl)
		if err != nil || name == "" {
			return
		}
		shape, err := analyzeDeclarator(dcl)
		if err != nil {
			w.fail(err)
			return
		}
		w.a.Scopes.declare(name, binding{
			Type:    Type{spellingOf(seq), shape.Stars},
			IsArray: shape.IsArray,
		})
	}
	walk(d)
}

// declareCondition records a condition-with-declaration (if/while/switch).
// Takes the statement context directly so both selection and iteration
// arms can call it.
func (w *walker) declareCondition(n antlr.ParserRuleContext) {
	// condition contexts appear under selection/iteration statements; walk
	// their subtree once for ConditionContext nodes.
	var walk func(nt antlr.Tree)
	walk = func(nt antlr.Tree) {
		if w.err != nil {
			return
		}
		if cond, ok := nt.(*parser.ConditionContext); ok {
			seq := cond.DeclSpecifierSeq()
			dcl := cond.Declarator()
			if seq != nil && dcl != nil && !hasAuto(seq) {
				if name, err := declaratorName(dcl); err == nil && name != "" {
					if shape, err := analyzeDeclarator(dcl); err == nil {
						w.a.Scopes.declare(name, binding{
							Type:    Type{spellingOf(seq), shape.Stars},
							IsArray: shape.IsArray,
						})
					}
				}
			}
			return // do not descend further
		}
		if pr, ok := nt.(antlr.ParserRuleContext); ok {
			for _, ch := range pr.GetChildren() {
				walk(ch)
			}
		}
	}
	if n != nil {
		walk(n)
	}
}

// handleDecl deduces an auto declaration inside a function body.
func (w *walker) handleDecl(ctx *parser.SimpleDeclarationContext) {
	seq := ctx.DeclSpecifierSeq()
	if seq == nil || !hasAuto(seq) {
		// plain declaration: record into current scope for later lookups
		if seq != nil {
			w.recordPlainDecl(ctx)
		}
		return
	}
	autoNode := findAutoToken(seq)
	if autoNode == nil {
		return
	}
	autoTok := autoNode.GetSymbol()
	idc := ctx.InitDeclarator()
	if idc == nil {
		w.fail(w.errAt(autoTok, errors.New("missing initializer")))
		return
	}
	d := idc.Declarator()
	if d == nil {
		w.fail(w.errAt(autoTok, errors.New("missing initializer")))
		return
	}
	shape, err := analyzeDeclarator(d)
	if err != nil {
		w.fail(w.errAt(autoTok, err))
		return
	}
	if shape.IsFuncPtr {
		w.fail(w.errAt(autoTok, errors.New("function-pointer declarator not supported")))
		return
	}
	init := idc.Initializer()
	if init == nil {
		w.fail(w.errAt(autoTok, errors.New("missing initializer")))
		return
	}
	base, err := w.baseType(init, shape.IsArray, autoTok)
	if err != nil {
		w.fail(w.errAt(autoTok, err))
		return
	}
	if base.Stars < shape.Stars {
		w.fail(w.errAt(autoTok, errors.New("insufficient pointer depth")))
		return
	}
	final := Type{base.Spelling, base.Stars - shape.Stars}
	name, _ := declaratorName(d)
	if name != "" {
		w.a.Scopes.declare(name, binding{Type: final, IsArray: shape.IsArray})
	}
	w.a.Subs[autoTok.GetTokenIndex()] = final.Text()
}

// baseType computes the deduced base type of an initializer (spec 6.3).
func (w *walker) baseType(init parser.IInitializerContext, isArray bool, autoTok antlr.Token) (Type, error) {
	boi := init.BraceOrEqualInitializer()
	if boi == nil {
		// paren initializer: auto i(12) — not in spec 6.3 table
		return Type{}, errUnsupportedExpr
	}
	if braced := boi.BracedInitList(); braced != nil {
		// bare brace form without '=' (auto a {1,2}) — treat as braced
		return w.bracedBase(braced, isArray)
	}
	clause := boi.InitializerClause()
	if clause == nil {
		return Type{}, errors.New("missing initializer")
	}
	if braced := clause.BracedInitList(); braced != nil {
		return w.bracedBase(braced, isArray)
	}
	ae := clause.AssignmentExpression()
	if ae == nil {
		return Type{}, errors.New("missing initializer")
	}
	if isArray {
		if isStringLitExpr(ae) {
			return Type{"char", 0}, nil
		}
		return Type{}, errUnsupportedExpr
	}
	return typeOf(ae, w.a)
}

// bracedBase deduces the element type from the FIRST element (spec 6.3).
func (w *walker) bracedBase(braced parser.IBracedInitListContext, isArray bool) (Type, error) {
	if !isArray {
		return Type{}, errUnsupportedExpr
	}
	list := braced.InitializerList()
	if list == nil {
		return Type{}, errors.New("empty initializer list")
	}
	first := list.InitializerClause(0)
	if first == nil {
		return Type{}, errors.New("empty initializer list")
	}
	if nested := first.BracedInitList(); nested != nil {
		return Type{}, errUnsupportedExpr
	}
	ae := first.AssignmentExpression()
	if ae == nil {
		return Type{}, errors.New("empty initializer list")
	}
	return typeOf(ae, w.a)
}

// isStringLitExpr reports whether the expression is a single string literal.
func isStringLitExpr(ae parser.IAssignmentExpressionContext) bool {
	node := antlr.Tree(ae)
	for depth := 0; depth < 32; depth++ {
		switch n := node.(type) {
		case *parser.LiteralContext:
			return strings.HasPrefix(n.GetText(), `"`)
		case *parser.PrimaryExpressionContext:
			if lits := n.AllLiteral(); len(lits) == 1 && n.IdExpression() == nil &&
				n.LeftParen() == nil && n.This() == nil && n.LambdaExpression() == nil {
				node = lits[0]
				continue
			}
			return false
		case *parser.PostfixExpressionContext:
			if p := n.PrimaryExpression(); p != nil && n.PostfixExpression() == nil &&
				n.LeftBracket() == nil && n.LeftParen() == nil {
				node = p
				continue
			}
			return false
		case *parser.UnaryExpressionContext:
			if p := n.PostfixExpression(); p != nil {
				node = p
				continue
			}
			return false
		case *parser.CastExpressionContext:
			if u := n.UnaryExpression(); u != nil && n.LeftParen() == nil {
				node = u
				continue
			}
			return false
		case *parser.MultiplicativeExpressionContext, *parser.AdditiveExpressionContext,
			*parser.ShiftExpressionContext, *parser.RelationalExpressionContext,
			*parser.EqualityExpressionContext, *parser.AndExpressionContext,
			*parser.ExclusiveOrExpressionContext, *parser.InclusiveOrExpressionContext,
			*parser.LogicalAndExpressionContext, *parser.LogicalOrExpressionContext,
			*parser.ConditionalExpressionContext, *parser.AssignmentExpressionContext,
			*parser.PointerMemberExpressionContext, *parser.ExpressionContext:
			ops := nonTerminals(n.(antlr.ParserRuleContext))
			if len(ops) == 1 {
				node = ops[0]
				continue
			}
			return false
		default:
			return false
		}
	}
	return false
}

// recordPlainDecl records a non-auto declaration inside a function body.
func (w *walker) recordPlainDecl(ctx *parser.SimpleDeclarationContext) {
	seq := ctx.DeclSpecifierSeq()
	idc := ctx.InitDeclarator()
	if idc == nil {
		return
	}
	d := idc.Declarator()
	if d == nil {
		return
	}
	name, err := declaratorName(d)
	if err != nil || name == "" {
		return
	}
	shape, err := analyzeDeclarator(d)
	if err != nil {
		w.fail(err)
		return
	}
	if shape.IsFuncPtr || shape.HasParams {
		return // nested function/func-ptr: not recorded (deferred)
	}
	w.a.Scopes.declare(name, binding{
		Type:    Type{spellingOf(seq), shape.Stars},
		IsArray: shape.IsArray,
	})
}

// findAutoToken returns the Auto terminal inside a declSpecifierSeq.
func findAutoToken(seq parser.IDeclSpecifierSeqContext) antlr.TerminalNode {
	if seq == nil {
		return nil
	}
	var found antlr.TerminalNode
	var walk func(n antlr.Tree)
	walk = func(n antlr.Tree) {
		if found != nil {
			return
		}
		if t, ok := n.(antlr.TerminalNode); ok {
			if t.GetSymbol().GetTokenType() == parser.SuckCParserAuto {
				found = t
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

// scanForbiddenAutos rejects every Auto token pass 2 did not consume
// (file scope, return type, parameters, decltype, class members, ...).
func scanForbiddenAutos(res *frontend.ParseResult, a *Analyzer) error {
	for _, tok := range res.Tokens {
		if tok.GetTokenType() != parser.SuckCParserAuto {
			continue
		}
		if _, ok := a.Subs[tok.GetTokenIndex()]; ok {
			continue
		}
		return &Error{
			File: a.Opts.FileName,
			Line: tok.GetLine(),
			Col:  tok.GetColumn() + 1,
			Msg:  "'auto' not allowed here",
		}
	}
	return nil
}
