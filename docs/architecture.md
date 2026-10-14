# SuckC Architecture

This document describes the internal architecture of the SuckC compiler —
the Go codebase that transpiles `*.suckc` sources into `.c` files. It is
written for contributors. For the language itself, see
[syntax.md](syntax.md); for user-facing docs see [index.md](index.md).

## 1. Overview

SuckC is a **source-to-source transpiler** (`*.suckc` → `.c`). It is not a
full compiler: there is no IR, no optimizer, no type checker in the C
sense, no runtime, and no ABI. Everything it does is syntax-level:
walk an ANTLR4 parse tree of the input, deduce just enough type
information to expand `auto`, and re-emit the original token text with a
small number of substitutions.

```
*.suckc ──► SuckC transpiler ──► .c ──► any vendor C compiler
                  │
            ANTLR4 grammar
            (internal/parser)
```

The roadmap is staged (README.md, "Roadmap"):

- **Phase 1 (current)** — full C/C++ grammar; output is *verbatim token
  round-trip* with only `auto` expansion, region-marker erasure, and
  `extern "C"` erasure. The emitted `.c` may still contain C++ constructs.
- **Phase 2** — trim the grammar to the target feature set; emit pure,
  portable C.
- **Later** — hardening against real vendor SDKs.

The architecture below reflects Phase 1: every design decision trades
future flexibility for what the current stage needs, and nothing more.

## 2. Design constraints

These constraints shape every component; they come from the project goals
(README.md, "Design Goals" and "What is SuckC?") and are load-bearing:

1. **Syntax-level only.** No symbol table, no semantic model persisted
   across stages. Consequences: the output must round-trip byte-for-byte
   unless a rule says otherwise, and type information exists only
   transiently inside `internal/sema`.
2. **Dependency-free output.** The generated `.c`/`.h` needs no runtime,
   no headers written by SuckC, no ABI support.
3. **Buildable without the ANTLR toolchain.** The generated Go parser is
   committed (internal/parser/generate.go). Only `cobra` + Go stdlib are
   dependencies (go.mod).
4. **Diagnostics must name the original source.** Vendor headers arrive
   preprocessed by `gcc -E` with line markers; errors must point at
   `vendor.h:34`, not at the preprocessed stream (see §5.3).
5. **One input file per invocation.** `cobra.ExactArgs(1)`
   (cmd/suckc/main.go:22); there is no project model, no build system
   integration. Output is written next to the source.

## 3. Package layout

```
cmd/suckc/           CLI: cobra command, flag parsing, pipeline wiring, file I/O
internal/
  parser/            ANTLR4 grammar (.g4) + generated Go parser (committed)
                     + .antlr/ (stale Java tool output, not used by the Go build)
  frontend/          Parse-tree access layer: ParseResult, cdecl extraction,
                     C-scope regions and line markers
  sema/              Semantic analysis: auto deduction + three validators
  backend/           Code generation: token substitution emitter
```

Total non-generated Go: ~4,400 lines across 20 files (~7,900 including
tests). The generated parser (internal/parser) is ~60,700 lines and is
not discussed further here except for its entry points.

### Dependency direction

Dependencies flow strictly downward; there are no cycles:

```
cmd/suckc ──► frontend ──► parser
     │            │
     ├──► sema ───┤        (sema reads *frontend.ParseResult; never calls frontend)
     └──► backend ┘        (backend reads *frontend.ParseResult + a map; never
                           imports sema — the contract is map[int]string)
```

Note the two deliberate asymmetries:

- `backend` does **not** import `sema`. The interface between them is a
  plain `map[int]string` (token index → replacement text), so backend
  stays independent of how types were deduced.
- `sema` does not own the parse tree either; it too consumes
  `*frontend.ParseResult`. `frontend` is the single source of parse-tree
  access for all downstream stages.

## 4. The compilation pipeline

```
frontend.ParseFile(file)                     cmd/suckc/main.go:34
   └─► *frontend.ParseResult                 (see §5.3; the sole handoff type)

   [ -d / --debug: dump parse tree, return ] cmd/suckc/main.go:39-42

sema.Analyze(result, sema.Options{...})      cmd/suckc/main.go:52
   └─► sema.Subs  (map[int]string)            token idx → replacement text

backend.Generate(result, subs)               cmd/suckc/main.go:57
   └─► string                                the generated .c text

os.WriteFile(<basename>.c, src, 0o644)       cmd/suckc/main.go:62-63
```

