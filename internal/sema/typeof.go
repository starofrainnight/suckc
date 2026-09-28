package sema

import (
	"errors"
	"fmt"

	"github.com/antlr4-go/antlr/v4"

	"github.com/starofrainnight/suckc/internal/parser"
)

var errUnsupportedExpr = errors.New("unsupported expression")

// typeOf deduces the type of an expression node (spec section 6.2).
// Errors carry no position; the auto-declaration walker prefixes
// file:line:col from the Auto token.
func typeOf(node antlr.Tree, a *Analyzer) (Type, error) {
	switch ctx := node.(type) {
	case *parser.ExpressionContext:
		// comma expression: deferred (spec 11)
		ops := nonTerminals(ctx)
		if len(ops) != 1 {
			return Type{}, errUnsupportedExpr
		}
		return typeOf(ops[0], a)

	case *parser.AssignmentExpressionContext:
		ops := nonTerminals(ctx)
		if len(ops) != 1 {
			// assignment-as-expression: not in the spec table
			return Type{}, errUnsupportedExpr
		}
		return typeOf(ops[0], a)

	case *parser.ConditionalExpressionContext:
		if !hasTerminalText(ctx, "?") {
			ops := nonTerminals(ctx)
			if len(ops) != 1 {
				return Type{}, errUnsupportedExpr
			}
			return typeOf(ops[0], a)
		}
		ops := nonTerminals(ctx)
		if len(ops) != 3 {
			return Type{}, errUnsupportedExpr
		}
		if _, err := typeOf(ops[0], a); err != nil {
			return Type{}, err
		}
		t1, err := typeOf(ops[1], a)
		if err != nil {
			return Type{}, err
		}
		t2, err := typeOf(ops[2], a)
		if err != nil {
			return Type{}, err
		}
		ct, err := commonType(t1, t2, a.Opts)
		if err != nil {
			return Type{}, fmt.Errorf("incompatible operand types '%s' and '%s'",
				t1.Text(), t2.Text())
		}
		return ct, nil

	case *parser.MultiplicativeExpressionContext, *parser.AdditiveExpressionContext,
		*parser.AndExpressionContext, *parser.ExclusiveOrExpressionContext,
		*parser.InclusiveOrExpressionContext:
		return foldBinary(ctx.(antlr.ParserRuleContext), a)

	case *parser.RelationalExpressionContext, *parser.EqualityExpressionContext,
		*parser.LogicalAndExpressionContext, *parser.LogicalOrExpressionContext:
		return foldPredicate(ctx.(antlr.ParserRuleContext), a)

	case *parser.ShiftExpressionContext:
		// shift is not in the spec 6.2 table: fail loud, but a bare chain
		// passes through this rule with no operator — only unwrap that.
		ops := nonTerminals(ctx)
		if len(ops) != 1 {
			return Type{}, errUnsupportedExpr
		}
		return typeOf(ops[0], a)

	case *parser.CastExpressionContext:
		if hasTerminalText(ctx, "(") {
			return Type{}, errUnsupportedExpr // C-style cast deferred (spec 11)
		}
		return onlyChildOfType(ctx, a)

	case *parser.PointerMemberExpressionContext:
		if hasTerminalText(ctx, ".*", "->*") {
			return Type{}, errUnsupportedExpr
		}
		return onlyChildOfType(ctx, a)

	case *parser.UnaryExpressionContext:
		return unaryTypeOf(ctx, a)

	case *parser.PostfixExpressionContext:
		return postfixTypeOf(ctx, a)

	case *parser.PrimaryExpressionContext:
		return primaryTypeOf(ctx, a)

	case *parser.LiteralContext:
		return deduceLiteralText(ctx.GetText(), a.Opts)

	default:
		return Type{}, errUnsupportedExpr
	}
}

func unaryTypeOf(u *parser.UnaryExpressionContext, a *Analyzer) (Type, error) {
	if u.Sizeof() != nil {
		if u.TheTypeId() != nil {
			return sizeType(a.Opts), nil
		}
		inner := u.UnaryExpression()
		if inner == nil {
			return Type{}, errUnsupportedExpr // sizeof... form
		}
		if _, err := typeOf(inner, a); err != nil {
			return Type{}, err
		}
		return sizeType(a.Opts), nil
	}
	if op := u.UnaryOperator(); op != nil {
		inner := u.UnaryExpression()
		if inner == nil {
			return Type{}, errUnsupportedExpr
		}
		t, err := typeOf(inner, a)
		if err != nil {
			return Type{}, err
		}
		switch op.GetText() {
		case "&":
			// The address of an object keeps its qualifiers, so a plain
			// name operand is read as declared instead of through the
			// value-read path, which drops cv.
			if lt, ok := lvalueType(inner, a); ok {
				return Type{lt.Spelling, lt.Stars + 1}, nil
			}
			return Type{t.Spelling, t.Stars + 1}, nil
		case "*":
			if t.Stars == 0 {
				return Type{}, errors.New("insufficient pointer depth")
			}
			return Type{t.Spelling, t.Stars - 1}, nil
		case "+", "-":
			return promote(t), nil
		case "!", "~":
			// operand validated above; result is always int (spec 6.2)
			return Type{"int", 0}, nil
		default:
			return Type{}, errUnsupportedExpr
		}
	}
	if u.PlusPlus() != nil || u.MinusMinus() != nil {
		return Type{}, errUnsupportedExpr // ++/-- deferred (spec 11)
	}
	if u.Alignof() != nil || u.NoExceptExpression() != nil ||
		u.NewExpression_() != nil || u.DeleteExpression() != nil {
		return Type{}, errUnsupportedExpr
	}
	if inner := u.PostfixExpression(); inner != nil {
		return typeOf(inner, a)
	}
	return Type{}, errUnsupportedExpr
}

