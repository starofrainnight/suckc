#include "Module.h"

namespace suckc {

class ModulePrivate {
  SUCKC_OBJECT_IMPL(Module)

public:
};

Module::Module() : dPtr_(new ModulePrivate(this)) {}

Module::~Module() { delete dPtr_; }

}