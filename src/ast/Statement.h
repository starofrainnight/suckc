#ifndef _9M_CA34F4FE_345F_4ACF_BF05_25DB9C851321
#define _9M_CA34F4FE_345F_4ACF_BF05_25DB9C851321

#include "Node.h"
#include <string>

namespace suckc {
namespace ast {
class Statement : public Node {
public:
  static const Type kThisNodeType = Type::Variable;

  Statement(Type type = kThisNodeType) : Node(type) {}
};

} // namespace ast
} // namespace suckc

#endif /* _9M_CA34F4FE_345F_4ACF_BF05_25DB9C851321 */
