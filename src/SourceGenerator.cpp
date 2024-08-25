#include "SourceGenerator.h"
#include "SourceContext.h"
#include <string>

namespace suckc {

class SourceGeneratorPrivate {
  SUCKC_OBJECT_IMPL(SourceGenerator)

public:
  SourceContext ctx;

  inline std::string getIndentText() const {
    return std::string(indent * 2, ' ');
  }

  std::vector<std::string> collectTokenText(antlr4::tree::ParseTree *node);
  std::vector<std::string> collectTokenText(antlr4::tree::TerminalNode *node);

  std::string getNodeRuleName(antlr4::tree::ParseTree *node);
  std::string getNodeStartTokenText(antlr4::tree::ParseTree *node);
  std::string getNodeTokensText(antlr4::tree::ParseTree *node);

  int indent = -1;
  SuckCParser *parser;
};

std::vector<std::string>
SourceGeneratorPrivate::collectTokenText(antlr4::tree::TerminalNode *node) {
  // This function collects the text of a terminal node.
  return {node->getText()};
}

std::vector<std::string>
SourceGeneratorPrivate::collectTokenText(antlr4::tree::ParseTree *node) {
  // This function collects the text of all tokens under a given non-terminal node.
  std::vector<std::string> tokens;

  for (auto child : node->children) {
    if (auto terminalNode = dynamic_cast<antlr4::tree::TerminalNode *>(child)) {
      auto childTokens = collectTokenText(terminalNode);

      tokens.insert(tokens.end(), childTokens.begin(), childTokens.end());
    } else if (auto ruleNode = dynamic_cast<antlr4::RuleContext *>(child)) {
      auto childTokens = collectTokenText(ruleNode);
      tokens.insert(tokens.end(), childTokens.begin(), childTokens.end());
    }
  }

  return tokens;
}

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
std::string
SourceGeneratorPrivate::getNodeStartTokenText(antlr4::tree::ParseTree *node) {
  if (node != nullptr) {
    antlr4::ParserRuleContext *ctx =
        dynamic_cast<antlr4::ParserRuleContext *>(node);
    if (ctx != nullptr) {
      // Get the start token of the context
      auto startToken = ctx->start;

      // Get the text associated with the start token
      auto nodeText = startToken->getText();

      return nodeText;
    }
  }

  return "";
}

std::string
SourceGeneratorPrivate::getNodeTokensText(antlr4::tree::ParseTree *node) {
  auto texts = collectTokenText(node);

  std::string text;
  for (auto i = 0; i < texts.size(); ++i) {
    text.append(texts[i]);
    text.append(" ");
  }

  return text;
}
SourceGenerator::SourceGenerator(SuckCParser *parser)
    : dPtr_(new SourceGeneratorPrivate(this)) {
  SUCKC_D();
  d->parser = parser;
}

SourceGenerator::~SourceGenerator() {}
std::any SourceGenerator::visitChildren(antlr4::tree::ParseTree *node) {
  SUCKC_D();

  ++d->indent;

#ifdef SUCKC_TRACE_NODE_TREE
  std::cout << d->getIndentText() << d->getNodeRuleName(node) << ": "
            << d->getNodeStartTokenText(node) << "\n";
#endif

  auto ret = SuckCBaseVisitor::visitChildren(node);

  --d->indent;

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
