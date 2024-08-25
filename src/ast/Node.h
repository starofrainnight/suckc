#ifndef _9M_D8D7E5C5_A37D_45A6_96E9_DC2A3E9C0907
#define _9M_D8D7E5C5_A37D_45A6_96E9_DC2A3E9C0907

#include <string>

namespace suckc {
namespace ast {

class Node {
public:
  Node();
  ~Node();

  void setName(const std::string &name) { this->name_ = name; }
  const std::string &getName() { return name_; }

  std::string name_;
};

} // namespace ast
} // namespace suckc

#endif /* _9M_D8D7E5C5_A37D_45A6_96E9_DC2A3E9C0907 */
