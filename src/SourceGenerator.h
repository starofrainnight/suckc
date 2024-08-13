#ifndef _9M_B7A6FE46_C563_4C35_A308_A3F9FAE49556
#define _9M_B7A6FE46_C563_4C35_A308_A3F9FAE49556

#include "SuckCBaseVisitor.h"
#include "TypeDefs.h"
#include <iostream>

namespace suckc {
class SourceGenerator : public SuckCBaseVisitor {
public:
  SourceGenerator();
  virtual ~SourceGenerator();

  virtual std::any visitChildren(antlr4::tree::ParseTree *node) override {
    std::cout << "Text: " << node->getText() << std::endl;
    return SuckCBaseVisitor::visitChildren(node);
  }
};

} // namespace suckc

#endif /* _9M_B7A6FE46_C563_4C35_A308_A3F9FAE49556 */
