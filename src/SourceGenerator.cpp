#include "SourceGenerator.h"
#include "ParserTreeHelper.h"
#include "SourceContext.h"
#include "ast/Literal.h"
#include <string>

namespace suckc {

class SourceGeneratorPrivate {
  SUCKC_OBJECT_IMPL(SourceGenerator)

public:
  SourceContext ctx;

  inline std::string getIndentText() const {
    return std::string(indent * 2, ' ');
  }

  void traceNode(antlr4::tree::ParseTree *node) {
    if (isEnabledDebug) {
      std::cout << getIndentText() << ParserTreeHelper::getNodeRuleName(node)
                << ": " << ParserTreeHelper::getNodeStartTokenText(node)
                << "\n";
    }
  }

  int indent = -1;
  SuckCParser *parser;
  bool isEnabledDebug;
};

SourceGenerator::SourceGenerator(SuckCParser *parser, bool isEnabledDebug)
    : dPtr_(new SourceGeneratorPrivate(this)) {
  SUCKC_D();
  d->parser = parser;
  d->isEnabledDebug = isEnabledDebug;
}

SourceGenerator::~SourceGenerator() {}
std::any SourceGenerator::visitChildren(antlr4::tree::ParseTree *node) {
  SUCKC_D();

  ++d->indent;

  d->traceNode(node);

  auto ret = SuckCParserBaseVisitor::visitChildren(node);

  --d->indent;

  return ret;
}

std::any SourceGenerator::visitFunctionDefinition(
    SuckCParser::FunctionDefinitionContext *ctx) {
  SUCKC_D();

  auto scope = d->ctx.getCurrentScope();
  auto func = std::make_shared<suckc::ast::Function>();
  auto name = ParserTreeHelper::getNodeStartTokenText(ctx->declarator());
  func->setName(name);
  (*scope)->addNode(name, func);

  d->ctx.beginScope(ScopeType::FunctionScope);
  auto ret = visitChildren(ctx);
  d->ctx.endScope();
  return ret;
}

// std::any
// SourceGenerator::visitItemDefinition(SuckCParser::ItemDefinitionContext *ctx) {
//   SUCKC_D();

//   return visitChildren(ctx);
// }

std::any SourceGenerator::visitCompoundStatement(
    SuckCParser::CompoundStatementContext *ctx) {
  SUCKC_D();
  d->ctx.beginScope(ScopeType::BlockScope);
  auto ret = visitChildren(ctx);
  d->ctx.endScope();
  return ret;
}

std::any SourceGenerator::visitSimpleDeclaration(
    SuckCParser::SimpleDeclarationContext *ctx) {
  return visitChildren(ctx);
}

std::any SourceGenerator::visitLiteral(SuckCParser::LiteralContext *ctx) {
  SUCKC_D();

  auto literal = std::make_shared<suckc::ast::Literal>();
  literal->setRuleContext(ctx);

  ++d->indent;

  d->traceNode(ctx);

  --d->indent;

  return std::any(literal);
}

// std::any SourceGenerator::visitBlockItem(SuckCParser::BlockItemContext *ctx) {
//   SUCKC_D();
//   d->ctx.beginScope(ScopeType::StatementScope);
//   auto ret = visitChildren(ctx);
//   d->ctx.endScope();
//   return ret;
// }

// std::any
// SourceGenerator::visitDeclaration(SuckCParser::DeclarationContext *ctx) {
//   SUCKC_D();
//   auto scope = d->ctx.getCurrentScope();
//   if (ctx->initDeclaratorList()) {
//     auto declList = ctx->initDeclaratorList();
//     for (auto decl : declList->initDeclarator()) {
//       auto var = std::make_shared<suckc::ast::Variable>();
//       auto name = d->getNodeStartTokenText(decl);
//       var->setName(name);
//       (*scope)->addNode(name, var);
//     }
//   } else {
//     auto lastSpecifier = ctx->declarationSpecifiers()->children.back();
//     auto var = std::make_shared<suckc::ast::Variable>();
//     auto name = d->getNodeStartTokenText(lastSpecifier);
//     (*scope)->addNode(name, var);
//   }

//   return visitChildren(ctx);
// }

} // namespace suckc
