package sema

import (
	"strings"
	"testing"

	"github.com/antlr4-go/antlr/v4"

	"github.com/starofrainnight/suckc/internal/frontend"
	"github.com/starofrainnight/suckc/internal/parser"
)

// deduceExpr parses `auto probe = <expr>;` inside a function body and
// returns typeOf applied to the initializer expression.
func deduceExpr(t *testing.T, setup func(*Analyzer), expr string) (Type, error) {
	t.Helper()
	src := "void f() {\n\tauto probe = " + expr + ";\n}\n"
	res, err := frontend.ParseSource(src)
	if err != nil {
		t.Fatalf("parse %q: %v", src, err)
	}
	a := newAnalyzer(Options{TargetBits: 32, FileName: "t.suckc"})
	if setup != nil {
		setup(a)
	}
	return typeOf(findProbeInit(t, res.Tree), a)
}

func findProbeInit(t *testing.T, tree antlr.Tree) antlr.Tree {
	t.Helper()
	var found antlr.Tree
	var walk func(n antlr.Tree)
	walk = func(n antlr.Tree) {
		if found != nil {
			return
		}
		if ctx, ok := n.(*parser.SimpleDeclarationContext); ok {
			if id := ctx.InitDeclarator(); id != nil {
				if d := id.Declarator(); d != nil {
					if name, _ := declaratorName(d); name == "probe" {
						if init := id.Initializer(); init != nil {
							if boi := init.BraceOrEqualInitializer(); boi != nil {
								if ic := boi.InitializerClause(); ic != nil {
									if ae := ic.AssignmentExpression(); ae != nil {
										found = ae
									}
								}
							}
						}
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
	walk(tree)
	if found == nil {
		t.Fatal("probe initializer not found in parse tree")
	}
	return found
}

func setupVars(a *Analyzer) {
	a.Scopes.push()
	a.Scopes.declare("i", binding{Type{"int", 0}, false})
	a.Scopes.declare("u", binding{Type{"unsigned int", 0}, false})
	a.Scopes.declare("d", binding{Type{"double", 0}, false})
	a.Scopes.declare("p", binding{Type{"int", 1}, false})
	a.Scopes.declare("arr", binding{Type{"int", 0}, true})
	a.Scopes.declare("s", binding{Type{"const char", 1}, false})
	a.Funcs["geti"] = Type{"int", 0}
	a.Funcs["getp"] = Type{"int", 1}
}

func TestTypeOfExpressions(t *testing.T) {
	cases := []struct {
		expr string
		want Type
	}{
		// 档1
		{"12", Type{"int", 0}},
		{"12 + 5", Type{"int", 0}},
		{"i + 1L", Type{"long", 0}},
		{"i + d", Type{"double", 0}},
		{"u + i", Type{"unsigned int", 0}},
		{"i & u", Type{"unsigned int", 0}},
		{"&i", Type{"int", 1}},
		{"*p", Type{"int", 0}},
		{"(i)", Type{"int", 0}},
		{"-i", Type{"int", 0}},
		{"geti()", Type{"int", 0}},
		{"geti() + 1", Type{"int", 0}},
		{"getp()", Type{"int", 1}},
		{"*getp()", Type{"int", 0}},
		// 档2
		{"i < d", Type{"int", 0}},
		{"i == u", Type{"int", 0}},
		{"i && d", Type{"int", 0}},
		{"i ? 1 : 2", Type{"int", 0}},
		{"i ? 1 : d", Type{"double", 0}},
		{"arr", Type{"int", 1}},   // array decay
		{"arr[0]", Type{"int", 0}}, // subscript on decayed array
		{"p[1]", Type{"int", 0}},   // subscript on pointer
		{"s", Type{"const char", 1}},
		{"\"abc\"", Type{"const char", 1}},
		{"'c'", Type{"char", 0}},
		{"sizeof(i)", Type{"size_t", 0}},
		{"!i", Type{"int", 0}},
		{"~i", Type{"int", 0}},
	}
	for _, c := range cases {
		got, err := deduceExpr(t, setupVars, c.expr)
		if err != nil {
			t.Errorf("typeOf(%q): %v", c.expr, err)
			continue
		}
		if got != c.want {
			t.Errorf("typeOf(%q) = %+v, want %+v", c.expr, got, c.want)
		}
	}
}

func TestTypeOfSizeofOn64(t *testing.T) {
	src := "void f() {\n\tauto probe = sizeof(i);\n}\n"
	res, err := frontend.ParseSource(src)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	a := newAnalyzer(Options{TargetBits: 64, FileName: "t.suckc"})
	setupVars(a)
	got, err := typeOf(findProbeInit(t, res.Tree), a)
	if err != nil {
		t.Fatalf("typeOf: %v", err)
	}
	if got != (Type{"size_t", 0}) {
		t.Errorf("sizeof on 64-bit = %+v, want size_t", got)
	}
}

func TestTypeOfErrors(t *testing.T) {
	cases := []struct {
		expr string
		want string
	}{
		{"nope", "unknown identifier 'nope'"},      // incl. NULL-style unknowns
		{"getu()", "unknown identifier 'getu'"},    // undeclared function call
		{"1 << 2", "unsupported expression"},       // shift: not in spec 6.2 table
		{"*i", "insufficient pointer depth"},       // deref non-pointer
		{"arr[0][0]", "insufficient pointer depth"}, // second subscript level (spec 5: no shape)
		{"i++", "unsupported expression"},          // ++ deferred (spec 11)
		{"i = 3", "unsupported expression"},        // assignment as expression
		{"p + 1", "incompatible operand types"},    // pointer arithmetic not in table
		{"i . x", "unsupported expression"}, // member access deferred (spec 11)
	}
	for _, c := range cases {
		_, err := deduceExpr(t, setupVars, c.expr)
		if err == nil {
			t.Errorf("typeOf(%q): expected error containing %q", c.expr, c.want)
			continue
		}
		if !strings.Contains(err.Error(), c.want) {
			t.Errorf("typeOf(%q) error = %q, want contains %q", c.expr, err, c.want)
		}
	}
}
