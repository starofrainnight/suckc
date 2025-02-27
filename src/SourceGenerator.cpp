#include "SourceGenerator.h"
#include "ParserTreeHelper.h"
#include "SourceContext.h"
#include "World.h"
#include "ast/Alias.h"
#include "ast/Expression.h"
#include "ast/FunctionVarBody.h"
#include "ast/IdExpression.h"
#include "ast/Literal.h"
#include "ast/TypeDeclaration.h"
#include <string>
#include <typeinfo>

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

std::any SourceGenerator::aggregateResult(std::any aggregate,
                                          std::any nextResult) {
  // If the next result has value, then return the next result, otherwise return previous result
  return nextResult.has_value() ? nextResult : aggregate;
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
  auto func = std::make_shared<ast::Function>();
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

  if (ctx->initDeclarator()) {
    auto variable = std::make_shared<ast::Variable>();

    variable->setRuleContext(ctx);

    // Get specifiers and attributes.
    auto specifiers = variable->getSpecifiers();
    {
      auto seq = ctx->declSpecifierSeq();
      if (seq) {
        for (auto specifierContext : seq->declSpecifier()) {
          auto specifier = std::make_shared<ast::Specifier>();
          specifier->setRuleContext(specifierContext);
          specifiers->push_back(specifier);
        }
      }
    }

    auto attributes = variable->getAttributes();
    {
      auto seq = ctx->attributeSpecifierSeq();
      if (seq) {
        for (auto attributeContext : seq->attributeSpecifier()) {
          auto attribute = std::make_shared<ast::Attribute>();
          attribute->setRuleContext(attributeContext);
          attributes->push_back(attribute);
        }
      }
    }

    auto value = visitChildren(ctx->initDeclarator()->declarator());
    if (value.type() == typeid(std::shared_ptr<ast::FunctionVarBody>)) {
      auto funcVarBody =
          std::any_cast<std::shared_ptr<ast::FunctionVarBody>>(value);
      variable->setName(funcVarBody->getName());
    } else {
      auto expr = std::any_cast<std::shared_ptr<ast::IdExpression>>(value);
      variable->setName(
          ParserTreeHelper::getNodeSource(expr->getRuleContext()));
    }

    // Added variable definition to the current scope
    auto scope = d->ctx.getCurrentScope();
    (*scope)->addNode(variable->getName(), variable);

    return value;
  } else {
    auto expr = std::make_shared<ast::Expression>();
    expr->setRuleContext(ctx);
    return std::any(expr);
  }
}

std::any SourceGenerator::visitLiteral(SuckCParser::LiteralContext *ctx) {
  SUCKC_D();

  auto literal = std::make_shared<ast::Literal>();
  literal->setRuleContext(ctx);

  d->traceNodeUnderNextIndent(ctx);

  return std::any(literal);
}

std::any SourceGenerator::visitPrimaryExpression(
    SuckCParser::PrimaryExpressionContext *ctx) {
  SUCKC_D();
  auto expr = std::make_shared<ast::Expression>();
  expr->setRuleContext(ctx);

  d->traceNodeUnderNextIndent(ctx);

  return std::any(expr);
}

std::any
SourceGenerator::visitIdExpression(SuckCParser::IdExpressionContext *ctx) {
  SUCKC_D();
  auto expr = std::make_shared<ast::IdExpression>();
  expr->setRuleContext(ctx);

  return std::any(expr);
}

std::any SourceGenerator::visitNoPointerDeclarator(
    SuckCParser::NoPointerDeclaratorContext *ctx) {
  SUCKC_D();

  if (ctx->noPointerDeclarator() && ctx->parametersAndQualifiers()) {
    auto expr = std::make_shared<ast::FunctionVarBody>();
    expr->setRuleContext(ctx);

    // FIXME: We not treat the parameters correctly
    auto nameValue = visitChildren(ctx->noPointerDeclarator());

    // Track the parameters either
    visitChildren(ctx->parametersAndQualifiers());

    if (typeid(std::shared_ptr<ast::IdExpression>) == nameValue.type()) {
      auto castedNameValue =
          std::any_cast<std::shared_ptr<ast::IdExpression>>(nameValue);
      expr->setName(
          ParserTreeHelper::getNodeSource(castedNameValue->getRuleContext()));
    }
    return std::any(expr);
  } else if (ctx->declaratorid()) {
    // FIXME: We must not ignores attributes!
    return visitChildren(ctx->declaratorid());
  } else {
    return visitChildren(ctx);
  }
}

