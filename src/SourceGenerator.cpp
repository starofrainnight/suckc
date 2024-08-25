#include "SourceGenerator.h"
#include "SourceContext.h"
#include <string>

namespace suckc {

class SourceGeneratorPrivate {
  SUCKC_OBJECT_IMPL(SourceGenerator)

public:
  SourceContext ctx;

  inline std::string getIndentText() const {
    return std::string(indenttion * 2, ' ');
  }

  std::string getNodeRuleName(antlr4::tree::ParseTree *node);

  int indenttion = -1;
  SuckCParser *parser;
};

std::string
SourceGeneratorPrivate::getNodeRuleName(antlr4::tree::ParseTree *node) {
  antlr4::ParserRuleContext *context =
      dynamic_cast<antlr4::ParserRuleContext *>(node);
  std::string name;

  if (context != nullptr) {
    // Get the rule index of this node
    int ruleIndex = context->getRuleIndex();

    // Get the parser from which you can retrieve rule names
    if (parser != nullptr) {
      // Retrieve the rule name using the rule index
      std::vector<std::string> ruleNames = parser->getRuleNames();
      std::string ruleName = ruleNames[ruleIndex];

      // Now you have the rule name, you can do something with it
      return ruleName;
    }
  }

  return name;
}

SourceGenerator::SourceGenerator(SuckCParser *parser)
    : dPtr_(new SourceGeneratorPrivate(this)) {
  SUCKC_D();
  d->parser = parser;
}

SourceGenerator::~SourceGenerator() {}
std::any SourceGenerator::visitChildren(antlr4::tree::ParseTree *node) {
  SUCKC_D();

  ++d->indenttion;

#ifdef SUCKC_TRACE_NODE_TREE
  std::cout << d->getIndentText() << d->getNodeRuleName(node) << "\n";
#endif

  auto ret = SuckCBaseVisitor::visitChildren(node);

  --d->indenttion;

  return ret;
}

std::any SourceGenerator::visitFunctionDefinition(
    SuckCParser::FunctionDefinitionContext *ctx) {
  SUCKC_D();

  d->ctx.beginScope(ScopeType::FunctionScope);
  auto ret = visitChildren(ctx);
  d->ctx.endScope();
  return ret;
}

std::any SourceGenerator::visitExternalDeclaration(
    SuckCParser::ExternalDeclarationContext *ctx) {
  SUCKC_D();
  return visitChildren(ctx);
}

} // namespace suckc
