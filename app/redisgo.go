package main

import (
	"fmt"
	"net"
)

var VALKEY = NewSys()

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		panic(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}

		go handle(conn)
	}
}

func handle(c net.Conn) {
	for {
		buf := make([]byte, 512)

		_, err := c.Read(buf)
		if err != nil {
			c.Close()
			return
		}

		l := NewLexer(buf)
		p := NewParser(c, l)
		fmt.Printf("\nORIGINAL: \n%s", PrettyPrintBytes(buf))
		p.parseRequest()
	}
}
