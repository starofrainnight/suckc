#ifndef _9M_B7A6FE46_C563_4C35_A308_A3F9FAE49556
#define _9M_B7A6FE46_C563_4C35_A308_A3F9FAE49556

#include "SuckCParserBaseVisitor.h"
#include "SuckTypes.h"
#include <iostream>

namespace suckc {
class SourceGenerator : public SuckCParserBaseVisitor {
  SUCKC_OBJECT_DECL(SourceGenerator);

public:
  SourceGenerator(SuckCParser *parser);
  virtual ~SourceGenerator();

  std::any visitChildren(antlr4::tree::ParseTree *node) override;

  std::any aggregateResult(std::any aggregate, std::any nextResult) override;

  // The entry node
  std::any
  visitTranslationUnit(SuckCParser::TranslationUnitContext *ctx) override;

  std::any
  visitFunctionDefinition(SuckCParser::FunctionDefinitionContext *ctx) override;
  // std::any
  // visitItemDefinition(SuckCParser::ItemDefinitionContext *ctx) override;

  /**
   * Visit statements that surround by "{" & "}"
   */
  std::any
  visitCompoundStatement(SuckCParser::CompoundStatementContext *ctx) override;

  std::any visitStatement(SuckCParser::StatementContext *ctx) override;

  /**
   * Visit statements that surround by ";"
   */
  // std::any visitBlockItem(SuckCParser::BlockItemContext *ctx) override;

  // std::any visitDeclaration(SuckCParser::DeclarationContext *ctx) override;
  std::any
  visitSimpleDeclaration(SuckCParser::SimpleDeclarationContext *ctx) override;

  std::any visitLiteral(SuckCParser::LiteralContext *ctx) override;
  std::any
  visitPrimaryExpression(SuckCParser::PrimaryExpressionContext *ctx) override;

  std::any visitIdExpression(SuckCParser::IdExpressionContext *ctx) override;

  std::any visitNoPointerDeclarator(
      SuckCParser::NoPointerDeclaratorContext *ctx) override;

  std::any
  visitPointerDeclarator(SuckCParser::PointerDeclaratorContext *ctx) override;

  // simpleDeclaration: declSpecifierSeq declSpecifierSeq initializer
  // auto v = 1;
  std::any
  visitDeclSpecifierSeq(SuckCParser::DeclSpecifierSeqContext *ctx) override;

  std::any visitDeclSpecifier(SuckCParser::DeclSpecifierContext *ctx) override;

  std::any visitAssignmentExpression(
      SuckCParser::AssignmentExpressionContext *ctx) override;

  std::any
  visitInitDeclarator(SuckCParser::InitDeclaratorContext *ctx) override;
  std::any visitInitializer(SuckCParser::InitializerContext *ctx) override;
  std::any visitBraceOrEqualInitializer(
      SuckCParser::BraceOrEqualInitializerContext *ctx) override;

  std::any visitFunctionPointerDeclarator(
      SuckCParser::FunctionPointerDeclaratorContext *ctx) override;

  std::any visitSimpleTypedefDeclarator(
      SuckCParser::SimpleTypedefDeclaratorContext *ctx) override;

  std::any
  visitTypedefDeclaration(SuckCParser::TypedefDeclarationContext *ctx) override;

  std::any
  visitAliasDeclaration(SuckCParser::AliasDeclarationContext *ctx) override;
};

} // namespace suckc

#endif /* _9M_B7A6FE46_C563_4C35_A308_A3F9FAE49556 */
