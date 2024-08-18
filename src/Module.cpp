#include "Module.h"

namespace suckc {

class ModulePrivate {
  SUCKC_OBJECT_IMPL(Module)

public:
};

Module::Module() : Scope(GlobalScope), dPtr_(new ModulePrivate(this)) {}

Module::~Module() {}

} // namespace suckc