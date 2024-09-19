#include "SourceGenerator.h"
#include "SuckCLexer.h"
#include "SuckCParser.h"
#include <antlr4-runtime.h>
#include <fstream>
#include <iostream>

/**
 * Print usage information
 */
void printUsage() { std::cout << "Usage: suckc <input file>" << std::endl; }

/**
 * Main entry point of the program.
 *
 * @param argc Number of arguments, representing the count of command line
 * arguments passed in.
 * @param argv Array of pointers to character arrays, each character array
 * representing a single command line argument.
 * @return The function always returns 0, indicating normal program termination.
 */
int main(int argc, char *argv[]) {
  // Read the file specific by first argument
  if (argc < 2) {
    printUsage();
    return 1;
  }

  std::cout << "Parsing file: " << argv[1] << std::endl;

  std::ifstream stream(argv[1]);
  if (!stream) {
    std::cerr << "Failed to open input file." << std::endl;
    return 1;
  }

  antlr4::ANTLRInputStream input(stream);
  SuckCLexer lexer(&input);
  antlr4::CommonTokenStream tokens(&lexer);
  SuckCParser parser(&tokens);
  // Entry point is translationUnit
  SuckCParser::TranslationUnitContext *tree = parser.translationUnit();

  suckc::SourceGenerator visitor(&parser);

  visitor.visit(tree);

  return 0;
}
