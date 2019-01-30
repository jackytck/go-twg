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

func doTheThing() error {
	return errors.New("error opening file: nat.txt")
}
