package note

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func mockNote() *Note {
	return New("body", mockTime())
}

func mockTime() time.Time {
	return time.Unix(1234567890, 0).Local()
}

func TestNew(t *testing.T) {
	// success
	note := New("body", mockTime())
	assert.Equal(t, "body", note.Body)
	assert.Equal(t, mockTime(), note.Time)
}

func TestNewParse(t *testing.T) {
	// success
	note := NewParse("1234567890 body")
	assert.Equal(t, "body", note.Body)
	assert.Equal(t, mockTime(), note.Time)
}

func TestString(t *testing.T) {
	// success
	text := mockNote().String()
	assert.Equal(t, "1234567890 body", text)
}
