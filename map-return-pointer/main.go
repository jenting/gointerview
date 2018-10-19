package main

import (
	"fmt"
)

type student struct {
	Name string
	Age  int
}

func a() map[string]*student {
	ret := make(map[string]*student)

	stus := []student{
		{"A", 24},
		{"B", 21},
		{"C", 23},
	}

	for _, stu := range stus {
		ret[stu.Name] = &stu
	}

	return ret
}

func main() {
	ret := a()

	for k, v := range ret {
		fmt.Printf("key: %s, val: %v\n", k, v)
	}

	// output
	//
	// key: A, val: &{C 23}
	// key: B, val: &{C 23}
	// key: C, val: &{C 23}
}
