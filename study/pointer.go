package main

import (
	"fmt"
)

func createPointer() *float64 {
	var myFloat = 99.5
	return &myFloat
}

func printPointer(myBoolPointer *bool) {
	fmt.Println(*myBoolPointer)
}

func double(number *int) {
	*number *= 2
}

func negate(myBoolean *bool) {
	*myBoolean = !*myBoolean
}

func main() {
	var myInt int
	var myIntPointer *int
	myInt = 42
	myIntPointer = &myInt
	fmt.Println(*myIntPointer)

	var myFloatPointer *float64 = createPointer()
	fmt.Println(*myFloatPointer)

	var myBool bool = true
	printPointer(&myBool)

	amount := 6
	double(&amount)
	fmt.Println(amount)

	truth := true
	negate(&truth)
	fmt.Println(truth)
	lies := false
	negate(&lies)
	fmt.Println(lies)
}