Stage order inside the CLI (cmd/suckc/main.go, `RunE`):

1. **Input check** (main.go:24-31) — `os.Stat`; missing file *or*
   directory → `file not found: <path>`.
2. **Parse** (main.go:34) — any frontend error aborts as
   `parse <path>: <err>`.
3. **Debug short-circuit** (main.go:39-42) — `-d` prints
   `Tree.ToStringTree(...)` and exits *before* flag validation.
4. **Flag validation** (main.go:44-50) — `--target-bits` must be 16, 32,
   or 64.
5. **Semantics** (main.go:52) — `sema.Analyze`.
6. **Generation** (main.go:57) — `backend.Generate`.
7. **Write output** (main.go:62) — `strings.TrimSuffix(file, ext) + ".c"`,
   mode `0o644`, next to the source.
8. **Report** (main.go:67) — `suckc: <in>: generated <out>`.

### CLI surface

| Flag | Type | Default | Meaning |
| --- | --- | --- | --- |
| `-d`, `--debug` | bool | `false` | Dump the ANTLR parse tree (`ToStringTree`), skip everything after parsing |
| `--target-bits` | int | `strconv.IntSize` (host) | Target integer width: 16, 32, or 64. Decides when a literal becomes `long` rather than `int` |

Plus cobra built-ins `--help` and `--version` (version string `0.1.0`,
main.go:21). Errors are returned from `RunE`; cobra prints them plus
usage to stderr and the process exits 1.

**Header (`.h`) emission is not implemented.** README's flow diagram
shows `.c / .h`, and `docs/syntax.md` mentions "the generated `.c`", but
the only file the CLI writes is the sibling `.c`. Treat `.h` as planned,
not present.

## 5. Components

### 5.1 `internal/parser` — grammar and generated code

Two grammar files, derived from the antlr4 C++14 grammar (MIT header
preserved in SuckCParser.g4:1-21):

- `SuckCLexer.g4` — 429 lines, ~147 token rules.
- `SuckCParser.g4` — 1089 lines, ~266 parser rules.

Regenerate after grammar edits (internal/parser/generate.go:12):

```bash
go generate ./internal/parser    # antlr4 -Dlanguage=Go -visitor -no-listener -o . <both .g4>
```

Three SuckC-specific changes relative to the upstream C++14 grammar:

1. **C-scope region tokens.** `CScopeBegin` / `CScopeEnd`
   (SuckCLexer.g4:31-33) match exactly
   `#pragma[ \t]+c_scope_begin[/\r]?`-shaped lines — trailing blanks and
   an optional CR are part of the marker so the lexer wins the
   longest-match contest against the generic `#pragma` rule (comment at
   SuckCLexer.g4:20-30). The parser rule
   `cImportRegion: CScopeBegin cCBlockItem* CScopeEnd`
   (SuckCParser.g4:331) is hooked into both `declaration`
   (SuckCParser.g4:271) and `statementSeq` (SuckCParser.g4:222), so a
   region can wrap declarations at file scope or statements inside a
   function body. An empty region produces a `Region` whose `StartTok` is
   greater than its `EndTok`.
2. **A C sub-grammar for region bodies** — 63 `cC*` rules
   (SuckCParser.g4:333-612, section comment 274-320) parsing vendor
   dialect C: multi-word type specifiers, bitfields, K&R-looking
   declarator lists, `__attribute__` / `__declspec` / `__asm__` /
   `typeof` spellings, `__packed`-style pointer modifiers, `extern "C"`
   (`cCLinkageSpecification`, SuckCParser.g4:518), anonymous struct
   members. Two decisions worth knowing (grammar comment 280-302):
   - Attributes are matched **positionally** — an attribute is only
     offered where a declarator cannot start — rather than by adding
     lexer rules, because new lexer rules would change tokenization for
     the whole language.
   - Function *definitions* are separated from *declarations*
     structurally (`cCDeclarator` vs `cCPlainDeclarator`, SuckCParser.g4:452-454),
     because ANTLR silently prefers the first viable alternative with no
     diagnostic: an order-based split would read `int C { int a; };` as a
     definition named `C` and say nothing.
