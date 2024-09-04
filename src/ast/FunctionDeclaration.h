#ifndef _9M_B67C6442_D5B2_489A_91EA_D8DB09F8B481
#define _9M_B67C6442_D5B2_489A_91EA_D8DB09F8B481

#include "Node.h"
namespace suckc {
namespace ast {
class FunctionDeclaration : public Node {
public:
  static const Type kNodeType = Type::FunctionDeclaration;

  FunctionDeclaration(Type type = kNodeType) : Node(type) {}
};

} // namespace ast
} // namespace suckc

#endif /* _9M_B67C6442_D5B2_489A_91EA_D8DB09F8B481 */
