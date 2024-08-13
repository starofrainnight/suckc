#ifndef _9M_B7A6FE46_C563_4C35_A308_A3F9FAE49556
#define _9M_B7A6FE46_C563_4C35_A308_A3F9FAE49556

#include "SuckCBaseVisitor.h"
#include "TypeDefs.h"
#include <iostream>

class DamnSourceGenerator : public SuckCBaseVisitor {
public:
  DamnSourceGenerator();
  virtual ~DamnSourceGenerator();

  virtual std::any visitChildren(antlr4::tree::ParseTree *node) override {
    std::cout << "Text: " << node->getText() << std::endl;
    return SuckCBaseVisitor::visitChildren(node);
  }
};

#endif /* _9M_B7A6FE46_C563_4C35_A308_A3F9FAE49556 */
