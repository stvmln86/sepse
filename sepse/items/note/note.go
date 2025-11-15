// Package note implements the Note type and methods.
package note

import (
	"fmt"
	"strings"
	"time"

	"github.com/stvmln86/sepse/sepse/tools/neat"
)

// Note is a single published message.
type Note struct {
	Body string
	Time time.Time
}

// New returns a new Note with a body and local Time object.
func New(body string, time time.Time) *Note {
	body = neat.Body(body)
	return &Note{body, time}
}

// NewParse returns a Note from a parsed string.
func NewParse(note string) *Note {
	unix, body, _ := strings.Cut(note, " ")
	tobj := neat.Time(unix)
	return New(body, tobj)
}

// String returns the Note as a string.
func (n *Note) String() string {
	unix := neat.Unix(n.Time)
	return fmt.Sprintf("%s %s", unix, n.Body)
}
