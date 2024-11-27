#ifndef _9M_ADE740F5_1942_429B_B460_5659B79F89C5
#define _9M_ADE740F5_1942_429B_B460_5659B79F89C5

#include "Expression.h"
namespace suckc {
namespace ast {
class Literal : public Expression {
public:
  static const Type kNodeType = Type::Literal;

  Literal(Type type = kNodeType) : Expression(type) {}
};

} // namespace ast
} // namespace suckc

#endif /* _9M_ADE740F5_1942_429B_B460_5659B79F89C5 */