func postfixTypeOf(p *parser.PostfixExpressionContext, a *Analyzer) (Type, error) {
	if p.LeftBracket() != nil {
		base, err := typeOf(p.PostfixExpression(), a)
		if err != nil {
			return Type{}, err
		}
		if idx := p.Expression(); idx != nil {
			if _, err := typeOf(idx, a); err != nil {
				return Type{}, err
			}
		}
		if base.Stars == 0 {
			return Type{}, errors.New("insufficient pointer depth")
		}
		return Type{base.Spelling, base.Stars - 1}, nil
	}
	if p.LeftParen() != nil && p.PostfixExpression() != nil {
		// function call: return type only, arguments unchecked (spec 6.2)
		name, ok := plainName(p.PostfixExpression())
		if !ok {
			return Type{}, errors.New("function-pointer call not supported")
		}
		if ret, ok := a.Funcs[name]; ok {
			return ret, nil
		}
		if _, isVar := a.lookupVar(name); isVar {
			return Type{}, errors.New("function-pointer call not supported")
		}
		return Type{}, fmt.Errorf("unknown identifier '%s'", name)
	}
	if p.Dot() != nil || p.Arrow() != nil {
		return Type{}, errUnsupportedExpr // member access deferred (spec 11)
	}
	if hasTerminalText(p, "++", "--") {
		return Type{}, errUnsupportedExpr
	}
	if p.SimpleTypeSpecifier() != nil || p.TypeNameSpecifier() != nil {
		return Type{}, errUnsupportedExpr // functional cast deferred
	}
	if p.Dynamic_cast() != nil || p.Static_cast() != nil ||
		p.Reinterpret_cast() != nil || p.Const_cast() != nil {
		return Type{}, errUnsupportedExpr
	}
	if prim := p.PrimaryExpression(); prim != nil && p.PostfixExpression() == nil {
		return typeOf(prim, a)
	}
	return Type{}, errUnsupportedExpr
}

// lvalueType returns the declared type of a plain name operand, keeping
// its qualifiers and applying array decay. It reports false when the
// operand is not a simple name, in which case the caller falls back to
// the ordinary expression type.
func lvalueType(node antlr.Tree, a *Analyzer) (Type, bool) {
	for {
		switch n := node.(type) {
		case *parser.PrimaryExpressionContext:
			if n.LeftParen() != nil || n.IdExpression() == nil || len(n.AllLiteral()) > 0 {
				return Type{}, false
			}
			name, ok := plainName(n)
			if !ok {
				return Type{}, false
			}
			b, found := a.lookupVar(name)
			if !found {
				return Type{}, false
			}
			if b.IsArray {
				return Type{b.Type.Spelling, b.Type.Stars + 1}, true
			}
			return b.Type, true
		case *parser.PostfixExpressionContext:
			if p := n.PostfixExpression(); p != nil {
				node = p
				continue
			}
			if p := n.PrimaryExpression(); p != nil {
				node = p
				continue
			}
			return Type{}, false
		case *parser.UnaryExpressionContext:
			if p := n.PostfixExpression(); p != nil {
				node = p
				continue
			}
			if e := n.UnaryExpression(); e != nil {
				node = e
				continue
			}
			return Type{}, false
		default:
			return Type{}, false
		}
	}
}

