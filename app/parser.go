package main

import "net"

const (
	I_PING = "PING"
	I_ECHO = "ECHO"
	I_SET  = "SET"
	I_GET  = "GET"
)

type Parser struct {
	lex  *Lexer
	conn net.Conn
}

func NewParser(c net.Conn, lex *Lexer) *Parser {
	return &Parser{
		lex:  lex,
		conn: c,
	}
}

func (p *Parser) parseRequest() {
	t := p.lex.Peek()

	if t.First == T_ARRAY {
		p.parseArray()
	} else {
		panic("PARSER ERR: unhandled request token")
	}
}

func (p *Parser) parseArray() {
	p.lex.GetToken()
	p.parseBulk()
}

func (p *Parser) parseBulk() {
	t := p.lex.Peek()

	if IsCommand(t.Second) {
		p.parseCommand()
	}
}

func (p *Parser) parseCommand() {
	t := p.lex.GetToken()

	switch t.Second {
	case I_PING:
		p.ping()
	case I_ECHO:
		p.echo()
	case I_GET:
  p.get()
	case I_SET:
		p.set()
	}
}
