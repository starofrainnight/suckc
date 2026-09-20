# SuckC Compiler

## Build & Run

```sh
go build ./...
go run ./cmd/suckc SuckC.suckc
go run ./cmd/suckc -d SuckC.suckc      # debug flag (stub this phase)
go test ./...
```

## Architecture

- Go rewrite in progress — cobra CLI at `cmd/suckc`; ANTLR4 grammar
  (`src/*.g4`) preserved for future frontend integration.

## Dependencies

- cobra (spf13) for the CLI
- Go stdlib for everything else

## Code Conventions

- Standard Go formatting (`gofmt`), stdlib `testing` for tests
- No CI, no linter scripts yet
