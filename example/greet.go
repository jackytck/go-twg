// Package example is a package for demostrating examples in source code.
package example

import "fmt"

type Demo struct{}

func (d *Demo) Hello() {
	fmt.Println("Hello")
}

// Hello prints out hello to the person provided.
func Hello(name string) (string, error) {
	return fmt.Sprintf("Hello, %s!", name), nil
}
