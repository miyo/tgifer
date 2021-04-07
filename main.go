package main

import (
	"os"
	"tgifer/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
