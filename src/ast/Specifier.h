#ifndef _9M_C37D7EDD_CBDD_43B9_81F7_84C8C5B835AC
#define _9M_C37D7EDD_CBDD_43B9_81F7_84C8C5B835AC

#include "Node.h"
namespace suckc {
namespace ast {

/**
 * Expressions will directly return fron visit()
 */
class Specifier : public Node {
public:
  static const Type kNodeType = Type::Specifier;

  Specifier(Type type = kNodeType) : Node(type) {}
};

} // namespace ast
} // namespace suckc

#endif /* _9M_C37D7EDD_CBDD_43B9_81F7_84C8C5B835AC */
