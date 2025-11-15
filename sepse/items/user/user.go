// Package user implements the User type and methods.
package user

import (
	"bufio"
	"fmt"
	"net"
)

// User is a single connected service client.
type User struct {
	Conn  net.Conn
	Lanes map[string]bool
}

// New returns a new User.
func New(conn net.Conn, lanes ...string) *User {
	user := &User{conn, make(map[string]bool)}
	for _, lane := range lanes {
		user.Lanes[lane] = true
	}

	return user
}

// Close closes the User's connection.
func (u *User) Close() error {
	if err := u.Conn.Close(); err != nil {
		return fmt.Errorf("cannot close User %q - %w", u.Conn.RemoteAddr(), err)
	}

	return nil
}

// Read returns a newline-ended string from the User.
func (u *User) Read() (string, error) {
	text, err := bufio.NewReader(u.Conn).ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("cannot read from User %q - %w", u.Conn.RemoteAddr(), err)
	}

	return text, nil
}

// Subscribe adds a lane name to the User.
func (u *User) Subscribe(lane string) {
	u.Lanes[lane] = true
}

// Unsubscribe removes a lane name from the User.
func (u *User) Unsubscribe(lane string) {
	delete(u.Lanes, lane)
}

// Write writes a formatted string to the User.
func (u *User) Write(text string, elems ...any) error {
	if _, err := fmt.Fprintf(u.Conn, text, elems...); err != nil {
		return fmt.Errorf("cannot write to User %q - %w", u.Conn.RemoteAddr(), err)
	}

	return nil
}

// WriteLane writes a formatted string to the User if the User contains a lane name.
func (u *User) WriteLane(lane, text string, elems ...any) error {
	if u.Lanes[lane] {
		return u.Write(text, elems...)
	}

	return nil
}
