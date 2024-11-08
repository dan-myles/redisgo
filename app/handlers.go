package main

import (
	"strconv"
	"time"
)

func (p *Parser) ping() {
	p.conn.Write([]byte("+PONG\r\n"))
}

func (p *Parser) echo() {
	echo := p.lex.GetToken().Second
	resp := RESPBulkFromString(echo)
	p.conn.Write(resp)
}

func (p *Parser) set() {
	key := p.lex.GetToken().Second
	val := p.lex.GetToken().Second
	VALKEY.Set(key, val)

	if p.lex.Peek().Second == "px" {
		p.lex.GetToken()
		raw := p.lex.GetToken().Second

		expr, _ := strconv.Atoi(raw)
		ms := time.Duration(expr) * time.Millisecond

		go func() {
			time.Sleep(ms)
			VALKEY.Delete(key)
		}()
	}

	p.conn.Write(RESPOk())
}

func (p *Parser) get() {
	key := p.lex.GetToken().Second
	val, ok := VALKEY.Get(key)

	if !ok {
		p.conn.Write(RESPNull())
		return
	}

	p.conn.Write(RESPBulkFromString(val))
}
