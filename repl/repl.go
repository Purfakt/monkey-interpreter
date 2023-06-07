package repl

import (
	"bufio"
	"fmt"
	"io"

	"monkey/lexer"
	"monkey/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		currentLine := lexer.New(line)

		for tok := currentLine.NextToken(); tok.Type != token.EOF; tok = currentLine.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
