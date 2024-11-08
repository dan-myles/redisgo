package main

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

	p.conn.Write(RESPOk())
}

func (p *Parser) get() {
	key := p.lex.GetToken().Second
	val := VALKEY.Get(key)

	p.conn.Write(RESPBulkFromString(val))
}
