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

// Name returns an uppercase whitespace-trimmed name string.
func Name(name string) string {
	name = strings.ToUpper(name)
	return strings.TrimSpace(name)
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
