package sema

import (
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/starofrainnight/suckc/internal/frontend"
	"github.com/starofrainnight/suckc/internal/parser"
)

// Argument type checking against C-region prototypes.
//
// C converts freely between arithmetic types, so demanding that the argument
// and the parameter canonicalise to the same type would reject almost every
// real call. What C does not convert is the shape: a pointer does not become an
// arithmetic type, an array does not become a scalar, and two struct or enum
// tags are never the same type. Those are the mismatches checked here.
//
// Everything is permissive on purpose. An argument whose type cannot be
// determined is skipped rather than guessed at, and two pointers of different
// pointee types are accepted even though C constrains them, because the type
// representation in this package canonicalises by spelling and over-strict
// comparison would produce false errors.

// checkCallAgainstC verifies one call site's argument shapes. args holds the
// argument subtrees in order; an argument whose type cannot be determined is
// skipped.
func (w *walker) checkCallAgainstC(name string, args []antlr.Tree, tokIdx int) {
	sig, ok := w.cFuncs[name]
	if !ok || sig.Unprototyped {
		return
	}
	// A call inside a C region is C code calling a C function; the region body
	// is opaque to this pass, so nothing inside it can reach here anyway.
	if w.res.InRegion(tokIdx) {
		return
	}
	for i, arg := range args {
		if i >= len(sig.Params) {
			// Arity is checkCCalls' job, and a variadic tail has no declared
			// parameter to check against.
			break
		}
		have, ok := w.argType(arg)
		if !ok {
			continue
		}
		base, stars := splitTrailingStars(sig.Params[i].TypeSpelling)
		want := w.resolver.resolve(base)
		want.Stars += stars
		if cArgCompatible(have, want, arg) {
			continue
		}
		where := ""
		if file, line := sourceAtToken(w.res, w.a, sig.TokIdx); file != "" {
			where = " (declared at " + file + ":" + strconv.Itoa(line) + ")"
		}
		w.fail(&CallError{
			File: w.a.Opts.FileName,
			Line: w.res.Tokens[tokIdx].GetLine(),
			Col:  w.res.Tokens[tokIdx].GetColumn() + 1,
			Name: name,
			Msg: "incompatible argument " + strconv.Itoa(i+1) + " in call to " +
				quoteName(name) + ": passing " + quoteName(have.display()) +
				" where " + quoteName(want.display()) + " is expected" + where,
		})
		return
	}
}

// cArgCompatible reports whether an argument of type arg may be passed where
// param is expected. argText is consulted only to recognise the null pointer
// constant, which is the one integer C converts to a pointer implicitly.
func cArgCompatible(arg, param canonType, argText antlr.Tree) bool {
	// C89 6.3.2.3: an integer constant with value 0 is a null pointer constant
	// and converts to any pointer type.
	if arg.Stars == 0 && isNullPointerConstant(argText) {
		return true
	}
	// void converts to void * and to no other object type.
	if arg.Stars == 0 && arg.Base == "void" {
		return param.Stars == 1 && param.Base == "void"
	}
	if arg.Stars != param.Stars {
		return false
	}
	if arg.Stars > 0 {
		// Compatible pointee types are a constraint C does enforce, but this
		// package canonicalises by spelling and a wrong "different" verdict
		// would reject valid calls, so the pointee is not compared.
		return true
	}
	// Both scalars. Arithmetic types interconvert freely; struct, union and
	// enum tags name distinct types that never match across different tags.
	if tag, ok := aggregateTag(arg.Base); ok {
		other, ok := aggregateTag(param.Base)
		return ok && tag == other
	}
	return true
}

// aggregateTag returns the leading "struct"/"union"/"enum" tag of a
// non-pointer type spelling, if it has one.
func aggregateTag(base string) (string, bool) {
	for _, kw := range []string{"struct ", "union ", "enum "} {
		if strings.HasPrefix(base, kw) {
			return base, true
		}
	}
	return "", false
}

func isZeroLiteral(text string) bool {
	body := strings.TrimRight(text, "uUlL")
	body = strings.ReplaceAll(body, "'", "")
	if strings.HasPrefix(body, "0x") || strings.HasPrefix(body, "0X") {
		body = body[2:]
	} else if strings.HasPrefix(body, "0b") || strings.HasPrefix(body, "0B") {
		body = body[2:]
	}
	if body == "" {
		return false
	}
	return strings.Trim(body, "0") == ""
}

