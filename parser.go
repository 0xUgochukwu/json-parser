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

	var lastToken Token

	for {
		pos, tok, _ := lexer.Lex()
		if pos.line == 1 && pos.column == 1 {
			if tok != OpenBrace {
				fmt.Println(tok, OpenBrace)
				fmt.Println("Invalid 1")
				os.Exit(1)
			}
		}
		if tok == EOF {
			if lastToken == CloseBrace {
				fmt.Println("Valid")
				os.Exit(0)
			} else {
				fmt.Println("Invalid 2")
				os.Exit(1)
			}
		}

		lastToken = tok
	}

	os.Exit(0)
}
