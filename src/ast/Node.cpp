#include "Node.h"

namespace suckc {
namespace ast {

Node::Node(Type type) : type_(type), ruleCtx_(nullptr) {}

Node::~Node() {}

} // namespace ast
} // namespace suckc
