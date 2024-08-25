#ifndef _9M_CE288385_B4D3_4126_83B9_114A31247C13
#define _9M_CE288385_B4D3_4126_83B9_114A31247C13

#include "Scope.h"
#include "SuckTypes.h"

namespace suckc {
class SourceContext {
  SUCKC_OBJECT_DECL(SourceContext);

public:
  SourceContext();
  ~SourceContext();

  const ScopeList &getScopes() const;
  size_t getScopeCount() const;
  ScopeList::iterator getCurrentScope() const;

  void beginScope(ScopeType scopeType);
  void endScope();
};

} // namespace suckc

#endif /* _9M_CE288385_B4D3_4126_83B9_114A31247C13 */
