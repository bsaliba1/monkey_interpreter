package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/bsaliba1/monkey_interpreter/lexer"
	"github.com/bsaliba1/monkey_interpreter/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		_ = scanner.Scan() // await content
		line := scanner.Text() // convert content to text
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() { // convert text to tokens
			fmt.Printf("%+v\n", tok)
		}
	}
}