3. **Upstream rules removed:** `templateDeclaration` /
   `explicitInstantiation` (SuckCParser.g4:263-265, commented out) and
   `namespaceDefinition` (SuckCParser.g4:267-268), reflecting the
   language changes documented in docs/syntax.md §2-3.

Deliberate grammar gaps, documented in the same comment block:

- `<<` / `>>` are **not lexer tokens** (only `<<=` / `>>=` are). SuckC
  expressions spell shifts as `Less Less` / `Greater Greater` token pairs
  (`shiftOperator`, SuckCParser.g4:156); `cCExpression` omits shift
  operators entirely because preprocessed headers spell shifts as macros
  that `cpp` has already expanded.
- A K&R list `f(a, b)` is parsed as two unnamed parameters whose types
  are the identifiers — a shape limitation the frontend's signature
  extraction mirrors (see §5.3).
- Preprocessor directives and vendor pragmas pass through as hidden
  tokens, so the backend can re-emit them verbatim.

The generated Go files are committed so the module builds without ANTLR.
`internal/parser/.antlr/` holds 8 Java-target artifacts from an earlier
default (Java) run — unused by the Go build.

### 5.2 Frontend — the parse-tree access layer (`internal/frontend`)

The frontend wraps ANTLR and everything that must be knowable about the
*text* (not the types): regions, line markers, and flat descriptions of
C-region declarations. It has three files.

**`parse.go` — parser wiring and the `ParseResult` type.**

```go
func ParseFile(path string) (*ParseResult, error)     // parse.go:40
func ParseSource(src string) (*ParseResult, error)    // parse.go:49
```

`parseSource` (parse.go:61-84):

1. `parser.NewSuckCLexer(antlr.NewInputStream(src))`
2. `lexer.RemoveErrorListeners()` + a custom `syntaxErrorListener`
   (shared by lexer and parser) — ANTLR's default stderr printing is
   removed so all diagnostics flow through the frontend
3. `antlr.NewCommonTokenStream(lexer, 0)`
4. `parser.NewSuckCParser(stream)` with the same listener swap
5. `tree := p.TranslationUnit()` — the only grammar entry rule used
6. Four collection passes over the tree/tokens:
   `collectRegions` → `collectLineMarkers` → `collectLinkageSpecs` →
   `collectCDecls`
7. Any error → return `nil` + `locatedError(firstError)` — **only the
   first error is ever reported**

**Architecturally notable: no ANTLR listener or visitor is used.**
Despite the committed listener/visitor files, every collector walks
`antlr.Tree` directly. (This also matches the `-no-listener` in the
generate directive; the committed listener `.go` files are stale
artifacts — see §9.)

Lexer errors are **fatal**, not warnings (rationale parse.go:56-60):
because emission is verbatim, silently dropping the offending characters
would copy damaged text into the output.

**`region.go` — C-scope regions, line markers, and the erase list.**

- `Region{StartTok, EndTok}` (region.go:15): the body *between* the two
  pragma markers (markers excluded), found by tree-walking
  `CImportRegionContext` (collectRegions, region.go:36).
  `InRegion(tokIdx)` (region.go:90) binary-searches the sorted list.
- `LineMarker{TokIdx, File, Line}` (region.go:21): the surviving
  `# N "file" ...` lines from `gcc -E`, matched by regex over
  **hidden-channel** tokens only (collectLineMarkers, region.go:66).
- `SourceAt(tokIdx)` (region.go:110) replays those markers to map a token
  index back to (file, line) in the *original* header — the mechanism
  behind §6's located diagnostics.
- `ToErase()` (region.go:182): the sorted token indices the backend
  blanks — both region markers (`StartTok-1`, `EndTok+1`; Keil/IAR reject
  unknown pragmas) plus every `extern "C"` wrapper token
  (`collectLinkageSpecs`, region.go:150). Line markers are deliberately
  **excluded** (rationale region.go:176-181: they carry accurate source
  positions for downstream compiler diagnostics).

**`cdecl.go` — flat descriptions of C-region declarations.**