std::any SourceGenerator::visitAssignmentExpression(
    SuckCParser::AssignmentExpressionContext *ctx) {
  SUCKC_D();
  auto expr = std::make_shared<ast::Expression>();
  expr->setRuleContext(ctx);

  d->traceNodeSourceUnderNextIndent(ctx);

  return std::any(expr);
}

std::any SourceGenerator::visitDeclSpecifierSeq(
    SuckCParser::DeclSpecifierSeqContext *ctx) {

  SUCKC_D();

  return visitChildren(ctx);
}

std::any SourceGenerator::visitPointerDeclarator(
    SuckCParser::PointerDeclaratorContext *ctx) {
  return visitChildren(ctx);
}

std::any
SourceGenerator::visitDeclSpecifier(SuckCParser::DeclSpecifierContext *ctx) {
  SUCKC_D();
  auto expr = std::make_shared<ast::Expression>();
  expr->setRuleContext(ctx);

  d->traceNodeUnderNextIndent(ctx);

  return std::any(expr);
}

std::any
SourceGenerator::visitInitDeclarator(SuckCParser::InitDeclaratorContext *ctx) {
  SUCKC_D();
  auto expr = std::make_shared<ast::Expression>();
  expr->setRuleContext(ctx);

  return visitChildren(ctx);
}

std::any
SourceGenerator::visitInitializer(SuckCParser::InitializerContext *ctx) {
  return visitChildren(ctx);
}

std::any SourceGenerator::visitBraceOrEqualInitializer(
    SuckCParser::BraceOrEqualInitializerContext *ctx) {
  return visitChildren(ctx);
}

std::any SourceGenerator::visitFunctionPointerDeclarator(
    SuckCParser::FunctionPointerDeclaratorContext *ctx) {
  SUCKC_D();

  visitChildren(ctx);

  auto decl = std::make_shared<ast::TypeDeclaration>();
  decl->setRuleContext(ctx);
  decl->setName(ctx->typedefName()->getText());

  return std::any(decl);
}

std::any SourceGenerator::visitSimpleTypedefDeclarator(
    SuckCParser::SimpleTypedefDeclaratorContext *ctx) {
  SUCKC_D();

  visitChildren(ctx);

  auto decl = std::make_shared<ast::TypeDeclaration>();
  decl->setRuleContext(ctx);
  decl->setName(ctx->typedefName()->getText());

  return std::any(decl);
}

std::any SourceGenerator::visitTypedefDeclaration(
    SuckCParser::TypedefDeclarationContext *ctx) {
  SUCKC_D();
  auto alias = std::make_shared<ast::Alias>(ast::Alias::AliasType::Typedef);
  auto value = visitChildren(ctx);
  auto decl = std::any_cast<std::shared_ptr<ast::TypeDeclaration>>(value);

  alias->setRuleContext(ctx);
  alias->setName(decl->getName());
  alias->setTypeDeclaration(decl);

  auto scope = d->ctx.getCurrentScope();
  auto found = (*scope)->findNode(alias->getType(), alias->getName());
  if ((*scope)->findNode(alias->getType(), alias->getName())) {
    std::cout << "WARN: Duplicate type name: " << alias->getName() << std::endl;
  }

  (*scope)->addNode(alias->getName(), alias);

  return alias;
}

std::any SourceGenerator::visitAliasDeclaration(
    SuckCParser::AliasDeclarationContext *ctx) {
  SUCKC_D();

  visitChildren(ctx);

  auto alias = std::make_shared<ast::Alias>(ast::Alias::AliasType::Using);

  alias->setRuleContext(ctx);
  alias->setName(ctx->Identifier()->getText());

  auto scope = d->ctx.getCurrentScope();
  auto found = (*scope)->findNode(alias->getType(), alias->getName());
  if ((*scope)->findNode(alias->getType(), alias->getName())) {
    std::cout << "WARN: Duplicate type name: " << alias->getName() << std::endl;
  }

  (*scope)->addNode(alias->getName(), alias);

  return alias;
}

} // namespace suckc
