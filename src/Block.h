#ifndef _9M_AD13CD6B_FD71_40E1_AA17_F3DA14C8F867
#define _9M_AD13CD6B_FD71_40E1_AA17_F3DA14C8F867

#include "Scope.h"

namespace suckc {

/**
 * A block statement code is a series statements that chain together.
 *
 * Normally a comma as a separator is used to separate block statements.
 *
 * For example:
 *
 * @code
 * // Block statement 1
 * a = func1(func2(b));
 *
 * // Block statement 2
 * b = func3(func4(c));
 * @endcode
 */
class Block : public Scope {
  SUCKC_OBJECT_DECL(Block);

public:
  Block();
  ~Block();
};

} // namespace suckc

#endif /* _9M_AD13CD6B_FD71_40E1_AA17_F3DA14C8F867 */
