package sema

import (
	"errors"
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"

	"github.com/starofrainnight/suckc/internal/parser"
)

var errNotConstant = errors.New("not a constant expression")

// maxConstDepth bounds the recursion of the constant check; deeper input
// is treated as non-constant rather than risking a stack overflow.
const maxConstDepth = 64

// checkFileConstant verifies that a file-scope auto's initializer is a
// C89 3.4 constant expression. Two forms are accepted outside the
// expression rules because C89 6.5.7 allows them in a static
// initializer: a braced list, and a string literal.
func checkFileConstant(init parser.IInitializerContext, a *Analyzer) error {
	if init == nil {
		return errors.New("missing initializer")
	}
	boi := init.BraceOrEqualInitializer()
	if boi == nil {
		return errUnsupportedExpr
	}
	if braced := boi.BracedInitList(); braced != nil {
		return checkConstBraced(braced, a, 0)
	}
	clause := boi.InitializerClause()
	if clause == nil {
		return errors.New("missing initializer")
	}
	if braced := clause.BracedInitList(); braced != nil {
		return checkConstBraced(braced, a, 0)
	}
	ae := clause.AssignmentExpression()
	if ae == nil {
		return errors.New("missing initializer")
	}
	if isStringLitExpr(ae) {
		return nil
	}
	return checkConstExpr(ae, a, 0)
}

func checkConstBraced(list parser.IBracedInitListContext, a *Analyzer, depth int) error {
	if list == nil || depth > maxConstDepth {
		return nil
	}
	il := list.InitializerList()
	if il == nil {
		return nil
	}
	for _, item := range il.AllInitializerClause() {
		if item == nil {
			return errNotConstant
		}
		if nested := item.BracedInitList(); nested != nil {
			return errUnsupportedExpr
		}
		ae := item.AssignmentExpression()
		if ae == nil {
			return errNotConstant
		}
		if isStringLitExpr(ae) {
			continue
		}
		if err := checkConstExpr(ae, a, depth+1); err != nil {
			return err
		}
	}
	return nil
}

func checkConstExpr(node antlr.Tree, a *Analyzer, depth int) error {
	if depth > maxConstDepth {
		return errNotConstant
	}
	switch n := node.(type) {
	case *parser.LiteralContext:
		// A string literal in an expression position is a value, not an
		// address; only the char-array and char-pointer forms are constant.
		if strings.Contains(n.GetText(), `"`) {
			return errNotConstant
		}
		return nil
	case *parser.UnaryExpressionContext:
		return checkConstUnary(n, a, depth)
	case *parser.PostfixExpressionContext:
		return checkConstPostfix(n, a, depth)
	case *parser.PrimaryExpressionContext:
		return checkConstPrimary(n, a, depth)
	case *parser.ConditionalExpressionContext:
		if n.Question() == nil {
			return checkOnlyChild(n, a, depth)
		}
		ops := nonTerminals(n)
		if len(ops) != 3 {
			return errNotConstant
		}
		for _, op := range ops {
			if err := checkConstExpr(op, a, depth+1); err != nil {
				return err
			}
		}
		return nil
	case *parser.AssignmentExpressionContext:
		if n.AssignmentOperator() != nil || n.ThrowExpression() != nil {
			return errNotConstant
		}
		return checkOnlyChild(n, a, depth)
	case *parser.ExpressionContext:
		// A comma expression is not a constant expression (C89 6.3.15).
		return checkOnlyChild(n, a, depth)
	case *parser.MultiplicativeExpressionContext, *parser.AdditiveExpressionContext,
		*parser.RelationalExpressionContext, *parser.EqualityExpressionContext,
		*parser.AndExpressionContext, *parser.ExclusiveOrExpressionContext,
		*parser.InclusiveOrExpressionContext, *parser.LogicalAndExpressionContext,
		*parser.LogicalOrExpressionContext:
		// A multi-type case leaves n as antlr.Tree, so the rule context has
		// to be recovered before it can be walked. The operator terminals
		// are skipped by nonTerminals, so what is left are the operands.
		chain, ok := n.(antlr.ParserRuleContext)
		if !ok {
			return errNotConstant
		}
		for _, op := range nonTerminals(chain) {
			if err := checkConstExpr(op, a, depth+1); err != nil {
				return err
			}
		}
		return nil
	case *parser.ShiftExpressionContext:
		if n.ShiftOperator(0) != nil {
			return errNotConstant
		}
		return checkOnlyChild(n, a, depth)
	case *parser.PointerMemberExpressionContext:
		return checkOnlyChild(n, a, depth)
	case *parser.CastExpressionContext:
		if n.TheTypeId() != nil {
			return errNotConstant
		}
		return checkOnlyChild(n, a, depth)
	}
	return errNotConstant
}

func checkOnlyChild(ctx antlr.ParserRuleContext, a *Analyzer, depth int) error {
	ops := nonTerminals(ctx)
	if len(ops) != 1 {
		return errNotConstant
	}
	return checkConstExpr(ops[0], a, depth+1)
}