func isNullPointerConstant(t antlr.Tree) bool {
	// Find the first token rather than insisting the node itself be a terminal:
	// an argument arrives wrapped in an initializerClause and half a dozen
	// expression rules, and the exact chain depth differs between a plain "0",
	// a suffixed "0u" and a hexadecimal "0x0". Stopping at the first token keeps
	// this correct for all three; a compound expression like "1 + 0" still lands
	// on the leading 1 and is rejected, which is the safe direction.
	var term antlr.TerminalNode
	var walk func(n antlr.Tree)
	walk = func(n antlr.Tree) {
		if term != nil || n == nil {
			return
		}
		if x, ok := n.(antlr.TerminalNode); ok && x != nil {
			term = x
			return
		}
		for i := 0; i < n.GetChildCount(); i++ {
			walk(n.GetChild(i))
		}
	}
	walk(t)
	if term == nil {
		return false
	}
	switch term.GetSymbol().GetTokenType() {
	case parser.SuckCParserDecimalLiteral, parser.SuckCParserHexadecimalLiteral,
		parser.SuckCParserOctalLiteral, parser.SuckCParserBinaryLiteral,
		parser.SuckCParserUserDefinedLiteral:
		// UserDefinedLiteral is still needed for suffixed integers. The lexer
		// once classified every numeric literal as it, because that rule was
		// declared first and tied with the numeric rules on length; moving it
		// after them fixed 0x0 and 0b1. It cannot fix "0u" or "12ll" that way:
		// the plain numeric rule stops before the suffix, so the user-defined
		// reading is strictly longer and longest-match wins regardless of
		// order. Type deduction never notices, because deduceLiteralText works
		// off the token text rather than the token type. isZeroLiteral still
		// rejects genuine user-defined literals such as "0_km".
		return isZeroLiteral(term.GetText())
	}
	return false
}

// argType determines the type of one argument expression, reporting false when
// the expression is beyond what this pass can type. Only literals, names, their
// parenthesised forms and unary &/* are handled; anything else is left alone.
func (w *walker) argType(t antlr.Tree) (canonType, bool) {
	// SuckC's expression grammar buries the interesting node under a chain of
	// one-child wrappers (initializerClause > assignmentExpression >
	// conditionalExpression > ... > idExpression). Unwrap before deciding,
	// otherwise every real argument looks unrecognised and is skipped.
	for t != nil && t.GetChildCount() == 1 {
		t = t.GetChild(0)
	}
	if t == nil {
		return canonType{}, false
	}
	// Check for a terminal before the child count: a terminal node has zero
	// children of its own, so a zero-child guard placed first would reject
	// every literal and every bare identifier, making all arguments invisible.
	if term, ok := t.(antlr.TerminalNode); ok && term != nil {
		// A bare identifier is a variable reference; anything else is a literal.
		// literalType cannot tell them apart -- deduceLiteralText rejects "s".
		if term.GetSymbol().GetTokenType() == parser.SuckCLexerIdentifier {
			return w.nameTypeFrom(term.GetText(), false)
		}
		return w.literalType(term.GetText())
	}
	if t.GetChildCount() == 0 {
		return canonType{}, false
	}
	// Only a genuine parenthesis wrapper may be entered at its middle child.
	// Any three-child node has a middle child, so the parens must be checked.
	if kids := t.GetChildren(); len(kids) == 3 &&
		tokenType(kids[0]) == parser.SuckCLexerLeftParen &&
		tokenType(kids[2]) == parser.SuckCLexerRightParen {
		return w.argType(kids[1])
	}
	if name, tokIdx := calleeName(t); name != "" {
		return w.nameType(name, tokIdx, 0)
	}
	return w.unaryArgType(t)
}

// literalType types a literal token with SuckC's dialect rules. The result is
// the type the argument has in SuckC, which is what the emitted C will declare
// for it.
func (w *walker) literalType(text string) (canonType, bool) {
	t, err := deduceLiteralText(text, w.a.Opts)
	if err != nil {
		return canonType{}, false
	}
	out := w.resolver.resolve(t.Spelling)
	out.Stars += t.Stars
	return out, true
}

// nameType resolves a name to a type, adding stars levels. A binding declared
// as an array decays to a pointer, which is what C does when passing it.
func (w *walker) nameType(name string, tokIdx, stars int) (canonType, bool) {
	b, ok := w.a.Scopes.lookup(name)
	if !ok {
		b, ok = w.a.FileVars[name]
	}
	if !ok {
		return canonType{}, false
	}
	if b.IsArray {
		stars++
	}
	out := w.resolver.resolve(b.Type.Spelling)
	out.Stars += b.Type.Stars + stars
	return out, true
}

