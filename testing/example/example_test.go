package example_test

import (
	"fmt"

	"github.com/hsiaoairplane/gotest/example"
)

func ExampleToLower() {
	fmt.Println(example.ToLower("HELLO"))
	fmt.Println(example.ToLower("hello"))

	// Output:
	// hello
	// hello
}
