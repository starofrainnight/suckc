#ifndef _9M_CE4DD4D1_30CF_48C2_8C55_41AEA7EDE624
#define _9M_CE4DD4D1_30CF_48C2_8C55_41AEA7EDE624

#include "Node.h"
namespace suckc {
namespace ast {

/**
 * Expressions will directly return fron visit()
 */
class Attribute : public Node {
public:
  static const Type kNodeType = Type::Attribute;

  Attribute(Type type = kNodeType) : Node(type) {}
};

} // namespace ast
} // namespace suckc

#endif /* _9M_CE4DD4D1_30CF_48C2_8C55_41AEA7EDE624 */
