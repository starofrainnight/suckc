
#include "DamnSourceContext.h"

class DamnSourceContextPrivate {
  SUCKC_OBJECT_IMPL(DamnSourceContext)

public:
};

DamnSourceContext::DamnSourceContext()
    : dPtr_(new DamnSourceContextPrivate(this)) {}

DamnSourceContext::~DamnSourceContext() { delete dPtr_; }
