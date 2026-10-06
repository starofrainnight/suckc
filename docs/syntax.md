# SuckC Syntax Rules

This document is the reference for the SuckC language syntax, written for
users of the language. It does **not** restate the full C89 grammar. Instead
it records only the rules that **differ from C89** — the extensions and
enhancements SuckC adds on top of its C89 base.

## 1. Base: C89

SuckC syntax is based on **C89 (ANSI C)**. Anything not explicitly overridden
or extended in this document follows the C89 grammar.

## 2. No references

C++ reference syntax (`&` / `&&` in declarations, ref-qualifiers on member
functions, by-reference lambda captures) is **not** supported. Use pointers
instead:

```c
// SuckC (invalid):
// void f(int& x) {}

// SuckC:
void f(int *x) {}
```

The `&` operator is still available in its C89 meanings: address-of and
bitwise AND.

## 3. No `class` keyword

SuckC has only the C `struct` keyword — the C++ `class` keyword is not part
of the language. Declare types with `struct`:

```c
// SuckC (invalid):
// class Uart { void reset(); };

// SuckC:
struct Uart { void reset(); };
```

Nothing is lost, because `struct` and `class` mean the same thing in C++.
`class` is not reserved either, so it is an ordinary identifier, exactly as in
C89.

Template type parameters are introduced by `typename`, which also covers the
`template <…>` form that `class` would otherwise have provided:

```c
// SuckC (invalid):
// template <class T> struct Box { T value; };

// SuckC:
template <typename T> struct Box { T value; };
```

## 4. `auto` type deduction

SuckC supports C++-`auto`-style variable declarations at **file scope** and
at **block scope** (including the init part of a `for` statement). The
transpiler deduces the type from the initializer and writes the concrete
type into the generated `.c` file:

```c
// SuckC:
auto i = 12;

// generated C:
int i = 12;
```

Rules:

- An initializer is required: `auto i;` is an error.
- Supported declarator forms: plain identifier (`auto p = e`), pointer
  (`auto *p = e`), and array (`auto a[] = {…}`, `auto a[10] = {…}`).
- Deduction covers literals, variable references, arithmetic and bitwise
  operators, comparisons, the ternary operator, address-of / dereference,
  array subscript, `sizeof`, and calls to functions declared in the same
  file (return type only — arguments are not type-checked).
- `auto` is **rejected** in function return types, in function parameters,
  in `decltype`, in struct members, and in range-for.
- A file-scope `auto` initializer must be a C89 constant expression
  (literals, `sizeof`, enum constants, addresses, and reads of `const`
  variables). Anything else is reported as
  `cannot deduce type for 'auto': not a constant expression`.
- A file-scope `auto` may reference a variable declared later in the file,
  but a reference cycle is reported as `circular dependency on 'NAME'`.
- Any deduction failure is a hard error reported as
  `file:line:col: cannot deduce type for 'auto': …` — there is no silent
  fallback.

The target integer width (which decides when a literal becomes `long`
versus `int`) is selected with the compiler flag `--target-bits 16|32|64`;
the default matches the host.

## 5. Single-declarator declarations

Unlike C89, a variable declaration may declare **at most one** variable:
`int a, b;` is not accepted. Declare them separately:

```c
// SuckC (invalid):
// int a, b;

// SuckC:
int a;
int b;
```

`for`-init declarations follow the same rule. (Struct member lists are
unaffected.)

## 6. Function definitions need a parameter list

A function definition must declare its parameters, even when there are none.
The braces alone are not enough to introduce a definition:

```c
// SuckC (invalid):
// int C { int a; };

// SuckC:
int C(void) { int a; return a; }
```

## 7. C-scope regions

Two pragmas delimit a region whose body is parsed with the C dialect rules
instead of the SuckC rules. Regions may appear at file scope and inside a
function body:

```c
#pragma c_scope_begin
#include <stdio.h>
void hello(void) { printf("hi\n"); }
#pragma c_scope_end

void myMain(void) {
    #pragma c_scope_begin
    hello();
    #pragma c_scope_end
    hello();
}
```

Marker rules:

- Exactly `#pragma`, one or more spaces or tabs, then `c_scope_begin` /
  `c_scope_end`. Trailing spaces, tabs and a CR (CRLF line endings) are part
  of the marker. Any *other* trailing content (a word, a comment) makes the
  line an ordinary preprocessor directive and the region stays unopened.
- Nested regions are a syntax error.
- An unterminated region is a syntax error.

The body is opaque to SuckC analysis: it is never treated as SuckC
declarations, and the same tokens outside a region keep their SuckC meaning.
Because the body is C rather than SuckC, `auto` and `const` inside a region
are C constructs — they are neither rejected nor used to deduce a type.

### What the body accepts

Declarations, in the C dialect: multi-word and vendor type specifiers,
pointer/array/function declarators, comma-separated init-declarator lists,
bitfields, `typedef`, enumerations, struct/union/enum definitions including
anonymous members, function definitions with K&R parameter lists, and
`extern "C" { ... }`. Function bodies use a C expression and statement
subgrammar. Vendor extensions pass through: `__attribute__` (both
spellings), `__declspec`, `__typeof__`, `typeof`, `__asm__` labels,
`__extension__`, `restrict`, and pointer modifiers such as `__packed`,
`__sptr`, `__w64`, `__ptr32`.

Constructs that are not C are rejected, not passed through: `namespace`,
`class`, and templates are syntax errors inside a region, exactly as they are
outside one. A declaration the C dialect does not cover is a syntax error —
the body is never silently reinterpreted as an opaque token run.

Two limitations are deliberate and current:

- The shift operators `<<` and `>>` are not available inside a region.
- A name spelled with a SuckC keyword token can be declared — glibc's
  `typedef int wchar_t;` is the case that motivates it — but not yet used as
  a type, because doing so needs a table of typedef names.

### Translation rules

- Both markers are dropped from the generated `.c`, so a downstream C
  compiler never sees them. So are the four wrapper tokens of an
  `extern "C" { ... }` specification, which is not valid C89. Only whole
  marker lines vanish; the newline that ended one stays, leaving a blank
  line. Nothing else in the body is rewritten — vendor extensions such as
  `__attribute__` and `#pragma pack(...)` pass through untouched, because
  removing them could change struct layout.
- A preprocessor line marker left by `gcc -E` (`# 42 "vendor.h" 1`) is
  passed through unchanged. It already names the file and line each
  surviving declaration came from, so it is more accurate than anything
  reachable by reconstruction, and the compiler that emitted it also
  consumes it: a diagnostic raised while compiling the generated `.c`
  is reported as `vendor.h:34`, not as a line of the `.c`. Diagnostics
  raised by suckc itself use the same table. Erasing a marker would only
  lose that, since the tokens the translation does drop leave their
  newline behind and so shift no line numbers.
- A lexical error aborts translation, inside a region as much as outside it.
  The lexer drops the offending characters, and verbatim emission would
  otherwise copy the damaged text into the `.c` without a word.

#### Dropping `extern "C"`

An `extern "C"` specification is removed from the generated `.c` rather than
emitted into it. The brace form loses all four wrapper tokens — `extern`,
the string literal, `{`, `}` — and the single-declaration form
`extern "C" int x;` loses the first two. The declarations between the
braces are kept. It applies whether the specification was written in SuckC
code or inside a region.

This is known, deliberate behaviour:

- `extern "C"` is not valid C89. Left in place, `gcc -std=c89` rejects it
  outright: `error: expected identifier or '(' before string constant`.
- Removing it changes no meaning. A declaration at file scope already has C
  linkage in C89, so `extern "C" { int x; }` and `int x;` declare exactly
  the same thing.
- Block scope is rejected by the parser before the difference could matter.
  `void f(void) { extern "C" int x; }` is a syntax error, so the wrapper is
  never dropped where it would silently change storage duration — removing
  `extern` there would leave `int x;`, an automatic local variable instead
  of an external declaration.
- Under `gcc -E -x c`, a header's own `extern "C"` has usually already gone:
  preprocessing as C makes the preprocessor drop `#ifdef __cplusplus`
  blocks. What reaches suckc is a specification written literally in the
  `.suckc`.

### Cross-region calls

Declarations inside a region are read, not just passed through: every
function prototype becomes a signature and every `typedef` a type name.
A call written in SuckC code to a function that only a region declares is
checked against that prototype's parameter count.

- `void f(void);` declares zero parameters. `int f();` declares an
  unspecified parameter list and therefore constrains nothing — a call to it
  is always accepted, whatever the argument count.
- A declaration with no prototype never conflicts with another declaration of
  the same function. Two prototypes that disagree on parameter count or on
  variadic-ness are an error.
- The diagnostic names the original header and line, not the preprocessed
  file: `too few arguments in call to 'vendor_init': 0 given, C region
  declares 1 (declared at vendor.h:34)`.
- A name you declare in SuckC code is not checked, even if a region also
  declares it. Redeclaring a vendor prototype in SuckC keeps working.

Argument types are checked by shape, not by equality. C converts freely
between arithmetic types, so requiring an exact type match would reject
nearly every real call. What is actually incompatible:

- A pointer is never compatible with an arithmetic type, and an array
  argument decays to a pointer first.
- An integer constant of value zero is a null pointer constant and satisfies
  any pointer parameter. `0x0` and `0u` qualify; `0xFF` does not.
- `void` is compatible only with `void *`.
- A tagged `struct`, `union` or `enum` only with the identical tag.
- Two pointers are always accepted. The pointee is deliberately not
  compared, so `char *` satisfies `const char *` and `void *`.
- An argument the checker cannot type — a compound expression, a call
  result, a name it cannot resolve — is skipped, never reported.

The check errs toward silence. Missing a bad call is recoverable; rejecting
a good one is not.

Type spellings are canonicalized before any of this, so equivalent spellings
compare equal: `unsigned long` is `long unsigned int` is `unsigned`, `int`
may be spelled with any legal length and signedness word, and `char`,
`signed char` and `unsigned char` stay three distinct types. A typedef name
in a region resolves to the type it aliases, which is why a `size_t`
parameter matches an `unsigned long` argument even though glibc defines it
as `long unsigned int`.

#### Shapes the checker cannot see

A statement-level call with a single identifier argument is ambiguous in the
grammar, and C resolves it by name lookup rather than by syntax:
`vendor_init(flags);` may mean "call `vendor_init` with `flags`" or "declare
`flags` of type `vendor_init`". SuckC has no typedef table at this point, so
a single-identifier argument is read as a declaration, and the argument is
invisible to both checks. A declaration is only read as a call when the type
name is a function some region actually declares; otherwise it stays a
declaration.

`vendor_init(&flags);` is lost the same way — the `&` reads as a reference
declarator. And `vendor_init(a, b);` with two identifier arguments does not
parse at all, because a SuckC declaration holds a single declarator. Making
these work needs a comma-separated declarator list in SuckC declarations, or
a pass that resolves the ambiguity from the collected signatures before
parsing statements.

A `typedef` name that happens to be spelled like a SuckC keyword —
`wchar_t`, `bool`, `constexpr` — can be declared inside a region but not used
as a type there.

### Maintenance rule

This document must stay in sync with the compiler. Whenever code changes
affect the language syntax — adding, removing, or modifying a rule —
this document must be updated in the same change. See `AGENTS.md`.