func primaryTypeOf(pr *parser.PrimaryExpressionContext, a *Analyzer) (Type, error) {
	if pr.LeftParen() != nil {
		inner := pr.Expression()
		if inner == nil {
			return Type{}, errUnsupportedExpr
		}
		return typeOf(inner, a)
	}
	if pr.IdExpression() != nil {
		name, ok := plainName(pr)
		if !ok {
			// `~x` parses as unqualified-id destructor form (Tilde className)
			// rather than unary ~; treat as bitwise complement: operand must
			// resolve, result is int (spec 6.2).
			if uq := pr.IdExpression().UnqualifiedId(); uq != nil &&
				uq.Tilde() != nil && uq.ClassName() != nil && pr.IdExpression().QualifiedId() == nil {
				operand := uq.ClassName().GetText()
				if _, found := a.lookupVar(operand); !found {
					return Type{}, fmt.Errorf("unknown identifier '%s'", operand)
				}
				return Type{"int", 0}, nil
			}
			return Type{}, errUnsupportedExpr // qualified id, operator(), ...
		}
		if b, found := a.lookupVar(name); found {
			if b.IsArray {
				// array decay: element type with Stars+1 (spec 6.2)
				return Type{b.Type.Spelling, b.Type.Stars + 1}, nil
			}
			// A value read yields the unqualified type, so `auto j = c;`
			// on a const c must not produce `const int j`.
			return Type{stripCV(b.Type.Spelling), b.Type.Stars}, nil
		}
		// An enum constant has type int (C89 3.1.1.1) and is not a name in
		// scope, so it is resolved from the collected enumerators.
		if a.Enums[name] {
			return Type{"int", 0}, nil
		}
		if _, isFn := a.Funcs[name]; isFn {
			// function used as value: function-pointer, deferred (spec 11)
			return Type{}, errUnsupportedExpr
		}
		return Type{}, fmt.Errorf("unknown identifier '%s'", name)
	}
	if lits := pr.AllLiteral(); len(lits) > 0 {
		if len(lits) > 1 {
			return Type{}, errUnsupportedExpr // string concatenation (spec 11)
		}
		return deduceLiteralText(lits[0].GetText(), a.Opts)
	}
	return Type{}, errUnsupportedExpr // this, lambda
}

// foldBinary folds arithmetic/bitwise chains with usual conversions.
func foldBinary(ctx antlr.ParserRuleContext, a *Analyzer) (Type, error) {
	ops := nonTerminals(ctx)
	if len(ops) == 0 {
		return Type{}, errUnsupportedExpr
	}
	left, err := typeOf(ops[0], a)
	if err != nil {
		return Type{}, err
	}
	for _, next := range ops[1:] {
		right, err := typeOf(next, a)
		if err != nil {
			return Type{}, err
		}
		combined, convErr := usualArith(left, right, a.Opts)
		if convErr != nil {
			return Type{}, fmt.Errorf("incompatible operand types '%s' and '%s'",
				left.Text(), right.Text())
		}
		left = combined
	}
	return left, nil
}

// foldPredicate validates every operand of a comparison/logical chain and
// yields int.
func foldPredicate(ctx antlr.ParserRuleContext, a *Analyzer) (Type, error) {
	ops := nonTerminals(ctx)
	if len(ops) == 0 {
		return Type{}, errUnsupportedExpr
	}
	if len(ops) == 1 {
		return typeOf(ops[0], a)
	}
	for _, o := range ops {
		if _, err := typeOf(o, a); err != nil {
			return Type{}, err
		}
	}
	return Type{"int", 0}, nil
}

func onlyChildOfType(ctx antlr.ParserRuleContext, a *Analyzer) (Type, error) {
	ops := nonTerminals(ctx)
	if len(ops) != 1 {
		return Type{}, errUnsupportedExpr
	}
	return typeOf(ops[0], a)
}

func sizeType(o Options) Type {
	return Type{"size_t", 0} // C89 size_t; o unused (kept for call-site symmetry)
}

// plainName extracts a bare identifier from a primary expression.
func plainName(n antlr.Tree) (string, bool) {
	switch t := n.(type) {
	case *parser.PrimaryExpressionContext:
		if t.IdExpression() == nil {
			return "", false
		}
		id := t.IdExpression()
		if id.UnqualifiedId() == nil || id.QualifiedId() != nil {
			return "", false
		}
		ident := id.UnqualifiedId().Identifier()
		if ident == nil {
			return "", false
		}
		return ident.GetText(), true
	case *parser.PostfixExpressionContext:
		// call callee is wrapped: postfixExpression -> primaryExpression
		if t.PostfixExpression() == nil && t.PrimaryExpression() != nil {
			return plainName(t.PrimaryExpression())
		}
		return "", false
	}
	return "", false
}

// nonTerminals returns the non-terminal children of a rule context.
func nonTerminals(n antlr.ParserRuleContext) []antlr.Tree {
	var out []antlr.Tree
	for _, ch := range n.GetChildren() {
		if _, ok := ch.(antlr.TerminalNode); ok {
			continue
		}
		out = append(out, ch)
	}
	return out
}

// hasTerminalText reports whether any direct terminal child has one of the
// given texts.
func hasTerminalText(n antlr.ParserRuleContext, texts ...string) bool {
	for _, ch := range n.GetChildren() {
		if t, ok := ch.(antlr.TerminalNode); ok {
			got := t.GetText()
			for _, want := range texts {
				if got == want {
					return true
				}
			}
		}
	}
	return false
}
