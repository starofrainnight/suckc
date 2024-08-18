
#include "SourceContext.h"
#include "Scope.h"
#include <deque>
#include <list>
#include <memory>
#include <string>

namespace suckc {

class SourceContextPrivate {
  SUCKC_OBJECT_IMPL(SourceContext)

public:
  ScopeList scopes;
};

SourceContext::SourceContext() : dPtr_(new SourceContextPrivate(this)) {}

SourceContext::~SourceContext() {}

size_t SourceContext::getScopeCount() const {
  SUCKC_D();
  return d->scopes.size();
}

ScopeList::iterator SourceContext::getCurrentScope() const {
  SUCKC_D();
  auto it = d->scopes.end();
  --it;
  return it;
}

void SourceContext::beginScope(ScopeType scopeType) {
  SUCKC_D();
  d->scopes.push_back(std::make_shared<Scope>(scopeType));
}

void SourceContext::endScope() {
  SUCKC_D();
  d->scopes.pop_back();
}

} // namespace suckc