#ifndef _9M_B7A6FE46_C563_4C35_A308_A3F9FAE49556
#define _9M_B7A6FE46_C563_4C35_A308_A3F9FAE49556

#include "SuckCBaseVisitor.h"
#include "SuckTypes.h"
#include <iostream>

namespace suckc {
class SourceGenerator : public SuckCBaseVisitor {
  SUCKC_OBJECT_DECL(SourceGenerator);

public:
  SourceGenerator(SuckCParser *parser);
  virtual ~SourceGenerator();

  std::any visitChildren(antlr4::tree::ParseTree *node) override;
  std::any
  visitFunctionDefinition(SuckCParser::FunctionDefinitionContext *ctx) override;
  std::any visitExternalDeclaration(
      SuckCParser::ExternalDeclarationContext *ctx) override;

  /**
   * Visit statements that surround by "{" & "}"
   */
  std::any
  visitCompoundStatement(SuckCParser::CompoundStatementContext *ctx) override;

  /**
   * Visit statements that surround by ";"
   */
  std::any visitBlockItem(SuckCParser::BlockItemContext *ctx) override;
};

} // namespace suckc

#endif /* _9M_B7A6FE46_C563_4C35_A308_A3F9FAE49556 */
