// Package parser holds the ANTLR4 grammar for the SuckC language and the
// Go code generated from it.
//
// Regenerate the parser after editing the grammar files:
//
//	go generate ./internal/parser
//
// The generated *.go files are committed so that the module builds without
// requiring the ANTLR toolchain.
package parser

//go:generate antlr4 -Dlanguage=Go -visitor -no-listener -o . SuckCLexer.g4 SuckCParser.g4
