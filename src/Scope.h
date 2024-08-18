#ifndef _9M_C5F79C4F_C2AE_4BAD_AAF7_D1357ABEA463
#define _9M_C5F79C4F_C2AE_4BAD_AAF7_D1357ABEA463

#include "TypeDefs.h"
#include <list>

namespace suckc {
enum ScopeType {
  GlobalScope,
  ClassScope,
  FunctionScope,
  BlockScope,
  StatementScope,
};

class Scope {
  SUCKC_OBJECT_DECL(Scope);

public:
  Scope(ScopeType type);
  ~Scope();

  ScopeType getType() const;
};

typedef std::list<std::shared_ptr<Scope>> ScopeList;

} // namespace suckc

#endif /* _9M_C5F79C4F_C2AE_4BAD_AAF7_D1357ABEA463 */
