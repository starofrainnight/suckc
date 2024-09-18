#include "SuckCParserBase.h"
#include "SuckCParser.h"

using namespace antlr4;

bool SuckCParserBase::IsPureSpecifierAllowed() {
  try {
    auto x = this->getRuleContext(); // memberDeclarator
    auto c = x->children[0]->children[0];
    auto c2 = c->children[0];
    auto p = c2->children[1];
    if (p == nullptr)
      return false;
    return typeid(*p) == typeid(SuckCParser::ParametersAndQualifiersContext);
  } catch (...) {
  }
  return false;
}