func checkConstUnary(u *parser.UnaryExpressionContext, a *Analyzer, depth int) error {
	if u.Sizeof() != nil || u.Alignof() != nil {
		return nil
	}
	if u.PlusPlus() != nil || u.MinusMinus() != nil || u.NoExceptExpression() != nil ||
		u.NewExpression_() != nil || u.DeleteExpression() != nil {
		return errNotConstant
	}
	if op := u.UnaryOperator(); op != nil {
		switch op.GetText() {
		case "&":
			return checkConstAddr(u.UnaryExpression(), a, depth+1)
		case "*":
			// Dereferencing keeps an lvalue, so &*&x chains stay constant.
			return checkConstAddr(u.UnaryExpression(), a, depth+1)
		}
		return checkConstExpr(u.UnaryExpression(), a, depth+1)
	}
	if p := u.PostfixExpression(); p != nil {
		return checkConstExpr(p, a, depth+1)
	}
	return errNotConstant
}

func checkConstPostfix(p *parser.PostfixExpressionContext, a *Analyzer, depth int) error {
	if p.LeftParen() != nil {
		return errNotConstant
	}
	if hasTerminalText(p, "++", "--") {
		return errNotConstant
	}
	if p.LeftBracket() != nil {
		if err := checkConstAddr(p.PostfixExpression(), a, depth+1); err != nil {
			return err
		}
		return checkConstExpr(p.Expression(), a, depth+1)
	}
	if p.BracedInitList() != nil || p.Dot() != nil || p.Arrow() != nil ||
		p.SimpleTypeSpecifier() != nil || p.TypeNameSpecifier() != nil ||
		p.Dynamic_cast() != nil || p.Static_cast() != nil ||
		p.Reinterpret_cast() != nil || p.Const_cast() != nil {
		return errNotConstant
	}
	if prim := p.PrimaryExpression(); prim != nil {
		return checkConstExpr(prim, a, depth+1)
	}
	return errNotConstant
}

func checkConstPrimary(pr *parser.PrimaryExpressionContext, a *Analyzer, depth int) error {
	if pr.This() != nil || pr.LambdaExpression() != nil {
		return errNotConstant
	}
	if lits := pr.AllLiteral(); len(lits) > 0 {
		if len(lits) > 1 {
			return errNotConstant
		}
		if strings.Contains(lits[0].GetText(), `"`) {
			return errNotConstant
		}
		return nil
	}
	if pr.LeftParen() != nil {
		return checkConstExpr(pr.Expression(), a, depth+1)
	}
	if pr.IdExpression() == nil {
		return errNotConstant
	}
	name, ok := plainName(pr)
	if !ok {
		return errNotConstant
	}
	if a.Enums[name] {
		return nil
	}
	b, found := a.lookupVar(name)
	if !found {
		return fmt.Errorf("unknown identifier '%s'", name)
	}
	if !b.IsConst {
		return errNotConstant
	}
	return checkConstRead(name, a, depth+1)
}

// checkConstRead verifies that a const variable's own initializer is
// itself constant, so a chain of const reads is accepted but a const
// whose value comes from a call is not.
func checkConstRead(name string, a *Analyzer, depth int) error {
	if depth > maxConstDepth {
		return errNotConstant
	}
	tree, ok := a.ConstInits[name]
	if !ok {
		return errNotConstant
	}
	clause, ok := tree.(*parser.InitializerClauseContext)
	if !ok {
		return errNotConstant
	}
	if braced := clause.BracedInitList(); braced != nil {
		return checkConstBraced(braced, a, depth+1)
	}
	ae := clause.AssignmentExpression()
	if ae == nil {
		return errNotConstant
	}
	if isStringLitExpr(ae) {
		return nil
	}
	return checkConstExpr(ae, a, depth+1)
}

// checkConstAddr checks the address constant forms of C89 3.4: &object,
// &array[i] and pointer arithmetic on such an address.
func checkConstAddr(node antlr.Tree, a *Analyzer, depth int) error {
	if depth > maxConstDepth {
		return errNotConstant
	}
	switch n := node.(type) {
	case *parser.PrimaryExpressionContext:
		if n.IdExpression() == nil || n.LeftParen() != nil ||
			len(n.AllLiteral()) > 0 || n.This() != nil || n.LambdaExpression() != nil {
			return errNotConstant
		}
		name, ok := plainName(n)
		if !ok {
			return errNotConstant
		}
		if _, found := a.lookupVar(name); !found {
			return fmt.Errorf("unknown identifier '%s'", name)
		}
		return nil
	case *parser.PostfixExpressionContext:
		if n.LeftBracket() == nil {
			// A postfix with no subscript is just a wrapped primary.
			if prim := n.PrimaryExpression(); prim != nil {
				return checkConstAddr(prim, a, depth+1)
			}
			return errNotConstant
		}
		if err := checkConstAddr(n.PostfixExpression(), a, depth+1); err != nil {
			return err
		}
		return checkConstExpr(n.Expression(), a, depth+1)
	case *parser.UnaryExpressionContext:
		// The operand of & is an lvalue, so its own constness is
		// irrelevant; only the address form is checked.
		if op := n.UnaryOperator(); op != nil {
			if op.GetText() == "*" {
				return checkConstAddr(n.UnaryExpression(), a, depth+1)
			}
			return errNotConstant
		}
		if p := n.PostfixExpression(); p != nil {
			return checkConstAddr(p, a, depth+1)
		}
		return errNotConstant
	}
	return checkConstExpr(node, a, depth)
}
