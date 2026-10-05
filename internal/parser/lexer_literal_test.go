package parser

import (
	"testing"

	"github.com/antlr4-go/antlr/v4"
)

// SuckCLexer.g4 declares UserDefinedLiteral (C++11 user-defined literals)
// after the plain numeric literal rules. That order is load-bearing: 0x0 and
// 0b1 match both readings at the same length, so the tie is broken by
// declaration order, and a user-defined literal declared first would swallow
// every hexadecimal and binary literal in the language. Nothing else notices,
// because type deduction works off token text rather than token type.
func TestHexAndBinaryLiteralsAreNotUserDefinedLiterals(t *testing.T) {
	tests := []struct {
		src  string
		want int
	}{
		{"0x0", SuckCLexerHexadecimalLiteral},
		{"0xFF", SuckCLexerHexadecimalLiteral},
		{"0b1010", SuckCLexerBinaryLiteral},
		{"0b0", SuckCLexerBinaryLiteral},
		// A genuine user-defined literal still lexes as one: the plain numeric
		// rule stops before the suffix, so that reading is strictly longer.
		{"1_km", SuckCLexerUserDefinedLiteral},
		{"1.5_km", SuckCLexerUserDefinedLiteral},
		// Known residual: a suffixed integer such as 0u or 12ll is still
		// classified as a user-defined literal, for the same longest-match
		// reason. Nothing depends on this today.
		{"0u", SuckCLexerUserDefinedLiteral},
		{"12ll", SuckCLexerUserDefinedLiteral},
	}
	for _, tt := range tests {
		t.Run(tt.src, func(t *testing.T) {
			lexer := NewSuckCLexer(antlr.NewInputStream(tt.src))
			lexer.RemoveErrorListeners()
			stream := antlr.NewCommonTokenStream(lexer, 0)
			stream.Fill()
			tokens := stream.GetAllTokens()
			if len(tokens) != 2 { // the literal plus EOF
				t.Fatalf("%q produced %d tokens, want 1", tt.src, len(tokens))
			}
			if got := tokens[0].GetTokenType(); got != tt.want {
				t.Errorf("%q tokenized as %s, want %s", tt.src,
					lexer.SymbolicNames[got], lexer.SymbolicNames[tt.want])
			}
		})
	}
}
