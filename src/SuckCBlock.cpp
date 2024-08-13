#include "SuckCBlock.h"

class SuckCBlockPrivate {
  SUCKC_OBJECT_IMPL(SuckCBlock)

public:
};
SuckCBlock::SuckCBlock() : dPtr_(new SuckCBlockPrivate(this)) {}

SuckCBlock::~SuckCBlock() { delete dPtr_; }
