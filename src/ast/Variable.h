#ifndef _9M_C3CD8AE8_9BD4_48D9_8A2F_9169296472E0
#define _9M_C3CD8AE8_9BD4_48D9_8A2F_9169296472E0

#include "Attribute.h"
#include "Node.h"
#include "Specifier.h"
#include <string>

namespace suckc {
namespace ast {
class Variable : public Node {
public:
  static const Type kNodeType = Type::Variable;

  Variable(Type type = kNodeType) : Node(type) {}

  std::list<std::shared_ptr<Specifier>> *getSpecifiers() {
    return &specifiers_;
  }
  std::list<std::shared_ptr<Attribute>> *getAttributes() {
    return &attributes_;
  }

  std::list<std::shared_ptr<Specifier>> specifiers_;
  std::list<std::shared_ptr<Attribute>> attributes_;
};

} // namespace ast
} // namespace suckc

#endif /* _9M_C3CD8AE8_9BD4_48D9_8A2F_9169296472E0 */
