// Package lane implements the Lane type and methods.
package lane

import (
	"fmt"
	"net"
	"sync"

	"github.com/stvmln86/sepse/sepse/items/note"
	"github.com/stvmln86/sepse/sepse/tools/neat"
)

// Lane is a single named message channel.
type Lane struct {
	Name  string
	Conns map[net.Conn]bool
	Mutex *sync.Mutex
}

// New returns a new Lane.
func New(name string, conns ...net.Conn) *Lane {
	name = neat.Name(name)
	cmap := make(map[net.Conn]bool)
	lane := &Lane{name, cmap, new(sync.Mutex)}
	for _, conn := range conns {
		lane.Add(conn)
	}

	return lane
}

// Add adds a Conn to the Lane.
func (l *Lane) Add(conn net.Conn) {
	l.Mutex.Lock()
	l.Conns[conn] = true
	l.Mutex.Unlock()
}

// Close closes all the Conns in the Lane.
func (l *Lane) Close() error {
	for conn := range l.Conns {
		if err := conn.Close(); err != nil {
			return fmt.Errorf("cannot close Lane - %w", err)
		}
	}

	return nil
}

// Remove removes a Conn from the Lane.
func (l *Lane) Remove(conn net.Conn) {
	l.Mutex.Lock()
	delete(l.Conns, conn)
	l.Mutex.Unlock()
}

// Send sends a Note to all Conns in the Lane.
func (l *Lane) Send(note *note.Note) error {
	for conn := range l.Conns {
		if _, err := fmt.Fprintf(conn, "%s %s\n", l.Name, note); err != nil {
			return fmt.Errorf("cannot send Note - %w", err)
		}
	}

	return nil
}
