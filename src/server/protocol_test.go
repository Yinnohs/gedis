package server

import (
	"testing"
)

const testRequest string = "*3\r\n$3\r\nSET\r\n$5\r\ntestkey\r\n$3\r\ntestvalue\r\n"

func TestProtocol(t *testing.T) {
	command, err := ParseCommand(testRequest)
}
