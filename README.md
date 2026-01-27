# Monkey Interpreter

An interpreter for the Monkey programming language, built by following "Writing an Interpreter in Go" by Thorsten Ball. Monkey is a simple, dynamically-typed language that supports integers, booleans, strings, arrays, hashes, and first-class functions.

The interpreter uses a lexer (found in the `lexer` folder) and a recursive descent parser using "top down operator precedence" (Pratt parsing).

## Usage

To start the REPL:

```bash
go run .
```

## Useful commands
- `go test -bench=. <folder_name>` - run all benchmarks in that folder (with unit tests)
- `go test -bench=. -run=^# <folder_name>` - run all benchmarks in that folder and skip unit tests
