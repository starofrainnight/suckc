#include "SuckCScope.h"

class SuckCScopePrivate {
  SUCKC_OBJECT_IMPL(SuckCScope)

public:
};

SuckCScope::SuckCScope() : dPtr_(new SuckCScopePrivate(this)) {}

SuckCScope::~SuckCScope() { delete dPtr_; }