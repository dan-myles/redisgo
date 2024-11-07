package main

import (
	"fmt"
	"net"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		panic(err)
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			panic(err)
		}

		go handle(c)
	}
}

func handle(c net.Conn) {
	for {
		buf := make([]byte, 256)

		_, err := c.Read(buf)
		if err != nil {
			fmt.Printf("there was an error: %s", err.Error())
		}

		c.Write([]byte("+PONG\r\n"))
	}
}
