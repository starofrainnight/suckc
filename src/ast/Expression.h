#ifndef _9M_B9D27349_56C1_4C0B_AE72_106F2AA66C04
#define _9M_B9D27349_56C1_4C0B_AE72_106F2AA66C04

#include "Node.h"
namespace suckc {
namespace ast {

/**
 * Expressions will directly return fron visit()
 */
class Expression : public Node {
public:
  static const Type kNodeType = Type::Expression;

  Expression(Type type = kNodeType) : Node(type) {}
};

} // namespace ast
} // namespace suckc

#endif /* _9M_B9D27349_56C1_4C0B_AE72_106F2AA66C04 */