```go
type CParam struct { TypeSpelling, Name string }        // cdecl.go:15
type CFuncSig struct {                                  // cdecl.go:26
    Name string; Params []CParam
    Variadic, Unprototyped bool
    TokIdx int
}
type CTypedef struct {                                  // cdecl.go:46
    Name, Base string; Stars int; FuncPtr bool
    Params []CParam; Variadic bool; TokIdx int
}
```

`collectCDecls` (cdecl.go:65) walks the tree but **stops at every
`CCExternalDeclarationContext`** — it never descends into function
bodies, so SuckC-level locals and K&R parameter declarations stay
invisible. Declarators are interpreted by one walk handling both parallel
hierarchies (`cCDeclarator` vs `cCPlainDeclarator`, declaratorTarget,
cdecl.go:151), which resolves `int (*fp)(int)` and
`int (*get(void))(int)` per C semantics. Distinctions the sema validators
depend on:

- `f()` → `Unprototyped` (unspecified parameter list, constrains
  nothing); lone unnamed `void` → empty parameter list.
- A pointer-returning function's signature uses its **own** parameters.
- Arrays of function pointers are not signatures; typedefs are never
  signatures.
- Type spellings are joined with single spaces (ANTLR `GetText` would
  produce `unsignedlong`).

### 5.3 `internal/sema` — semantic analysis

One entry point, seven exported identifiers (sema/sema.go):

```go
type Options struct { TargetBits int; FileName string }        // sema.go:15
type Subs map[int]string                                      // sema.go:29
type Error struct { File string; Line, Col int; Msg string }  // sema.go:32
type CallError struct { File, Line, Col, Name, Msg string }    // callcheck.go:25
type Analyzer struct{...}   // shared cross-pass state          // collect.go:11
type Type struct { Spelling string; Stars int }                 // type.go:13
func Analyze(res *frontend.ParseResult, opts Options) (Subs, error)  // sema.go:46
```

`Analyze` runs **five stages** (a more honest count than README's
"two-pass"):

```
pass1                collect.go:56     TU signatures: file-scope decls, enums,
                                       const initializers, queued file autos
   │
pass1b               file_auto.go:16   deduce file-scope autos, dependency-ordered
   │
pass2                walk.go:30        body walk: scopes, block autos,
                                       argument-shape checks
   │
scanForbiddenAutos   walk.go:345       every Auto token not consumed = error
   │
checkCCalls          callcheck.go:46   arity check vs C-region prototypes
```

**`pass1` (collect.go)** walks file scope, stopping at function
definitions (records name → return type; parameters and function-pointer
returns are skipped) and simple declarations. Declarations are routed:
`auto` → queued as pending (but `decltype(auto)` is deliberately not, so
the forbidden-scan reports it), enums → constant-name set, const-declared
→ initializer tree saved for the C89 constant gate, function-typed
declarator → signature.

**`pass1b` (file_auto.go)** deduces queued file-scope autos in **lazy
dependency order** with a state machine that doubles as cycle detection
(`circular dependency on 'p'`, file_auto.go:41), so a file-scope `auto`
may reference one declared later. Depth capped at 64.

**`pass2` (walk.go)** is an ANTLR `ParseTreeWalker` listener over the
tree, maintaining a scope stack (scope.go: `binding{Type, IsArray,
IsConst}` frames, innermost-first lookup, inner frame shadowing silently
wins). On entering a function it pushes a frame and declares parameters;
compound statements and `for`/`if`-condition declarations push frames
too; on exit it pops. Declarations inside bodies run the shared
`deduceAuto` (deduce.go:25) — the single deduction path for both scopes:

1. require an initializer (`auto i;` is an error)
2. reject function-pointer declarators
3. **file scope only:** the C89 3.4 constant gate (constcheck.go)
4. type the initializer expression (`typeOf`, typeof.go:17)
5. subtract the declarator's pointer depth
6. emit `Subs[autoTokenIndex] = finalType.Text()` and record the binding

The emitted type deliberately **excludes** the declarator's stars:
`auto *p = &a;` substitutes `int`, the `*` already being in the source.
The constant gate accepts only C89 3.4 forms — literals, enum constants,
`sizeof`, addresses, recursive reads of `const` variables — and rejects
calls, assignments, commas, casts, and `++`/`--`. Errors are wrapped as
`*sema.Error` with the `auto` token's position (sema.go:39-41).

