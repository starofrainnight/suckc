#ifndef _9M_FF5E45F4_FCC7_4418_85AB_C81E4151DA2F
#define _9M_FF5E45F4_FCC7_4418_85AB_C81E4151DA2F

#include "Node.h"
#include "TypeDeclaration.h"
#include <memory>

namespace suckc {
namespace ast {

class Alias : public Node {
public:
  static const Type kNodeType = Type::Alias;

  enum class AliasType {
    Typedef,
    Using,
  };

  Alias(AliasType aliasType, Type type = kNodeType)
      : Node(type), aliasType_(aliasType) {}

  /**
   * Set type declaration
   * 
   * @param typeDecl
   */
  void setTypeDeclaration(const std::shared_ptr<TypeDeclaration> &typeDecl) {
    typeDeclaration_ = typeDecl;
  }

  /**
   * Get type declaration
   */
  std::shared_ptr<TypeDeclaration> getTypeDeclaration() {
    return typeDeclaration_;
  }

  AliasType getAliasType() { return aliasType_; }

  std::shared_ptr<TypeDeclaration> typeDeclaration_;
  AliasType aliasType_;
};

} // namespace ast
} // namespace suckc

#endif /* _9M_FF5E45F4_FCC7_4418_85AB_C81E4151DA2F */
