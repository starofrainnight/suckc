#ifndef _9M_FB7782EF_9349_4884_9191_2935E214CA3B
#define _9M_FB7782EF_9349_4884_9191_2935E214CA3B

#include "Node.h"

namespace suckc {
namespace ast {
class Struct : public Node {
public:
  static const Type kThisNodeType = Type::Struct;

  Struct(Type type = kThisNodeType) : Node(type) {}
};

} // namespace ast
} // namespace suckc

#endif /* _9M_FB7782EF_9349_4884_9191_2935E214CA3B */
