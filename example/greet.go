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

func Page(checkins map[string]bool) {
	for name, checkIn := range checkins {
		if !checkIn {
			fmt.Printf("Paging %s: please see the front desk to check in.\n", name)
		}
	}
}
