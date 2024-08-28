#ifndef _9M_D8D7E5C5_A37D_45A6_96E9_DC2A3E9C0907
#define _9M_D8D7E5C5_A37D_45A6_96E9_DC2A3E9C0907

#include <string>

namespace suckc {
namespace ast {

class Node {
public:
  enum class Type {
    Variable,
    Function,
    FunctionDeclaration,
    Struct,
    Count_,
  };

  Node(Type type);
  ~Node();

  void setName(const std::string &name) { this->name_ = name; }
  const std::string &getName() { return name_; }

  Type getType() { return type_; }

  std::string name_;
  Type type_;
};

} // namespace ast
} // namespace suckc

#endif /* _9M_D8D7E5C5_A37D_45A6_96E9_DC2A3E9C0907 */
