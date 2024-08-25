#ifndef _9M_C3CD8AE8_9BD4_48D9_8A2F_9169296472E0
#define _9M_C3CD8AE8_9BD4_48D9_8A2F_9169296472E0

#include "Node.h"
#include <string>

namespace suckc {
namespace ast {
class Variable : public Node {
public:
  void setType(const std::string &type) { type_ = type; }
  const std::string getType() { return type_; }

  std::string type_;
};

} // namespace ast
} // namespace suckc

#endif /* _9M_C3CD8AE8_9BD4_48D9_8A2F_9169296472E0 */
