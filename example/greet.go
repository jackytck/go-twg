package example

import "fmt"

// Hello prints out hello to the person provided.
func Hello(name string) (string, error) {
	return fmt.Sprintf("Hello, %s!", name), nil
}
