#include "SourceGenerator.h"
#include "ParserTreeHelper.h"
#include "SourceContext.h"
#include "World.h"
#include "ast/Expression.h"
#include "ast/IdExpression.h"
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
    if (World::getInstance()->isEnabledDebug()) {
      std::cout << getIndentText() << ParserTreeHelper::getNodeRuleName(node)
                << ": " << ParserTreeHelper::getNodeStartTokenText(node)
                << "\n";
    }
  }

  void traceNodeUnderNextIndent(antlr4::tree::ParseTree *node) {
    ++indent;
    traceNode(node);
    --indent;
  }

  void traceNodeSourceUnderNextIndent(antlr4::tree::ParseTree *node) {
    if (World::getInstance()->isEnabledDebug()) {
      ++indent;
      std::cout << getIndentText() << ParserTreeHelper::getNodeRuleName(node)
                << ": " << ParserTreeHelper::getNodeSource(node) << "\n";
      --indent;
    }
  }

  int indent = -1;
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

  ++d->indent;

  d->traceNode(node);

  auto ret = SuckCParserBaseVisitor::visitChildren(node);

  --d->indent;

  return ret;
}

std::any SourceGenerator::visitTranslationUnit(
    SuckCParser::TranslationUnitContext *ctx) {
  SUCKC_D();

  visitChildren(ctx);

  if (World::getInstance()->isEnabledDebug()) {
    std::cout << std::endl << "Final Context:" << std::endl;
    d->ctx.print();
  }
  return std::any();
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

std::any SourceGenerator::visitStatement(SuckCParser::StatementContext *ctx) {
  SUCKC_D();
  d->ctx.beginScope(ScopeType::StatementScope);
  auto ret = visitChildren(ctx);
  d->ctx.endScope();
  return ret;
}

std::any SourceGenerator::visitSimpleDeclaration(
    SuckCParser::SimpleDeclarationContext *ctx) {
  SUCKC_D();

  auto variable = std::make_shared<suckc::ast::Variable>();
  variable->setRuleContext(ctx);
  // variable->setValueType();

  return visitChildren(ctx);
}

std::any SourceGenerator::visitLiteral(SuckCParser::LiteralContext *ctx) {
  SUCKC_D();

  auto literal = std::make_shared<suckc::ast::Literal>();
  literal->setRuleContext(ctx);

  d->traceNodeUnderNextIndent(ctx);

  return std::any(literal);
}

std::any SourceGenerator::visitPrimaryExpression(
    SuckCParser::PrimaryExpressionContext *ctx) {
  SUCKC_D();
  auto expr = std::make_shared<suckc::ast::Expression>();
  expr->setRuleContext(ctx);

  d->traceNodeUnderNextIndent(ctx);

  return std::any(expr);
}

std::any SourceGenerator::visitAssignmentExpression(
    SuckCParser::AssignmentExpressionContext *ctx) {
  SUCKC_D();
  auto expr = std::make_shared<suckc::ast::Expression>();
  expr->setRuleContext(ctx);

  d->traceNodeSourceUnderNextIndent(ctx);

  return std::any(expr);
}

std::any SourceGenerator::visitDeclSpecifierSeq(
    SuckCParser::DeclSpecifierSeqContext *ctx) {

  return visitChildren(ctx);
}

std::any
SourceGenerator::visitDeclSpecifier(SuckCParser::DeclSpecifierContext *ctx) {
  SUCKC_D();
  auto expr = std::make_shared<suckc::ast::Expression>();
  expr->setRuleContext(ctx);

  d->traceNodeUnderNextIndent(ctx);

  return std::any(expr);
}

std::any
SourceGenerator::visitInitDeclarator(SuckCParser::InitDeclaratorContext *ctx) {
  SUCKC_D();
  auto expr = std::make_shared<suckc::ast::Expression>();
  expr->setRuleContext(ctx);

  d->traceNodeSourceUnderNextIndent(ctx);

  return std::any(expr);
}

std::any
SourceGenerator::visitInitializer(SuckCParser::InitializerContext *ctx) {
  SUCKC_D();
  auto expr = std::make_shared<suckc::ast::Expression>();
  expr->setRuleContext(ctx);

  d->traceNodeSourceUnderNextIndent(ctx);

  return std::any(expr);
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
