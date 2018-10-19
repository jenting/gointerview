package main

import "fmt"

type structA struct {
}

func (a *structA) methodA() {
	fmt.Println("structA methodA")
	a.methodB()
}

func (a *structA) methodB() {
	fmt.Println("structA methodB")
}

type structB struct {
	structA
}

func (b *structB) methodB() {
	fmt.Println("structB methodB")
}

func main() {
	b := structB{}
	b.methodA()

	// output
	//
	// structA methodA
	// structA methodB
}
