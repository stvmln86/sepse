// Package neat implements value sanitisation and parsing functions.
package neat

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Body returns a whitespace-trimmed body string.
func Body(body string) string {
	return strings.TrimSpace(body)
}

// Lane returns an uppercase whitespace-trimmed lane name string.
func Lane(lane string) string {
	lane = strings.ToUpper(lane)
	return strings.TrimSpace(lane)
}

// Time returns a local Time object from a Unix UTC string.
func Time(unix string) (time.Time, error) {
	unix = strings.TrimSpace(unix)
	uint, err := strconv.ParseInt(unix, 10, 64)
	if err != nil {
		return time.Unix(0, 0), fmt.Errorf("cannot parse time %q - invalid format", unix)
	}

	return time.Unix(uint, 0).Local(), nil
}

// Unix returns a Unix UTC string from a local Time object.
func Unix(tobj time.Time) string {
	uint := tobj.Unix()
	return strconv.FormatInt(uint, 10)
}
