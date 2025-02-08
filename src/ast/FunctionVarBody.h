#ifndef _9M_F9323CBE_543A_4799_A904_220574417E30
#define _9M_F9323CBE_543A_4799_A904_220574417E30

#include "Node.h"
namespace suckc {
namespace ast {
class FunctionVarBody : public Node {
public:
  static const Type kNodeType = Type::FunctionVarBody;

  FunctionVarBody(Type type = kNodeType) : Node(type) {}
};

} // namespace ast
} // namespace suckc

#endif /* _9M_F9323CBE_543A_4799_A904_220574417E30 */
