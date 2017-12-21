package log

import (
	"testing"
)

type testLogger struct{}

func (t *testLogger) Log(v ...interface{}) {
	return
}

func (t *testLogger) Logf(format string, v ...interface{}) {
	return
}

func testLog(l Logger) {
	l.Log("test\n")
}

func testLogf(l Logger) {
	l.Logf("%s\n", "test")
}

func TestLogger(t *testing.T) {
	l := new(testLogger)
	testLog(l)
	testLogf(l)
}

func TestNoOpLogger(t *testing.T) {
	l := new(noOpLogger)
	testLog(l)
	testLogf(l)
}