// unaryArgType handles `&x` and `*x`, which change the type's shape rather than
// its spelling.
func (w *walker) unaryArgType(t antlr.Tree) (canonType, bool) {
	kids := t.GetChildren()
	if len(kids) != 2 {
		return canonType{}, false
	}
	op := tokenType(kids[0])
	if op != parser.SuckCParserAnd && op != parser.SuckCParserStar {
		return canonType{}, false
	}
	// Find the operand's name without going through calleeName, which reports a
	// call rather than a bare name.
	var walk func(n antlr.Tree) string
	walk = func(n antlr.Tree) string {
		if term, ok := n.(antlr.TerminalNode); ok && term != nil &&
			term.GetSymbol().GetTokenType() == parser.SuckCLexerIdentifier {
			return term.GetText()
		}
		for i := 0; i < n.GetChildCount(); i++ {
			if name := walk(n.GetChild(i)); name != "" {
				return name
			}
		}
		return ""
	}
	name := walk(kids[1])
	if name == "" {
		return canonType{}, false
	}
	// The operator token index is not available here, so the name is resolved
	// without a diagnostic anchor; a name that fails to resolve just skips.
	return w.nameTypeFrom(name, op == parser.SuckCParserAnd)
}

func (w *walker) nameTypeFrom(name string, addStar bool) (canonType, bool) {
	out, ok := w.nameType(name, 0, 0)
	if !ok {
		return canonType{}, false
	}
	if addStar {
		out.Stars++
	}
	return out, true
}

// exprCallArgs returns the callee name and the argument subtrees of a call
// spelled as an expression. The shape rules are the ones callOf applies; this
// is callOf with the arguments kept instead of counted.
func exprCallArgs(pe *parser.PostfixExpressionContext) (string, []antlr.Tree, int, bool) {
	kids := pe.GetChildren()
	n := len(kids)
	if n < 3 || tokenType(kids[n-1]) != parser.SuckCParserRightParen {
		return "", nil, 0, false
	}
	var args []antlr.Tree
	switch head := kids[n-2].(type) {
	case *parser.ExpressionListContext:
		if head == nil || head.InitializerList() == nil {
			return "", nil, 0, false
		}
		args = initializerClauses(head.InitializerList())
	case *antlr.TerminalNodeImpl:
		if head.GetSymbol().GetTokenType() != parser.SuckCParserLeftParen {
			return "", nil, 0, false
		}
	default:
		return "", nil, 0, false
	}
	name, tokIdx := calleeName(kids[0])
	if name == "" {
		return "", nil, 0, false
	}
	return name, args, tokIdx, true
}

// declCallArgs returns the callee name and argument subtrees of the
// statement-level call `f(a, b);`, which this grammar parses as a declaration
// with no type specifier. The conditions are declarationShapedCall's.
func (w *walker) declCallArgs(sd *parser.SimpleDeclarationContext) (string, []antlr.Tree, int, bool) {
	init := sd.InitDeclarator()
	if init == nil || init.Declarator() == nil {
		return "", nil, 0, false
	}
	if name, tokIdx, ok := typedefNameSpelling(sd.DeclSpecifierSeq()); ok {
		if _, isCFunc := w.cFuncs[name]; isCFunc {
			return name, []antlr.Tree{init.Declarator()}, tokIdx, true
		}
	}
	if hasTypeSpecifier(sd.DeclSpecifierSeq()) {
		return "", nil, 0, false
	}
	name, tokIdx := callTargetName(init.Declarator())
	if name == "" {
		return "", nil, 0, false
	}
	if init.Initializer() == nil {
		if !hasParameterList(init.Declarator()) {
			return "", nil, 0, false
		}
		return name, nil, tokIdx, true
	}
	kids := init.Initializer().GetChildren()
	if len(kids) != 3 {
		return "", nil, 0, false
	}
	list, ok := kids[1].(*parser.ExpressionListContext)
	if !ok || list == nil || list.InitializerList() == nil {
		return "", nil, 0, false
	}
	return name, initializerClauses(list.InitializerList()), tokIdx, true
}

func initializerClauses(il parser.IInitializerListContext) []antlr.Tree {
	all := il.AllInitializerClause()
	out := make([]antlr.Tree, 0, len(all))
	for _, c := range all {
		out = append(out, c)
	}
	return out
}

// cFuncLookup keys the region signatures for argument checking. It applies the
// same duplicate-declaration merge as indexCFuncs but reports no conflict: the
// arity pass already reports that, and reporting it twice would be noise.
//
// It is deliberately a separate function rather than a shared one with
// indexCFuncs so that the arity check, which has tests of its own, keeps its
// current behaviour.
func cFuncLookup(res *frontend.ParseResult) map[string]frontend.CFuncSig {
	index := make(map[string]frontend.CFuncSig, len(res.CFuncs))
	for _, sig := range res.CFuncs {
		if prev, ok := index[sig.Name]; ok {
			if !sigConflict(prev, sig) && (prev.Unprototyped || sig.Unprototyped) {
				prev.Unprototyped = true
				index[sig.Name] = prev
			}
			continue
		}
		index[sig.Name] = sig
	}
	return index
}
