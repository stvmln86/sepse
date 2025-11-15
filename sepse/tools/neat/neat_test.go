package neat

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	mockTime = time.Unix(1234567890, 0).Local()
)

func TestBody(t *testing.T) {
	// success
	body := Body("\tBody.\n")
	assert.Equal(t, "Body.", body)
}

func TestLane(t *testing.T) {
	// success
	lane := Lane("\tlane\n")
	assert.Equal(t, "LANE", lane)
}

func TestTime(t *testing.T) {
	// success
	tobj, err := Time("\t1234567890\n")
	assert.Equal(t, mockTime, tobj)
	assert.NoError(t, err)

	// failure - cannot parse
	tobj, err = Time("nope")
	assert.Zero(t, tobj.Unix())
	assert.EqualError(t, err, `cannot parse time "nope" - invalid format`)
}

func TestUnix(t *testing.T) {
	// success
	unix := Unix(mockTime)
	assert.Equal(t, "1234567890", unix)
}
