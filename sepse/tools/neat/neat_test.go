package neat

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBody(t *testing.T) {
	// success
	body := Body("\tBody.\n")
	assert.Equal(t, "Body.", body)
}

func TestName(t *testing.T) {
	// success
	name := Name("\tname\n")
	assert.Equal(t, "NAME", name)
}

func TestTime(t *testing.T) {
	// setup
	want := time.Unix(1234567890, 0).Local()

	// success
	tobj := Time("1234567890")
	assert.Equal(t, want, tobj)
}

func TestUnix(t *testing.T) {
	// setup
	tobj := time.Unix(1234567890, 0).Local()

	// success
	unix := Unix(tobj)
	assert.Equal(t, "1234567890", unix)
}
