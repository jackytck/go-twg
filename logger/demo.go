package logger

import (
	"errors"
	"log"
	"os"
)

// DemoV1 is the original implementation.
func DemoV1() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	err := doTheThing()
	if err != nil {
		logger.Println("error in doTheThing():", err)
	}
}

// DemoV2 demonstrates injecting logger into it as args.
func DemoV2(logger *log.Logger) {
	err := doTheThing()
	if err != nil {
		logger.Println("error in doTheThing():", err)
	}
}

// DemoV3 injects as a function.
func DemoV3(logFn func(...interface{})) {
	err := doTheThing()
	if err != nil {
		logFn("error in doTheThing():", err)
	}
}

// Logger represents the dependent interface.
type Logger interface {
	Println(...interface{})
	Printf(string, ...interface{})
}

// DemoV4 injects with interface.
// logger := log.New(...)
// DemoV4(logger)
func DemoV4(logger Logger) {
	err := doTheThing()
	if err != nil {
		logger.Println("error in doTheThing():", err)
		logger.Printf("error in doTheThing(): %v\n", err)
	}
}

// Thing injects as a struct.
// Then no need to pass around the Logger.
type Thing struct {
	Logger interface {
		Println(...interface{})
		Printf(string, ...interface{})
	}
}

// DemoV5 uses struct as dependency.
func (t Thing) DemoV5() {
	err := doTheThing()
	if err != nil {
		t.Logger.Println("error in doTheThing():", err)
		t.Logger.Printf("error in doTheThing(): %v\n", err)
	}
}

func doTheThing() error {
	return errors.New("error opening file: nat.txt")
}
