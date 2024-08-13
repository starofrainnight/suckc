#include "SuckCModule.h"

class SuckCModulePrivate {
  SUCKC_OBJECT_IMPL(SuckCModule)

public:
};

SuckCModule::SuckCModule() : dPtr_(new SuckCModulePrivate(this)) {}

SuckCModule::~SuckCModule() { delete dPtr_; }