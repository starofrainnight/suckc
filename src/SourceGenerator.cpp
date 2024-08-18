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

} // namespace suckc
