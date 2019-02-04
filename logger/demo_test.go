package logger_test

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/jackytck/twg/logger"
)

type fakeLogger struct {
	sb strings.Builder
}

func (f *fakeLogger) Println(v ...interface{}) {
	fmt.Fprintln(&f.sb, v...)
}

func (f *fakeLogger) Printf(format string, v ...interface{}) {
	fmt.Fprintf(&f.sb, format, v...)
}

func TestDemoV3(t *testing.T) {
	var fl fakeLogger
	logger.DemoV3(fl.Println)
	want := "error in doTheThing():"
	got := fl.sb.String()
	if !strings.Contains(got, want) {
		t.Errorf("Logs = %q; want substring %q", got, want)
	}
}

func TestDemoV4(t *testing.T) {
	var fl fakeLogger
	logger.DemoV4(&fl)
	want := "error in doTheThing():"
	got := fl.sb.String()
	if !strings.Contains(got, want) {
		t.Errorf("Logs = %q; want substring %q", got, want)
	}
}

func TestDemoV5(t *testing.T) {
	var sb strings.Builder
	testLogger := log.New(&sb, "", 0)
	thing := logger.Thing{
		Logger: testLogger,
	}
	thing.DemoV5()

	want := "error in doTheThing():"
	got := sb.String()
	if !strings.Contains(got, want) {
		t.Errorf("Logs = %q; want substring %q", got, want)
	}
}

func TestDemoV5b(t *testing.T) {
	var testLogger fakeLogger
	thing := logger.Thing{
		Logger: &testLogger,
	}
	thing.DemoV5()

	want := "error opening file: nat.txt"
	got := testLogger.sb.String()
	if !strings.Contains(got, want) {
		t.Errorf("Logs = %q; want substring %q", got, want)
	}
}

func TestDemoV7(t *testing.T) {
	var testLogger fakeLogger
	thing := logger.ThingV2{
		Logger: &testLogger,
	}
	thing.DemoV7()

	want := "error opening file: nat.txt"
	got := testLogger.sb.String()
	if !strings.Contains(got, want) {
		t.Errorf("Logs = %q; want substring %q", got, want)
	}
}
