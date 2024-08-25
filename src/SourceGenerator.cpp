#include "SourceGenerator.h"
#include "SourceContext.h"
namespace suckc {

class SourceGeneratorPrivate {
  SUCKC_OBJECT_IMPL(SourceGenerator)

public:
  SourceContext ctx;
};

SourceGenerator::SourceGenerator() : dPtr_(new SourceGeneratorPrivate(this)) {}

SourceGenerator::~SourceGenerator() {}
std::any SourceGenerator::visitChildren(antlr4::tree::ParseTree *node) {
  return SuckCBaseVisitor::visitChildren(node);
}

std::any SourceGenerator::visitFunctionDefinition(
    SuckCParser::FunctionDefinitionContext *ctx) {
  SUCKC_D();

  d->ctx.beginScope(ScopeType::FunctionScope);

  std::cout << "got function: " << ctx->getText() << std::endl;
  auto ret = visitChildren(ctx);
  d->ctx.endScope();
  return ret;
}

std::any SourceGenerator::visitExternalDeclaration(
    SuckCParser::ExternalDeclarationContext *ctx) {
  SUCKC_D();
  std::cout << "Extern Declaration: " << std::hex
            << (uintptr_t)ctx->functionDefinition() << std::endl;
  return visitChildren(ctx);
}

} // namespace suckc
