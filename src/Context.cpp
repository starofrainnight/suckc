
#include "Context.h"

class ContextPrivate {
  SUCKC_OBJECT_IMPL(Context)

public:
};

Context::Context() : dPtr_(new ContextPrivate(this)) {}

Context::~Context() { delete dPtr_; }
