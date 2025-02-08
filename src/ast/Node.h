#ifndef _9M_D8D7E5C5_A37D_45A6_96E9_DC2A3E9C0907
#define _9M_D8D7E5C5_A37D_45A6_96E9_DC2A3E9C0907

#include <antlr4-runtime.h>
#include <list>
#include <string>

namespace suckc {
namespace ast {

class Node {
public:
  enum class Type {
    Variable,
    Function,
    FunctionDeclaration,
    FunctionVarBody,
    Struct,
    Expression,
    Literal,
    Specifier,
    Attribute,
    Count_,
  };

  Node(Type type);
  ~Node();

  void setName(const std::string &name) { name_ = name; }
  const std::string &getName() { return name_; }

  Type getType() { return type_; }

  antlr4::RuleContext *getRuleContext() { return ruleCtx_; }
  void setRuleContext(antlr4::RuleContext *ctx) { ruleCtx_ = ctx; }

  std::list<Node *> *getChildren() { return &children_; };

  virtual std::string toSource();

  std::string name_;
  Type type_;
  antlr4::RuleContext *ruleCtx_;
  std::list<Node *> children_;
};

} // namespace ast
} // namespace suckc

#endif /* _9M_D8D7E5C5_A37D_45A6_96E9_DC2A3E9C0907 */
