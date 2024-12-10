#ifndef _9M_F4257AE3_72AC_4EE3_A9A2_10BF012C0116
#define _9M_F4257AE3_72AC_4EE3_A9A2_10BF012C0116

#include <antlr4-runtime.h>
#include <string>
#include <vector>

namespace suckc {
class ParserTreeHelper {
public:
  static std::vector<std::string>
  collectTokenText(antlr4::tree::ParseTree *node);

  static std::vector<std::string>
  collectTokenText(antlr4::tree::TerminalNode *node);

  static std::string getNodeRuleName(antlr4::tree::ParseTree *node);
  static std::string getNodeStartTokenText(antlr4::tree::ParseTree *node);

  /**
   * @brief Get the token text of a parse tree node
   * 
   * This function collects and concatenates the text content of all tokens 
   * under the given parse tree node. It is useful when generating source 
   * code, especially when considering the original text content of nodes.
   * 
   * @param node A pointer to the parse tree node. This is the starting node
   * from which we collect token text.
   * @return std::string Returns a string containing the concatenated text
   * content of all tokens under the node, separated by spaces.
   */
  static std::string getNodeTokensText(antlr4::tree::ParseTree *node);
  static std::string getNodeSource(antlr4::tree::ParseTree *node);
};

} // namespace suckc

#endif /* _9M_F4257AE3_72AC_4EE3_A9A2_10BF012C0116 */
