

// For compilers that support precompilation, includes "wx/wx.h".
#include "wx/wxprec.h"

// for all others, include the necessary headers (this file is usually all you
// need because it includes almost all "standard" wxWidgets headers)
#ifndef WX_PRECOMP
#include "wx/wx.h"
#endif

#include "SourceGenerator.h"
#include "SuckCLexer.h"
#include "SuckCParser.h"
#include "World.h"
#include <antlr4-runtime.h>
#include <fstream>
#include <glog/logging.h>
#include <iostream>
#include <wx/app.h>
#include <wx/cmdline.h>

static const wxCmdLineEntryDesc cmdLineDesc[] = {
    {wxCMD_LINE_SWITCH, "h", "help", "show this help message",
     wxCMD_LINE_VAL_NONE, wxCMD_LINE_OPTION_HELP},
    {wxCMD_LINE_SWITCH, "d", "debug", "debug the ast tree", wxCMD_LINE_VAL_NONE,
     0},
    {wxCMD_LINE_PARAM, nullptr, nullptr, "SRC_FILE", wxCMD_LINE_VAL_STRING, 0},
    wxCMD_LINE_DESC_END,
};

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
  wxApp::CheckBuildOptions(WX_BUILD_OPTIONS_SIGNATURE, "program");

  wxInitializer initializer;
  if (!initializer) {
    fprintf(stderr, "Failed to initialize the wxWidgets library, aborting.");
    return -1;
  }

  bool isDebug = false;

  google::InitGoogleLogging(argv[0]);
  FLAGS_logtostderr = true;

  wxCmdLineParser cmdLineParser(cmdLineDesc, argc, argv);
  auto ret = cmdLineParser.Parse();

  if (-1 == ret) {
    return -2;
  }

  if (argc == 1) {
    // If there were no command-line options supplied, emit a message
    // otherwise it's not obvious that the sample ran successfully
    return -3;
  }

  auto fpath = cmdLineParser.GetParam(0);

  std::ifstream stream(fpath.c_str().AsChar());
  if (!stream) {
    LOG(ERROR) << "Failed to open input file!";
    return -4;
  }

  antlr4::ANTLRInputStream input(stream);
  SuckCLexer lexer(&input);
  antlr4::CommonTokenStream tokens(&lexer);
  SuckCParser parser(&tokens);
  auto world = suckc::World::getInstance();

  world->registerParser(&parser);
  world->setEnabledDebug(cmdLineParser.Found("debug"));

  // Entry point is translationUnit
  SuckCParser::TranslationUnitContext *tree = parser.translationUnit();

  suckc::SourceGenerator visitor(&parser);

  visitor.visit(tree);

  return 0;
}
