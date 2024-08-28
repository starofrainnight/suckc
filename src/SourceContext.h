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

  std::shared_ptr<ast::Variable> findVariable(const std::string &name) {
    return std::static_pointer_cast<ast::Variable>(
        findNode(ast::Variable::kThisNodeType, name));
  }

  std::shared_ptr<ast::Function> findFunction(const std::string &name) {
    return std::static_pointer_cast<ast::Function>(
        findNode(ast::Function::kThisNodeType, name));
  }

  std::shared_ptr<ast::Node> findNode(const ast::Node::Type &type,
                                      const std::string &name);
  void addNode(const std::string &name, const std::shared_ptr<ast::Node> &var);
};

} // namespace suckc

#endif /* _9M_CE288385_B4D3_4126_83B9_114A31247C13 */
