package server

import "fmt"

// const ProtocolTerminator string = "CRLF"
// var binaryProtocolTerminator []byte = []byte{13, 10}
// Write the protocol

type Command struct {
}

func ParseCommand(msg string) (Command, error) {
	t := rune(msg[0])
	fmt.Println(t)

	switch t {
	case '*':
		fmt.Println(msg[:1])
	}
	return Command{}, nil
}
