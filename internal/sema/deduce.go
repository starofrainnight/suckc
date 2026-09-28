package sema

import (
	"errors"

	"github.com/antlr4-go/antlr/v4"

	"github.com/starofrainnight/suckc/internal/parser"
)

// errAt wraps err with the source position of tok, so every diagnostic
// is reported as file:line:col.
func errAt(a *Analyzer, tok antlr.Token, err error) error {
	return &Error{
		File: a.Opts.FileName,
		Line: tok.GetLine(),
		Col:  tok.GetColumn() + 1,
		Msg:  err.Error(),
	}
}

// deduceAuto is the single deduction path for both file and block scope;
// fileScope adds the C89 3.4 constant gate and records the result in
// FileVars instead of the scope stack.
func deduceAuto(a *Analyzer, ctx *parser.SimpleDeclarationContext, fileScope bool) error {
	autoNode := findAutoToken(ctx.DeclSpecifierSeq())
	if autoNode == nil {
		return nil
	}
	autoTok := autoNode.GetSymbol()
	initDecl := ctx.InitDeclarator()
	if initDecl == nil || initDecl.Declarator() == nil || initDecl.Initializer() == nil {
		return errAt(a, autoTok, errors.New("missing initializer"))
	}
	d := initDecl.Declarator()
	shape, err := analyzeDeclarator(d)
	if err != nil {
		return errAt(a, autoTok, err)
	}
	if shape.IsFuncPtr {
		return errAt(a, autoTok, errors.New("function-pointer declarator not supported"))
	}
	if fileScope {
		if err := checkFileConstant(initDecl.Initializer(), a); err != nil {
			return errAt(a, autoTok, err)
		}
	}
	base, err := baseType(a, initDecl.Initializer(), shape.IsArray, autoTok)
	if err != nil {
		return errAt(a, autoTok, err)
	}
	if base.Stars < shape.Stars {
		return errAt(a, autoTok, errors.New("insufficient pointer depth"))
	}
	// The star of a declarator such as "auto *p" stays in the source, so
	// the substituted text is the type without it, while p itself has the
	// full type.
	final := Type{Spelling: base.Spelling, Stars: base.Stars - shape.Stars}
	a.Subs[autoTok.GetTokenIndex()] = final.Text()
	name, err := declaratorName(d)
	if err != nil || name == "" {
		return nil
	}
	b := binding{
		Type:    base,
		IsArray: shape.IsArray,
		IsConst: isConst(ctx.DeclSpecifierSeq()),
	}
	if fileScope {
		if a.FileVars == nil {
			a.FileVars = map[string]binding{}
		}
		a.FileVars[name] = b
	} else {
		a.Scopes.declare(name, b)
	}
	return nil
}
