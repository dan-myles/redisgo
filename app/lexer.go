package main

import "fmt"

var (
	String  byte = '+'
	Bulk    byte = '$'
	Integer byte = ':'
	Array   byte = '*'
	Error   byte = '-'
	CR      byte = '\r'
	LF      byte = '\n'
)

const (
	T_STRING      = "STRING"
	T_BULK_STRING = "BULK_STRING"
	T_INTEGER     = "INTEGER"
	T_ARRAY       = "ARRAY"
	T_ERROR       = "ERROR"
	T_EOF         = "EOF"
)

type Token Pair[string, string]

type Lexer struct {
	buf []byte
}

func NewLexer(b []byte) *Lexer {
	return &Lexer{
		buf: b,
	}
}

func (l *Lexer) GetToken() Token {
	if len(l.buf) == 0 {
		panic("PROTOCOL ERR: unexpected end of buffer")
	}

	t := l.buf[0]
	l.buf = l.buf[1:]

	switch t {
	case Array:
		return l.array()
	case Bulk:
		return l.bulk()
	default:
		fmt.Printf("HORRIBLE THING IN GET TOKEN")
		return Token{First: T_EOF, Second: ""}
	}
}

func (l *Lexer) Peek() Token {
	originalBuf := l.buf
	token := l.GetToken()
	l.buf = originalBuf

	return token
}

func (l *Lexer) array() Token {
	return Token{First: T_ARRAY, Second: l.readUntilCRLF()}
}

func (l *Lexer) bulk() Token {
	l.readUntilCRLF()
	return Token{First: T_BULK_STRING, Second: l.readUntilCRLF()}
}

func (l *Lexer) readUntilCRLF() string {
	for i := 0; i < len(l.buf)-1; i++ {
		if l.buf[i] == CR && l.buf[i+1] == LF {
			raw := string(l.buf[:i])
			l.buf = l.buf[i+2:]
			return raw
		}
	}

	panic("PROTOCOL ERR: CRLF ending could not be found")
}
