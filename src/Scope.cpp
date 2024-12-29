#include "Scope.h"
#include "ast/Function.h"
#include <unordered_map>

namespace suckc {

class ScopePrivate {
  SUCKC_OBJECT_IMPL(Scope)

public:
  typedef std::unordered_map<std::string, std::shared_ptr<ast::Node>> NodeMap;

  ScopeType type;
  NodeMap nodes[static_cast<int>(ast::Node::Type::Count_)];
};

Scope::Scope(ScopeType type) : dPtr_(new ScopePrivate(this)) {
  SUCKC_D();
  d->type = type;
}

Scope::~Scope() {}

void Scope::print() const {
  SUCKC_D();

  std::cout << "SCOPE_BEGIN(" << d->type << ")" << std::endl;
  for (auto &nodeMap : d->nodes) {
    for (auto &node : nodeMap) {
      std::cout << node.first << ": " << node.second->toSource() << std::endl;
    }
  }
  std::cout << "SCOPE_END" << std::endl;
}

ScopeType Scope::getType() const {
  SUCKC_D();
  return d->type;
}

std::shared_ptr<ast::Node> Scope::findNode(const NodeType type,
                                           const std::string &name) {
  SUCKC_D();
  auto nodeMap = &d->nodes[static_cast<int>(type)];
  auto it = nodeMap->find(name);
  if (it != nodeMap->end())
    return it->second;
  else
    return nullptr;
}

void Scope::addNode(const std::string &name,
                    const std::shared_ptr<ast::Node> &node) {
  SUCKC_D();
  auto &nodeMap = d->nodes[static_cast<int>(node->getType())];
  nodeMap[name] = node;
}

} // namespace suckc