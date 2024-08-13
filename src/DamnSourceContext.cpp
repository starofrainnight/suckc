
#include "DamnSourceContext.h"
#include "Scope.h"
#include <list>
#include <memory>
#include <string>

namespace suckc {

class DamnSourceContextPrivate {
  SUCKC_OBJECT_IMPL(DamnSourceContext)

public:
  std::list<std::shared_ptr<Scope>> scopeStack;
};

DamnSourceContext::DamnSourceContext()
    : dPtr_(new DamnSourceContextPrivate(this)) {}

DamnSourceContext::~DamnSourceContext() { delete dPtr_; }

}