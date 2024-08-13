#include "SuckCCode.h"
#include <list>
#include <string>

class SuckCCodePrivate {
  SUCKC_OBJECT_IMPL(SuckCCode)

public:
  // Declarations
  std::list<std::string> decls;
  //   std::list<std::string>
};

SuckCCode::SuckCCode() : dPtr_(new SuckCCodePrivate(this)) {}

SuckCCode::~SuckCCode() {}
