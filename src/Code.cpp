#include "Code.h"
#include <list>
#include <string>

namespace suckc {

class CodePrivate {
  SUCKC_OBJECT_IMPL(Code)

public:
  // Declarations
  std::list<std::string> decls;
  //   std::list<std::string>
};

Code::Code() : dPtr_(new CodePrivate(this)) {}

Code::~Code() {}

} // namespace suckc
