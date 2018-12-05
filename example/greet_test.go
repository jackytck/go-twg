package example

import "fmt"

func ExampleHello() {
	greeting, err := Hello("Nat")
	if err != nil {
		panic(err)
	}
	fmt.Println(greeting)

	// Output:
	// Hello, Nat!
}
