#include "Scope.h"

namespace suckc {

class ScopePrivate {
  SUCKC_OBJECT_IMPL(Scope)

public:
};

Scope::Scope() : dPtr_(new ScopePrivate(this)) {}

Scope::~Scope() { delete dPtr_; }

} // namespace suckc