#ifndef _9M_CE046455_186A_42C5_8EE7_4382BAEC7538
#define _9M_CE046455_186A_42C5_8EE7_4382BAEC7538

#include "Node.h"
namespace suckc {
namespace ast {
class TypeDeclaration : public Node {
public:
  static const Type kNodeType = Type::TypeDeclaration;

  TypeDeclaration(Type type = kNodeType) : Node(type) {}
};

} // namespace ast
} // namespace suckc

#endif /* _9M_CE046455_186A_42C5_8EE7_4382BAEC7538 */
