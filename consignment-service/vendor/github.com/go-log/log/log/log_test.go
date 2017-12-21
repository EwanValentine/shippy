package log

import (
	"testing"

	"github.com/go-log/log"
)

func testLog(l log.Logger) {
	l.Log("test\n")
}

func testLogf(l log.Logger) {
	l.Logf("%s", "test\n")
}

func TestLogLogger(t *testing.T) {
	l := new(logLogger)
	testLog(l)
	testLogf(l)
}
