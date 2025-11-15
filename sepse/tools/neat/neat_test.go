package neat

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func mockTime() time.Time {
	return time.Unix(1234567890, 0).Local()
}

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
	// success
	tobj := Time("\t1234567890\n")
	assert.Equal(t, mockTime(), tobj)
}

func TestUnix(t *testing.T) {
	// success
	unix := Unix(mockTime())
	assert.Equal(t, "1234567890", unix)
}
