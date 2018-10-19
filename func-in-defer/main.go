package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	a := 1
	b := 2
	defer calc("10", a, calc("1", a, b))
	b = 0
	defer calc("20", a, calc("2", a, b))
	a = 0

	// output
	//
	// 1 1 2 3
	// 2 1 0 1
	// 20 1 1 2
	// 10 1 3 4
}
