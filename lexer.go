package main

import (
	"bufio"
	"io"
)

type Token int

const (
	EOF = iota
	OpenBrace
	CloseBrace
)

type Position struct {
	line   int
	column int
}

type Lexer struct {
	pos    Position
	reader *bufio.Reader
}

func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{
		pos:    Position{line: 1, column: 0},
		reader: bufio.NewReader(reader),
	}
}

func (l *Lexer) ResetPosition() {
	l.pos.line++
	l.pos.column = 0
}

func (l *Lexer) Lex() (Position, Token, string) {
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return l.pos, EOF, ""
			}
			panic(err)
		}
		l.pos.column++

		switch r {
		case '\n':
			l.ResetPosition()
		case '{':
			return l.pos, OpenBrace, "{"
		case '}':
			return l.pos, CloseBrace, "}"
		}

	}
}
