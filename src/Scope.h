#ifndef _9M_C5F79C4F_C2AE_4BAD_AAF7_D1357ABEA463
#define _9M_C5F79C4F_C2AE_4BAD_AAF7_D1357ABEA463

#include "SuckTypes.h"
#include "ast/Function.h"
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
  typedef ast::Node::Type NodeType;

  Scope(ScopeType type);
  ~Scope();

  ScopeType getType() const;

  std::shared_ptr<ast::Node> findNode(const NodeType type,
                                      const std::string &name);

  /**
   * @brief Find a node in scope
   * 
   * @param name Variable name
   * @return Variable or nullptr
   */
  auto findFunction(const std::string &name) -> std::shared_ptr<ast::Function> {
    return std::static_pointer_cast<ast::Function>(
        findNode(NodeType::Function, name));
  }

  auto findVariable(const std::string &name) -> std::shared_ptr<ast::Variable> {
    return std::static_pointer_cast<ast::Variable>(
        findNode(NodeType::Function, name));
  }

  void addNode(const std::string &name, const std::shared_ptr<ast::Node> &node);
};

typedef std::list<std::shared_ptr<Scope>> ScopeList;

} // namespace suckc

#endif /* _9M_C5F79C4F_C2AE_4BAD_AAF7_D1357ABEA463 */
