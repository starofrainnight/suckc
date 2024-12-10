#include "ParserTreeHelper.h"
#include "SuckCParser.h"
#include "World.h"

namespace suckc {
std::vector<std::string>
ParserTreeHelper::collectTokenText(antlr4::tree::TerminalNode *node) {
  // This function collects the text of a terminal node.
  return {node->getText()};
}

std::vector<std::string>
ParserTreeHelper::collectTokenText(antlr4::tree::ParseTree *node) {
  // This function collects the text of all tokens under a given non-terminal node.
  std::vector<std::string> tokens;

  for (auto child : node->children) {
    if (auto terminalNode = dynamic_cast<antlr4::tree::TerminalNode *>(child)) {
      auto childTokens = collectTokenText(terminalNode);

      tokens.insert(tokens.end(), childTokens.begin(), childTokens.end());
    } else if (auto ruleNode = dynamic_cast<antlr4::RuleContext *>(child)) {
      auto childTokens = collectTokenText(ruleNode);
      tokens.insert(tokens.end(), childTokens.begin(), childTokens.end());
    }
  }

  return tokens;
}

std::string ParserTreeHelper::getNodeRuleName(antlr4::tree::ParseTree *node) {
  antlr4::ParserRuleContext *context =
      dynamic_cast<antlr4::ParserRuleContext *>(node);
  if (context != nullptr) {
    // Get the rule index of this node
    int ruleIndex = context->getRuleIndex();

    // Get the parser from which you can retrieve rule names
    if (context != nullptr) {
      // Retrieve the rule name using the rule index
      std::vector<std::string> ruleNames =
          suckc::World::getInstance()->getParser()->getRuleNames();
      std::string ruleName = ruleNames[ruleIndex];

      // Now you have the rule name, you can do something with it
      return ruleName;
    }
  }

  return "";
}
std::string
ParserTreeHelper::getNodeStartTokenText(antlr4::tree::ParseTree *node) {
  if (node != nullptr) {
    antlr4::ParserRuleContext *ctx =
        dynamic_cast<antlr4::ParserRuleContext *>(node);
    if (ctx != nullptr) {
      // Get the start token of the context
      auto startToken = ctx->start;

      // Get the text associated with the start token
      auto nodeText = startToken->getText();

      return nodeText;
    }
  }

  return "";
}

std::string ParserTreeHelper::getNodeTokensText(antlr4::tree::ParseTree *node) {
  auto texts = collectTokenText(node);

  std::string text;
  for (auto i = 0; i < texts.size(); ++i) {
    text.append(texts[i]);
    text.append(" ");
  }

  return text;
}

std::string ParserTreeHelper::getNodeSource(antlr4::tree::ParseTree *node) {

  // Traverse the tree to find the node of interest
  // For simplicity, we will just use the root node in this example
  if (node != nullptr) {
    antlr4::ParserRuleContext *ctx =
        dynamic_cast<antlr4::ParserRuleContext *>(node);
    if (ctx != nullptr) {
      // Get the start and stop tokens
      antlr4::Token *startToken = ctx->start;
      antlr4::Token *stopToken = ctx->stop;

      // Get the TokenStream from the lexer
      antlr4::TokenStream *tokenStream =
          World::getInstance()->getParser()->getTokenStream();

      // Get the source code block between the start and stop tokens
      std::string sourceCodeBlock = tokenStream->getText(antlr4::misc::Interval(
          startToken->getTokenIndex(), stopToken->getTokenIndex()));

      return sourceCodeBlock;
    }
  }
  return "";
}

} // namespace suckc