`typeOf` (typeof.go) recurses over the ANTLR expression grammar:
literals via shape parsing (literal.go), names via scope → file vars,
`?:` via commonType, arithmetic chains folded with the C89 6.3.1.8
subset (`usualArith`, type.go:150) using the target-width limit tables
(`intLimits`/`longLimits`, type.go:55/66, driven by `Options.TargetBits`
— this is where `--target-bits 16` makes `40000` a `long` while 32/64
make it an `int`), `sizeof` → `size_t`, function calls → the return type
collected in pass1.

**The three validators:**

| Validator | When | What it checks | Deliberate bias |
| --- | --- | --- | --- |
| `constcheck` (constcheck.go) | inside `deduceAuto`, file scope only | initializer is a C89 constant expression | reject |
| `argcheck` (argcheck.go) | inside pass2 (needs the live scope stack) | call arguments vs C-region prototypes, **by shape only** | err toward silence |
| `callcheck` (callcheck.go) | after pass2 (needs whole-TU shadow names) | arity vs C-region prototypes | err toward silence |

Shape checking (argcheck.go) compares only what C itself never converts:
integer constant `0` is a null pointer constant for any pointer
parameter; `void` matches only `void *`; pointer star counts must match;
tagged structs match only the identical tag. Pointee types are
deliberately **not** compared (`char *` satisfies `const char *`), and
anything the checker cannot type is skipped, never reported — "missing a
bad call is recoverable; rejecting a good one is not."

The grammar has no call rule, so call sites are recognized by shape:
`callOf` (callcheck.go:359) handles `name(args);` as a postfix
expression; `declarationShapedCall` (callcheck.go:228) handles the
statement-level shape that parses as a typeless declaration; both ignore
names declared in SuckC code (redeclaring a vendor prototype in SuckC
keeps working). Conflicting prototypes for one function are an error;
unprototyped (`f()`) declarations are skipped. All spellings are
normalized before comparison (canon.go:22): `unsigned long`,
`long unsigned int`, and `unsigned` compare equal, while `char`,
`signed char`, and `unsigned char` stay distinct; region typedefs are
expanded to what they alias, which is why a `size_t` parameter matches
an `unsigned long` argument.

