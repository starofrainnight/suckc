# SuckC Naming Conventions

This document is the reference for naming identifiers in SuckC, written for
users of the language. It defines **one** naming style for SuckC and applies it
to every kind of identifier, so that code written by different people — or by
different vendors' SDKs — reads the same way.

Naming is a **convention layer**, not part of the grammar: the compiler accepts
any identifier you write.

## 1. Base: lowerCamelCase

Identifiers use **lowerCamelCase**:

```c
int retryCount;
void startUart();
```

## 2. No visibility by capitalization

Visibility is expressed only by the `public` / `private` / `protected` access
specifiers of a class, or not at all at file scope. So capitalization is free
to carry meaning inside the word, and nothing else:

```c
class Uart {
public:
  void reset();        // public because of `public:`, not because of `R`
private:
  int errorCount_;     // private because of `private:`; `_` marks membership
};
```

## 3. Package names

Package names are **short, lowercase, singular**, and carry no underscores:

`driver`, `uart`, `crc`, `parser`, `ringbuffer` — not `Drivers`, `UART`,
`ring_buffer`, `parsers`.

Keep them short enough to type in one breath; they qualify every name they
touch.

## 4. Acronyms are ordinary words

Treat an acronym as one ordinary word and camelCase it: `Http`, `Url`, `Io`,
`Pwm`, `Spi`, `Adc` — never `HTTP`, `URL`, `IO`. The same word shape applies at
both ends of the name:

```c
// SuckC (invalid):
// void parseURL(char *url);
// int HTTPport;
// HttpServer server;

// SuckC:
void parseUrl(char *url);
int httpPort;
HttpServer server;
```

## 5. Short names win

Prefer the shortest name that stays unambiguous in its scope:

- One or two letters are fine in a small scope: `i`, `n`, `buf`, `len`.
- Inside a method, `count` beats `retryCount` unless the shorter name hides
  something.
- Avoid names that restate the type: `int byteCount` (not `int byteCountInt`),
  `bool isEnabled` is acceptable because the `is` carries meaning.

## 6. Types

Class, struct, enum, and typedef / alias names start with a capital and use
PascalCase:

```c
class Uart;
class RingBuffer;
enum PinState;
typedef unsigned int ByteCount;
```

## 7. Functions and methods

lowerCamelCase, starting with the verb:

```c
void start();
int readByte();
bool parseFrame(int *data, int length);
```

## 8. Member variables

Member variables take a **trailing underscore**:

```c
class Uart {
  int baud_;
  int errorCount_;
};
```

The underscore is the only marker of membership, because capitalization is
already spent on the word itself (rule 2) and on the type convention (rule 6).
It marks *membership*, not visibility — a `_` member can be `public`.

## 9. Local variables and parameters

lowerCamelCase, no underscore:

```c
void setTimeout(int timeoutMs) {
  int elapsed = 0;
}
```

## 10. Constants

Named constants — `const` variables, enum constants, and macro definitions —
use **camelCase**, not `SHOUTING_CASE`:

```c
const int MaxRetryCount = 3;
enum { PinCount = 8 };
#define SectorSize 512
```

Spell the constant out rather than abbreviating it; `SectorSize` reads better
than `SecSz` and costs nothing at compile time.

## 11. Interfaces

A type that exists only to provide **one** operation is named after that
operation with an `-er` suffix: `Reader`, `Writer`, `Closer`, `Encoder`,
`Comparer`.

SuckC has no `interface` keyword; the rule applies to any type whose whole
purpose is a single operation, typically an abstract class.

## 12. Errors

- Error **values** — a named condition callers compare against — take an `Err`
  prefix: `ErrTimeout`, `ErrOverflow`, `ErrNotFound`.
- Error **types** take an `Error` suffix: `TimeoutError`, `ParseError`,
  `IoError`.

The prefix/suffix split mirrors the type/value split: `ErrTimeout` is a value
you compare, `TimeoutError` is a type you catch.

## 13. Getters and setters

A getter is the bare field name; the setter adds `set`. The pair is
semi-symmetric on purpose — the getter reads as a property, the setter reads as
an action:

```c
class Uart {
  int baud_;

public:
  int baud();
  void setBaud(int baud);
};
```

Only add a getter/setter when the field is not simply public data. Do not name
them `getBaud()` — the `get` prefix adds nothing.

## 14. File names

File names are lowercase, short, and underscore-free, matching the package they
belong to: `uart.suckc`, `ringbuffer.suckc`, `crc.suckc`.

## Summary

| Kind | Style | Example |
| --- | --- | --- |
| Package | lowercase, singular, short | `uart`, `ringbuffer` |
| Class / struct / enum / typedef | PascalCase | `RingBuffer`, `PinState` |
| Function / method | lowerCamelCase | `readByte()` |
| Member variable | lowerCamelCase + `_` | `baud_` |
| Local / parameter | lowerCamelCase | `timeoutMs` |
| Constant / enum constant / macro | camelCase | `MaxRetryCount` |
| One-operation type | `<Operation>er` | `Reader` |
| Error value | `Err` + camelCase | `ErrTimeout` |
| Error type | PascalCase + `Error` | `ParseError` |
| Getter / setter | `name()` / `setName()` | `baud()` / `setBaud()` |
| File | lowercase, no underscore | `uart.suckc` |

## Maintenance rule

This document must stay in sync with the compiler. Whenever a code change adds,
removes, or renames a language construct that these rules cover — a new
declaration form, a new keyword — this document must be updated in the same
change. See `AGENTS.md`.