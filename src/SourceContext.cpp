#include "SourceContext.h"
#include "Scope.h"
#include "ast/Node.h"
#include <deque>
#include <list>
#include <memory>
#include <string>

namespace suckc {

class SourceContextPrivate {
  SUCKC_OBJECT_IMPL(SourceContext)

public:
  ScopeList scopes;
};

SourceContext::SourceContext() : dPtr_(new SourceContextPrivate(this)) {
  beginScope(ScopeType::GlobalScope);
}

SourceContext::~SourceContext() { endScope(); }

const ScopeList &SourceContext::getScopes() const {
  SUCKC_D();
  return d->scopes;
}

size_t SourceContext::getScopeCount() const {
  SUCKC_D();
  return d->scopes.size();
}

ScopeList::iterator SourceContext::getCurrentScope() const {
  SUCKC_D();
  auto it = d->scopes.end();
  --it;
  return it;
}

void SourceContext::beginScope(ScopeType scopeType) {
  SUCKC_D();
  d->scopes.push_back(std::make_shared<Scope>(scopeType));
}

void SourceContext::endScope() {
  SUCKC_D();
  d->scopes.pop_back();
}

void SourceContext::print() {
  SUCKC_D();

  std::cout << "CONTEXT_BEGIN" << std::endl;
  for (auto it = d->scopes.begin(); it != d->scopes.end(); ++it) {
    (*it)->print();
  }
  std::cout << "CONTEXT_END" << std::endl;
}

std::shared_ptr<ast::Node> SourceContext::findNode(const ast::Node::Type &type,
                                                   const std::string &name) {
  SUCKC_D();
  for (auto it = d->scopes.rbegin(); it != d->scopes.rend(); ++it) {
    auto node = (*it)->findNode(type, name);
    if (node != nullptr) {
      return node;
    }
  }

  return nullptr;
}

void SourceContext::addNode(const std::string &name,
                            const std::shared_ptr<ast::Node> &node) {
  (*getCurrentScope())->addNode(name, node);
}

} // namespace suckc