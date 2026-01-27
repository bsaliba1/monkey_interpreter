# Notes

## REPL (Read-Eval-Print Loop)

Also known as a language shell. A simple interactive programming environment that:
1. **Reads** an input
2. **Evaluates** it
3. **Prints** the output
4. **Loops** back to read the next input

## Lexer

A lexical analyzer (lexer) takes an input string containing source code and transforms it into tokens with types.

## Parser

A parser takes source code as input (either as text or, in our case, tokens) and turns it into a usable data structure—usually a hierarchical structure called an Abstract Syntax Tree (AST).

### Viewing S-expressions in Ruby

To see the symbolic expression tree in Ruby:

```ruby
require 'ripper'

code = <<~RUBY
  [1, 2, 3].each do |n|
    puts n
  end
RUBY

pp Ripper.sexp(code)
```
