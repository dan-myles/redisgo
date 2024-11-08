package main

import (
	"bytes"
	"strconv"
)

type Pair[T, U any] struct {
	First  T
	Second U
}

func PrettyPrintBytes(data []byte) string {
	var result string

	for i := 0; i < len(data); i++ {
		if data[i] == '\r' && i+1 < len(data) && data[i+1] == '\n' {
			result += `\r\n`
			i++
		} else if data[i] == '\n' {
			result += `\n`
		} else if data[i] == '\r' {
			result += `\r`
		} else {
			result += string(data[i])
		}
	}

	return result
}

func IsCommand(bulk string) bool {
	switch bulk {
	case I_ECHO:
		return true
	case I_PING:
		return true
	default:
		return false
	}
}

func RESPFromString(raw string) []byte {
	var buf bytes.Buffer

	buf.WriteString("$")
	buf.WriteString(strconv.Itoa(len(raw)))
	buf.WriteString("\r\n")
	buf.WriteString(raw)
	buf.WriteString("\r\n")

	return buf.Bytes()
}
