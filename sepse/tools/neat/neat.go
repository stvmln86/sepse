// Package neat implements value sanitisation functions.
package neat

import (
	"strconv"
	"strings"
	"time"
)

// Body returns a whitespace-trimmed body string.
func Body(body string) string {
	return strings.TrimSpace(body)
}

// Lane returns an uppercase whitespace-trimmed lane string.
func Lane(lane string) string {
	lane = strings.ToUpper(lane)
	return strings.TrimSpace(lane)
}

// Time returns a local Time object from a Unix UTC string.
func Time(unix string) time.Time {
	unix = strings.TrimSpace(unix)
	uint, _ := strconv.ParseInt(unix, 10, 64)
	return time.Unix(uint, 0).Local()
}

// Unix returns a Unix UTC string from a local Time object.
func Unix(tobj time.Time) string {
	uint := tobj.Unix()
	return strconv.FormatInt(uint, 10)
}
