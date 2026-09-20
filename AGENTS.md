# SuckC Compiler

## Overview

See [README.md](README.md) for project description, design goals, roadmap,
current status, and build/run instructions.

## Dependencies

- cobra (spf13) for the CLI
- Go stdlib for everything else

## Code Conventions

- Standard Go formatting (`gofmt`), stdlib `testing` for tests
- No CI, no linter scripts yet
- Reuse: prefer existing libraries, built-in functionality, and logic already
  in the project over writing new code from scratch
- Readability and maintainability: favor simple, self-documenting code that
  a new contributor can understand and modify with confidence