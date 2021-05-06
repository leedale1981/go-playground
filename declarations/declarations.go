package main

import "fmt"

// constant block declaration with iota
const (
	first  = iota
	second = iota + 1
)

func main() {
	var third int = first + 1 // explicit type declaration
	fourth := "Fourth"        // implicit type declaration
	fmt.Println(first, second, third, fourth)

	// Explicit pointer declaration & assignment
	var ptr *string = new(string)
	*ptr = "Hello"
	fmt.Println(ptr)  // print address of pointer
	fmt.Println(*ptr) // print value of pointer

	address := &ptr
	fmt.Println(*address) // print address of pointer
}
