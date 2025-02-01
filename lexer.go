package main

import (
	"bufio"
	"fmt"
	"io"
)

type Token int

const (
	EOF = iota
	OpenBrace
	CloseBrace
	String
	Colon
	Comma
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

func (l *Lexer) ReadString() string {
	var s string

	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				panic("EOF in string")
			}

		}
		l.pos.column++
		if r == '"' {
			break
		}
		s += string(r)
	}

	return s
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
		case '"':
			return l.pos, String, l.ReadString()
		case ':':
			return l.pos, Colon, ":"
		case ',':
			return l.pos, Comma, ","
		default:
			if r != ' ' && r != '\t' && r != '\r' {
				panic(fmt.Sprintf("Unexpected character: %c", r))
			}
		}

	}
}
