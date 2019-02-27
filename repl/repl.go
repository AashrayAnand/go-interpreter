package repl

import (
	"bufio"
	"fmt"
	"go-interpreter/lexer"
	"go-interpreter/token"
	"io"
)

const PROMPT = "go-interpreter>> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		// endless loop

		// print prompt
		fmt.Printf(PROMPT)

		// scan input, if no input, exit REPL
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		// get scanned text
		line := scanner.Text()
		// create new lexer, calling New function from
		// lexer package, which creates lexer with
		// scanned line as input
		lexer := lexer.New(line)

		// iterate through the tokens, while we have not reached an EOF token
		for tok := lexer.NextToken(); tok.Type != token.EOF; tok = lexer.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
