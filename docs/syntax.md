# SuckC Syntax Rules

This document is the reference for the SuckC language syntax, written for
users of the language. It does **not** restate the full C89 grammar. Instead
it records only the rules that **differ from C89** — the extensions and
enhancements SuckC adds on top of its C89 base.

## 1. Base: C89

SuckC syntax is based on **C89 (ANSI C)**. Anything not explicitly overridden
or extended in this document follows the C89 grammar.

### Maintenance rule

This document must stay in sync with the compiler. Whenever code changes
affect the language syntax — adding, removing, or modifying a rule —
this document must be updated in the same change. See `AGENTS.md`.