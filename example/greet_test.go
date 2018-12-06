package example

import "fmt"

func Example() {
	fmt.Println("Package level")

	// Output:
	// Package level
}

func Example_second() {
	fmt.Println("Package level second")

	// Output:
	// Package level second
}

func ExampleDemo() {
	fmt.Println("Struct level")

	// Output:
	// Struct level
}

func ExampleDemo_Hello() {
	fmt.Println("Method level")

	// Output:
	// Method level
}

func ExampleDemo_Hello_second() {
	fmt.Println("Method level second")

	// Output:
	// Method level second
}

func ExampleHello() {
	greeting, err := Hello("Nat")
	if err != nil {
		panic(err)
	}
	fmt.Println(greeting)

	// Output:
	// Hello, Nat!
}
