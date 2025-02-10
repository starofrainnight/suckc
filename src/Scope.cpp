#include "Scope.h"
#include "ast/Function.h"
#include <list>
#include <unordered_map>

namespace suckc {

class ScopePrivate {
  SUCKC_OBJECT_IMPL(Scope)

public:
  typedef std::unordered_map<std::string, std::shared_ptr<ast::Node>> NodeMap;
  typedef std::list<std::string> NodeList;

  ScopeType type;
  NodeMap nodeMaps[static_cast<int>(ast::Node::Type::Count_)];
  NodeList keyLists[static_cast<int>(ast::Node::Type::Count_)];
};

Scope::Scope(ScopeType type) : dPtr_(new ScopePrivate(this)) {
  SUCKC_D();
  d->type = type;
}

Scope::~Scope() {}

void Scope::print() const {
  SUCKC_D();

  std::cout << "SCOPE_BEGIN(" << d->type << ")" << std::endl;
  for (int i = 0; i < static_cast<int>(ast::Node::Type::Count_); ++i) {
    auto &nodeMap = d->nodeMaps[i];
    auto &keyList = d->keyLists[i];
    for (auto &key : keyList) {
      auto node = nodeMap[key];
      std::cout << "NODE(" << (int)node->getType() << "):"
                << "name(" << node->getName()
                << "):source: " << node->toSource() << std::endl;
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
  auto nodeMap = &d->nodeMaps[static_cast<int>(type)];
  auto it = nodeMap->find(name);
  if (it != nodeMap->end())
    return it->second;
  else
    return nullptr;
}

void Scope::addNode(const std::string &name,
                    const std::shared_ptr<ast::Node> &node) {
  SUCKC_D();
  auto &nodeMap = d->nodeMaps[static_cast<int>(node->getType())];
  auto &keyList = d->keyLists[static_cast<int>(node->getType())];
  if (nodeMap.find(name) == nodeMap.end()) {
    keyList.push_back(name);
  }
  nodeMap[name] = node;
}

} // namespace suckc