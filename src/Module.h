#ifndef _9M_F7D55327_EE07_4AF1_8366_F7D91D79C13C
#define _9M_F7D55327_EE07_4AF1_8366_F7D91D79C13C

#include "Scope.h"

namespace suckc {

class Module : public Scope {
  SUCKC_OBJECT_DECL(Module);

public:
  Module(/* args */);
  ~Module();
};

} // namespace suckc

#endif /* _9M_F7D55327_EE07_4AF1_8366_F7D91D79C13C */
