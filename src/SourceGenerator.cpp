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

  int indenttion = -1;
  SuckCParser *parser;
};

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
  // Cast the node to a ParserRuleContext to access the rule context
  antlr4::ParserRuleContext *context =
      dynamic_cast<antlr4::ParserRuleContext *>(node);

  if (context != nullptr) {
    // Get the rule index of this node
    int ruleIndex = context->getRuleIndex();

    // Get the parser from which you can retrieve rule names
    auto parser = d->parser;

    if (parser != nullptr) {
      // Retrieve the rule name using the rule index
      std::vector<std::string> ruleNames = parser->getRuleNames();
      std::string ruleName = ruleNames[ruleIndex];

      // Now you have the rule name, you can do something with it
      std::cout << d->getIndentText() << ruleName << "\n";
    }
  }

#endif

  std::any ret = SuckCBaseVisitor::visitChildren(node);
#ifdef SUCKC_TRACE_NODE_TREE

#endif
  --d->indenttion;

  return ret;
}

std::any SourceGenerator::visitFunctionDefinition(
    SuckCParser::FunctionDefinitionContext *ctx) {
  SUCKC_D();

  d->ctx.beginScope(ScopeType::FunctionScope);

  // std::cout << "got function: " << ctx->getText() << std::endl;
  auto ret = visitChildren(ctx);
  d->ctx.endScope();
  return ret;
}

std::any SourceGenerator::visitExternalDeclaration(
    SuckCParser::ExternalDeclarationContext *ctx) {
  SUCKC_D();
  // std::cout << "Extern Declaration: " << std::hex
  //           << (uintptr_t)ctx->functionDefinition() << std::endl;
  return visitChildren(ctx);
}

} // namespace suckc
