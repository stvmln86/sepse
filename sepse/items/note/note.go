// Package note implements the Note type and methods.
package note

import (
	"fmt"
	"regexp"
	"time"

	"github.com/stvmln86/sepse/sepse/tools/neat"
)

// Note is a single published message.
type Note struct {
	Body string
	Time time.Time
}

// Regex is the regular expression used to parse Notes.
var Regex = regexp.MustCompile(`^(\d+) (.+)$`)

// New returns a new Note.
func New(body string, time time.Time) *Note {
	body = neat.Body(body)
	return &Note{body, time}
}

// Parse returns a new Note from a parsed string.
func Parse(text string) (*Note, error) {
	regx := Regex.FindStringSubmatch(text)
	if len(regx) == 0 {
		return nil, fmt.Errorf("cannot parse Note %q", text)
	}

	body := neat.Body(regx[2])
	tobj := neat.Time(regx[1])
	return New(body, tobj), nil
}

// String returns the Note as a string.
func (n *Note) String() string {
	unix := neat.Unix(n.Time)
	return fmt.Sprintf("%s %s", unix, n.Body)
}
