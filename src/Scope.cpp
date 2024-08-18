#include "Scope.h"

namespace suckc {

class ScopePrivate {
  SUCKC_OBJECT_IMPL(Scope)

public:
  ScopeType type;
};

Scope::Scope(ScopeType type) : dPtr_(new ScopePrivate(this)) {
  SUCKC_D();
  d->type = type;
}

Scope::~Scope() {}

ScopeType Scope::getType() const {
  SUCKC_D();
  return d->type;
}

} // namespace suckc