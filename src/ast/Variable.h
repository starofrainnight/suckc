#ifndef _9M_C3CD8AE8_9BD4_48D9_8A2F_9169296472E0
#define _9M_C3CD8AE8_9BD4_48D9_8A2F_9169296472E0

#include "Node.h"
#include <string>

namespace suckc {
namespace ast {
class Variable : public Node {
public:
  static const Type kThisNodeType = Type::Variable;

  Variable(Type type = kThisNodeType) : Node(type) {}

  void setValueType(const std::string &type) { valueType_ = type; }
  const std::string getValueType() { return valueType_; }

  std::string valueType_;
};

} // namespace ast
} // namespace suckc

#endif /* _9M_C3CD8AE8_9BD4_48D9_8A2F_9169296472E0 */
