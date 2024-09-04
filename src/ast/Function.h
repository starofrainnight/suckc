#ifndef _9M_B91DAF96_D66F_49E0_A0CA_A1DB8EAFAFFE
#define _9M_B91DAF96_D66F_49E0_A0CA_A1DB8EAFAFFE

#include "Node.h"
namespace suckc {
namespace ast {
class Function : public Node {
public:
  static const Type kNodeType = Type::Function;

  Function(Type type = kNodeType) : Node(type) {}
};

} // namespace ast
} // namespace suckc

#endif /* _9M_B91DAF96_D66F_49E0_A0CA_A1DB8EAFAFFE */
