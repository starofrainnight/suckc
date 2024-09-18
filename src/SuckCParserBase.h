#pragma once

#include <antlr4-runtime.h>

class SuckCParserBase : public antlr4::Parser {
public:
  SuckCParserBase(antlr4::TokenStream *input) : Parser(input) {}
  bool IsPureSpecifierAllowed();
};
