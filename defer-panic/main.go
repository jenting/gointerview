package main

import "fmt"

func main() {
	deferFunc()
}

func deferFunc() {
	defer a()
	defer b()
	defer c()

	panic("I am panic.")

	// output
	//
	// c
	// b
	// a
	// panic: I am panic.
}

func a() {
	fmt.Println("a")
}

func b() {
	fmt.Println("b")
}

func c() {
	fmt.Println("c")
}