**Shape blindness the checker accepts** (docs/syntax.md "Shapes the
checker cannot see"): `f(s);` with a single identifier argument is
grammar-ambiguous (call vs declaration) and SuckC has no typedef table,
so it reads as a declaration and the argument is invisible; `f(&flags);`
is lost the same way. Fixing this needs a declarator-list pass or a
typedef table — a noted future stage, not a bug to paper over here.

### 5.4 `internal/backend` — code generation

One exported function (backend/generate.go:33):

```go
func Generate(res *frontend.ParseResult, subs map[int]string) (string, error)
```

The emitter (generate.go:65) is an ANTLR listener with a single
overridden method, `VisitTerminal` (generate.go:84):

1. Build the effective substitution map: caller `subs` **plus**
   `res.ToErase()` entries mapped to `""` (generate.go:40-46).
2. Walk the tree with `antlr.NewParseTreeWalker().Walk(e, res.Tree)`.
3. In `VisitTerminal`: skip EOF; backfill every hidden-channel token
   since the previous terminal through `write()` — this is what
   preserves comments, blank lines, and preprocessor lines byte-for-byte;
   emit the substitution if there is one (regardless of token channel —
   that is how visible pragma markers vanish while hidden gcc line
   markers survive), else the token's verbatim text.
4. Flush trailing hidden tokens, return the buffer.

There is no AST, no pretty-printer, and no formatting stage; the output
is verbatim token text with substitutions applied. Tests pin this
round-trip property directly: `TestGenerateRoundTrip` asserts generated
output **equals input** for a source with no substitutions
(generate_test.go:9).

## 6. Diagnostics model

Two error types with different formats, because `Error`'s wording is
auto-deduction-specific:

- `*sema.Error` — `file:line:col: cannot deduce type for 'auto': <reason>`
  (sema.go:39-41). Reasons include `missing initializer`,
  `insufficient pointer depth`, `unknown identifier 'x'`,
  `not a constant expression`, `wide character literal not supported`.
- `*sema.CallError` — generic `file:line:col: <message>`, e.g.
  `too many arguments in call to 'f': 2 given, C region declares 1
  (declared at hdr.h:4)` and
  `incompatible argument 1 in call to 'put': passing 'char *' where
  'int' is expected (declared at hdr.h:4)`.

Positions come from three mechanisms, in order of specificity:
`frontend.ParseResult.SourceAt` maps a token index through surviving
`gcc -E` line markers so C-region diagnostics name the original header
and line; diagnostics raised by suckc itself use the same table;
frontend syntax errors fall back to `line %d:%d` when no markers exist
(locatedError, parse.go:98-112).

Every stage aborts on the first error — there is no error recovery and
no accumulation anywhere in the pipeline. Two known asymmetries, stated
here rather than hidden: the circular-dependency error from pass1b is a
plain `fmt.Errorf` (no `*Error` wrapper, no position), and syntax errors
report only `el.errors[0]`.

## 7. Testing strategy

Each package pins the behavior its layer owns:

| Package | What is pinned |
| --- | --- |
| `cmd/suckc` | end-to-end: generated `.c` byte-identical to input (Phase 1 identity), `-d` tree dump, flag validation, `auto i = 12` → `int i = 12`, `--target-bits 16` vs 32 (`long` vs `int` for `40000`), file-scope pointer deduction |
| `internal/frontend` | valid/invalid parses, first-error reporting, `class` lexes as a plain identifier, function definitions require a parameter list; signature/typedef extraction table (variadic, unprototyped, func-ptr, K&R); region bounds/nesting rejection; real `gcc -E` fixtures for line markers |
| `internal/sema` | deduction table across declarator forms and scopes, target-width effects, constant-gate accept/reject, arity and argument-shape matrices, canonicalization equivalences, scope shadowing |
| `internal/backend` | byte round-trip, comment/whitespace preservation, substitution application, region-marker and `extern "C"` erasure, line-marker survival |

Fixtures are deliberately *real* toolchain output: `internal/frontend/testdata/region_line_markers.c` is genuine `gcc -E -x c` output (with regeneration commands documented in region_test.go), and `testdata/lex_err_region.c` pins that a lexer error inside a region is reported as `bad_vendor.h:3:11`.

## 8. Known limitations of the current architecture

Surfaced per project Rule 12 (fail loud); these are properties of
Phase 1, not surprises to be discovered later:

1. **No `.h` emission.** The CLI writes exactly one sibling `.c`
   (main.go:62-63); README's "`.c` / `.h`" describes the target, not the
   current state.
2. **No AST.** All stages share the ANTLR parse tree; there is no
   normalized representation to transform against. Phase 2 (trimming to
   the target language) will need one or a bigger emitter.
3. **Five stages, not two passes.** README describes the sema pipeline
   as "two-pass"; pass1b (file-scope auto deduction with dependency
   ordering) is really a third pass, and the forbidden-scan and
   call-check are post-passes.
4. **Generated-code drift:** the generate directive says `-no-listener`,
   but `suckc_parser.go` (and separate committed listener files) were
   produced by a listener-producing run. Regenerating per the directive
   would drop those files. Harmless today (unused) but worth aligning.
5. **Code comments cite a "spec" and a "Stage 2b"** (e.g. sema/canon.go,
   SuckCParser.g4:310) that are not in the repository; the numbering
   refers to design work outside `docs/`. Do not cite them as repo
   artifacts.
6. **Region grammar gaps are intentional** (SuckCParser.g4:305-311): no
   shift operators inside C regions (macros are pre-expanded anyway),
   K&R lists read as unnamed parameters, single-identifier calls are
   invisible to the checks.

## 9. Build, run, test, regenerate

```bash
go build ./...
go run ./cmd/suckc SuckC.suckc      # transpile (writes SuckC.c)
go run ./cmd/suckc -d SuckC.suckc   # dump the parse tree instead
go test ./...
go generate ./internal/parser       # regenerate the Go parser from .g4
```

## 10. Related documents

- [index.md](index.md) — documentation index
- [syntax.md](syntax.md) — SuckC language syntax rules (the user-facing
  specification these components implement; must stay in sync with
  grammar changes)
- [naming-conventions.md](naming-conventions.md) — identifier naming
- [../README.md](../README.md) — project goals, roadmap, current status
