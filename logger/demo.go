package logger

import (
	"errors"
	"log"
	"os"
	"sync"
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

// DemoV6 demonstrates default logger.
func DemoV6(logger Logger) {
	if logger == nil {
		logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	}
	err := doTheThing()
	if err != nil {
		logger.Println("error in doTheThing():", err)
		logger.Printf("error in doTheThing(): %v\n", err)
	}
}

// ThingV2 injects as a struct with default logger getter.
type ThingV2 struct {
	Logger interface {
		Println(...interface{})
		Printf(string, ...interface{})
	}

	once sync.Once
}

func (t *ThingV2) logger() Logger {
	t.once.Do(func() {
		if t.Logger == nil {
			t.Logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
		}
	})
	return t.Logger
}

// DemoV7 uses struct as dependency.
func (t *ThingV2) DemoV7() {
	err := doTheThing()
	if err != nil {
		t.logger().Println("error in doTheThing():", err)
		t.logger().Printf("error in doTheThing(): %v\n", err)
	}
}
func doTheThing() error {
	return errors.New("error opening file: nat.txt")
}
