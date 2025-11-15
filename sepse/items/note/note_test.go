package note

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	mockTime = time.Unix(1234567890, 0).Local()
)

func TestNew(t *testing.T) {
	// success
	note := New("body", mockTime)
	assert.Equal(t, "body", note.Body)
	assert.Equal(t, mockTime, note.Time)
}

func TestParse(t *testing.T) {
	// success
	note, err := Parse("1234567890 body")
	assert.Equal(t, "body", note.Body)
	assert.Equal(t, mockTime, note.Time)
	assert.NoError(t, err)

	// failure - cannot parse Note
	note, err = Parse("nope")
	assert.Nil(t, note)
	assert.EqualError(t, err, `cannot parse Note "nope"`)
}

func TestString(t *testing.T) {
	// setup
	note := New("body", mockTime)

	// success
	text := note.String()
	assert.Equal(t, "1234567890 body", text)
}
