package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("No File Passed")
		os.Exit(2)
	}

	jsonFile, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	lexer := NewLexer(jsonFile)

	// parser

	for {
		_, tok, val := lexer.Lex()
		fmt.Println(val)
		if tok == EOF {
			break
		}
	}
}
