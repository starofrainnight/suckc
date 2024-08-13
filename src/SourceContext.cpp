
#include "SourceContext.h"
#include "Scope.h"
#include <list>
#include <memory>
#include <string>

namespace suckc {

class SourceContextPrivate {
  SUCKC_OBJECT_IMPL(SourceContext)

public:
  std::list<std::shared_ptr<Scope>> scopeStack;
};

SourceContext::SourceContext() : dPtr_(new SourceContextPrivate(this)) {}

SourceContext::~SourceContext() { delete dPtr_; }

} // namespace suckc