// Package conn implements connection input/output functions.
package conn

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

// Read returns a newline-ended string from a Conn.
func Read(conn net.Conn, tout time.Duration) (string, error) {
	conn.SetReadDeadline(time.Now().Add(tout))
	text, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("cannot read from %q - %w", conn.RemoteAddr(), err)
	}

	return text, nil
}

// Write writes a newline-ended string into a Conn.
func Write(conn net.Conn, tout time.Duration, text string) error {
	conn.SetWriteDeadline(time.Now().Add(tout))
	if _, err := fmt.Fprintf(conn, "%s\r\n", text); err != nil {
		return fmt.Errorf("cannot write into %q - %w", conn.RemoteAddr(), err)
	}

	return nil
}
