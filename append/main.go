package main

import "fmt"

func main() {
	a := make([]int, 5)
	a = append(a, 1, 2, 3)
	fmt.Println(a)

	// output
	//
	// [0 0 0 0 0 1 2 3]
}
