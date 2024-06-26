package srslog

import (
	"fmt"
	"os"
	"time"
)

// Formatter is a type of function that takes the consituent parts of a
// syslog message and returns a formatted string. A different Formatter is
// defined for each different syslog protocol we support.
type Formatter func(p Priority, hostname, app, tag, content string) string

// DefaultFormatter is the original format supported by the Go syslog package,
// and is a non-compliant amalgamation of 3164 and 5424 that is intended to
// maximize compatibility.
func DefaultFormatter(p Priority, hostname, app, tag, content string) string {
	timestamp := time.Now().Format(time.RFC3339)
	msg := fmt.Sprintf("<%d> %s %s %s[%d]:%s %s",
		p, timestamp, hostname, app, os.Getpid(), tag, content)
	return msg
}

// UnixFormatter omits the hostname, because it is only used locally.
func UnixFormatter(p Priority, hostname, app, tag, content string) string {
	timestamp := time.Now().Format(time.Stamp)
	msg := fmt.Sprintf("<%d>%s %s[%d]:%s %s",
		p, timestamp, app, os.Getpid(), tag, content)
	return msg
}

// RFC3164Formatter provides an RFC 3164 compliant message.
func RFC3164Formatter(p Priority, hostname, app, tag, content string) string {
	timestamp := time.Now().Format(time.Stamp)
	msg := fmt.Sprintf("<%d>%s %s %s[%d]:%s %s",
		p, timestamp, hostname, app, os.Getpid(), tag, content)
	return msg
}

// if string's length is greater than max, then use the last part
func truncateStartStr(s string, max int) string {
	if (len(s) > max) {
		return s[len(s) - max:]
	}
	return s
}

// RFC5424Formatter provides an RFC 5424 compliant message.
func RFC5424Formatter(p Priority, hostname, app, tag, content string) string {
	timestamp := time.Now().Format(time.RFC3339)
	pid := os.Getpid()
	msg := fmt.Sprintf("<%d>%d %s %s %s %d %s - %s",
		p, 1, timestamp, hostname, app, pid, tag, content)
	return msg
}
