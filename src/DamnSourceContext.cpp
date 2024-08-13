
#include "DamnSourceContext.h"
#include "SuckCScope.h"
#include <list>
#include <memory>
#include <string>

class DamnSourceContextPrivate {
  SUCKC_OBJECT_IMPL(DamnSourceContext)

public:
  std::list<std::shared_ptr<SuckCScope>> scopeStack;
};

DamnSourceContext::DamnSourceContext()
    : dPtr_(new DamnSourceContextPrivate(this)) {}

DamnSourceContext::~DamnSourceContext() { delete dPtr_; }
