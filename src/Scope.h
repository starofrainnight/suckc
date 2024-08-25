#ifndef _9M_C5F79C4F_C2AE_4BAD_AAF7_D1357ABEA463
#define _9M_C5F79C4F_C2AE_4BAD_AAF7_D1357ABEA463

#include "SuckTypes.h"
#include "ast/Variable.h"
#include <list>

namespace suckc {
enum ScopeType {
  GlobalScope,
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

  /**
   * @brief Find variable in scope
   * 
   * @param name Variable name
   * @return Variable or nullptr
   */
  std::shared_ptr<suckc::ast::Variable> findVariable(const std::string &name);
  void addVariable(const std::string &name,
                   const std::shared_ptr<suckc::ast::Variable> &var);
};

typedef std::list<std::shared_ptr<Scope>> ScopeList;

} // namespace suckc

#endif /* _9M_C5F79C4F_C2AE_4BAD_AAF7_D1357ABEA463 */
