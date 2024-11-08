package main

import (
	"fmt"
	"net"
	"os"
)

var VALKEY = NewSys()

func main() {
	port := "6379"
	if len(os.Args) >= 3 {
		port = os.Args[2]
	}

	l, err := net.Listen("tcp", "0.0.0.0:"+port)
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
