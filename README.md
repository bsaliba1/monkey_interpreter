# Monkey Interpreter

An interpreter for the Monkey programming language, built by following "Writing an Interpreter in Go" by Thorsten Ball. Monkey is a simple, dynamically-typed language that supports integers, booleans, strings, arrays, hashes, and first-class functions.

The interpreter uses a lexer (found in the `lexer` folder) and a recursive descent parser using "top down operator precedence" (Pratt parsing).

## Usage

To start the REPL:

```bash
go run .
```

## Useful commands
- `go test -bench=. <folder_name>` - to run all benchmarks in that folder **(with unit tests)**
- `go test -bench=. -run=^# <folder_name>` - to run all benchmarks in that folder and **skip unit tests**

## Notes
### Read–Eval–Print Loop (REPL) a.k.a Language Shell
Simple interactive programming program that "reads" an input, "evaluates" it, and "prints" the output, and then loops.


### Lexers
- lexical Analyser (Lexer): takes an input string containing source code and turns it into tokens with types

### Parser
- Parser: take source code as input (either as text or in our case tokens) and turns it into a useable data structure -- usually a hierarchical structure or in this case an abstract sytax tree (AST)
- To see the sybolic expression tree in ruby you can do:

```ruby
require 'ripper'

# <<~RUBY is a heredoc that allows for formatted strings ("~" removes preceding whitespace)
code = <<~RUBY
  [1, 2, 3].each do |n|
    puts n
  end
RUBY

pp Ripper.sexp(code)
```

-
