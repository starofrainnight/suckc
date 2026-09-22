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

### Maintenance rule

This document must stay in sync with the compiler. Whenever code changes
affect the language syntax — adding, removing, or modifying a rule —
this document must be updated in the same change. See `AGENTS.md`.