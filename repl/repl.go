package repl

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"tgifer/commands"
	"tgifer/lexer"
	"tgifer/parser"
)

func Start(in io.Reader, out io.Writer, args []string) {
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

	if len(args) == 0 {
		bytes, err := json.MarshalIndent(program, "", " ")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(bytes))
		return
	}

	if args[0] == "strings" {
		ret := commands.GetStrings(program)
		for _, r := range ret {
			fmt.Println(r)
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
