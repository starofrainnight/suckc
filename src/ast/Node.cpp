#include "Node.h"
#include "ParserTreeHelper.h"

namespace suckc {
namespace ast {

Node::Node(Type type) : type_(type), ruleCtx_(nullptr) {}

Node::~Node() {}

std::string Node::toSource() {
  return ParserTreeHelper::getNodeSource(ruleCtx_);
}

} // namespace ast
} // namespace suckc
