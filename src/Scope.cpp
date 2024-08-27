#include "Scope.h"
#include "ast/Function.h"
#include <unordered_map>

namespace suckc {

class ScopePrivate {
  SUCKC_OBJECT_IMPL(Scope)

public:
  typedef std::unordered_map<std::string, std::shared_ptr<suckc::ast::Variable>>
      VariableMap;

  typedef std::unordered_map<std::string, std::shared_ptr<suckc::ast::Function>>
      FunctionMap;

  ScopeType type;
  VariableMap variables;
  FunctionMap funcs;
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

std::shared_ptr<suckc::ast::Function>
Scope::findFunction(const std::string &name) {
  SUCKC_D();
  auto it = d->funcs.find(name);
  if (it != d->funcs.end())
    return it->second;
  else
    return nullptr;
}

void Scope::addFunction(const std::string &name,
                        const std::shared_ptr<suckc::ast::Function> &func) {
  SUCKC_D();
  d->funcs[name] = func;
}

} // namespace suckc