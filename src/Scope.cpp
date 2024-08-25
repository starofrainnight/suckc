#include "Scope.h"
#include <unordered_map>

namespace suckc {

class ScopePrivate {
  SUCKC_OBJECT_IMPL(Scope)

public:
  typedef std::unordered_map<std::string, std::shared_ptr<suckc::ast::Variable>>
      VariableMap;

  ScopeType type;
  VariableMap variables;
};

Scope::Scope(ScopeType type) : dPtr_(new ScopePrivate(this)) {
  SUCKC_D();
  d->type = type;
}

Scope::~Scope() {}

ScopeType Scope::getType() const {
  SUCKC_D();
  return d->type;
}

std::shared_ptr<suckc::ast::Variable>
Scope::findVariable(const std::string &name) {
  SUCKC_D();

  auto it = d->variables.find(name);
  if (it != d->variables.end())
    return it->second;
  else
    return nullptr;
}

void Scope::addVariable(const std::string &name,
                        const std::shared_ptr<suckc::ast::Variable> &var) {
  SUCKC_D();
  d->variables[name] = var;
}

} // namespace suckc