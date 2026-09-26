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

## 3. `auto` type deduction

SuckC supports C++-`auto`-style variable declarations at **block scope**
(including the init part of a `for` statement). The transpiler deduces the
type from the initializer and writes the concrete type into the generated
`.c` file:

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
- `auto` is **rejected** at file scope, in function return types, in
  function parameters, in `decltype`, in class members, and in range-for.
- Any deduction failure is a hard error reported as
  `file:line:col: cannot deduce type for 'auto': …` — there is no silent
  fallback.

The target integer width (which decides when a literal becomes `long`
versus `int`) is selected with the compiler flag `--target-bits 16|32|64`;
the default matches the host.

## 4. Single-declarator declarations

Unlike C89, a variable declaration may declare **at most one** variable:
`int a, b;` is not accepted. Declare them separately:

```c
// SuckC (invalid):
// int a, b;

// SuckC:
int a;
int b;
```

`for`-init declarations follow the same rule. (Class member lists are
unaffected.)

### Maintenance rule

This document must stay in sync with the compiler. Whenever code changes
affect the language syntax — adding, removing, or modifying a rule —
this document must be updated in the same change. See `AGENTS.md`.