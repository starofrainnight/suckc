# SuckC

SuckC is a C-based language with C++-style modern syntax, and a transpiler that
translates `*.suckc` sources into plain C (`.c` / `.h`) files.

It aims to be a universal, cross-compiler, cross-chip C-like language for
embedded and microcontroller development — one codebase that runs across
vendors, SDKs, and MCU toolchains.

## What is SuckC?

SuckC starts from the C language and enriches it with carefully selected
C++-style features — adding and trimming syntax to keep the C-like simplicity
while gaining part of the expressiveness of modern C++.

The compiler is **syntax-level only**: it performs pure syntax and
syntax-sugar transformation, and generates **dependency-free** `.c` and `.h`
output. There is no runtime, no ABI, no library requirement.

## Motivation

Most embedded/MCU vendors ship C89 compilers that:

- support only C89 (or a partial superset of it),
- differ from each other in syntax and extensions,
- vary widely in quality,
- are **never updated**.

This fragmentation locks the codebase to the worst common denominator. SuckC
breaks the lock: because the output is plain C, it works with **any** vendor
compiler and SDK — while the source language itself can keep evolving,
independent of vendor support.

## Design Goals

- **Full C compatibility** — SuckC code interoperates with existing C code;
  the transpiler only performs syntax-level and syntax-sugar conversion.
- **Cross-compiler, cross-chip** — generated `.c` / `.h` files compile with
  every mainstream embedded C compiler (C89 and up).
- **One codebase, many targets** — write once, build for any MCU / embedded
  platform.
- **Evolving syntax** — the language iterates at the transpiler level; no
  vendor cooperation required.
- **Simple and C-like** — C++-style ergonomics without C++ complexity.

## How It Works

```
*.suckc  ──►  SuckC transpiler  ──►  .c / .h  ──►  any vendor C compiler
                    ▲
               ANTLR4 grammar
```

The frontend is generated from ANTLR4 grammar files; the backend emits
self-contained C sources.

## Roadmap

### Phase 1 — Full C/C++ grammar (current)

- Implement the complete C/C++ syntax based on the ANTLR4 grammar.
- Transpile any `*.suckc` into `.c` and `.h`.
- Note: in this phase the emitted `.c` may still contain C++ constructs
  (whether or not that is reasonable). This is intentional — later phases
  refine the output into pure C.

### Phase 2 — Trim to the target language

- Gradually trim the grammar down to the desired SuckC feature set —
  C core plus the selected C++-style features.
- Ensure the output is pure C, portable across embedded compilers.

### Later — Hardening

- Embedded-oriented validation against real vendor SDKs and compilers.

## Current Status

- **Implementation**: originally written in C++ (ANTLR4 + C++20); the C++
  codebase has been removed and the Go rewrite is underway. The CLI scaffold
  (cobra) is complete; transpilation is not implemented yet.
- **Frontend**: ANTLR4 grammar — `src/SuckCLexer.g4`, `src/SuckCParser.g4`
  (based on the antlr4 C++14 grammar), preserved for future frontend
  integration.
- **Backend**: planned — AST built from the parse tree; scopes tracked via
  `SourceContext`; code emitted by the source generator.

## Build & Run

```bash
go build ./...
go run ./cmd/suckc SuckC.suckc
```

Debug (stub this phase):

```bash
go run ./cmd/suckc -d SuckC.suckc
```

Tests:

```bash
go test ./...
```