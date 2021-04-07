package repl

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"tgifer/lexer"
	"tgifer/parser"
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	var s = ""
	for {
		scanned := scanner.Scan()
		if !scanned {
			break
		}
		s += scanner.Text() + "\n"
	}
	l := lexer.New(s)
	p := parser.New(l)
	program := p.ParseProgram()

	bytes, err := json.MarshalIndent(program, "", " ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bytes))
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
