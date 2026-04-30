# SuckC Compiler

## Build & Run

```sh
cmake -B build -DCMAKE_BUILD_TYPE=Debug
cmake --build build
./build/suckc SuckC.suckc
./build/suckc -d SuckC.suckc          # debug: print AST tree
GLOG_logtostderr=1 ./build/suckc ...  # verbose glog output
```

## Architecture

- **Entry point**: `src/main.cpp` -- wxWidgets `wxInitializer` + `wxCmdLineParser` for CLI
- **Grammar**: `src/SuckCLexer.g4` + `src/SuckCParser.g4` (C++14/C11-based, via ANTLR4)
- **Visitor**: `suckc::SourceGenerator` extends generated `SuckCParserBaseVisitor` to walk parse tree and build AST
- **AST nodes**: `src/ast/` -- `Node` base class with `Variable`, `Function`, `TypeDeclaration`, `Alias`, `Expression`, `Literal`, `Struct`, `Specifier`, `Attribute`
- **Scopes**: `Scope` (4 types: Global, Function, Block, Statement), managed as a stack by `SourceContext`
- **Singleton**: `World::getInstance()` holds the parser reference and debug flag

## Dependencies

ANTLR4 runtime, wxWidgets (base, core), glog, Abseil. All system-installed; no package manager config.

ANTLR4 binary is found via PATH by `cmake/FindAntlr4Ex.cmake` (uses `cmake/fake-java.sh` shim). Lexer + parser sources are auto-generated during build into `build/antlr4_generated_src/`.

## Code Conventions

- PIMPL idiom via `SUCKC_OBJECT_DECL` / `SUCKC_OBJECT_IMPL` macros (see `SuckTypes.h`)
- 2-space indent, 80 column limit (`.clang-format`)
- C++20 standard
- UUID-style include guards (`_9M_UUID`)
- No tests, no CI, no linter scripts

## Key Tips

- Regenerating ANTLR sources: just rebuild (`cmake --build build` handles it)
- Grammar is based on the antlr4 C++14 grammar (MIT License, by Camilo Sanchez / Martin Mirchev)
- Include dirs needed at edit time: `src/`, `build/antlr4_generated_src/suckc_lexer/`, `build/antlr4_generated_src/suckc_parser/`
